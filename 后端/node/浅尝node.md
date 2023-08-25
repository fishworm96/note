# Node

## Node.js全局对象global

在浏览器中全局对象是window,在Node中全局对象是global。

Node中全局对象下有以下方法，可以在任何地方使用，global可以省略

```
console.log()	在控制台中输出
setTimeout()	设置超时定时器
clearTimeout()	清除超时定时器
setInterval		设置间歇定时器
clearInterval	清楚间歇定时器
```

### Node.js中模块化开发规范

Node.js规定一个Javascript文件就是一个模块，模块内部定义的变量和函数默认情况下在外部无法得到

模块内部可以使用**exports对象进行成员导出**，使用**require方法**导入其他模块

```node
exports.变量 = 变量	//导出
```

```
require.变量 = 变量	//导入
```

#####  模块成员导出的另一种方式

```
module.exports.变量 = 变量;
```

exports是module.exports的别名(地址引用关系)，导出对象最终以module.exports为准

#### 系统模块fs文件操作

f:file文件,s:system系统，文件操作系统

```
const fs = require('fs');
```

##### 读取文件内容:错误优先的回调函数

```
fs.readFile('文件路径/文件名称',['文件编码'], callback);
```

```
例：
//读取上一级css目录下的中的base.css
fs.readFile('../css/base.css','utf-8' (err, doc) => {
	//如果文件读取错误 参数err的值为错误对象 否则err的值为null
	//doc参数为文件内容
	if(err == null) {
		//在控制台中输出文件内容
		console.log(doc);
	}
});
```



##### 写入文件内容：错误日志

```node
fs.writeFile('文件路径/文件名称', '数据', callback);
```

```
例：
const content = '<h3>正在使用fs.writeFile写入文件内容</h3>';
fs.writeFile('../index.html', content, err => {
	if (err != null) {
		console.log(err);
		return;
	}
	console.log('文件写入成功');
})
```

#### 系统模块path路径操作

##### 路径拼接语法

```
path.join('路径','路径', ...)
例：
//导入path模块
const path = require('path');
//路径拼接
let finialpath = path.join('itcast', 'a', 'b', 'c.css');
//输出结果 itcast\a\b\c.ss
console.log(finialpath)
```

#### 相对路径vs绝对路径

大多数情况下使用绝对路径，因为相对路径有时候相对的是命令行工具的当前工作目录

在读取文件或者设置文件路径时都会选择绝对路径

使用__dirname获取当前文件夹所在的绝对路径

```
例：
const fs = require('fs');
const path = require('path');
fs.readFile(path.join(__dirname, '01.helloworld.js'), 'utf8', (err, doc) => {
	console.log(err)
	console.log(doc)
});
```

### 第三方模块

#### 获取第三方模块

下载：npm install 模块名称

```
npm install formidable
```

卸载：npm uninstall packgae 模块名称

```
npm uninsatll formidable
```

全局安装与本地安装

命令行工具：全局安装

库文件：本地安装

#### 第三方模块nodemon(第三方模块名称)

nodemon是一个命令行工具，用以辅助项目开发

在Node.js中，每次修改文件都要在命令行工具中重新执行该文件，非常繁琐

使用步骤

1.使用npm install nodemon -g下载他

```
npm install nodemon -g 
```

2.在命令行工具中使用nodemon命令替代node命令执行文件

```
调用： nodemon .\01.helloworld.js(这是路径)
```

#### 第三方模块nrm

nrm(npm registry manager): npm下载地址切换工具

npm默认的下载地址在国外，国内下载速度慢

使用步骤

1.使用npm install nrm -g下载它

2.查询可用下载地址列表nrm ls

3.切换npm下载地址nrm use 下载地址名称

#### 第三方模块Gulp

##### Gulp使用

1.使用npm Install gulp下载gulp库文件

2.在项目根目录下建立gulpfile.js文件

3.重构项目的文件夹结构src目录防止源代码文件dist目录放置构建后文件

4.在gulpfile.js文件中编写任务

