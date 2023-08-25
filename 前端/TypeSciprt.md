# TypeScript

## 1.搭建环境

`
npm init //初始化
`

```
packname: //包的名字
version: //版本号
description: //描述
entry point: //入口文件，默认为index.js
test command: //测试指令
git repository: //git源
keywords: //关键字
author: //作者
license: //协议，常用(MIT)
```

`npm init -y //简便写法`

构建目录：

├─typings	//ts模块声明文件
├─src		//项目中的文件
|  ├─utils	//业务相关可复用的方法
|  ├─tools	//跟业务无关纯工具函数
|  ├─config	//配置文件，可能会修改的配置抽离出来
|  ├─assets	//静态资源
|  |   ├─img	//图片
|  |   ├─font	//字体文件
|  ├─api	//可复用接口方法
├─build	//打包上线的配置、本地开发的配置，一般为webpack配置

依赖：`npm install typescript tslint -g`

`tsc --init //初始化typecsript`

环境：webpack4	安装依赖: `npm install webpack@4.29.6 webpack-cli@3.2.3 webpack-dev-server -D`

安装依赖: `npm install ts-loader@5.3.3 -D`

安装依赖: `npm install cross-env -D`

安装依赖: `npm install clean-webpack-plugin html-webpack-plugin@3.2.0 -D`

安装依赖: `npm install typescript`

配置webpack.config.js文件

`build下创建webpack.config.js文件` `src下创建index.ts文件:项目编译的入口文件` ``

```
//webpack.config.js文件下
const HtmlWebpackPlugin = require('html-webpack-plugin')
const { CleanWebpackPlugin } = require('clean-webpack-plugin')

module.exports = {
  // 项目编译的入口文件
  entry: "./src/index.ts",
  // 指定项目输出文件
  output: {
    // 第一个选项fielname
    filename: "main.js",
  },
  resolve: {
    // 自动解析文件拓展，这个选项是一个数组
    extensions: ['.js', '.ts', '.tsx']
  },
  module: {
    // 指定后缀文件的处理
    rules: [{
      // 指定匹配的文件后缀，是一个正则表达式，这里匹配的是ts活tsx后缀的文件
      test: /\.tsx?$/,
      // 使用ts-loader处理
      use: 'ts-loader',
      // 排除文件，编译不去处理的文件
      exclude: /node_modules/
    }]
  },
  // source-map，在调试时定位到代码。只在开发模式时使用，打包时不适用能加快打包速度和文件大小。
  devtool: process.env.NODE_ENV === 'production' ? false : 'inline-source-map',
  // 
  devServer: {
    // 本地开发环境基于哪个文件夹作为根目录运行的
    contentBase: './dist',
    // 在控制台打印的信息
    stats: 'errors-only',
    // 不启动压缩
    compress: false,
    // 域名
    host: 'localhost',
    // 端口
    port: 8080
  },
  //插件
  plugins: [
    // 清理指定文件夹
    new CleanWebpackPlugin({
      cleanOnceBeforeBuildPatterns: ['./dist']
    }),
    // 指定编译的模板html文件
    new HtmlWebpackPlugin({
      template: './src/template/index.html'
    })
  ]
}
```

```
//package.json文件下
"scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    //应用webpack打包的指令配置，使用cross-env依赖传递development环境
    "start": "cross-env NODE_ENV=development webpack-dev-server --config ./build/webpack.config.js",
    //配置打包指令，使用webpack打包。
    "build": "cross-env NODE_ENV=prodction webpack --config ./build/webpack.config.js"
  },
```

## 2.基础类型

- 布尔值
- 数值
- 字符串
- 数组
- 元组
- 枚举值
- any
- void
- null和undefined
- never
- object
- 类型断言

在src下创建example文件夹，在文件夹下创建basic-type.ts。example存放示例代码

### 布尔类型

```
// let bool: boolean = false
let bool: boolean
bool = true
// bool = 123 //如果不是布尔类型就会报错
```
### 数值类型
```
// 数值类型
let num: number = 123
// num = 'abc' //报错
//ts支持 二进制 八进制 十进制 十六进制
num = 0b111011
num = 0o173
num = 0x7b
```
### 字符串类型
```
// 字符串类型
let str: string
str = 'abc'
str = `数值是${num}`  //123
```
### 数组类型
```
// 数组类型
//[1, 2, 3]
//写法1
let arr: number[] //arr数组都是number类型
//arr = [5, 'b']  //报错，不能将string类型分配给number
```
```
// 写法二
let arr2: Array<number> //指定数组类型，里面只能放number
let arr3: (string | number)[] //联合类型，[]代表数组类型。这里要加()不然会被判定为不是string类型就是number类型
arr3 = [1, 'a']
```
### 元组类型
```
// 元组类型
// 类似数组的类型。和数组的区别，数组可以随意长度随意类型。元组为固定长度固定类型
let tuple: [string, number, boolean] //指定元组类型，需要一一对应的填写
tuple = ['a', 1, false] //如果类型不同会报错，元组指定了定义的length，如果多写就会报错。超出的长度称为越界元素
```
### 枚举类型
```
// 枚举类型   //
enum Roles {  //习惯枚举值大写开头
  SUPER_ADMIN = 1,  //枚举的值会有一个默认的序列号，可以更改序列号。
  ADMIN = 4,
  USER  //5
}
```
### any类型
```
// any类型  任何类型
let value: any
value = 123
value = 'asd'
value = false
const arr4: any[] = [1, 'a']
```
### void类型
```
// void类型 和any类型相反
const consoleText = (text: string): void => {
  console.log(text)
}
let v: void //void类型可以赋值undef和null，需要在关闭严格检查 "strict"
v = undefined
v = null
consoleText('123')
```
### null和undefined
```
//null和undefined
let u: undefined  
u = undefined
let n: null
n = null

num = undefined
num = null  //如果开了严格检查就不能赋给num，只能赋给定义了undefined和null类型的变量
```
### never类型
```
// never类型 永远不存在的类型 
const errorFunc = (message: string): never => {
  throw new Error(message)  //报错的类型是never类型
}
const infiiteFunc = (): never => {
  while(true){} //不可能又返回值的类型就是never类型
}
// never类型是任何类型的子类型，never可以赋值给任何类型，任何类型不能赋值给never类型
let neverVariable = (() => {
  while(true){}
})
// neverVariable = 123 //报错
```
### object类型
```
// object类型
let obj = {
  name: 'fish'
}
let obj2 = obj
obj2.name = 'Z' //内存的引用
console.log(obj)  //Z
function getObject(obj: object): void {
  console.log(obj)
}
//getObject(123) //报错
getObject({obj2}) //只能传入对象
```
### 类型断言
```
// 类型断言
//使用场景：ts没有我们了解一个值的类型，让ts不做类型检查，我们自己知道它是什么类型。将某个值强行指定为一个类型
const getLength = (target: string | number): number => {
  if ((<string>target).length || (target as string).length === 0) {
    return (<string>target).length  //如果传入的是字符串是有长度的 //返回值也要定义类型断言
  } else {
    return target.toString().length //如果传入的是数值先转换为字符串
  }
}
//第一种写法 
//(<string>target).length
//第二种写法
//(target as string).length 在react中只jsx中只能使用这种方法
getLength(123)
```

## 3.ES6-symbol

- 基础
- 作为属性名
- 属性名的遍历
- symbol.for和symbol.keyFor
- 11个内置symbol值

使用Symbol

```
 const s = Symbol()
 console.log(s)
```


出现`无法调用各类型缺少调用签名的表达式`，需要在tsconfig.json中使用lib并添加es6或es2015

```
//tsconfig.json文件下
    "lib": [
      "es6","DOM"	//如果还需要使用dom就添加dom这个库
    ]
```

### 基础

Symbols是不可改变且唯一的。

```
const s = Symbol()
console.log(s)

const s2 = Symbol()
console.log(s2)

console.log(s === s2) //false

const s3 = Symbol('zzz')
console.log(s3)

const s4 = Symbol(123)  //传入的数字将toString转换为字符串

const s5 = Symbol({a: 'a'}) //在ts中报错，只能传入number类型和string类型。//在浏览器中返回值是Symbol([Object Object])对象转为字符串

s4 + 12 //symbol不能做运算
console.log(s4.toString())  //字符串形式的Symbol(123)
console.log(Boolean(s4))  //true
```

### 作为属性名

像字符串一样，symbols也可以被用做对象属性的键。

```
let prop = 'name'
const info = {
  // name: 'zz',
  [prop]: 'zz',
  [`my${prop}`]: 'zz'
}
console.log(info) //{name: 'zz}
```

### 属性名的遍历

```
const s5 = Symbol('name')
const info2 = {
  [s5]: 'zzz',
  age: 18,
  sex: 'man'
}
console.log(info2)  //{Symbol(name): "zzz"}
info2[s5] = 'haha'
//info2.s5 = 'xxx' //无法修改
console.log(info2)  //{Symbol(name): "haha"}

for (const key in info2) {
  console.log(key)  //age sex，不会打印Symbol值
}

console.log(Object.keys(info2)) //["age", "sex"]，不会打印Symbol值

console.log(Object.getOwnPropertyNames(info2)) //["age", "sex"]，不会打印Symbol值

console.log(JSON.stringify(info2))  //{"age":18,"sex":"man"}，不会打印Symbol值

console.log(Object.getOwnPropertySymbols(info2))  //[Symbol(name)] ，返回Symbol值

console.log(Reflect.ownKeys(info2))   //["age", "sex", Symbol(name)]。es6的Reflect方法，返回所有类型组成的数组。
```

### symbol.for和symbol.keyFor

```
//静态方法
//Symbol.for()方法 
//会拿创建的字符串在全局中搜索是否有相同的Symbol值，如果有直接返回字符串创建的这个值，如果没有就会创建一个新的Symbol值
const s6 = Symbol.for('zzz')
const s7 = Symbol.for('zzz')
const s8 = Symbol.for('xxx')
//s6 === s7 //true
//s7 === s8 //false

//Symbol.keyFor()方法
//接收一个参数，是使用Symbol.for()创建的值，返回Symbol.for()全局注册的标识
console.log(Symbol.keyFor(s6))  //zzz

```

### 11个内置的Symbol值

`Symbol.hasInstance`

```
Symbol.hasInstance  //决定一个构造器对象是否认可一个对象时它的实例
// instanceof 使用关键字判断一个对象是否是它的实例的执行的方法，即确定一个对象实例的原型链上是否有原型。
const obj1 = {
  [Symbol.hasInstance] (otherObj) {
    console.log(otherObj) //{a: "a"}
  }
}
console.log({a: 'a'} instanceof <any> obj1) //false
```

`Symbol.isConcatSpreadabl`

```
Symbol.isConcatSpreadabl 可读写的布尔值，当一个数组设为true时在concat中不会扁平化。
//Symbol.isConcatSpreadabl 默认值为undefined
let arr6 = [1, 2]
console.log([].concat(arr6, [3, 4]))  //[1, 2, 3, 4]

arr6[Symbol.isConcatSpreadable] = false
console.log([].concat(arr6, [3, 4]))  //[Array(2), 3, 4]
```

`Symbol.species`

```
//Symbol.species 创建衍生对象的构造函数
class C extends Array {
  constructor (...args) {
    super(...args)
  }

  static get [Symbol.species] () {
    return Array
  }

  getName () {
    return 'aaa'
  }
}
const c = new C(1, 2, 3)
const a = c.map(item => item + 1)
console.log(a)  //[2, 3, 4]
console.log(a instanceof C) //false
console.log(a instanceof C) //如果不使用[Symbol.species] 则为true
console.log(a instanceof Array) //true
```

`Symbol.match`

```
//Symbol.match //方法，被String.prototype.match调用。正则表达式用来匹配字符串。可以用作计算作为返match方法的返回值
let obj3 = {
  [Symbol.match] (string) {
    console.log(string.length) //4
  },
  [Symbol.split] (string) {
    console.log('split', string.length) //split 4
  }
}
'abcd'.match(<RegExp>obj3)
'abcd'.split(<any>obj3)

//Symbol.replace  //类似match
//Symbol.search //类似match
//Symbol.split  //类似match
```

`Symbol.iterator`

```
Symbol.iterator // 返回对象默认的迭代器
const arr8 = [1, 2, 3]
const iterator = arr8[Symbol.iterator] ()
console.log(iterator) //Array Iterator {} 返回对象，对象下有一个next()方法
console.log(iterator.next())  //{value: 1, done: false}
console.log(iterator.next())  //{value: 2, done: false}
console.log(iterator.next())  //{value: 3, done: false}
console.log(iterator.next())  //没有数据了 {value: undefined, done: true}
```

`Symbol.toProimitive `

```
//Symbol.toProimitive //方法，被ToPrimitive抽象操作调用。把对象转换为相应的原始值。
let obj9: unknown = {
  [Symbol.toPrimitive] (type) {
    console.log(type)
  }
}
// const res = (obj9 as number)++  //number
// const res = `abc${obj9}`  //default
```

`Symbol.toStringTag`

