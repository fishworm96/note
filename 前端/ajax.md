# Ajax基础

## Ajax的应用场景

1.页面上拉加载更多数据

2.列表数据无刷新分页

3.表单项离开焦点数据证

4.搜索框提示文字下拉列表

## Ajax的运行环境

Ajax技术需要运行在网站环境中才能生效，当前课程会使用Node创建的服务器作为网站服务器

## Ajax运行原理

Ajax相当于浏览器发送请求与接受相应的代理人，以实现在不影响用户浏览页面的情况下，局部更新页面数据，从而提高用户体验。

## Ajax的实现步骤

1.创建Ajax对象

```
var xhr = new XMLHttpRequest();
```

2.告诉Ajax请求地址以及请求方式

```
xhr.open('get/post', 'http://www.example.com');
```

3.发送请求

```
xhr.send();
```

4.获取服务器端与客户端的响应数据

```
xhr.onload = function () {
	console.log(xhr.responseText);
}
```

## 服务器端响应的数据格式

在真实的项目中，服务器端大多数情况下会以JSON对线作为响应数据的格式。当客户端拿到相应数据时，要将JSON数据和HTML字符串进行拼接，然后将拼接的结果展示在页面中。

在http请求与响应的过程中，无论是请求参数还是响应内容，如果是对象类型，最终都会被转换为对象字符串进行输出。

```
JSON.parse() //将json字符串转换为json对象
```

## 请求参数传递

传统网站表单提交

```
<form method="get" action="http://www.example.com">
	<input type="text" name='username'/>
	<input type="password" name="password">
</form>
<!- http://www.example.com?username=zhangsan&password=123456 -->
```

GET请求方式

```
xhr.open('get', 'http://www.example.com?name=zhangsan&age=20');
```

POST请求方式

```
xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded')
xhr.send('name=zhangsan&age=20');
```

## 请求参数格式

1.application/x-www-form-urlencoded

```
name=zhangsan&age=20&sex=男
```

2.application/json

```
{naem:'zhangsan', age='20', sex:'男'}
```

在请求头中指定Content-Type属性的值是application/json，告诉服务器端当前请求参数的格式是json

```
JSON.stringify() //将json对象转换为json字符串
```

注意：get请求是不能提交json对象数据类型的，传统网站的表单提交也是不支持json对象数据格式的。

## 获取服务器端的响应

### Ajax状态码

在创建ajax对象，配置ajax对象，发送请求，以及接收完服务器端响应数据，这个过程中的每一个步骤都会对应一个数值，这个数值就是ajax状态码

0：请求未初始化(还没有调用open())

1：请求已经建立，但是还没有发送(还没有调用send())

2：请求已经发送

3：请求正在处理中，通常响应中已经有部分数据可以用了

4：响应已经完成，可以获取并使用服务器的响应了

```
xhr.readyState //获取Ajax状态码
```

两种获取服务器端响应方式的区别

| 区别描述               | onload事件 | onreadystatechange事件 |
| ---------------------- | ---------- | ---------------------- |
| 是否兼容IE低版本       | 不兼容     | 兼容                   |
| 是否需要判断Ajax状态码 | 不需要     | 需要                   |
| 被调用次数             | 一次       | 多次                   |

## Ajax错误处理

1.网络通常，服务器端能接收到请求，服务器端返回的结果不是预期结果

可以判断服务器端的返回码，分别进行处理。xhr.status获取http状态码

2.网络通常，服务器端没有接收到请求，返回404状态码

检查请求地址是否错误

3.网络通常，服务器端能接收到请求，服务器端返回500状态码

服务器端错误，找后端程序员进行沟通

4.网络中断，请求无法发送到服务器端

会触发xhr对象下面的onerror事件，在onerror事件处理函数中对错误进行处理

## 低版本IE浏览器的缓存问题

问题：在低版本的 IE 浏览器中，Ajax 请求有严重的缓存问题，即在请求地址不发生变化的情况下，只有第一次请求会真正发送到服务器端，后续的请求都会从浏览器的缓存中获取结果。即使服务器端的数据更新了，客户端依然拿到的是缓存中的旧数据。

解决方案：在请求地址的后面加请求参数，保证每一次请求中的请求参数的值不相同。 

```
xhr.open('get', 'http://www.example.com?t'+ Math.random());
```

## **同步异步概述**

**同步**

l 一个人同一时间只能做一件事情，只有一件事情做完，才能做另外一件事情。

l 落实到代码中，就是上一行代码执行完成后，才能执行下一行代码，即代码逐行执行。

```
 console.log('before');

 console.log('after');


```

**异步**

l一个人一件事情做了一半，转而去做其他事情，当其他事情做完以后，再回过头来继续做之前未完成的事情。

l落实到代码上，就是异步代码虽然需要花费时间去执行，但程序不会等待异步代码执行完成后再继续执行后续代码，而是直接执行后续代码，当后续代码执行完成后再回头看异步代码是否返回结果，如果已有返回结果，再调用事先准备好的回调函数处理异步代码执行的结果。