5.在命令行工具中执行gulp任务

##### Gulp中提供的方法

gulp.src():获取任务要处理的文件

gulp.dest():输出文件

gulp.task():建立gulp任务

gulp.watch():监控文件的变化

```
const gulp = require('gulp');
//使用gulp.task()方法建立任务
gulp.task('first', () => {
	//获取要处理的文件
	gulp.src('./src/css/base.css')
	//将处理后的文件输出到dist目录
	.pipe(gulp.dest('./dist/css'));
});
```

#### Gulp插件

gulp-htmlmin:html文件压缩

gulp-csso:压缩css

gulp-babel:javascript语法转化

gulp-less:less语法转化

gulp-uglify:压缩混淆javascript

gulp-file-include公共文件包含

browsersync浏览器实时同步

```
1.下载模块
2.调用
//引用gulp模块 
const gulp = require('gulp');
const htmlmin = require('gulp-htmlmin');
var fileinclude = require('gulp-file-include');
const { createBrotliCompress } = require('zlib');
const less = require('gulp-less');
const csso = require('gulp-csso')
const babel = require('gulp-babel');
const uglify = require('gulp-uglify');

//使用gulp.task建立任务
//1.任务的名称
//2.任务的回调函数
gulp.task('first', () => {
    //1.使用gulp
    console.log('第一次使用gulp任务')
    gulp.src('./src/css/base.css')
        .pipe(gulp.dest('./dist/css'));
    cb();
});

//html任务
//1.html文件中代码的压缩操作
//2.抽取html文件中的公共代码
gulp.task('htmlmin', () => {
    gulp.src('./src/*.html')
        .pipe(fileinclude())
        //压缩html文件中的代码
        .pipe(htmlmin({ collapseWhitespace: true }))
        .pipe(gulp.dest('dist'));
});

//css任务
//1.less语法转换
//2.css代码压缩
gulp.task('cssmin', () => {
    //选择css目录先的所有less文件以及css文件
    gulp.src(['./src/css/*.less', './src/css/*.css'])
        //将less语法转化为css语法
        .pipe(less())
        //将css代码进行压缩
        .pipe(csso())
        //将处理的结果进行输出
        .pipe(gulp.dest('dist/css'))
});

//js任务
//1.ES6代码转换
//2.代码压缩
gulp.task('jsmin', () => {
    gulp.src('./src/js/*.js')
        .pipe(babel({
            //它可以判断当前代码的运行环境 将代码转换为当前运行环境支持的代码
            presets: ['@babel/env']
        }))
        //压缩js代码
        .pipe(uglify())
        .pipe(gulp.dest('dist/js'))
})

//复制文件夹
gulp.task('copy', () => {

    gulp.src('./src/images/*')
        .pipe(gulp.dest('dist/images'));

    gulp.src('./src/lib/*')
        .pipe(gulp.dest('dist/lib'))
})

//构建任务
gulp.task('default', ['htmlmin', 'cssmin', 'jsmin', 'copy']);
```

##### package.json文件的作用

项目描述文件，记录了当前项目信息，例如项目名称、版本、作者、github地址、当前项目依赖了那些第三方模块等。

使用npm init -y命令生成。 这里的-y是yes的意思，不填写信息全部填写默认值

```
{
  "name": "description",	//项目名称
  "version": "1.0.0",	//项目版本
  "description": "",	//项目描述
  "main": "index.js",	//项目的主入口文件
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },	//存储的命令的别名
  "keywords": [],	//关键字
  "author": "",		//项目的作者
  "license": "ISC"	//项目的协议	ISC开放源代码
  "dependencies": {	//依赖
    "formidable": "^1.2.2",
    "mime": "^2.4.7"
  }
}

```

#### 项目依赖

在项目的开发阶段和线上运营阶段，都需要依赖的第三方包，成为项目依赖

使用npm install 包名称命令下载的文件会默认被添加到package.json文件的dependencies字段中

```
{
	"dependencies": {
		"jquery":"^3.3.1"
	}
}
```