```
Symbol.toStringTag //方法，被内置方法Object.prototype.toString调用。返回创建对象时默认的字符串描述。
let obj11 = {
  // [Symbol.toStringTag]: 'zzz'
  get [Symbol.toStringTag] () {
    return 'zzz'
  }
}
console.log(obj11.toString()) //都返回的时[object zzz]
```

`Symbol.unscopables`

```
Symbol.unscopables  //对象，它自己拥有的属性会被with作用域排除在外。
const obj22 = {
  a: 'a',
  b: 'b'
}
// with (obj22) {
//   console.log(a)  //a
//   console.log(b)  //b
// }
console.log(Array.prototype[Symbol.unscopables])  
// {copyWithin: true, entries: true, fill: true, find: true, findIndex: true, …}
// copyWithin: true
// entries: true
// fill: true
// find: true
// findIndex: true
// flat: true
// flatMap: true
// includes: true
// keys: true
// values: true
```

## 4.接口

- 基本用法
- 可选属性
- 多余属性检查
- 绕开多余属性检查
- 只读属性
- 函数类型
- 索引类型
- 继承接口
- 混合类型接口

使用`tslint --init`初始化配置，使用vscode需要安装tslint相应的插件。

在example中创建interface.ts，如果引入interface时报错，需要在tslint.json的rule中添加`"quotemark": [false]`

```
    "rules": {
      "quotemark": [false]，
      "semicolon": [false],
      //从来不需要前缀
      "interface-name": [true, "never-prefix"]
    },
```

如果提示要添加分号可以，可以使用`"semicolon": [false]`关闭报错。

在settings.json中配置`  "tslint.autoFixOnSave": true`可以修复部分ts的报错

### 基本用法

```
interface NameInfo {
  firstName: string,
  lastName: string
}

const getFullName = ({ firstName, lastName }: NameInfo): string // 定义返回类型 => { 
  return `${firstName}${lastName}]`
}

console.log(getFullName({
  firstName: 'haha',
  lastName: "12" // 报错
}))
```

### 可选属性

```
interface Vegetable {
  color?: string, // '?'代表可选属性，
  type: string,
}

const getVegetables = ({ color, type }: Vegetable) => {
  return `A ${color ? (color + '') : ''}${type}`
}

console.log(getVegetables({
  // color: 'red',	//如果不设置可选参数，并且不传入color就会报错。
  type: 'tomato'
}))
```

### 多余属性检查

```
interface Vegetable {
  color?: string,
  type: string,
}

const getVegetables = ({ color, type }: Vegetable) => {
  return `A ${color ? (color + '') : ''}${type}`
}

console.log(getVegetables({
  // color: 'red',
  type: 'tomato',
  size: 2 // 传入的参数多了一个则会报错
}))
```

### 绕开多余属性检查

第一种解决方法，使用类型断言as Vegetable

```
interface Vegetable {
  color?: string,
  type: string,
}

const getVegetables = ({ color, type }: Vegetable) => {
  return `A ${color ? (color + '') : ''}${type}`
}

console.log(getVegetables({
  // color: 'red',
  type: 'tomato',
  size: 2
}as Vegetable))	//这里使用了类型断言
```

第二种方法使用索引签名[prop: string]: any

```
interface Vegetable {
  color?: string,
  type: string,
  [prop: string]: any	// 这里使用了索引签名
}

const getVegetables = ({ color, type }: Vegetable) => {
  return `A ${color ? (color + '') : ''}${type}`
}

console.log(getVegetables({
  // color: 'red',
  type: 'tomato',
  size: 2
}))
```

第三种方法，使用类型兼容性。 意思vegetableInfo传进来的参数多了无所谓，本身有的参数必须有。

```
interface Vegetable {
  color?: string,
  type: string,
}

const getVegetables = ({ color, type }: Vegetable) => {
  return `A ${color ? (color + '') : ''}${type}`
}

const vegetableInfo = ({
  color: 'red',
  type: 'tomato',
  size: 2
})

console.log(getVegetables(vegetableInfo))	//这里使用了类型兼容性
```

### 只读属性

```
interface Vegetable {
  color?: string,
  // 只读属性 只需在前面添加一个readonly
  readonly type: string,
}

let vegetableObj: Vegetable = {	//使用Vegetable接口
  type: 'tomato'
}

vegetableObj.type = 'carrot' //报错，type不能修改

```

对象结构

```
interface ArrInter {
  0: number,
  readonly 1: string
}

let arr1: ArrInter = [1, '2']
```

### 函数类型

```
type AddFunc = (num1: number, num2: number) => number

const add: AddFunc = (n1, n2) => n1 + n2  // 这里的参数名字不必相同，只需位置和个数对应就行。

const add: AddFunc = (n1, n2) => `${n1}, ${n2}` //报错，string不能将类型分配给number
```

### 索引类型

```
interface RoleDic {
  [id: number]: string
}
const role1: RoleDic = {
  0: 'super_admin' 
  'a': 'super_admin' //报错，索引只能是数值类型
}
```

```
interface RoleDic {
  [id: string]: string
}
const role2: RoleDic = {
  a: 'sper_admin',
  1: 'admin'  // 如果索引是数值会隐式转换为字符串
}
```

### 继承接口

```
interface Vegetables {
  color: string
}
interface Tomato extends Vegetables {
  radius: number
}
interface Carrot {
  length: number
}
// 这里就这里继承了color，所以传参时必须添加color
const tomato: Tomato = {
  radius: 1,
  color: 'red'
}
```

### 混合类型接口

```
interface Counter {
  (): void,	//不是任何类型
  count: number
}

const getCounter = (): Counter => {
  const c = () => {c.count++}
  c.count = 0
  return c
}
const counter: Counter = getCounter()
counter()
console.log(counter.count)  // 1
counter()
console.log(counter.count)  // 2
counter()
console.log(counter.count)  // 3
```

## 5.函数

**函数类型**
- 为函数定义类型

- 完整的函数类型

- 使用接口定义函数类型

- 使用类型别名

**参数**

- 可选参数

- 默认参数

- 剩余参数

**重载**

在example中创建funciton.ts

### 函数类型

#### 为函数定义类型

```
// ES5以前的方法
function add1(arg1: number, arg2: number): number {
  return arg1 + arg2
}
// ES6方法
const add = (arg1: number, arg2: number) => arg1 + arg2
```

#### 完整的函数类型

```
let add: (x: number, y: number) => number // 一个完整的函数是一个括号，一个箭头，一个返回的number
add = (arg1: number, arg2: number): number => arg1 + arg2
add = (arg1: string, arg2: number) => arg1 + arg2 // 这里arg1的类型不兼容，要定义成number
let arg3 = 3
add = (arg1: string, arg2: number) => arg1 + arg2 + arg3 // arg3不需要定义到函数里
```

#### 使用接口定义函数类型

```
type Add = (x: number, y: number, z: number) => number
```

#### 使用类型别名

```
type istring = string
let addFunc: Add
addFunc = (arg1: number, arg2: number) => arg1 + arg2
```

### 参数

#### 可选参数

```
type AddFunciton = (arg1: number, arg2: number, arg3?: number) => number
let addFunction: AddFunciton
addFunction = (x: number, y: number) => x + y
addFunction = (x: number, y: number, z: number) => x + y + z
```

#### 默认参数

```
let addFunction = (x: number, y = 3) => x + y
console.log(addFunction(1, 2))  //  3
console.log(addFunction(1, 'a'))  //  报错
```

#### 剩余参数

应用场景，有时候函数的参数个数不一定，需要挨个进行处理。在es5中是使用arguments。

```
// es5
function handleData () {
  if (arguments.length === 1) return arguments[0] * 2
  else if (arguments.length === 2) return arguments[0] * arguments[0]
  else return Array.prototype.slice.apply(arguments).join('_')  // arguments是一个伪数组，这里把arguments转为数组。
}
handleData(2) // 4
handleData(2, 3) // 6
handleData(2, 3, 4) // "2_3_4"

// es6标准的方法
const handleData = (...args) => {console.log(args)}
handleData(1) // [1]
handleData(1, 2) // [1, 2]
handleData(1, 2, 3) // [1, 2, 3]

// ts中
const handleData = (arg1: number, ...args: number[]) => {
  // ...
}
```

### 重载

```
function handleData(x: string): string[]
function handleData(x: number): number[]  // 这两个是函数重载，下面的是函数实体。函数重载只能用于函数，不能用于接口或别名。
function handleData(x: any): any {
  if (typeof x === 'string') {
    return x.split('')
  } else {
    return x.toString().split('').map((item) => Number(item))
  }
}
console.log(handleData('abc')) // ["a", "b", "c"]
console.log(handleData(123))  // [1, 2, 3]

```

## 6.泛型

- 简单使用
- 泛型变量
- 泛型类型
- 泛型约束
- 在泛型约束中使用类型参数

泛型的意义：支持多种数据，同时支持类型结构的检查。

在example下创建generics.ts

### 简单使用

```
// 在定义前面添加<>，里面可以添加任意字母作为泛型变量，习惯写法使用大写。
const getArray = <T>(value: T, times: number = 5): T[] => {
  return new Array(times).fill(value)
}
// 调用泛型变量
console.log(getArray<number>(2, 3))	// [2, 2, 2]
console.log(getArray<number>('abc', 4).map((item) => item.length))	// 报错，这样就可以帮助检查出错误。
```

### 泛型变量

```
// 这里用到2个泛型变量，中间用','隔开。
const getArray = <T, U>(param1: T, param2: U, times: number): [T, U][] =>{
  return new Array(times).fill([param1, param2])
}
getArray<number, string>(1, 'a', 3).forEach((item) => {
  console.log(item[0])	// 1 1 1
  console.log(item[1])	// a a a
})
```

### 泛型类型

泛型定义函数类型

```
let getArray: <T>(arg: T, times: number) => T[]
getArray = (arg: any, times: number) => {
  return new Array(times).fill(arg)
}
getArray(123, 3).map((item) => item.length) // 报错,number类型没有length属性。
```

类型别名

```
type GetArray = <T>(arg: T, times: number) => T[]
let getArray: GetArray = (arg: any, times: number) => {
  return new Array(times).fill(arg)
}
```

接口定义泛型类型

```
interface GetArray {
// 在前面添加泛型变量
  <T>(arg: T, times: number): T[]
}

// 把泛型变量提升到最外侧，接口里定义的都可以使用泛型变量了。 
interface GetArray<T> {
  (arg: T, times: number): T[]
  array: T[]
}
```

### 泛型约束

泛型约束就是对范型变量的一个条件限制。

```
interface ValueWithLength {
// 有length属性的数据，类型为number
  length: number
}
// 范型变量这里继承了接口
const getArray = <T extends ValueWithLength>(arg: T, times): T[] => {
  return new Array(times).fill(arg)
}

getArray([1, 2], 3)
getArray(123, 3)	// 没有length属性，报错。
getArray('123', 3)
// 还可以这样使用
getArray({
  length: 2
}, 3)
```

### 在泛型约束中使用类型参数

```
// 使用keyof索引类型，返回一个对象上所有属性的数组。K相当于联合类型。
const getProps = <T, K extends keyof T>(object: T, propName: K) => {
  return object[propName]
}

const objs = {
  a: 'a',
  b: 'b'
}

getProps(objs, 'a')
getProps(objs, 'c') // 如果不定义范型变量约束将没有报错。定义后将报错，提示输入错误。
```

## 7.ES6中的类

- ES5和ES6实现创建实例
- constructor方法
- 类的实例
- 取值函数和存值函数
- class表达式
- 静态方法
- 实现属性其他写法
- 实现私有方法

### ES5实现创建实例

```
function Point (x, y) {
  this.x = x
  this.y = y
}
Point.prototype.getPosition = function () {
  return '( ' + this.x + ' , ' + this.y + ')'
}
var p1 = new Point(2, 3)
console.log(p1)
console.log(p1.getPosition())
var p2 = new Point(4, 5)
console.log(p2)
console.log(p2.getPosition()) // Point指向要创建的实例，当使用new操作符就会创建一个空的对象，然后给这个对象设置属性x和y。最后把对象返回赋值给p1。getPostion方法是创建在原型对象上的，当使用new操作符时就会继承这getPosition()方法。
```

### ES6实现创建实例

```
class Point {
  constructor(x, y) {
    this.x = x
    this.y = y
    // return {a: 'a'}
  }
  getPosition() {
    return `(${this.x}, ${this.y})`
  }
}
var p1 = new Point(2, 3)
console.log(p1)	// Point {x: 2, y: 3}
```

每个类都有一个constructor构造函数，如果别添加constructor就默认创建一个空的constructor构造函数。这个函数的返回值就是这个实例的对象。也可以用return自己返回东西。如果自己定义了类的结果，那么这个类就不是类的实例了。

### 取值函数和存值函数

再es5中使用

```
var info = {
  _age: 10,
  set age (newValue) {
    if (newValue >10) {
      console.log('大于10')
    } else {
      console.log('小于10')
    }
  },
  get age () {
    console.log('打印get') // 打印get
    return this._age
  }
}
console.log(info.age) // 10
info.age = 11 // 大于10
info.age = 9 // 小于10
```