```
 console**.**log**(**'before'**);**

 **setTimeout****(**

  **()** **=>** **{** console**.**log**(**'last'**);**

 **},** 2000**);**

 console**.**log**(**'after'**);
```

# AJax异步编程

## Ajax封装

问题：发送一次请求代码过多，发送多次请求代码冗余且负责

解决方案：将请求代码封装到函数中，发送请求时调用函数即可

```
ajax({
	type: 'get',
	url: 'http://localhost:4000/first',
	success: function (data) {
		console.log(data);
	}
})
```

# 模板引擎

## 模板渲染

```
<script type="text/html" id="tpl">
<div>
	<span>{{name}}</span>
	<span>{{age}}</span>
</div>
<scipt>
```

```
//将特定模板与特定数据进行拼接
const html = template('tpl', {
	name: '张三',
	age: 20
});
```

## 模板语法

### 原文输出

如果数据中携带HTML标签，默认情况下，模板引擎不会解析标签，会将其转义后原文输出。

```
<h2>{{@ value}}</h2>
```

### 条件判断

```
{{if 条件}}...{{/if}}
{{if v1}}...{{else if v2}}...{{/if}}
```

```
{{if 条件}}
	<div>条件成立 显示我</div>
{{else}}
	<div>条件不成立 显示我</div>
</if>
```

### 循环

```
{{each target}}
	{{$index}} {{$value}}
{{/each}}
```

### 导入模块变量

```
<div>$imports.dataFormat(time)</div>
```

```
template.defaults.imports.变量名 = 变量值;
$imports.变量名称
```

```
function dateFormat (未格式化的原始时间) {
	return '已经格式化好的当前时间'
}
template.defaults.imports.dateFormat = dateFormat;
```

```
例：
<script type="text/html" id="tp;"><div>当前时间是: {{$imports.dateFormat(date)}}</div></script>
<script>
        window.onload = function () {    
            //将方法导入到模板中
            template.defaults.imports.dateFormat = dateFormat;
            //这是告诉模板引擎将模板id为tpl的模板和data数据对象惊醒拼接
            var html = template('tpl', {
                data: new Date()
            });
            document.getElementById('container').innerHTML = html;
            //特定格式
            function dateFormat(date) {
                return date.getFullYear() + '年' + (date.getMoth() + 1) + '月' + date.getDate() + '日';
            }
        }
</script>
```





# FormData对象的使用

## FormData对象的使用

1.准备HTML对象的使用

```
<form id="form">
	<input type="text" name="username" />
	<input type="password" name="password">
	<input type="button">
</form>
```

2.将HTML表单转化为formdata对象

```
var form = document.getElementById('form');
var formData = new Formata(form);
```

3.提交表单对象

```
xhr.send(formData);
```

## FormData 对象的实力方法

1.获取表单对象中属性的值

```
formData.get('key');
```

2.设置表单对象中属性的值

```
formData.set('key', 'value');
```

```
//存在表单值就覆盖原有值
console.log(formData.set('username', 'itcast'));
//不存在表单值就创建
console.log(formData.set('age', 100));
```

3.删除表单对象中属性的值

```
formData.delete('key');
```

```
//删除用户输入的值
formData.delete('password');
```

4.想表单对象中追加属性值

```
formData.append('key', 'value');
```

```
//创建空的表单对象
var f = new formData();
f.append('sex', '男');
console.log(f.get('sex'));
```

注意：set方法与append方法的区别是，在属性名已经存在的情况下，set会覆盖已有键名的值，append会保留两个值。只打印最后一个值

## FormData二进制

```
<input type="file" id="file" />
```

```
var file = document.query.Selector('#file');
//当用户选择文件的时候
file.onchange = function () {
	//创建空表单对象
	var formData = new FomrData();
	//将用户选择的二进制文件追加到表单对象中
	formData.append('attrName', this.files[0]);
	//配置ajax对象,请求方式必须为post
	xhr.open('post', 'www.example.com');
	xhr.send(fomrData);
}
```

FormData文件上传进度展示

```
//当用户选择文件的时候
file.onchange = function () {
	//文件上传过程中持续触发onprogress事件
	xhr.upload.onprogress = function (ev) {
		//当前上传文件大小/文件总大小 再将结果转换为百分数
		//将结果赋值给进度条的宽度属性
		bar.style.width = (ev.loaded/ ev.total) * 100 + '%';
	}
}
```

## FormData文件上传图片即时预览

在我们将图片上传到服务器端以后，服务器端通常都会将图片地址做为响应数据传递到客户端，客户端可以从响应数据获取图片地址，然后将图片再显示在页面中。

## 使用JSONP解决同源限制问题

jsonp是json with padding 的缩写，它不属于Ajax请求，但它可以模拟Ajax请求。

1.将不同源的服务器端请求地址写在script标签的src属性中

```
<script src="www.example.com"></script>
<script src="http://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
```

2.服务器端响应数据必须是一个函数的调用，真正要发送给客户端的数据需要作为函数调用的参数。