#### 开发依赖

在项目的开发阶段需要依赖，线上运营阶段不需要依赖的第三方包，成为开发依赖

使用npm install包名--save-dev命令将包添加到package.json文件的devDependencies字段中

```
{
	"devDependencies": {
		"gulp": "^3.9.1"
	}
}
```

通过npm install不管是开发依赖还是项目依赖都会全部下载

npm install --production只会下载项目依赖

npm install --save-dev只会下载开发依赖

##### package-lock.json文件的作用

锁定包的版本，确保再次下载时不会因为包版本不同而产生问题

加快下载速度，因为该文件中已经记录了项目所依赖第三方包的树状结构和包的下载地址，重新安装时只需下载即可，不需要做额外的工作。

##### Node.js加载机制

```
当模块拥有路径但没有后缀时
require('./find.js');
require('./find');
1.require方法根据模块路径查找模块，如果是完整路径，直接引入模块。
2.如果模块后缀省略，先找同名js文件再找同名js文件夹
3.如果找到了同名文件夹，找文件夹中的index.js
4.如果文件夹中没有index.js就会去当前文件夹中的package.js文件中查找main选项中的入口文件
5.如果找指定的入口文件不存在或者没有指定入口文件就会报错，模块没有被找到

当模块没有路径且没有后缀时
require('find');
1.Node.js会假设它是系统模块
2.Node.js会去node_modules文件夹中
3.首先看是否有该名字的js文件
4.再看是否有该名字的文件夹
5.看文件夹里面是否有index.js
6.如果没有index.js查看该文件夹中的package.json中的main选项确定模块入口文件
```

##### 创建web服务器

```
//用于创建网站服务器的模块
const http = require('http');
//app对象就是网站服务器对象
const app = http.createServer();
//当客户端有请求来的时候
app.on('request', (req, res) => {
    res.end('<h1>hi user</h1>');
});
//监听端口
app.listen(4000);
console.log('网站服务器启动成功');
```

### http协议

#### 请求报文

##### 1.请求方式(request method)

get请求数据

post发送数据

##### 2.请求地址(Request URL)

```
app.on('request', (req, res) => {
	req.headers	//获取请求报文
	req.url	//获取请求地址
	req.method	//获取请求方法
});
```

#### 响应报文

##### 1.HTTP状态码

```
200请求成功

404请求的资源没有被找到

500服务器端错误

400客户端请求有语法错误
```

##### 2.内容类型

```
text/html

text/css

application/javascrtip

image/jpeg

application/json
```

#### GET请求参数

url里的内置模块parse(解析)返回一个对象 

1.要解析的url地址

2.将查询参数解析成对象形式

ulr.parse(req.url, true);

```
let { query, pathname } = url.parse(req.url, true);

if (pathname == '/index' || pathname == '/') {
    res.end('<h2>欢迎来到首页</h2>');
} else if (pathname == '/list') {
    res.end('welcome to listpage');
} else {
    res.end('not found');
}
```



#### POST请求参数

参数被放置在请求体中进行传输

获取post参数需要使用data事件和end事件

使用querystring系统模块将参数转换为对象格式

```
//导入系统模块querystring用于将HTTP参数转换为对象格式
const querysting = require('querysting');
app.on('require', (req, res) => {
	let postData = '';
	//监听参数传输事件
	req.on('data', (chunk) => postData += chunk;);
	//监听参数传输完毕事件
	req.on('end', () => {
		console.log(querystring.parse(postData));
	});
});
```

#### 路由

路由是指客户端请求地址与服务端程序代码的对应关系

```
//当客户端发来请求的时候
app.on('request', (req, res) => {
	//获取客户端的请求路径
	let { pathname } = url.parse(req.url);
	if (pathname == '/' || pathname =='/index') {
		res.end('欢迎来到首页');
	} else if (pathname == '/list') {
		res.end('欢迎来到列表也');
	} else {
		res.end('抱歉，您访问的页面不存在');
	}
})
```