在es6中使用set/get

```
class Info {
  constructor (age) {
    this.age = age
  }
  set age (newAge) {
    console.log('new age is:' + newAge)
    this._age = newAge
  }
  get age () {
    return this._age
  }
}
const infos = new Info(10)  // new age is:10
infos.age = 9
console.log(info.age) // new age is:9
```

### class表达式

```
// 匿名函数
const func = function () {}	
// 命名函数
function func () {}	

// 命名类
class Infos {
  constructor () {}
}
// 匿名类
const Infos = class {
   constructor () {}
}
const testInfo = new Infos()
```

### 静态方法

```
function testFunc () {}
console.log(testFunc.name)  //testFunc，每个函数都有一个name属性。
class Point {
  constructor(x, y) {
    this.x = x
    this.y = y
  }
  getPosition () {
    return `(${this.x}, ${this.y})`
  }
  static getClassName () {
    return Point.name // 每一个类都有一个name属性
  }
}
const p = new Point(1, 2)
console.log(p.getPosition()) // (1,2)
// console.log(p.getClassName()) // 报错，不是函数。类的静态方法是继承不了的。
console.log(Point.getClassName()) //Point 解决类的静态方法无法继承。使用类自身调用，返回类的name值。
```

### 实现属性其他写法

```
function testFunc () {}
console.log(testFunc.name)  //testFunc，每个函数都有一个name属性。
class Point {
  // z = 0 // 这里可以直接定义，或者再constructor里定义
  constructor(x, y, z) {
    this.x = x
    this.y = y
    this.z = z
  }
  getPosition () {
    return `(${this.x}, ${this.y})`
  }
  static getClassName () {
    return Point.name // 每一个类都有一个name属性
  }
}
```

### 静态属性

```
class Point {
  constructor () {
    this.x = 0
  }
  // static y = 2 还没有通过。
}
Point.y = 2
const p = new Point()
console.log(p.x) // 0
console.log(p.y) // undefined，实现静态属性。
```

###  实现私有方法

```
// es6暂时不支持私有方法，需要使用一些技巧
class Point {
  func1 () {

  }
  // 前面添加_来给其他调用者提示是私有方法
  _func2 () {

  }
}
```

封装成模块，就不能_func2()直接调用

```
const _func2 = () => {}
class Point {
  func1 () {
    _func2.call(this)
  }
}
const p = new Point()
p._func2()  // 无法调用

使用symbol值来实现私有方法
a.js
const func1 = Symbol('func1')
export default class Point {
  static [func1] () {

  }
}
// b.js
import Point from './a.js'
const p = new Point()
console.log(p) // 无法调用
```

私有属性

```
class Point {
  // #为私有属性，暂未发布。
  #ownProp = 12
}
```

new.target属性

一般用于构造函数中，返回new命令构造函数。

```
function Point() {
  // 代表构造函数
  console.log(new.target)
}
const p = new Point() // 返回构造函数
const p2 = Point() // undefied

class Point {
  constructor () {
    // console.log(new.target)
  }
}
const p3 = new Point() // 返回构造函数

class Parent {
  constructor () {
    // console.log(new.target)
    if (new.target === Parent) {
      throw new Error('不能实例化')
    }
  }
}
class Childe extends Parent {
  constructor () {
    super()
  }
}
const c = new Childe() // 打印父类的构造函数
const c = new Parent() // 报错，不能实例化
const c = new Childe() // 顺利运行
```

## 8.es6中的类(进阶)

- es5中的继承
- es6中的类的继承
- Object.getPrototypeOf
- super
  - 作为函数
  - 作为对象
- 类的prototype属性和`__proto__`属性

### es5中的继承

```
function Food () {
  this.type = 'food'
}
Food.prototype.getType = function () {
  return this.type
}
function Vegetable (name) {
  this.name = name
}
Vegetable.prototype = new Food() // 继承实例
const tomato = new Vegetable('tomato')
console.log(tomato.getType()) // food
```

### es6类的继承

```
class Parent {
  constructor (name) {
    this.name = name
  }
  getName () {
    return this.name
  }
  static getName () {
    return this.name
  }
}
class Child extends Parent {
  constructor (name, age) {
    super(name) // 只有调用了super才能使用this
    this.age = age
  }
}
const c = new Child('zz', 10)
console.log(c) // Child {name: "zz", age: 10}
console.log(c.getName()) // zz
console.log(c instanceof Child) // true
console.log(c instanceof Parent) // true 即是子类的实例，又是父类的实例。并且继承了父类的静态方法。
console.log(Child.getName()) // Child。父类继承给了子类所以这里答应的是子类的类名
```

### Object.getPrototypeOf

```
这个方法能狗获得一个构造函数的原型对象
class Parent {
  constructor (name) {
    this.name = name
  }
  getName () {
    return this.name
  }
  static getName () {
    return this.name
  }
}
class Child extends Parent {
  constructor (name, age) {
    super(name) // 只有调用了super才能使用this
    this.age = age
  }
}
console.log(Object.getPrototypeOf(Child) === Parent) // true Child的原型对象就是Parent这个类
```

### super

#### 作为函数

作为函数时代表父类的构造函数constructor，es6标准中constructor必须调用一次super函数，如果没有参数可以不传参数。继承的子类里的constructor必须调用super不然就不能使用this。super方法只能再constructor中调用，其他地方调用会报错。如果父类中使用了this，将指向子类。父类constructor中定义的所有this的参数都会添加到子类身上。

#### 作为对象

```
Parent {
  constructor () {
    this.type = 'parent'
  }
  getName () {
    return this.type
  }
}
Parent.getType = () => {
  return 'is Parent'
}
class Child extends Parent {
  constructor () {
    super()
    console.log('constructor: ' + super.getName())
  }
  getParentName () {
    console.log('getParentName: ' + super.getName())
  }
  static getParentType () {
    console.log('getParentType:' + super.getType())
  }
}
// super在普通方法中
const c = new Child() // constructor: parent 这里的super指代的是父类的原型对象的引用。
c.getParentName() // getParentName: parent
// c.getParentType() // 报错，子类调用的是父类的原型对象，子类不能直接调用父类本身。
// super在静态方法中
Child.getParentType() // getParentType:is Parent 这样就会调用Paret.getType
理解super
class Parent {
  constructor() {
    this.name = 'parent'
  }
  print () {
    console.log(this.name)
  }
}
class Child extends Parent {
  constructor () {
    super()
    this.name = 'child'
  }
  childPrint () {
    super.print() // 这里使用super调用父类的print方法
  }
}
const c = new Child()
c.childPrint() // child，这里调用父类中的print，父类的this指向子类。
// super当作函数调用或当作一个对象并且访问函数或方法。不要只使用super不做操作。
```

### 类的prototype属性和`__proto__`属性

```
let obj = new Object()
console.log(obj.__proto__ === Object.prototype) // true
```

子类的`__proto__`指向父类本身

子类prototype属性的`__proto__`指向父类的prototype属性

实例的`__proto__`属性的`__proto__`指向父类实例的`__proto__`

#### 在es5中原生构造是没法继承的

Boolean

Number

String

Array

Date

Function

RegExp

Error

Object

#### 在es6中支持原生构造函数的继承

```
class CustomArray extends Array {
  constructor (...args) {
    super(...args)
  }
}
const arr = new CustomArray(3) // CustomArray(3) [empty × 3]，返回3个空为的数组
arr.fill('-')
console.log(arr) // CustomArray(3) ["-", "-", "-"]
console.log(arr.join('+')) // -+-+-
```

```
class CustomArray extends Array {
  constructor (...args) {
    super(...args)
  }
}
const arr = new CustomArray(3, 4, 5)
console.log(arr.join('+')) // 3+4+5
```

## 9.TS中的类

- 基础
- 修饰符
- readonly修饰符
- 参数属性
- 静态属性
- 可选类属性
- 存储器
- 抽象类
- 实例类型
- 补充知识

### 基础

```
class Point {
  public x: number
  public y: number
  constructor(x: number, y: number) {
    this.x = x
    this.y = y
  }
  getPosition () {
    return `(${this.x},${this.y})`
  }
}
const point = new Point(1, 2)
console.log(point) // Point {x: 1, y: 2}

class Parent {
  public name: string
  constructor(name: string) {
    this.name = name
  }
}
class Child extends Parent {
  constructor(name: string) {
    super(name)
  }
}
```

### 修饰符

#### public

公共的，可以通过创建实例访问的。就是类外部可以访问的属性和方法。

#### private

私有的，修饰的属性在类的外面没办法访问。

```
class Parent {
  private age: number
  constructor(age: number) {
    this.age = age
  }
  // 只能够在类的内部使用age属性
}
const p = new Parent(18)
// console.log(p.age) // 报错，不能够在外部使用
// console.log(Parent.age) // 报错，在实例上访问不到，在类上也访问不到
class Child extends Parent {
  constructor(age: number) {
    super(age)
    console.log(super.age) // 报错，使用继承也是无法访问私有修饰符定义的类
  }
}
```

#### protected

受保护的修饰符,和private相似,protected在继承该类的子类时可以访问

```
class Parent {
  protected age: number
  // 给父类的constructor添加protected只能给子类继承使用
  protected constructor(age: number) {
    this.age = age
  }
  protected getAge() {
    return this.age
  }
  // 只能够在类的内部使用age属性
}
// 父类的constructor将不能创建实例,只能同过子类创建实例.
const p = new Parent(18)
console.log(p.age) // 报错，不能够在外部使用
console.log(Parent.age) // 报错，在实例上访问不到，在类上也访问不到
```

```
class Child extends Parent {
  constructor(age: number) {
    super(age)
    // 通过super只能拿到基类的公共方法和受保护方法,属性不能访问到.
    // console.log(super.age)
    // 继承的方法可以使用
    // console.log(super.getAge())
  }
}
// 只能同过子类创建实例
const child = new Child(19)
```

### readonly修饰符

```
class UserInfo {
  // readonly只读属性,如果不添加修饰符会自动添加Public修饰符
  public readonly name: string
  constructor(name: string) {
    this.name = name
  }
}
const userinfo = new UserInfo('zz')
console.log(userinfo) // UserInfo {name: "zz"}
userinfo.name = 'ss' // 报错,因为它时只读属性.
```

### 参数属性

```
class A {
    constructor(name: string) {}
  }
  const a1 = new A('zz')
  // constructor中没有定义,直接调用返回空对象
console.log(a1) // A {}

class A {
  constructor(public name: string) {}
}
const a1 = new A('zz')
// 使用修饰符,调用后能够打印出定义的值.
console.log(a1) // A {name: "zz"}
```

### 静态属性

```
class Parent {
  public static age: number = 18
  public static getAge() {
    return Parent.age
  }
  constructor() {}
}
const p = new Parent()
console.log(p.age) // 报错,无法访问类的静态属性
console.log(Parent.age) // 18
```

```
class Parent {
  public static getAge() {
    return Parent.age
  }
  private static age: number = 18
  constructor() {}
}
const p = new Parent()
console.log(p.age) // 报错
console.log(Parent.age) // 报错,现在age是私有属性.只能在Parent中访问
```

### 可选类属性

```
class Info {
  public name: string
  public age?: number
  constructor(name: string, age?: number, public sex?: string) {
    this.name = name
    this.age = age
  }
}
const info1 = new Info('zzz')
console.log(info1) // Info {sex: undefined, name: "zzz", age: undefined} 这里设置了age但是没有传参数所以为undefined
const info3 = new Info('zzz', 18)
console.log(info3) // Info {sex: undefined, name: "zzz", age: 18} sex使用了public但是没有传参数所以这里显示undefined
const info4= new Info('zzz', 18, 'man')
console.log(info4) // Info {sex: "man", name: "zzz", age: 18}
```

### 存储器

```
class Info {
  public name: string
  public age?: number
  private _infoStr: string
  constructor(name: string, age?: number, public sex?: string) {
    this.name = name
    this.age = age
  }
  // 取值器
  get infoStr() {
    return this._infoStr
  }
  // 存值器
  set infoStr(value) {
    console.log(`setter: ${value}`)
    this._infoStr = value
  }
}
const info1 = new Info('zzz', 18, 'man')
info1.infoStr = 'zzz: 18' // 使用存值器 打印:setter: zzz: 18
console.log(info1) // 使用取值器 打印:Info {sex: "man", name: "zzz", age: 18, _infoStr: "zzz: 18"}
```

### 抽象类

```
// 抽象类一般用来被其他类继承而不自己创建实例
abstract class People {
  constructor(public name: string) {}
  // 抽象类和类内部定义的方法都使用abstract定义
  public abstract printName(): void
}
// const p1 = new People() // 不能直接创建抽象类的实例
class Man extends People {
  constructor(name: string) {
    super(name)
    this.name = name
  }
  // 非抽象类“Man”不会实现继承自“People”类的抽象成员“printName”,如果要继承就是实现.
  public printName() {
    console.log(this.name)
  }
}
const m = new Man('zzz')
m.printName() // 调用方法打印 zzz
```

