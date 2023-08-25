---
title: 封装并发布一个库到npm
date: 2021-08-25 14:27:16
tags: typescript、npm
---

# 封装并发布一个TS库到npm

## 1.初始化

### 1.初始化package.json

创建文件夹，这里使用test-array-map。

使用`npm init`初始化一个package.json，如果想默认初始化可以使用`npm init -y`。这里自己配置。

```
{
  "name": "zzz-array-map", // 包的名字
  "version": "1.0.0", // 版本
  "description": "array map funtion use ts", // 描述
  "main": "./dist/test-array-map.js", // 内容的路径
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  }, // 配置的执行脚本
  "keywords": [
    "typescript"
  ], // 关键字
  "author": "zzz", // 作者
  "license": "MIT" // 开源协议
}
```

### 2.初始化typescript

使用`tsc --init`初始化tsconfig.json。

```
// tsconfig.json
// 修改一下配置
{
"declaration": true // 打包后是否生成声明文件，设置为true后不仅会生成一个.d.ts文件还会生成一个js文件。
"outDir": "./dist" // 输出目录，编译后输出的目录放到dist文件夹下。
},
// 打包后会排除掉一下文件夹或者文件。
  "exclude": [
    "./dist",
    "./example"
  ]

```

因为要使用TS，这里就需要安装TS依赖。使用`npm install typescript -D`安装TS依赖。

## 2.创建库文件

### 1.创建纯js中使用的库文件

创建test-array-map.ts文件，模拟原生js的map方法。

```
// map方法有三个参数，第一个参数是当前遍历项。第二个参数是索引。第三个参数是遍历后返回的数组。
// 这里array模拟需要遍历的数组。回调函数里的三个参数就是map方法里的三个参数。
// 遍历的数组可以是任意类型的数组所以这里定义any类型，回调函数里第一项就是传入的每一项所以也是任意类和传入的数组类型相同所以也使用any。索引就是数值类型。第三个参数就是传入数组的引用也用any的数组。回调函数的结果返回就是any类型，arrayMap的返回值类型也是any类型的数组。
const arrayMap = (array: any[], callback: (item: any, index: number, arr: any[]) => any): any[] => {
	// 初始索引值
  let i = -1
  // 数组长度
  const len = array.length
  // 空数组放结果
  const resArray = []
  // 遍历数组
  while(++i < len) {
  // 执行完遍历返回的结果
    resArray.push(callback(array[i], i, array))
  }
  // 返回结果数组
  return resArray
}
export = arrayMap
```

在终端中使用`tsc`编译文件，就会在dist文件夹中生成编译后的文件。

编译结束后测试一下，`cd dist`文件中使用`node`来执行文件。

```
const arrayMap = require('./test-array-map')
arrayMap([1, 2], (item) => {
	return item + 1
})
```

现在在纯js的项目就可以直接使用了。

### 2.改良文件

创建测试文件夹example，文件夹下创建test.ts文件

```
// test.ts文件下
import arrayMap = require('../dist/test-array-map')
console.log(arrayMap([1, 2], (item) => {
  return item + 2
}))
```

终端中使用`cd example`切换至example文件夹下，使用`tsc test.ts`编译文件。

使用`node test.ts`执行文件。打印出来说明执行成功了。

现在改良文件

```
// 现在传入的类型和回调函数第一个参数的类型是一样，索引这里定义泛型变量T。回调函数的第三个参数arr，是传入参数的引用，我们不希望它可以更改所以这里使用内置的只读属性并且元素类型为T。回调函数返回的类型应该和方法返回的类型一样，因为方法是由回调函数组成的数组，所以这里使用了U泛型变量。
const arrayMap = <T, U>(array: T[], callback: (item: T, index: number, arr: ReadonlyArray<T>) => U): U[] => {
  let i = -1
  const len = array.length
  const resArray = []
  while(++i < len) {
    resArray.push(callback(array[i], i, array))
  }
  return resArray
}
export = arrayMap
```

```
import arrayMap = require('../dist/test-array-map')
arrayMap([1, 2], (item) => {
  return item + 2
}).forEach((item) => {
  console.log(item.length)
})
```

按理说数值类型是不会拥有length属性，所以改良以后再重新执行一下。

切换回上级目录`cd ..`，重新编译一下`tsc`，先再它就会报错。

```
import arrayMap = require('../dist/test-array-map')
arrayMap([1, 2], (item) => {
  return item + 2
}).forEach((item) => {
  item.toFixed()
})
```

### 3.上传至npm

#### 1.上传

在test-array-map路径下创建.npmignore文件用来过滤不想上传的文件或文件夹。现在只想上传dist文件夹，node_modules是默认不上传的，使用`/example/`这里就过滤掉example文件夹。

执行`npm publish`上传包，输入自己的npm账号、密码、与邮箱。

#### 2.测试npm包

在package.json中的name与包名字不能一样。所以修改一下。使用`npm install (包名字)`，就可以大功告成了。

#### 3.更新包

在package.json中修改version就可以修改包的版本了，再使用`npm publish`更新包。使用`npm install (包名字)@latest`可以下载最新的包的版本。