```
//1.引入系统模块http
//2.创建网站服务器
//3.为网站服务器对象题那家请求事件
//4.实现路由功能
//  1.获取客户端的请求方式
//  2.获取客户端的请求地址
const http = require('http');
const url = require('url');

const app = http.createServer();

app.on('request', (req, res) => {
    //获取请求方式
    const method = req.method.toLowerCase();
    //获取请求地址
    const pathname = url.parse(req.url).pathname;

    res.writeHead(200, {
        'content-type': 'text/html;charset=utf8'
    });

    if (method == 'get') {
        if (pathname == '/' || pathname == '/index') {
            res.end('欢迎来到首页');
        } else if (pathname == '/list') {
            res.end('欢迎来到列表页');
        } else {
            res.end('您访问的页面不存在')
        }

    } else if (method == 'post') {

    }
});
app.listen(4000);
console.log('服务器启动成功');
```

#### 静态资源

服务器端不需要处理，可以直接响应给客户端的资源就是静态资源，例如css，javascript，image文件。

#### 动态资源

相同的请求地址不同的响应资源，这种资源就是动态资源

```
www.index.cn/article?id=1
www.index.cn/article?id=2
```

#### 同步API，异步API

同步API可以从返回值中拿到API执行的结果，但是异步API是不可以的

同步API从上到下一次执行，情面代码会阻塞后面代码的执行

异步API不会等待API执行完成后再向下执行代码

#### Promise

promise出现的目的是解决Node.js异步编程中回调地狱的问题

```
let promise = new promise((resolve, reject) => {
	setTimeout(() => {
		if(true) {
			resolve({name: '张三'})
		} else {
			reject('失败了')
		}
	}, 2000);
});
promise.then(result => console.log(result));//{name: '张三'}
		.catch(error => console.log(error)); //失败了
```

```
const fs = require('fs');
let p1 = new Promise((resolve, reiect) => {
	fs.readFile('./1.txt', 'utf8' (err, result) => {
		resolve(result)
	})
});
let p2 = new Promise((resolve, reiect) => {
	fs.readFile('./2.txt', 'utf8' (err, result) => {
		resolve(result)
	})
});
let p3 = new Promise((resolve, reiect) => {
	fs.readFile('./3.txt', 'utf8' (err, result) => {
		resolve(result)
	})
});
//链式调用编程
p1().then((r1) =>{
	console.log(r1);
	return p2();
})
.then((r2) => {
	console.log(r2);
	return p3();
})
.then((r3) => {
	console.log(r3)
});
```

#### 异步函数

异步函数是异步语法的终极解决方案，它可以让我们讲异步代码携程同步的形式，让代码不再有回调函数嵌套，使代码变得更清晰明了。

```
const fn = async () => {};
async function fn() {}
```

##### async关键字

1.普通函数定义前面加async关键字 普通函数编程异步函数

2.异步函数默认返回promise对象

3.在异步函数内部使用return关键字进行结果返回 结果会被包裹的promise对象中 return关键字代替了resolve方法

4.在异步函数内部使用throw关键字抛出程序错误

5.调用异步函数在链式调用then方法获取异步函数执行结果

6.调用异步函数在链式调用catch方法获取异步函数执行的错误信息

```
//await关键字
//1.它只能出现在异步函数中
//2.await promise 它可以暂停异步函数的执行 等待promise对象返回结果再向下执行
//3.await promise await后面只能写promise对象 写其他类型的API是不可以的

// async function fn() {
//     throw '发生错误'
//     return 123;
// }

// fn().then(function(data) {
//     console.log(data);
// }).catch(function(err) {
//     console.log(err);
// })

async function p1() {
    return 'p1';
}
async function p2() {
    return 'p2';
}
async function p3() {
    return 'p3';
}
async function run() {
    let r1 = await p1()
    let r2 = await p2()
    let r3 = await p3()
    console.log(r1)
    console.log(r2)
    console.log(r3)
}
```

## 模板引擎

### art-template模板引擎