```
// TS在2.0以后的版本在abstract可以标记类和类里面的方法,还可以标记类里面的存取器.
// abstract class People {
//   abstract _name: string
//   abstract get insideName(): string
//   // 存值器不能标记返回值类型
//   abstract set insideName(value: string)
// }
// class P extends People {
//   public _name: string
//   public insideName: string
// }
// 抽象方法和抽象存取器都不能包含实际的代码块,只需标记它的属性名、方法名、方法参数和返回值类型就可以了.存值器函数不需要标返回值类型,如果标记了返回值类型就会报错.
```

### 实例类型

```
class People {
  constructor(public name: string) {}
}
let p2: People = new People('zzz')
class Animal {
  constructor(public name: string) {}
}
p2 = new Animal('zxc')
```

### 补充知识

```
interface FoodInterface {
  // 接口里的代码块可以用逗号、分号、或者直接换行隔开。
  type: string,
  name: string;
  age: number
}
// 使用implements实现接口
class FoodClass implements FoodInterface {
  // 里面要和定义的接口相同不然会报错。
  type: string
  name: string
  age: number
}
```

```
// 接口检测是类定义的实例
class FoodClass implements FoodInterface {
  public static type: string // 实例上没有所以使用静态属性会报错
  name: string
  age: number
}
```

```
// 接口也可以继承类，接口继承了类以后只继承成员不包括实现。就是只继承成员和成员类型。接口还会继承private和protected
class A {
  protected name: string
}
interface I extends A {}
class B extends A implements I {
  // B使用I里继承的name时需要定义name并且要继承A，因为A的name属性是受保护只允许子类中使用。
  public name: string
}
```

```
// 泛型中使用类类型
const create = <T>(c: new() => T): T=> {
  return new c()
}
class Infos {
  public age: number
}
create<Infos>(Infos)
```

调用create传入一个类，返回的是类创建的实例。参数c里的new()表示调用类的构造函数它类型就是类创建实例后的类型。return出来的是用new创建的传进来的实例。调用new()构造函数后返回的类型是T。

`create(Infos)`默认使用T类型，也可以自己使用`create<Infos>(Infos)`来使用类类型。

## 10.枚举

数字枚举

反向映射

字符串枚举

异构枚举

枚举成员类型和联合枚举类型

运行时的枚举

const enum

### 数字枚举

基本用法，枚举值会递增

```
enum Status {
  Uploading,
  Success,
  Failed,
}
console.log(Status.Uploading) // 0
console.log(Status.Success) // 1
console.log(Status.Failed) // 2
```

自定义枚举值

```
const test = 1
const getIndex = () => {
  return 2
}
enum Status {
  Uploading = 11,
  // 这里使用了变量作为枚举值
  Success = getIndex(),
  Failed , // 前一个使用了常量或者计算就要给一个初始值，它不会自增，并且报错。
}
console.log(Status.Uploading) // 11
console.log(Status.Success) // 2
console.log(Status.Failed) // 报错
```

### 反向映射

字段名会映射值，值也会映射字段名

```
const test = 1
const getIndex = () => {
  return 2
}
enum Status {
  Uploading = 11,
  Success = getIndex(),
  Failed = 5,
}
console.log(Status) // {2: "Success", 5: "Failed", 11: "Uploading", Uploading: 11, Success: 2, Failed: 5}
```

### 字符串枚举

字符串枚举要求每个字段的值都是字符串字面量，或者该枚举值中另一个字符串枚举成员。

```
enum Message {
  Error = 'error',
  Success = 'success',
  Failed = Error,
}
console.log(Message.Error) // error 所以枚举可以使用字符串常量或者枚举成员
```

### 异构枚举

```
// 即包含数字的值又包含字符串的值，不建议使用。
enum Result {
  Faild = 0,
  Success = 'success'
}
```

### 枚举成员类型和联合枚举类型

```
1.不带初始值的枚举成员 enum E { A }
2.值为字符串字面量 enum E { A = 'a' }
3.值为数值的字面量或者负数字面量或正数 enum E { A = -1 }
满足以上一个条件的枚举值或者成员就可以作为类型来使用。
```

#### 1.使用枚举成员

```
enum Animals {
  Dog = 1,
  Cat = 2,
}
interface Dog {
  type: Animals.Dog
}
const dog: Dog = {
  type: Animals.Dog,
  // type: Animals.Cat, // 报错，因为接口没有定义Cat类型
}
```

#### 2.联合枚举类型

string | number 联合类型就是用竖线隔开，标识既可以使用string也可以使用number

```
enum Status {
  Off,
  On,
}
interface Light {
  status: Status
}
const light: Light = {
  status: Status.Off
  // status: Animals.Dog // 报错
}
```

### 运行时的枚举

一个枚举值编译后就是一个真实的对象，运行时使用枚举值，相当于一个丰富的对象。接口这些在编译后是没有实际的东西，就不能在运行时使用。

### const enum

使用const

```
const enum Animals {
  Dog = 1,
}
const dog = Animals.Dog
编译后 var dog = 1 /* Dog */;
```

不使用const

```
enum Animals {
  Dog = 1,
}
const dog = Animals.Dog
// 编译后返回的是一个对象
var Animals;
(function (Animals) {
    Animals[Animals["Dog"] = 1] = "Dog";
})(Animals || (Animals = {}));
var dog = Animals.Dog;

```

## 11.类型推论和兼容性

- 类型推论
  - 基础
  - 多类型联合
  - 上下文类型
- 类型兼容性
  - 基础
  - 函数兼容性
    - 参数个数
    - 参数类型
    - 返回值类型
    - 可选参数和剩余参数
    - 参数双向协变
    - 函数重载
- 枚举
- 类
- 泛型

### 1.类型推论

#### 1.1基础

```
let name1 = 'zz'
// name1 = 123 // 声明了字符串类型，再赋值数字类型就会报错。
```

#### 1.2多类型联合

```
let arr5 = [1, 'a']
arr5 = [2, 'b', 3] // 可以
// arr5 = [2, 'b', false] // 报错，类型会推断出你要的类型，默认完成泛型定义。
```

#### 1.3上下文类型

```
  // TypeScript类型推论也可能按照相反的方向进行。 这被叫做“按上下文归类”。
  window.onmousedown = (mouseEvent) => {
    console.log(mouseEvent.button);  // <- Error
};
// 这个例子会得到一个类型错误，TypeScript类型检查器使用Window.onmousedown函数的类型来推断右边函数表达式的类型。 因此，就能推断出 mouseEvent参数的类型了。 如果函数表达式不是在上下文类型的位置， mouseEvent参数的类型需要指定为any，这样也不会报错了。
```

### 2.类型兼容性

#### 2.1基础

```
interface InfoInterface {
  name: string,
  info: { age: number}
}
let infos: InfoInterface
const infos1 = {name: 'zz', info: { age: 18}}
const infos2 = { age: 18 }
const infos3 = { name: 'zz', age: 18, info: { age: 18} }
infos = infos1 // 兼容性检测时深层次的递归检测，就算有嵌套也会检测到，如果info里的age改为其他类型的值就会报错。
// infos = infos2 // 报错，因为被赋值的值里必须要有name属性
// infos = infos3 // 多的属性没有关系
```

### 3函数兼容性

#### 3.1参数个数

```
let x = (a: number) => 0
let y = (b: number, c: string) => 0
y = x
x = y // 报错，赋值的值必须小于被复制的值，就是y赋给x的个数要小于x的个数。
```

#### 3.2参数类型

```
const x = (a: number) => 0
const y = (b: string) => 0
x = y // 参数的类型不对应
```
#### 3.3返回值类型

```
let x = (): string | number => 0
let y = (): string => 'a'
let z = (): boolean => false
x = y // 可以
y = x // y只能返回string类型，所以报错了
y = z // 报错，返回值类型不一样
```

#### 3.4可选参数和剩余参数

```
const getSun = (arr: number[], callback: (...args: number[]) => number): number => {
  return callback(...arr)
}
const res = getSun([1, 2, 3], (...args: number[]): number => args.reduce((a, b) => a + b, 0))
const res2 = getSun([1, 2, 3], (arg1: number, arg2: number, arg3: number): number => arg1 + arg2 + arg3)
// 当要被赋值的参数中有可选参数赋值的函数可以用任意个数参数代替，但是类型需要对应。就是res、res2中里都是number类型，不能是其他类型。剩余参数可以看作无数个可选参数。
```

#### 3.5可选参数和剩余参数

```
const getSun = (arr: number[], callback: (...args: number[]) => number): number => {
  return callback(...arr)
}
const res = getSun([1, 2, 3], (...args: number[]): number => args.reduce((a, b) => a + b, 0))
const res2 = getSun([1, 2, 3], (arg1: number, arg2: number, arg3: number): number => arg1 + arg2 + arg3)
// 当要被赋值的参数中有可选参数赋值的函数可以用任意个数参数代替，但是类型需要对应。就是res、res2中里都是number类型，不能是其他类型。剩余参数可以看作无数个可选参数。
```

#### 3.6参数双向协变

```
let funcA = (arg: number | string): void => { }
let funcB = (arg: number): void => { }
funcA = funcB
funcB = funcA
// funcB可以赋值给funcA，反之也可以，因为他们都有一个arg属性为number类型。这就是双向协变

```

#### 3.7函数重载

```
function merge(arg1: number, arg2: number): number
function merge(arg1: string, arg2: string): string
function merge(arg1: any, arg2: any): any {
  return arg1 + arg2
}
function sum(arg1: number, arg2: number): number
function sum(arg1: any, arg2: any): any {
  return arg1 + arg2
}
let func = merge // 现在func的类型就是接口，里面有2个函数重载
// func = sum // 报错，缺一个函数重载。通过类型推论里面有2个函数重载的函数要包括number类型和string类型。使用sum函数重载赋值，它只有一种函数重载，就会赋值错误。他们就会不兼容
```

### 4.枚举

```
enum StatusInterface {
  On,
  Off,
}
enum AnimalEnum {
  Dog,
  Cat,
}
let s = StatusInterface // 编译器就会推断出s就是StatusInterface的On的枚举成员，和数值类型兼容。
s = AnimalEnum.Dog // 不兼容，数字枚举类型只能与数字类型兼容，在不同值之间是不兼容的
```

### 5.类

```
类与对象字面量和接口差不多，但有一点不同：类有静态部分和实例部分的类型。 比较两个类类型的对象时，只有实例的成员会被比较。 静态成员和构造函数不在比较的范围内。
class AnimalClass {
  feet: number;
  constructor(name: string, numFeet: number) { }
}

class Size {
  feet: number;
  constructor(numFeet: number) { }
}

let z: AnimalClass;
let s: Size;

z = s;  // OK
s = z;  // OK
  类的私有成员和受保护成员
private
protected
类的私有成员和受保护成员会影响兼容性。 当检查类实例的兼容时，如果目标类型包含一个私有成员，那么源类型必须包含来自同一个类的这个私有成员。 同样地，这条规则也适用于包含受保护成员实例的类型检查。 这允许子类赋值给父类，但是不能赋值给其它有同样类型的类。
class ParentClass {
  private age: number
  constructor() {}
}
class ChildrenClass extends ParentClass {
  constructor() {
    super()
  }
}
class OtherClass {
  private age: number
  constructor() {}
}
// 子类可以赋值给父类的类型的值
const children: ParentClass = new ChildrenClass()
// const other: ParentClass = new OtherCLass() // 报错，类型具有私有属性
```

### 6.泛型

```
interface Data<T> {}
let data1: Data<number>
let data2: Data<string>
// 这里它赋值的是一个空对象，所以没问题
data1 = data2
```

```
interface Data<T> {
  data: T
}
let data1: Data<number>
let data2: Data<string>
  // 给接口里定义了内容，再赋值就会报错。
data1 = data2 // 报错
```

### 12.高级类型(1)

- 交叉类型
- 联合类型
- 类型保护
- null和undefined
- 类型保护和类型断言
- 类型别名
- 字面量类型
- 枚举成员类型
- 可辨识联合

#### 1.交叉类型

交叉类型就是取多个类型的并集用&符号定义

```
const mergeFunc = <T, U>(arg1: T, arg2: U): T & U=> {
  // 使用类型断言
  let res = {} as T & U
  res = Object.assign(arg1, arg2)
  return res
}
```

`Object.assign()`方法合并2个对象，res即包含了T类型又包含了U类型，就可以使用交叉类型`T & U`来处理返回值。可以理解为与

#### 2.联合类型

`type1 | type2 | tpye3 `,只用是其中一种类型就可以。
设置传入的类型

```
const getLengthFunc = (content: string|number): number => {
  if (typeof content === 'string') { return content.length}
  else { return content.toString().length}
}
// 只能传入字符串或者数字类型
console.log(getLengthFunc(123))
console.log(getLengthFunc('abc'))
// console.log(getLengthFunc(false)) // 报错
```

#### 3.类型保护

有些数据是要再代码运行起来才知道结果

```
const valueList = [123, 'abc']
const getRandomValue = () => {
  const number1 = Math.random() * 10
  if (number1 < 5) { return valueList[0] }
  else { return valueList[1] }
}
const item = getRandomValue()
console.log(item) // 随机abc和123，无法确认是什么类型
```