```
const data= 'fn({naem:"张三", age:"20"})';
res.send(data);
```

3.在客户端全局作用域下定义函数fn

```
function fn (data) {}
```

4.在fn函数内部对服务器端返回的数据进行处理

```
function fn (data) { console.log(data);}
```

## JSONP代码优化

1.客户端需要将函数名称传递到服务端

2.将script请求的发送变成动态请求

3.封装jsonp函数，方便请求发送

## CORS跨域资源共享

CORS：全称Cross-orgin resource sharing，即跨域资源共享，它允许浏览器想跨域服务器端发送Ajax请求，克服了只能同源使用的限制。

| 浏览器端 | 请求/请求头(origin)                       | 服务器端 |
| -------- | ----------------------------------------- | -------- |
|          | 响应/响应头(Access-Control-Access-Origin) |          |

```
origin: http://localhost:4000
```

```
Access-Control-Allow-Origin: 'http://localhost:4000'
Access-control-Allow-Origin: '*'
```

## withCredentials属性

在使用Ajax技术发送跨域请求时，默认情况下不会在请求中携带cookie信息。

withCredentials: 指定在涉及到跨域请求时，是否携带cookie信息，默认值为false

Access-Control-Allow-Credentials：true允许客户端发送请求时携带cookie

$.ajax()方法概述

作用：发送Ajax请求。

```
$.ajax({
	type: 'get',
	url: 'http://www.example.com',
	data: { name: 'zhangsan', age: '20'},
	contentType: 'application/x-www-form-urlencoded',
	beforeSend: function () {
		return false
	},
	success: function (response) {},
	error: function () {}
});
{
	data: 'name=zhangsan&age=20'
}
{
	contentType: 'application/json'
}
JSON.stringify({name: 'zhangsan', age:'20'})
```

### serialize方法

作用：将表单中的数据自动拼接成字符串类型的参数

```
var params = $('#form').serialize();
//name=zhangsan&age=30
```

## Jquery中Ajax全局事件

### 全局时间

当页面中有ajax请求

```
.ajaxStart()	//当请求开始发送时触发
.ajaxComplete	//当请求完成时触发
```

### 第三方插件NProgress

```
<link rel='stylesheet' href='nprogress.css'/>
<script src='nprogress.js'></script>
```

```
NProgress.start();	//进度条开始运动
NProgress.done();	//进度条结束运动
```

```
例：
	$(document).on('ajaxStart', function () {
		NProgress.start()
	})
	
	$(document).on('ajaxComplete', function () {
		NProgress.done()
	})
```

## RESTful风格的API

GET:	获取数据

POST：	添加数据

PUT：	更新数据

DELETE：	删除数据

### RESTful API的实现

GET：	http://www.example.com/user	获取用户列表数据

POST：	http://www.example.com/user	创建(添加)用户数据

GET：	http://www.example.com/user/1	获取用户ID为1的用户信息

PUT：	http://www.example.com/user/1	修改用户ID为1的用户信息

DELETE：	http://www.example.com/user/1	删除用户ID为1的用户信息

```
例：
//获取用户列表信息
app.get('/users', (req, res) => {
	res.send('当前是获取用户列表信息的路由');
});

//获取某一个用户信息的路由
app.get('/users/:id', (req, res) => {
	//获取客户端传递过来的用户id
	const id = req.params.id;
	res.send('当前我们是在获取id为$(id)用户信息');
})

//删除某一个用户
app.delete('/users/:id', (req, res) => {
	//获取客户端传递过来的用户id
	const id = req.params.id;
	res.send('当前我们是在删除id为${id}用户信息')
})

//修改某一个用户
app.put('/users/:id', (req, res) => {
	//获取客户端传递过来的用户id
	const id = req.params.id;
	res.send('当前我们是在修改id为${id}用户信息')
})
```

```
在$.ajax页面实现
例:
//获取用户信息
$.ajax({
	type: 'get',
	url: '/users',
	success: function (response) {
		console.log(response)
	}
})

//获取id为1的用户信息
$.ajax({
	type: 'get',
	url: '/users/1',
	success: function (response) {
		console.log(response)
	}
})

//删除id为1的用户信息
$.ajax({
	type: 'delete',
	url: '/users/1',
	success: function (response) {
		console.log(response)
	}
})
```

## XML

全称是extensible markup language, 代表可扩展标记语言，它的作用是传输和存储数据

```
客户端例：
function () {
	var xhr = new XMLHttpRequest();
	xhr.open('get', '/xml');
	xhr.send();
	xhr.onload = function () {
		//xhr.responseXML 获取服务器端返回的xml数据
		var xmlDocument = xhr.responseXML;
		var title = xmlDocument.getElementsByTagName('title')[0].innerHTML;
		container.innerHTML = title;
	}
}
```

```
服务器端例：
app.get('/xml', (req, res) => {
	res.header('content-type', 'html/xml');
	res.send('<message><title>消息标题</title></message>')
})
```