```
1.在命令行工具中使用npm install art-template命令进行下载
2.使用const template = require('art-template')引入模板引擎
3.告诉模板引擎要拼接的数据和模板在哪 const html = template('模板路径',数据);
```

```
//导入模板引擎模块
const template = require('art-template');
//将特定模板与特定数据进行拼接
const html = template('./views/index.art', {
	data: {
		name: '张三',
		age: 20
	}
});
```

### 模板语法

art-template同时支持两种模板语法：标准语法和原始语法

标准语法可以让模板更容易读写，原始语法具有强大的逻辑处理能力

标准语法：{{数据}}

原始语法：<%=数据 %>

#### 输出

将某项数据输出在模板中，标准语法和原始语法如下：

标准语法：{{ 数据 }}

原始语法：<%= 数据 %>

```
<!-- 标准语法 -->
<h2>{{value}}</h2>
<h2>{{a ? b : c}}</h2>
<h2>{{a + b}}</h2>

<!-- 原始语法 -->
<h2><%= value %></h2>
<h2><%= a ? b : c %></h2>
<h2><%= a + b %></h2>
```

#### 原文输出

如果数据中携带HTML标签，默认模板引擎不会解析，会将其转义后输出。

标注语法：{{@ 数据 }}

原始语法：<%- 数据 %>

#### 条件判断

在模板中可以根据条件来决定先好似哪块HTML代码

```
<!-- 标准语法 -->
{{if 条件 }} ... {{/if}}
{{if v1}} ... {{else if v2}} ... {{/if}}
<!-- 原始语法 -->
<% if (value) { %> ... <% } %>
<% if (v1) { %> ... <} else if (v2) { %> ... <} %>
```

#### 循环

标准语法：{{each 数据}} {{/each}}

原始语法：<% for() {%> <% } %>

```
<!-- 标准语法 -->
{{each target}}
	{{$index}} {{$value}}
{{/each}}
<!-- 原始语法 -->
<% for (var i = 0; i < target.length; i++) { %>
	<%= i %> <%= target[i] %>
<% } %>
```

#### 子模板

使用子模板可以将网站公共区块(头部、底部)抽离岛单独的文件中。

标准语法：{{include '模板路径'}}

原始语法：<% include('模板路径') %>

```
<!-- 标准语法 -->
{{include './header.art'}}
<!-- 原始语法 -->
<% include(./header.art) %>
```

#### 模板继承

使用模板继承可以将网站HTML骨架抽离到单独的文件中，其他页面模板可以继承骨架文件

```
<!doctype html>
<html>
	<head>
	<meat charset="utf-8">
	<title>HTML骨架模板</title>
	{{block 'head'}} ... {{/block}}
	</head>
	<body>
	{{block 'content'}} ... {{/block}}
	</body>
</html>
```

```
<!-- index.art 首页模板 -->
{{extend './layout.art'}}
{{block 'head'}} <link rel="stylesheet" href="custom.css"> {{</block>}}
{{block 'content'}} <p> 首页模板 </p> {{/block}}
```

#### 模板配置

1.向模板中导入变量 template.defaults.imports.变量名 = 变量值;

2.设置模板根目录 template.defaults.root = 模板目录

3.设置模板默认后缀 template.defaults.extname = '.art'



#### 第三方模块router

功能：实现路由

使用步骤

1.获取路由对象

2.调用路由对象提供的方法创建路由

3.启用路由，是路由生效

```
const gerRouter = require('router')
const router = getRouter();
router.get('/add', (req, res) => {
	res.end('Hello World')
})
server.on('request', (req, res) => {
	router(req, res)
})
```

#### 第三方模块server-static

功能：实现静态资源访问服务

步骤

1.引入serve-static模块获取创建静态资源服务功能的方法

2.调用方法创建静态资源服务并指定静态资源服务目录

3.启用静态资源服务功能

```
const serverStaic = require('serve-static')
const serve = serveStatic('public')
server.on('require', () => {
	serve(req, res) 
})
server.listen(3000)
```