```
// 这里会报错，ts会推断出是number类型或者string类型。
if (item.length) {
  console.log(item.length)
} else {
  console.log(item.toFixed)
}
```

```
// 简单使用可以使用类型断言，但是如果又很多数据就要使用类型保护。
if ((item as string).length) {
  console.log((item as string).length)
} else {
  console.log((item as number).toFixed)
}
function isString(value: number | string): value is string {
  return typeof value === 'string'
}
```

```
// 这样只用第一行判断了之后就不用判断，使用类型保护，ts就会推断出使用的是什么类型。
if (isString(item)) {
  console.log(item.length)
} else {
  console.log(item.toFixed)
}
```

```
// 使用typeof类型保护。如果只是简单的一句，就可以只用typeof来判断.typeof只能使用===或者!==来判断。typeof判断的类型只能是string/number/boolean/symbol中的一种。
if (typeof item === 'string') {
  console.log(item.length)
} else {
  console.log(item.toFixed)
}
```

```
// 使用instanceof类型保护。判断一个实例是不是某个构造函数创建的实例，或者某个类创建的实例。
class CreateByClass1 {
  public age = 18
  constructor() {}
}
class CreateByClass2 {
  public name = 'lison'
  constructor() {}
}
function getRandomItem() {
  return Math.random() < 0.5 ? new CreateByClass1() : new CreateByClass2()
}
const item1 = getRandomItem()
// 使用类型保护
if (item1 instanceof CreateByClass1) {
  console.log(item1.age)
} else {
  console.log(item1.name)
}
```

#### 4.null/numdefined

```
// 赋值操作时，声明的变量会默认先赋值undefined。在ts中时会默认使用联合类型。
let values: string | undefined = '123'
// 定义传入的参数y为可选参数，这里ts就会推断出y是一个number和undefined的联合类型。
const sunFunc = (x: number, y?: number) => {
  return x + (y || 0)
}
```

#### 5.类型保护和类型断言

当参数和属性为联合类型或者any类型这种不唯一的类型时，就需要使用类型保护来做一些判断。

```
const getLengthFunctoin = (value: string | null): number => {
  // if (value === null) {return 0}
  // else { return value.length}
  // 简写
  return (value || '').length
}
```

在tsconfig.json中开启strictNullChecks严格模式后，有些时候编译器无法在我们声明变量前知道这个值是否为null，所以需要手动使用类型断言。

```
function getSplicedStr(num: number | null): string {
  function getRes(prefix: string) {
    return prefix + num!.toFixed().toString()
  }
  num = num || 0.1
  return getRes('zzz-')
}
```

函数执行，如果num没有值就为0.1，有值就为那个值。但是这里`return prefix + num.toFixed().toString()`里的num就会报错，显示num可能为空值。这里就需要使用类型断言，在num后面添加感叹号!。表示不为空值。

```
console.log(getSplicedStr(1.2)) // zzz-1
```

#### 6.类型别名

##### 6.1定义类型别名

```
type TypeString = string
let str: TypeString
// 类型别名也可以使用泛型
type PostionType<T> = { x: T, y: T }
const postion1: PostionType<number> = {
  x: 1,
  y: -1
}
const postion2: PostionType<string> = {
  x: 'left',
  y: 'top'
}
```

```
// 使用类型别名的时候还可以在属性中使用自己
type Childs<T> = {
  current: T,
  child?: Childs<T>
}
```

```
// 类型别名嵌套
let ccc: Childs<string> = {
  current: 'first',
  child: {
    current: 'second',
    child: {
      current: 'third',
      // child: 'test' // 报错
    }
  }
}
```

为接口起别名时不能使用extends implements

```
// 接口有时可以起到类型别名的作用
type Alias = {
  num: number
}
interface Interface {
  num: number
}
let _alias: Alias = {
  num: 123
}
let _interface: Interface = {
  num: 321
}
_alias = _interface
```

在使用implements时使用接口，当类型别名为implements起别名的时候是不行的;当无法通过接口使用联合类型并且需要使用元组类型的时候就使用类型别名。

#### 7.字面量类型

##### 7.1字符串字面量

```
type Name = 'zzz'
// const name3: Name = 'haha' // 报错，这里的zzz是当作类型来使用的。
```

```
// 联合类型使用字符串
type Direction = 'north' | 'east' | 'south' | 'west'
function getDirectionFirstLetter(direction: Direction) {
  return direction.substr(0, 1)
}
console.log(getDirectionFirstLetter('north')) // 取到第一个字母n
// console.log(getDirectionFirstLetter('zzz')) // 报错
```

##### 7.2数值字面量

````
type Age = 18
interface InfoInterface {
  name: string
  age: Age
}
const _info: InfoInterface = {
  name: 'zzz',
  age: 18
  // age: 20 // 这里类型就是18，如果写20就会报错。
}
````

#### 8.枚举成员类型

能作为类型的枚举要符合三个条件

一、不带初始值的枚举成员。

二、值为成员里面的值为字符串字面量。

三、值为数值字面量或者带有负号的数值字面量。

这三个条件满足一条，它的枚举成员或者枚举值都可以作为类型来使用。

##### 8.1可辨识联合

单例类型多值枚举成员类型和字面量类型。可以把单例类型、联合类型、类型保护和类型别名这几种类型进行合并，创建可辨识联合的高级类型。也可以称作标签联合或代数数据类型。

1.具有普通的单例类型属性— 可辨识的特征。

2.一个类型别名包含了那些类型的联合— 联合。

3.此属性上的类型保护。

```
interface Square {
  kind: "square";
  size: number;
}
interface Rectangle {
  kind: "rectangle";
  width: number;
  height: number;
}
interface Circle {
  kind: "circle";
  radius: number;
}
```

首先我们声明了将要联合的接口。 每个接口都有 kind属性但有不同的字符串字面量类型。 kind属性称做 可辨识的特征或 标签。 其它的属性则特定于各个接口。 注意，目前各个接口间是没有联系的。 下面我们把它们联合到一起：

`type Shape = Square | Rectangle | Circle;`

现在我们使用可辨识联合:

```
function area(s: Shape) {
  switch (s.kind) {
      case "square": return s.size * s.size;
      case "rectangle": return s.height * s.width;
      case "circle": return Math.PI * s.radius ** 2;
  }
}
```

##### 8.2完整性检查

```
function area(s: Shape): number { // error: returns number | undefined
  switch (s.kind) {
    case "square": return s.size * s.size;
    case "rectangle": return s.height * s.width;
  }
}
```

现在去掉了"circle"，这里的number就会报错。当遗漏了一种，如果传进来的是circle，这2个分支都不会进入。最后就会返回一个undefined。而指定的类型是number，如果没开strictNullChecks则会兼容，ts默认为联合类型number | undefined。如果开了就不能赋值。这种方法在旧的代码中支持不好，因为没有strictNullChecks。

###### 8.2.1第二种使用never的方式

```
function assertNever(x: never): never {
  throw new Error("Unexpected object: " + x);
}
function area(s: Shape): number {
  switch (s.kind) {
    case "square": return s.size * s.size;
    case "rectangle": return s.height * s.width;
    case "circle": return Math.PI * s.radius ** 2
    default: return assertNever(s)
  }
}
```

这里， assertNever检查 s是否为 never类型—即为除去所有可能情况后剩下的类型。 如果你忘记了某个case，那么 s将具有一个真实的类型并且你会得到一个错误。 这种方式需要你定义一个额外的函数，但是在你忘记某个case的时候也更加明显。

### 13.高级类型(2)

- this类型
- 索引类型
  - 索引类型查询操作符
  - 索引访问操作符
- 映射类型
  - 基础
  - 由映射类型进行推断
  - 增加或移除特定修饰符
  - keyof和映射类型2.9的升级
  - 元组和数组上的映射类型
- 条件类型
  - 基础
  - 分布式条件类型
  - 条件类型的类型推断-infer
  - TS预定义条件类型

#### 1. this类型

```
class Counters {
  constructor(public count: number = 0) {}
  public add(value: number) {
    this.count += value
    return this
  }
  public subtract(value: number) {
    this.count -= value
    return this
  }
}
let counter1 = new Counters(10)
// console.log(counter1.add(3).subtract(2))
// 使用return返回this实例，就可以使用链式调用。

class PowCounter extends Counters {
  constructor(public count: number = 0) {
    super(count)
  }
  public pow(value: number) {
    this.count = this.count ** value
    return this
  }
}
let powCounter = new PowCounter(2)
console.log(powCounter.pow(3).add(1).subtract(3))
// 在ts会对方法返回来的this进行判断，判断它是继承来的this返回的实例。继承来的this就不会报错，可以使用链式调用。
```

#### 2. 索引类型

##### 2.1 索引类型查询操作符

keyof

```
interface InfoInterfaceAdvanced {
  name: string;
  age: number;
}
let infoProp: keyof InfoInterfaceAdvanced // infoProp就成为联合类型
infoProp = 'name'
infoProp = 'age'
// infoProp = 'sex' // 报错
// 使用keyof的接口就是name和age组成的联合类型。

```

```
// 和泛型结合使用，ts可以检查使用了动态属性的代码
function getValue<T, K extends keyof T>(obj: T, names: K[]): T[K][] {
  return names.map((item) => obj[item])
}
const infoObj = {
  name: 'zzz',
  age: 18
}
let infoValues: (string | number)[] = getValue(infoObj, ['name', 'age'])
console.log(infoValues)
// keyof T就是获取T类型里面的所有属性名组成的联合类型，K extends keyof T，K的类型就只能是T里面的所有属性名组成的字符串联合类型。
```

##### 2.2 索引访问操作符

```
interface InfoInterfaceAdvanced {
  name: string;
  age: number;
}
let infoProp: keyof InfoInterfaceAdvanced // infoProp就成为联合类型
infoProp = 'name'
infoProp = 'age'

type NameTypes = InfoInterfaceAdvanced['name']

function getProperty<T, K extends keyof T>(o: T, name: K): T[K] {
  return o[name]
}
```

结合接口

```
interface Objs<T> {
  [key: string]: T
}
const objs1: Objs<number> = {
  age: 18
}
let keys: keyof Objs<number>
```

```
interface Type {
  a: never;
  b: never;
  c: string;
  d: number;
  e: undefined;
  f: null;
  g: object;
}
type Test = Type[keyof Type] // 返回类型不为never属性名
```

#### 3. 映射类型

##### 3.1 基础

```
interface Info1 {
  age: number
}
interface ReadonlyType {
  readonly age: number
}
```

现在想让age变为值读类型，就需要给age前面加readoly，但是如果需要给很多个属性名添加时就要用到映射。

```
interface Info1 {
  age: number
  name: string
  sex: string
}
// 简单映射类型
type ReadonlyType<T> = {
  readonly [P in keyof T]: T[P]
}
// 可选属性方式
type ReadonlyType<T> = {
  readonly [P in keyof T]?: T[P]
}
```

ts在内部实现使用for...in，keyof T就是传进来所有的属性名，P in就是对所有属性名的一个遍历，P就是每次遍历的属性名赋值到P变量上。T[P]就是结果类型的这个值。

```
type ReadOnlyInfo1 = ReadonlyType<Info1>
let info11: ReadOnlyInfo1 = {
  age: 18,
  name: 'zzz',
  sex: 'man'
}
info11.age = 20 // 报错，age为只读属性
```

ts已经内置了Readonly和Partial，可以直接拿来用

```
type ReadonlyPerson = Readonly<Info1> // 这里就全是只读参数

type PartialPerson = Partial<Info1> // 这里就全是可选参数
```

内置Pick和Record

```
// Pick的用法
interface Info2 {
  name: string;
  age: number;
  address: string;
}
const info5: Info2 = {
  name: 'zzz',
  age: 18,
  address: 'shanghai'
}
function pick<T, K extends keyof T>(obj: T, keys: K[]): Pick<T, K> {
  const res: any = {}
    keys.map(key => {
      res[key] = obj[key]
    })
    return res
}
const nameAndAddress = pick(info5, ['name', 'address'])
console.log(nameAndAddress) // {name: "zzz", address: "shanghai"}
```

Record将一个对象的每个属性转为其他值

```
// Record用法
function mapObject<K extends string | number, T, U>(obj: Record<K, T>, f: (x: T) => U): Record<K, U> {
  const res: any = {}
  for (const key in obj) {
    res[key] = f(obj[key])
  }
  return res
}
const names = {0: 'hello', 1: 'world', 2: 'bye'}
const lengths = mapObject(names, s => s.length)
console.log(lengths) // {0: 5, 1: 5, 2: 3}
```

Readonly， Partial和 Pick是同态的，但 Record不是。 因为 Record并不需要输入类型来拷贝属性，所以它不属于同态。

##### 3.2 由映射类型进行推断

```
type Proxy<T> = {
  get(): T;
  set(value: T): void;
}
type Proxify<T> = {
  [P in keyof T]: Proxy<T[P]>
}
function Proxify<T>(obj: T): Proxify<T> {
  const result = {} as Proxify<T>
  for (const key in obj) {
    result[key] = {
      get: () => obj[key],
      set: (value) => obj[key] = value,
    }
  }
  return result
}
let props = {
  name: 'zzz',
  age: 18
}
let proxyProps = Proxify(props)
proxyProps.name.set('zz')
console.log(proxyProps.name.get()) // zz
```

拆包，由映射类型推断出原始类型。

```
function unproxify<T>(t: Proxify<T>): T {
  const result = {} as T
  for (const k in t) {
    result[k] = t[k].get()
  }
  return result
}
let originalProps = unproxify(proxyProps)
console.log(originalProps)
```

##### 3.3 增加或移除特定修饰符

```
interface Info1 {
  age: number
  name: string
  sex: string
}
type ReadonlyType<T> = {
  readonly [P in keyof T]?: T[P]
}
type ReadOnlyInfo1 = ReadonlyType<Info1>
let info11: ReadOnlyInfo1 = {
  age: 18,
  name: 'zzz',
  sex: 'man'
}
type RemoveReadonlyInfo2<T> = {
  // 通过减号把传进来的readonly去掉，同理可以减去可选属性。
  -readonly [P in keyof T]-?: T[P]
}
type Info1WidthoutReadonly = RemoveReadonlyInfo2<ReadOnlyInfo1>
```

##### 3.4 keyof和映射类型2.9的升级

ts在2.9版本中支持用number和symbol命名的属性

```
// keyof
const stringIndex = 'a'
const numberIndex = 1
const symbolIndex = Symbol()
type Objs2 = {
  [stringIndex]: string,
  [numberIndex]: number,
  [symbolIndex]: symbol
}
type keysType = keyof Objs2

// 映射类型
// 映射类型对于number类型和symbol类型的支持
type ReadonlyTypes<T> = {
  readonly [P in keyof T]: T[P]
}
let objs3: ReadonlyTypes<Objs2> = {
  a: 'aa',
  1: 11,
  [symbolIndex]: Symbol(),
}
// objs3.a = 'bb' // 报错，因为它是只读属性
```

##### 3.5 元组和数组上的映射类型

元组和数组上的映射类型，ts在3.1的版本中在元组和数组的映射类型会生成新的元组和数组，并不会创建一个新的类型。这个类型会存在push、pop等方法和数组属性。

```
type MapToPromise<T> = {
  [K in keyof T]: Promise<T[K]>
}
type Tuple = [number, string, boolean]
type promiseTuple = MapToPromise<Tuple>
let Tuple: promiseTuple = [
  new Promise((resolve, reject) => resolve(1)),
  new Promise((resolve, reject) => resolve('a')),
  new Promise((resolve, reject) => resolve(false))
]

// unknown
// 1.任何类型都可以赋值给unknown类型
let value1: unknown
value1 = 'a'
value1 = 123

// 2.如果没有类型断言或基于控制流的类型细化时，unknown不可以赋值给其他类型，此时他只能赋值给unknown和any。
let value2: unknown
// let value3: string = value2 // 报错
value1 = value2

// 3.如果没有类型断言或基于控制流的类型细化，不能再他上面进行任何操作。
let value4: unknown
// value4 += 1 // 报错

// 4.unknown与任何其他类型组成的交叉类型，最后都等于其他类型。
type type1 = string & unknown // string类型
type type2 = number & unknown // number类型
type type3 = unknown & unknown // unknownn类型
type type4 = string[] & unknown // string[]类型

// 5.unknown与任何其他类型(处理any)组成的联合类型，都等于unknown类型。
type type5 = unknown | string // unknown类型
type type6 = any | unknown  // any类型
type type7 = number[] | unknown // unknown类型

// 6.never类型时unknown的子类型
// 使用条件类型
type type8 = never extends unknown ? true : false

// 7.keyof unknown等于类型never
type type9 = keyof unknown

// 8.只能对unknown进行等或不等操作，不能进行其他操作
// value1 === value2
// value1 !== value2
// value1 += value2 // 报错

// 9.unknown类型的值不能访问他的属性、作为函数调用和作为类创建实例
let value10: unknown
// value10.age // 报错
// value10() // 报错
// new value10() // 报错

// 10.使用映射类型时如果遍历的时unknown类型，则不会映射任何属性
type Types1<T> = {
  [P in keyof T]: number
}
type type11 = Types1<any> // [x: string]: number
type type12 = Types1<unknown> // 空
```

#### 4. 条件类型

条件类型语法上像三元操作符，它是以条件表达式类型关系的检测，然后再两种类型中选择一个。

##### 4.1 基础

```
// T extends U ? X : Y
type Types2<T> = T extends string ? string : number
let index: Types2<'a'> // 判断是否是字符串的子类型
// let index: Types2<123> // 数值的子类型
```

##### 4.2 分布式条件类型

判断T是否为any的子类型，否则返回never类型

```
type TypeName<T> = T extends any ? T : never
// 组成联合类型
type Type3 = TypeName<string | number>
```
官方例子

```
type TypeName<T> =
  T extends string ? string :
  T extends number ? number :
  T extends boolean ? boolean :
  T extends undefined ? undefined :
  T extends () => void ? void :
  object
type Type4 = TypeName<() => void> // 返回函数类型
type Type5 = TypeName<string[]> // 传入字符串类型的数组，都不符合返回Object类型
type Type6 = TypeName<(() => void) | string[]> // 返回object类型或者void类型
// 分布式条件实际利用
type Diff<T, U> = T extends U ? never : T
type Test = Diff<string | number | boolean, undefined | number>
```
条件类型和映射类型结合

```
type Type7<T> = {
  [K in keyof T]: T[K] extends Function ? K : never
}[keyof T]
interface Part {
  id: number;
  name: string;
  subparts: Part[];
  undatePart(newName: string): void
}
type Test1 = Type7<Part>
```

接口Part里有4个字段，其中undatePart的字段是一个函数，Type7映射类型涉及到映射类型(K in typeof T)、条件访问类型(T[K] extends Function ? K : never)和索引类型(keyof T)。K in keyof T用于遍历所有属性名，它的值使用了条件类型T[K]代表了当前属性名的属性值。判断是否为函数类型，是就返回属性名字符串，不是就返回never。通过索引访问keyof T不为never的属性名。最后返回不为never的属性名。在Part接口中只有undatePart为函数类型，其他不为函数的类型都返回never，在Test1中，只显示了函数字面量为undatePart。

##### 4.3 条件类型的类型推断-infer

不使用infer的写法

```
type Type8<T> = T extends any[] ? T[number] : T
// 判断三级变量T是否为一个数组，如果是获取它的索引访问类型，通过传一个number类型的索引，获取到它的值的类型。如果不是一个函数直接返回T
type Test2 = Type8<string[]>
// Test2使用Type8字符串数组元素
type Test4 = Type8<string>
// 直接返回字符串
// 如果传进去的是数组就是使用数组元素的类型，没有传入数组就直接使用类型。

```

使用infer

```
// tslint:disable-next-line:array-type
type Type9<T> = T extends Array<infer U> ? U : T
type Test5 = Type9<string[]>
// 如果传进来的是字符串数组，就走前面返回字符串元素，如果是字符串类型就返回字符串类型。
```

##### 4.4 TS预定义条件类型

Exclude<T, U>

```
// Exclude表示第一个参数类型里和第二个参数类型里不相同的参数类型。
type Type10 = Exclude<'a' | 'b' | 'c', 'a' | 'b'> // c


```

Extract<T, U>

```
// 选取出T中可以赋值给U的类型
type Type11 = Extract<'a' | 'b' | 'c', 'c'> // c
```

NonNullable<T>

```
// 可以从T中去掉null和undefined
type Type12 = NonNullable<string | number | null | undefined> // string | number
```

ReturnType<T>

```
// 获取函数类型返回值类型
type Type13 = ReturnType<() => string> // string
type Type14 = ReturnType<() => void> // void
```

InstanceType<T>

```
// 获取构造函数类型的实例类型
class AClass {
  constructor() {}
}
type T1 = InstanceType<typeof AClass> // 使用typeof获取类型 T1的类型就是AClass
type T2 = InstanceType<any> // T2为any类型
type T3 = InstanceType<never> // T3为never类型
// type T4 = InstanceType<string> // 报错，不满足构造函数类型。
```

### 14. ES6和Node.js的模块

- ES6的模块
  - export
  - import
  - export default
  - import和export的复合写法
- Node.js的模块
  - exports
  - module.exports

#### 1.ES6的模块

##### 1.1export

```
const name = 'zzz'
const age = '18'
const address = 'hanghzhou'
// 导出
export { name, age, address}
// 导出函数
export function func () {}
// 导出类
export class A {}
```

```
// a.js
// 导出重命名
function func1 () {}
class B {}
const b = ''
export {
  func1 as Function1,
  B as ClassB,
  b as StringB
}
```

```
错误的写法，直接导出一个具体的值
export 'zzz' // 错误

错误的写法，直接导出常量
const name = 'zzz'
export name
```

export导出的值与其对应的值是动态绑定的

```
// b.js
export let time = new Date()
setInterval(() => {
  time = new Date()
}, 1000)
// export是会自动提升到全局上下文的最顶部
```

```
// index.js
import { time } from './b'
setInterval(() => {
  console.log(time)
}, 1000)
每1秒打印一次时间
```

##### 1.2imoprt

```
export const name = 'zzz'
export const age = 18
export const info = {
  name: 'zzz',
  age: 18
}
```

```
// index.js
import { name as nameProp, age, info} from './c'
console.log(age)
// 还可以起别名
console.log(nameProp)
// 引入对象
console.log(info)
```

```
// d.js
document.title = 'zzz'
```

```
// index.js
// 直接引入执行，不用再起别名。
import './d'
```

import还有提升的效果,import是静态执行的在编译阶段就会执行，并且需要添加路径。

```
// e.js
export function getName () {
  return 'zzzz'
}
```

```
// index.js
// import提升，使用import引入，再js编译时会自动将它提升到顶部。
getName()
import { getName } from './e'
错误写法
import { 'get' + 'name'} from './e'
```

```
// 引入合并
import { name } from './c'
import { age } from './'
js会自动编译成这样
import { name, age } from './c
```

```
// 使用*引入文件内所内容
import * as info from './a'
console.log(info) // 打印出a文件下的所有内容
```

##### 1.3export default

```
// d.js
// 一个模块只能使用一次export default
export default function func() {}
```

```
// index.js
// 调用，默认导出引入后名字可以不一样
import funcitonName from './d'
funcitonName()
```

default还可以使用别名

```
// 导出
export { func as default} from './d'
// 引入
import { default as func } from './d'
```

如果想把一个模块引入再导出可以这样

```
import func from './d'
export default func
// 简写
export { default as func } from './d'
// 引入
import {func} from './d'
```

export default 可以直接接一个值，但是export不能。

```
// d.js
export default 'zz'
```

```
// index.js
import str from './d'
console.log(str) // zz
```

既存在export又存在export default的情况

```
// d.js
let sex = 'man'
export let name = 'zz'
export default sex
```

```
import sex, {name} from './d'
console.log(sex, name) // man zz
```

##### 1.4import和export的符合写法

```
import { name, age } from './a'
export {name, age}
简写
export { name, age } from './a'
console.log(name, age) // 报错
```

```
// 别名导出
export { name as nameProp } from './a'
// 整体导出
export * from './a'
// 默认导出
export { name as defualt } from './a'
import { name } from './a'
// 简写
export default name
```

import是静态编译的，如果需要按需加载就没有办法。所以加入了import()方法，但是没有正式使用。webpack已经实现了。

```
const status = 1
if (status) {
  import ('./a')
} else {
  import ('./b')
}
```

#### 2.Node.js的模块

##### 2.1exports

nodejs模块，nodejs是遵循commonjs

```
// 导出export对象
exports.name = 'zzz'
exports.age = 18
// 引入
const name = require('./a')
```

##### 2.2module.exports

```
// a.js
module.exports = funcitn () {
  console.log('zzz')
}
```

```
// b.js
// 引入,直接调用。
const name = require('./a')
name()
```

### 15.模块和命名空间

- 模块
  - export
  - import
  - export default
  - export = 和 import xx = require()
- 命名空间
  - 定义和使用
  - 拆分为多个文件
- 别名
- 模块解析
  - 相对和非相对模块导入
  - 模块解析策略
  - 模块解析配置项

#### 1.模块

##### 1.1export

```
// a.ts
export interface FuncInterface {
  name: string;
  (arg: number): string
}
export class ClassC {
  constructor() {}
}
class ClassD {
  constructor() {}
}
export { ClassD }
export { ClassD as ClassNameD }
export * from './b'
export { name } from './b'
// 导出别名别名
export { name as NameProp } from './b'
```

##### 1.2import

```
// b.ts
export const name = 'zzz'
```

```
// index.ts
import { name } from './b'
或者
import * as info from './b'
// 或者别名
import { name as NameProp} from './b'
console.log(info) // {name: "zzz", __esModule: true}
```

引入上面a.ts导出的内容

```
// index.ts
import * as AData from './a'
console.log(AData)
// 打印内容，这里不会显示接口的内容。
ClassC: ƒ ClassC()
ClassD: ƒ ClassD()
ClassNameD: ƒ ClassD()
NameProp: (...)
name: (...)
__esModule: true
get NameProp: ƒ ()
get name: ƒ ()
[[Prototype]]: Object
```

```
// index.ts
// 直接引入，不能调用里面的逻辑。
import './a'
```

##### 1.3export default

```
// c.ts
export default 'zzz'
```

```
// index.ts
import name from './c'
console.log(name) // zzz
```

##### 1.4export = 和 import xx = require()

ts为了解决commonjs、amd和es6的导入不兼容添加了export = 和import x = require()

```
// c.ts
const name = 'zzz'
export = name
```

```
// index.ts
import name = require('./c')
console.log(name) // zzz
```

#### 2.命名空间

##### 2.1定义和使用

```
namespace Validation {
  const isLetterReg = /^[A-Za-z]+$/
  export const isNumberReg = /^[0-9]+$/
  export const checkLetter = (text: any) => {
    return isLetterReg.test(text)
  }
}
```

##### 2.2拆分为多个文件

如果需要再tsc里引入命名空间需要先再里面标记三斜线

```
/// <refernce path="space.ts">
// src/index.js为输出的路径，src/ts-modeuls/index.ts为需要编译的文件路径
tsc --outFile src/index.js src/ts-modules/index.ts
```

#### 3.别名

```
// 别名
namespace Shapes {
  export namespace Polygons {
    export class Triangle {}
    export class Squaire {}
  }
}

// 起别名
import polygons = Shapes.Polygons

// 使用别名，别名可以解决深层次嵌套的问题。
let triangle = new polygons.Triangle()
```

#### 4.模块解析

##### 4.1相对和非相对模块导入

```
相对到入 / 根目录 ./ 当前目录 ../上级目录

文件寻找机制，可以省略.ts后缀，它会先找目录下的.ts文件，如果没有再找.d.ts。

对于非相对模块编译器会依次从内向内遍历查找，会从引入的地方逐级的往外查找，会先查找.ts文件，如果没有就查找.d.ts文件。如果没有找到就向上级目录继续找这个的ts文件，如果没有找到，再找.d.ts文件。
```

##### 4.2模块解析策略

`模仿nodejs运行时的解析，在编译阶段定位模块的文件。区别在于nodejs查找.js文件，ts查找.ts、.tsx和.d.ts三种文件。nodejs会查找package.json文件，在里面查找main里定义的文件入口。而ts在package.json里查找tyeps，如果里面又定义就去找它指定的是什么文件。比如要找一个模块，它首先会去node_modules里找.ts文件，如果没有就找.tsx文件，如果没有就去找.d.ts文件，如果还没找到就去找package.json文件。在package.json里找types里定义的文件，如果没有就去找node_modules里模块文件夹下的index.ts文件，如果还没找到就去找index.tsx文件，如果没又找到就去找index.d.ts文件。还没找到就往上一级的node_modules文件里继续按流程查找。`

##### 4.3模块解析配置项

```
tsconfig.json文件下
baseUrl: "./"要求在运行时将模块放到一个文件里，这些模块可能放在各个文件里，但是构建工具会把他们集中放在一起。通过baseUrl来告诉编译器查找模块都放在哪里。相对模块不受baseUrl影响，相对模块时根据引入的相对路径查找的。
paths:"" 用来设置路径映射，
比如引入第三方文件 "path": {"*":["node_modules/@types", "src/styp"]}，*就是匹配所有，就会在这两个文件里引入。paths是相对于baseUrl的，使用了paths就必须配置baseUrl。如果baseUrl就是当前根目录就可以只写一个点"."
rootDirs 指定路径列表，在构建时编译器会将路径列表中的路径的列表放到一个文件夹中
"rootDirs": ["src/module", "src/core"] 会将这两个文件输出到同一个文件夹中。
noreResolve 只引入指定模块
中终端中使用tsc index.ts ./a.ts --noResolve，这样就只会引入./a.ts文件。
```

### 16.声明合并

- 补充知识
- 合并接口
- 合并命名空间
- 不同类型合并
  - 命名空间和函数
  - 命名空间和枚举

#### 1.补充知识

在 TypeScript 中，声明至少在以下三组之一中创建实体：命名空间、类型或值。命名空间创建声明创建一个命名空间，其中包含使用点符号访问的名称。类型创建声明就是这样做的：它们创建一个类型，该类型对声明的形状可见并绑定到给定的名称。最后，创建值的声明会创建在输出 JavaScript 中可见的值。

```
interface InfoInter {
  name: string
}
interface InfoInter {
  age: number
}
let infoInter: InfoInter
// 这里会提示少一个age字段，因为同名的接口会合并成一个接口。
// infoInter = {
//   name: 'zz'
// }
infoInter = {
  name: 'zz',
  age: 18
}
```

#### 2.合并接口

```
最简单也可能是最常见的声明合并类型是接口合并。在最基本的层面上，合并将两个声明的成员机械地连接到一个同名的接口中。
```

```
interface InfoInter {
  name: string
  getRes(inpupt: string): number
}
interface InfoInter {
  name: string
  getRes(input: number): string
}
let infoInter: InfoInter
infoInter = {
  name: 'zzz',
  getRes(text: any): any {
    if (typeof text === 'string') { return text.length}
    else { return String(text)}
  }
}
console.log(infoInter.getRes('123')) // 3 打印字符串长度
console.log(infoInter.getRes(123)) // 123 打印数值
console.log(infoInter.getRes(123).toFixed) // 报错
console.log(infoInter.getRes('123').length) // 报错
```

#### 3.合并命名空间

```
namespace Validations {
  // 如果这里的numberReg不加export那么就找不到numberReg，无法打印
  export const numberReg = /^[0-9]+$/
  export const checkNumber = () => {}
}
namespace Validations {
  // 不加export无法打印
  console.log(numberReg)
  export const checkLetter = () => {}
}
```

#### 4.不同类型合并

同名的类要定义再同名的命名空间前面

```
class Validations {
  constructor() {}
  public checkType() {}
}
```

定义同名的命名空间

```
namespace Validations {
  export const numberReg = /^[0-9]+$/
}
console.dir(Validations) // 它下面就会有numberReg属性。原型上会有checkType方法
```

##### 4.1命名空间和函数

函数的定义要放在命名空间前面

```
function countUp () {
  countUp.count++
}
// 定义同名的命名空间
namespace countUp {
  export let count = 0
}
console.log(countUp.count) // 0
countUp()
console.log(countUp.count) // 1
countUp()
console.log(countUp.count) // 2
```

##### 4.2命名空间和枚举

可以通过命名空间拓展枚举的内容

没有命名空间和枚举先后顺序

```
enum Colors {
  red,
  green,
  blue
}
namespace Colors {
  export const yellow = 3
}
// 只有添加了yellow为3，没有3为yellow的。
console.log(Colors) // {0: "red", 1: "green", 2: "blue", red: 0, green: 1, blue: 2, yellow: 3}
```

### 17.装饰器

- 基础
  - 装饰器定义
  - 装饰器工厂
  - 装饰器组合
  - 装饰器求值
  - 类装饰器
  - 方法装饰器
  - 访问装饰器
  - 属性装饰器
  - 参数装饰器

#### 1.基础

随着 TypeScript 和 ES6 中 Classes 的引入，现在存在一些需要额外功能来支持注释或修改类和类成员的场景。装饰器提供了一种为类声明和成员添加注释和元编程语法的方法。装饰器是JavaScript的[第 2 阶段提案](https://github.com/tc39/proposal-decorators)，可作为 TypeScript 的实验性功能使用。

`注意装饰器是一项实验性功能，可能会在未来版本中更改。`

要启用对装饰器的实验性支持，您必须`experimentalDecorators`在命令行或在您的`tsconfig.json`：

命令行：

```
tsc --target ES5 --experimentalDecorators
```

**tsconfig.json**：

```
{
  "compilerOptions": {
    "target": "ES5",
    "experimentalDecorators": true
  }
}
```

##### 1.1装饰器定义

装饰器是一种特殊种类的声明可被附连到一个类声明，方法，访问器，属性，或参数。装饰器使用形式@expression, where expressionmust 计算一个将在运行时调用的函数，其中包含有关装饰声明的信息。

```
function setProp (target) {
  // ...
}
```

##### 1.2装饰器工厂

它的返回值就是一个函数，返回值的函数就是装饰器的调用函数。

```
function setProp () {
  // 函数的调用
  return function (target) {
    // ...
  }
}
// 装饰器的调用
@setProp()
```

##### 1.3装饰器组合

@setProp()
@setName()
@setAge()
target
装饰器工厂的执行顺序是从上至下。
装饰器的调用时从下至上
定义装饰器工厂

```
function setName() {
  console.log('get setName')
  // 定义装饰器
  return (target) => {
    console.log('setName')
  }
}
function setAge () {
  console.log('get setAge')
  return (target) => {
    console.log('setAge')
  }
}
// 类装饰器
// 因为是装饰器工厂，所有要用()调用
@setName()
@setAge()
class ClassDec {

}
// 打印了4句话1、get setName 2、get setAge 3、setAge 4、setName，按照这样的顺序执行。
```

##### 1.4装饰器求值

对于如何应用应用于类内部各种声明的装饰器，有一个明确定义的顺序：

Parameter Decorators，然后是Method、Accessor或Property Decorators应用于每个实例成员。

Parameter Decorators，然后是Method、Accessor或Property Decorators应用于每个静态成员。

参数装饰器应用于构造函数。

类装饰器应用于类。

##### 1.5类装饰器

类装饰就是一个类声明之前声明。类装饰器应用于类的构造函数，可用于观察、修改或替换类定义。类装饰器不能在声明文件或任何其他环境上下文中使用（例如在declare类上）。

类装饰器的表达式将在运行时作为函数调用，装饰类的构造函数作为其唯一参数。

如果类装饰器返回一个值，它将用提供的构造函数替换类声明。

注意 如果您选择返回一个新的构造函数，您必须注意维护原始原型。在运行时应用装饰器的逻辑不会为您执行此操作。

```
let sign = null
// 定义装饰器工厂
function setName (name: string) {
  return (target: new() => any) => {
    sign = target
    console.log(target.name)
  }
}
// 定义类
@setName('zzz')
class ClassDes {
  constructor() {}
}
console.log(sign === ClassDes) // true
console.log(sign === ClassDes.prototype.constructor) // true
```

通过类的装饰器可以修改类的原型对象和构造函数

```
function addName(constructor: new() => any) {
  constructor.prototype = 'zzz'
}
@addName
class ClassD {}
// 原本d下面不存再name属性，通过接口来声明合并来添加到原型对象上。
interface ClassD {
  name: string
}
const d = new ClassD()
console.log(d.name) // zzz
```

通过类覆盖原型链的操作

```
function classDecorator<T extends new(...args: any[]) => {}>(target: T) {
  return class extends target {
    newProty = 'new property'
    public hello = 'override'
  }
}
@classDecorator
class Greeter {
  property = 'property'
  public hello: string
  constructor(m: string) {
    this.hello = m
  }
}
console.log(new Greeter('world')) // 这个新创建的类就会使用装饰器classDecorator里定义的类，覆盖Greeter里的属性和方法。
```

```
function classDecorator(target: any) {
  return class {
    newProty = 'new property'
    public hello = 'override'
  }
}
@classDecorator
class Greeter {
  property = 'property'
  public hello: string
  constructor(m: string) {
    this.hello = m
  }
}
console.log(new Greeter('world')) // 现在greeter里只有classDecorator里定义的内容，替换掉了Greeter里的内容。
```

##### 1.6方法装饰器

方法装饰只是一个方法声明之前声明。装饰器应用于方法的属性描述符，可用于观察、修改或替换方法定义。方法装饰器不能用于声明文件、重载或任何其他环境上下文（例如在declare类中）。

方法装饰器的表达式将在运行时作为函数调用，带有以下三个参数：

1、静态成员的类的构造函数，或实例成员的类的原型。

2、成员的姓名。

3、成员的属性描述符。



`js中的属性装饰符`

对象可以设置属性，如果属性值为函数，这么称它为方法。每一个属性和方法在定义时都存在三个描述符。

```
configurable // 可配置
writable // 可写
enumerable // 可枚举
```

```
interface ObjWithAnyKeys {
  [key: string]: any
}
let obj1: ObjWithAnyKeys = {}
Object.defineProperty(obj1, 'name', {
  value: 'zzz',
  // 设置不可写
  writable: false,
  configurable: true,
  enumerable: true
})
console.log(obj1.name) // zzz
obj1.name = 'test'
console.log(obj1.name) // zzz，不可写
```

ts中

```
function enumerable(bool: boolean): any {
  // 当方法装饰器装饰的是静态成员时，装饰的是类的构造函数。如果装饰的类的实例成员时，是类的原型对象。
  return (target: any, propertyName: string, descriptor: PropertyDescriptor) => {
    console.log(target, propertyName)
    descriptor.enumerable = bool
  }
}
function enumerable(bool: boolean): any {
  return (target: any, propertyName: string, descriptor: PropertyDescriptor) => {
    // 替换
    return {
      value () {
        return 'not age'
      },
      enumerable: bool
    }
    }
}
class ClassF {
  constructor(public age: number) {}
  // 传入false，使它不可枚举。
  // @enumerable(false)
  // 可枚举
  @enumerable(true)
  public getAge () {
    return this.age
  }
}
const classF = new ClassF(18)
// for (const key in classF) {
//   // 不可枚举打印
//   // console.log(key) // age
//   // 可枚举打印
//   console.log(key) // age 、getAge
// }
console.log(classF.getAge()) // not age
```

##### 1.7访问装饰器

访问器装饰器是在访问器声明之前声明的。访问器装饰器应用于访问器的属性描述符，可用于观察、修改或替换访问器的定义。不能在声明文件或任何其他环境上下文（例如在declare类中）中使用访问器装饰器。

访问器装饰器的表达式将在运行时作为函数调用，具有以下三个参数：

1、静态成员的类的构造函数，或实例成员的类的原型。

2、成员的姓名。

3、成员的属性描述符。

如果访问器装饰器返回一个值，它将用作成员的属性描述符。

```
function enumerable(bool: boolean) {
  return (target: any, propertyName: string, descriptor: PropertyDescriptor) => {
    descriptor.enumerable = bool
  }
}
class ClassG {
  private _name: string
  constructor(name: string) {
    this._name = name
  }
  // get和set为一个整体，只需要添加一个装饰器就可以，如果两个都添加就会报错。
  // 不可枚举
  // @enumerable(false)
  // 可枚举
  @enumerable(true)
  get name () {
    return this._name
  }
  set name (name) {
    this._name = name
  }
}
const classG = new ClassG('zzz')
for (const key in classG) {
  // 不可枚举打印
  console.log(key) // _name
  // 可枚举打印
  console.log(key) // _name 、 name
}
```

##### 1.8属性装饰器

属性声明之前声明。属性装饰器不能在声明文件或任何其他环境上下文中使用（例如在declare类中）。

属性装饰器的表达式将在运行时作为函数调用，带有以下两个参数：

1、静态成员的类的构造函数，或实例成员的类的原型。

2、成员的姓名。

```
function printPropertyName (target: any, propertyName: string) {
  console.log(propertyName) // name
}
class ClassH {
  @printPropertyName
  public name: string
}
```

##### 1.9参数装饰器

参数装饰只是一个参数声明之前声明。参数装饰器应用于类构造函数或方法声明的函数。参数装饰器不能用在声明文件、重载或任何其他环境上下文中（例如在declare类中）。



参数装饰器的表达式将在运行时作为函数调用，带有以下三个参数：

1、静态成员的类的构造函数，或实例成员的类的原型。

2、成员的姓名。

3、函数参数列表中参数的序数索引。

```
function required(target: any, propertyName: string, index: number) {
  console.log(`修饰的是${propertyName}的第${index + 1}个参数`)
}
class ClassI {
  public name: string = 'zzz'
  public age: number = 18
  public getInfo(prefix: string,@required infoType: string): any {
    return prefix + ' ' + this[infoType]
  }
}
interface ClassI {
  [key: string]: string | number | Function
}
const classI = new ClassI()
classI.getInfo('hihi', 'age') // 修饰的是getInfo的第2个参数
```

### 18.混入

- 对象的混入
- 类的混入

混入就是把两个对象或者类的内容合到一起，实现功能复用。

#### 1.对象的混入

```
interface ObjectA {
  a: string
}
interface ObjectB {
  b: string
}
let Aa: ObjectA = {
  a: 'a'
}
let Bb: ObjectB = {
  b: 'b'
}
let AB = Object.assign(Aa, Bb) // 类型推断出AB是ObjectA & ObjectB的交叉类型
console.log(AB) // {a: "a", b: "b"}
```

#### 2.类的混入

```
class ClassAa {
  public isA: boolean
  public funcA() {}
}
class ClassBb {
  public isB: boolean
  public funcB() {}
}
// 继承多个类用逗号隔开
class ClassAB implements ClassAa, ClassBb {
  public isA: boolean = false
  public isB: boolean = false
  public funcA: () => void
  public funcB: () => void
  constructor() {}
}
function mixins(base: any, from: any) {
  from.forEach((fromItem) => {
    // getOwnPropertyNames方法会保留类的自身属性和方法去掉继承过来的属性和方法
    Object.getOwnPropertyNames(fromItem.prototype).forEach((key) => {
      console.log(key)
      base.prototype[key] = fromItem.prototype[key]
    })
  })
}
mixins(ClassAB, [ClassAa, ClassBb]) // constructor funcA constructor funcB
const ab = new ClassAB()
console.log(ab) // ClassAB {isA: false, isB: false}
// ClassAB {isA: false, isB: false}
// isA: false
// isB: false
// [[Prototype]]: Object
// funcA: ƒ ()
// funcB: ƒ ()
// constructor: ƒ ClassBb()
// [[Prototype]]: Object
```

### 19.其他重要更新

- async函数以及promise
- 动态导入表达式
- 弱类型探测
- ...展开操作符

#### 1.async函数以及promise

promise

```
function getIndex (bool) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log(1)
      if (bool) { resolve('a')}
      else { reject(Error('error'))}
    }, 1000)
  })
}
getIndex(false).then((res) => {
  console.log(res)
}).catch((error) => {
  console.log(error)
})
```

使用async asait

```
function getIndex (bool) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log(1)
      if (bool) { resolve('a')}
      else { reject(Error('error'))}
    }, 1000)
  })
}
async function asyncFunction () {
  try {
    const res = await getIndex(true)
  console.log(res)
  } catch (error) {
    console.log(error)
  }
}
asyncFunction()
```

例子

```
interface Res {
  data: {
    [key: string]: any
  }
}
namespace axios {
  export function post(url: string, config: object): Promise<Res> {
    return new Promise((resolve, rejcet) => {
      setTimeout(() => {
        const res: Res = { data: {}}
        if (url === '/login') { res.data.user_id = 111}
        else { res.data.role = 'admin'}
        console.log(2)
        resolve(res)
      })
    })
  }
}
interface LoginInfo {
  user_name: string
  password: string
}
async function loginReg({ user_name, password}:LoginInfo) {
  try {
    console.log(1)
    const res = await axios.post('login', {
      data: {
        user_name,
        password
      }
    })
    console.log(3)
    return res
  } catch(error) {
    throw new Error(error)
  }
}
async function getRoleReg(user_id: number) {
  try {
    const res = await axios.post('/user_roles', {
      data: {
        user_id,
      }
    })
    return res
  } catch (error) {
    throw new Error(error)
  }
}
loginReg({ user_name: 'zzz', password: '123'}).then((res) => {
  const { data: { user_id } } = res
  getRoleReg(user_id).then(res => {
    const { data: { role }} = res
    console.log(role)
  })
})
```

#### 2.动态导入表达式

```
async function getTime (format: string) {
  const moment = await import ('moment')
  return moment.default().format(format)
}
  // 动态导入只会再调用的时候再使用
getTime('L').then((res) => {
  console.log(res)
})
```

#### 3.弱类型探测

任何可选属性都被标记为弱类型

```
interface ObjIn {
  name?: string
  age?: number
}
let objIn = {
  sex: 'man'
}
function printInfo (info: ObjIn) {
  console.log(info)
}
// printInfo(objIn) // 报错，不具有相同的属性
printInfo(objIn as ObjIn) // 使用类型断言让他就是这个类型
```

#### 4....展开操作符

泛型中使用

```
function mergeOptions<T, U extends string> (op1: T, op2: U) {
  return { ...op1, op2}
}
// console.log(mergeOptions({ a: 'a'}, 'name')) // {a: "a", op2: "name"}
function getExcludeProp<T extends { props: string}>(obj: T) {
  const { props, ...rest } = obj
  return rest
}
const obj = {
  props: 'something',
  name: 'zzz',
  age: 18
}
console.log(getExcludeProp(obj)) // {name: "zzz", age: 18}
```

### 20.声明文件

- 识别已有js库的类型
  - 全局库
  - 模块化库
  - UMD库
- 处理库声明文件
  - 模块插件或UMD插件
  - 修改全局的模块
  - 使用依赖
  - 快捷外部模块声明

#### 1.识别已有js库的类型

判断库的类型：

```
在文档中既可以使用<script src="xx"></script>引入，也可以import或者require方式引入就是UMD库
```

```
// 比如jquery就是全局库
$(function () {})
仅能使用<script src="xx"></script>引入即为全局库
在源码中有使用var、使用if判断是否有document或者赋值给window变量有可能就是全局库
```

##### 1.1、全局库

###### 写一个简单的全局库

创建handle-title.sj文件

```
// handle-title.js

function setTitle(title) {
  document && (document.title = title)
}
function getTitle () {
  return document ? document.title : ''
}
let documentTitle = getTitle()
```

安装webpack的一个插件`npm install copy-webpack-plugin@5.0.2 -D`

###### 修改一下webpack配置

```
// build/webpack.config.js

  plugins: [
    new CopyWebpackPlugin([{
      // 复制的路径
      from: path.join(__dirname, '../src/modules/handle-title.js'),
      // 输出的路径
      to: path.resolve(__dirname, '../dist')
    }])
  ]
```

在html文件引入

```
<script src="./handle-title.js"></script>
```

使用`npm run build`打包

在handle-title.js中打印一下测试是否引用成功

现在需要使用它，创建一个ts文件。

使用

```
// decalration-files.ts
console.log(documentTitle)
```

在ts中测试，它会报错cannot find name 'documentTitle'。因为还需要再配置一下tsconfig.json文件

再tsconfig.json中查找include添加，如果没有就自己添加include字段。

```
  "include": [
  // 默认情况下不写它会把所有.ts、.d.ts文件都引入。
  // 这里只指定src下的所有.ts和.d.ts文件
    "./src/**/*.ts",
    "./src/**/*.d.ts"
  ]
```

###### 编写声明

现在要给handle-title.js写声明文件。再src下创建globals.d.ts文件

```
// globals.d.ts
// 它是全局的使用declare来修饰
declare function setTitle(title: string | number): void

declare function getTitle(): string

declare let documentTitle: string
```

再次调试打印的ts文件

```
// decalration-files.ts
setTitle('zzz')
console.log(documentTitle) // 现在title就变为zzz了
console.log(getTitle()) // zzz
```

##### 1.2、模块化库

判断模块化库

使用export或module.exports等导出语句的库就是模块化库

模板声明文件：

1. module-function.d.ts // 导入后可以直接当函数使用

2. module.class.d.ts // 当类创建实例使用

3. module.d.ts // 既不能当函数调用又不能当类创建实例

##### 1.3、UMD库

判断UMD库，UMD库将全局库和模块化库进行了合并。它会首先模块中有没有模块加载器，判断`typeof define === 'funciton'`有define方法的。`typeof module === 'object' module.exports`有module方法的。

#### 2.处理库声明文件

##### 2.1、模块插件或UMD插件

比如现在有一个库判断有没有声明文件，使用`npm install @types/arr-diff(这是库的名字) -D`，如果安装成功了说明有这个库的声明文件。提示+ @types/(库名)@版本，说明库的声明文件安装成功了。

##### 2.2、修改全局的模块
###### 给修改全局的文件写声明文件

现在模拟一下，创建add-methods-to-string.js文件

```
// add-methods-to-string.js
String.prototype.getFirstLetter = function () {
  return this[0]
}
```

这是一个挂在到原型上的全局的函数。现在帮它写一个声明文件。

```
// globals.d.ts
// 还是再globals.d.ts文件下书写
interface String {
  getFirstLetter(): string
}
```

使用这个插件啊

```
// declaretion-files.ts
// 引入插件
import '../modules/add-methods-to-string.js'
const name = 'zzz'
console.log(name.getFirstLetter()) // z
```
##### 2.3、使用依赖

有些库会依赖其他的库，比如依赖node下的fa、path等等。所以再定义库的声明文件的时候，声明对其他库的依赖，从而加载其他库的声明。如果是依赖全局的库可以使用`/// <reference types="(某个全局库)">`，然后再引入这个库文件，比如引入moment`import * as moment from 'moment'`。

如果全局库依赖的是某个UMD库，也可以使用`/// <reference types="(某个全局库)">`来引入。

注意点：

1. 防止命名冲突：在写全局命名声明的时候，在全局命名返回大量定义类型，有时会导致冲突。建议相关的命名放在命名空间里。
2. ES6的模块的插件影响：一些开发者在为库开发了插件在位于原有库的基础上添加更多的功能，这往往需要修改原有的库导出的模块。ES6导出的模块是不允许修改的，但是在commonjs和其他加载器中是允许的。
3. ES6模块的调用：在使用一些库的时候，引入的模块可以做为函数直接调用。ES6顶层模块对象是一个对象，它不能作为函数调用。

##### 2.4、快捷外部模块声明

不想为模块写声明文件。比如为moment写一个快捷声明就可以在typings文件夹下创建一个moment文件夹在moment文件夹下创建index.d.ts文件

```
// index.d.ts
declare module 'moment'
```

