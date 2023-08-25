# VUE

## VUE概述

vue：渐进式JavaScript框架

声明式渲染 ->组件系统->客户端路由->集中式状态管理->项目构建

## 基本使用

### 1.实例参数分析

el：元素的挂在位置(值可以是css选择器或者DOM元素)

data：模型数据(值是一个对象)

### 2.插值表达式用法

```
{{}}
```

将数据填充到HTML标签中

插值表单式支持基本的计算操作

### 3.Vue代码运行原理分析

概述编译过程的概念(Vue语法->原生语法)

Vue代码->Vue框架->原生JS代码

## Vue模板语法

### 指令

#### 1.v-cloak

插值表达式存在的问题： “闪动”

如何解决该问题：使用v-cloak指令

原理：先隐藏，替换好值之后再显示最终的值

```
<style type="text/css">
	[v-cloak]{
		display: none;
	}
</style>

<div v-cloak>{{msg}}</div>
```

#### 2.数据绑定指令

##### v-text填充纯文本

相比插值表达式更简洁

```
例：
	<div v-text="msg"></div>	//返回msg的内容
```

##### v-html填充html片段

存在安全问题

本网站内部数据可以使用，来自第三方的数据不可以用

```
例：
data: {msg1: '<h1>HTMl</h1>'}
<div v-html="msg1"></div>		//HTML
```

##### v-pre填充原始信息

显示原始信息，跳过编译过程

```
<div v-pre>{{msg}}</div>		//{{msg}}
```

#### 3.数据响应式

##### 响应式：

1.html5中的响应式(屏幕尺寸的变化导致样式的变化)

2.数据的响应式(数据的变化导致页面内容的变化)

##### 数据绑定：

将数据填充到标签中

##### v-once只编译一次

显示内容之后不再具有响应式功能

#### 4.双向数据绑定

##### 1.v-model

```
例：	//当'msg'改变时{{msg}}的内容会改成'msg'的内容
	<div>{{msg}}</div>
	<div>
	<input type="text" v-model='msg'>
	</div>
```

##### 2.MVVM设计思想

M(model)

V(view)

VM(View-Model)

#### 5.事件绑定

##### 1.处理事件方法

###### v-on与v-on简写

```
例：
	data:{num=0}
	<div>{{num}}</div>
	<button v-on:click='num++'>点击</button>	//普通方式
	<button @click='num++'>点击1</button>		//简写
```

##### 2.事件函数的调用方式

###### 直接绑定函数名与调用函数

```
例：
	data:{num=0}，methods:{handle: function() {this.num++}}	//this指向VUE
	<div>{{num}}</div>
	<button @click='handle'>点击2</button>	//直接绑定函数名
	<button @click='handle'>点击3</button>	//与调用函数
```

##### 3.事件函数参数传递

###### 普通参数和事件参数

事件绑定-参数传递
1.如果事件直接绑定函数名称，那么默认会传递事件对象作为事件函数的第一个参数
2.如果事件绑定函数调用，那么事件对象必须作为最后一个参数显示传递，并且事件对象的名称必须是$event

```
例：
	methods: {
		handle1: function(event) {
			console.log(event.target.innerHTML)
		},
		handle2: function(p, p1, event) {
			conosle.log(p, p1)
			console.log(event.target.innerHTML)
		}
	}
<button v-on:click='handle'>点击1</button>		//普通参数
<button v-on:click='handle(123, 456, $event)'>点击1</button>	//事件参数
```

##### 4.事件修饰符

###### .stop阻止冒泡

```
<button v-on:click.stop='handle1'>点击1</button>
原生stopPropagetion();
```

###### .prevent阻止默认行为

```
<a href="http://www.baidu.com" v-on:click='handle1'>百度</a>
原生preventDefault();
```

##### 5.按键修饰符

###### .enter回车键

```
<input v-on:keyup.enter='submit'>
```

.delete删除键

```
<input v-on:keyup.delete='handle'>
```

##### 6.自定义按键修饰符

###### 全局config.keyCodes对象

自定义按键修饰符名字是自定义的，但是对应的值必须是按键对应event.keyCode值

```
Vue.config.keyCodes.a = 65
```

##### 7.属性绑定

###### 1.vue动态处理属性

v-bind和缩写

```
<a v-bind:href='url'>跳转</a>	//基本写法
<a :href='url'>跳转</a>		//缩写
```

###### 2.v-model的底层实现

```
methods:{
	hendle: function (event) {
		//使用输入域中的最新的数据覆盖原来的数据
		this.msg = event.target.value;
	}
}
双向绑定三种实现方法
<input type='text' v-bind:value="msg" v-on:input='handle'>
<input type='text' v-bind:value="msg" v-on:input='msg=$event.target.value'>
<input type='text' v-model='msg'>
```

##### 8.样式绑定

###### 1.class样式处理

对象语法

```
<div v-bind:class="{ active: isActive}"></div>
```

```
例：
	<div id="app">
		<div v-bind:class='{active: isActive,error: isError}'>测试样式</div>
		<button v-on:click='handle'>切换</button>
	</div>
	data: {
		isActive: true,
		isError: true
	},
	methods: {
		handle: function () {
			//控制isActive的值在true和false之间进行切换
			this.isActive = !this.isActive,
			this.isError = !this.isError;
		}
	}
```

数组语法

```
<div v-bind:class="[activeClass, errorClass]"></div>
```

```
例：
	<div id="app">
		<div v-bind:class='[activeClass, errorClass]'>测试样式</div>
		<button v-on:click='handle'>切换</button>
	</div>
	data: {
		activeClass: 'active',
		errorClass: 'error'
	},
	methods: {
		handle: function () {
			this.activeClass = '';
			this.errorClass = '';
		}
	}
```

###### 2.style样式处理

对象语法

```
<div v-bind:style="{color: activeColor, fontSize: fontSize}"></div>
```

数组语法

```
<div v-bind:style={baseStyles, overridingStyles}></div>
```

##### 9.分支循环结构

###### 1.分支结构

v-if

v-else

v-else-if

v-show	v-show的原理：控制元素样式是否显示display:none

###### 2.v-if与v-show的区别

v-if控制元素是否渲染到页面

v-show控制元素是否显示(已经旋绕到了页面)

###### 3.循环结构

v-for遍历数组

```
<li v-for='item in list'>{{item}}</li>
<li v-for='(item, index) in list'>{{item}} + '---' + {{index}}</li>
```

key的作用:帮助vue区分不同的元素，从而提高性能

```
<li :key='item.id' v-for='(item, index) in list'>{{item}} + '---' {{index}}</li>
```

v-for遍历对象

```
<div v-for='(value, key, index) in object'></div>
```

v-if和v-for结合使用

```
<div v-if='value==12' v-for='(value, key, index) in object'></div>
```

## 表单操作

### 1.基于Vue的表单操作

input	单行文本

textarea	多行文本

select	下拉多选

radio	单选框

checkbox	多选框

### 2.表单域修饰符

number：转化为数值

trim：去掉开始和结尾的空格

lazy：将input事件切换为change事件

```
<input v-model.number="age" type="number">
```

### 自定义指令

#### 自定义指令的语法规则(获取元素焦点)

```
Vue.directive('focus' {
	inserted: function (el) {
		//获取元素的焦点
		el.focus();
	}
})
```

##### 自定义指令用法

```
<input type="text" v-focus>
```

#### 带参数的自定义指令(改变元素背景色)

```
Vue.directive('color', {
	inserted: function (el, binding) {
		el.style.backgroundColor = binding.value.color;
	}
})
```

##### 指令的用法

```
<input type="text" v-color='{color:"orange"}'>
```

#### 局部指令

```
directives: {
	focus: {
		//指令的定义
		inserted: function (el) {
			el.focus
		}
	}
}
```

### 计算属性

#### 计算属性的用法

```
computed {
	reversedMessage: function () {
		return this.msg.split('').reverse().join('')
	}
}
```

#### 计算属性与方法的区别

计算属性是基于它们的依赖进行缓存的

方法不存在缓存

### 侦听器

#### 1.侦听器的应用场景

数据变化时执行异步或开销较大的操作

#### 2.侦听器的用法

```
watch: {
	firstName: function (val) {
		// val表示变化之后的值
		this.fullName = val + this.lastName;
	},
	lsttName: function (val) {
		this.fullName = this.firstName + val;
	}
}
```

### 过滤器

#### 自定义过滤器

```
Vue.filter('过滤器名称', function (value) {
	//过滤器业务逻辑
	return
})
```

##### 过滤器的使用

```
<div>{{msg | upper}}</div>
<div>{{msg | upper | lower}}</div>
<div v-bind:id="id | formatId"></div>
```

```
例：
	<input type="text" v-model='msg'>
	<div>{{msg | upper}</div>
	
	Vue.filter('upper', function (value) {
	return val.charAt(0).toUpperCase() + val.slice(1);
})
	data: {msg: ''}
```

#### 局部过滤器

```
filters: {
	capitalize: function () {
	}
}
```

#### 带参数的过滤器

```
Vue.filter('format', function (value, arg) {
	//value就是过滤器传递过来的参数
})
```

##### 过滤器的使用

```
<div>{{data | format(`yyyy-MM-dd`)}}</div>
```

### 生命周期

#### 主要阶段

挂载(初始化相关属性)

1.beforeCreate

2.created

3.beforeMount

4.mounted

更新(元素或组件的变更操作)

1.beforeUpdate

2.updated

销毁(销毁相关属性)

1.beforeDestroy

2.destroyed

##### 2.Vue实例的产生过程

1.beforeCreate 在实例初始化之后，数据观测和事件配置之前被调用。

2.created 在实例创建完成后被立即调用。

3.beforeMount 在挂在开始之前被调用。

4.mounted el被新创建的vm.$el替换，并挂在到实例上去之后调用该钩子。

5.beforeUpdat 数据更新时调用，发生在虚拟DOM打补丁之前。

6.update 由于数据更改导致的虚拟DOM重新渲染和打补丁，在这之后会调用该钩子。

7.beforeDestroy 实例销毁之前调用。

8.destroyed 实例销毁后调用

### 数组更新测试

#### 1.变异方法(修改原有数据)

push()	方法可向数组的末尾添加一个或多个元素，并返回新的长度。

pop()	方法用于删除并返回数组的最后一个元素。

shift()	方法用于把数组的第一个元素从其中删除，并返回第一个元素的值。

unshift()	方法可向数组的开头添加一个或更多元素，并返回新的长度。

splice()	方法向/从数组中添加/删除项目，然后返回被删除的项目。

sort()	方法用于对数组的元素进行排序。

reverse()	方法用于颠倒数组中元素的顺序。

#### 2.替换数组(生成新的数组)

filter()	方法创建一个新的数组，新数组中的元素是通过检查指定数组中符合条件的所有元素。

concat()	方法用于连接两个或多个数组。

slice()	方法可从已有的数组中返回选定的元素。

#### 3.修改响应式数据

Vue.set(vm.items, indexOfltem, newValue)

vm.$set(vm.items, indexOfltem, newValue)

1.参数一表示要处理的数组名称

2.参数二表示要处理的数组的索引

3.参数三表示要处理的数组的值

## 组件注册

### 全局组件注册语法

```
Vue.component(组件名称, {
	data: 组件数据,		//data必须是函数
	template: 组件模块内容
})
```

```
//定义一个名为button-counter的心组件
Vue.component('button-counter', {
	data: function () {		//data必须是函数
		return {
			count: 0
		}
	},
	template: '<button v-on:click="count++">点击了{{ count }}次.</button>'
})
```

#### 组件注册注意事项

1.data必须是函数

2.组件模板必须是单个跟元素

3.组件模板内容可以是模板字符串

模板字符串需要浏览器提供支持(ES6语法)

### 组件用法

```
<div id="app">
	<button-counter></button-counter>
</div>
```

#### 组件命名方式

短横线方式

```
Vue.component('my-component', {内容})
```

驼峰式

```
Vue.component('MyComponent', {内容})
如果使用驼峰式命名组件，那么在使用组件的时候，只能在字符串模板中用驼峰的方式使用组件，但是在普通的标签模板中，必须使用短横线的方式使用组件。
```

### 局部组件注册

```
var ComponentA = {}	//抽取出来的内容
var ComponentB = {}
var ComponentC = {}
new Vue({
	el: '#app',
	components: {
		'component-a': ComponentA,
		'component-b': ComponentB,
		'component-c': ComponentC,
	}
})
```

局部组件只能在注册他的父组件中使用

#### 父组件向子组件传值

##### 1.组件内部通过props接收传递过来的值

```
Vue.componetnt('menu-item', {
	props: ['title'],
	template: '<div>{{ title }}</div>'
})
```

##### 2.父组件通过属性将值传递给子组件

```
<menu-item title="来自父组件的数据"></menu-item>
<menu-item :title="title"></menu-item>
```

##### 3.props属性名规则

在props中使用驼峰形式，模板中需要使用短横线形式

字符串形式的模板中没有这个限制

```
Vue.component('menu-item', {
	//在Javascript中是驼峰形式
	props: ['menuTitle'],
	template: '<div>{{menuTitle}}</div>'
})
<!-在html中短横线方式的->
<menu-item menu-title="nihao"></menu-item>
```

props传递数据原则：单向数据流

#### 子组件向父组件传递

##### 1.子组件通过自定义事件向父组件传递信息

```
<button v-on:click='$emit("enlarge-text")'>扩大字体</button>
```

##### 2.父组件监听子组件的事件

```
<menu-item v-on:enlarge-text='fontSize += 0.1'></menu-item>
```

##### 3.子组件通过自定义事件向父组件传递信息

```
<button v-on:click='$emit("enlarge-text", 0.1)'>扩大字体</button>
```

##### 4.父组件监听子组件的事件

```
<ment-item v-on:enlarge-text='fontSize += $event'></menu-item>
```

#### 非父子组件间传值

##### 1.单独的时间中心管理组件间的通信

```
var eventHub = new Vue()
```

##### 2.监听事件与销毁

```
eventHub.$on('add-todo', addTodo)
eventHub.$off('add-todo')
```

##### 3.触发事件

```
eventHub.$emit('add-todo', id)
```

### 组件插槽的作用

父组件向子组件传递内容

#### 1.插槽位置

```
Vue.component('alert-box', {
	template: `
		<div class="demo-alert-box">
		<strong>Error!</strong>
		<slot><slot>	//固定语法
		</div>
	`
})
```

#### 2.插槽内容

```
<alert-box>Something bad happened.</alert-box>
```

### 具名插槽用法

#### 1.插槽定义

```
<div class="container">
	<header>
		<slot name="header"></slot>
	</header>
	<main>
	<slot><slot>
	</main>
    <footer>
    	<slot name="footer"></slot>
    <footer>
</div>
```

#### 2.插槽内容

```
<base-layout>
	<h1 slot="header">标签内容</h1>
	<p>主要内容1</p>
	<p>主要内容2</p>
	<p slot="footer">地步内容</p>
</base-layout>
```

### 作用域插槽

#### 1.插槽定义

```
<ul>
	<li v-for="item in list" v-bind:key="item.id">
		<slot v-bind:item="item">
			{{item.name}}
		</solt>
	</li>
</ul>
```

#### 2.插入内容

```
<fruit-list v-bind:list="list">
	<template slot-scope="slotProps">		//slot-scope关键字
		<strong v-if="slotProps.item.current">
			{{slotProps.item.text}}
		</strong>
	</template>
</fruit-list>
```

## Promise基本用法

实例化Promise对象，构造函数中传递函数，该函数中用于处理异步任务

resolve和reiect两个参数用于处理成功和失败两种情况，并通过p.then获取处理结果

```
var p = new Promise(function(resolve, reject) {
	//成功时调用 
	resolve()
	//失败时调用
    reject()
});
p.then(function(ret) {
	//从resolve得到正常结果
}), function (ret) {
	//从reject得到错误信息
};
```

### 基于Promise处理Ajax请求

#### 1.原生Ajax

```
function queryData() {
	return new Promise(function(resolve, reject) {
		var xhr = new XHLHttpRequest();
		xhr.onreadystatechange = function () {
			if(xhr.readyState !=4) return;
			if(xhr.status == 200) {
				resolve(xhr.responseText)
			}else{
				reject('出错了')
			}
		}
		xhr.open('get', '/data');
		xhr.send(null)
	})
}
```

#### 2.发送多次ajax请求

```
queryData ()
	.then(function (data) {
		return queryData();
	})
	.then(function (data) {
		return queryData();
	})
	.then(function (data) {
		return queryData;
	})
```

#### then参数中的函数返回值

##### 1.返回Promise实例对象

返回的该实例对象会调用下一个then

##### 2.返回普通值

返回的普通值会直接传递给下一个then，通过then参数中的函数的参数接收该值。

#### Pormise常用的API

##### 1.实例方法

p.then()得到异步任务的正确结果

p.catch()获取异常信息

p.finally()成功与否都会执行(暂时不是正式标准)

```
queryData()
.then(function (data){
	console.log(data)
})
.catch(function (data) {
	console.log(data)
})
.finally(function() {
	console.log(finished)
})
```

##### 2.对象方法

Promise.all()并发处理多个异步任务，所有惹怒我都执行完成才能得到结果

Promise.race()并发处理多个异步任务，只要有一个任务完成就能得到结果

```
Promise.all([p1, p2, p3]).then(resutl) => {
	console.log(result)
}
Promise.race([p1, p2, p3]).then(result) => {
	console.log(result)
}
```

## fetch概述

#### 1.基本特性

更加简单的数据获取方式，功能更加强大，更灵活，可以看作是xhr的升级版

基于Promise实现

#### 2.语法结构

```
fetch(url).then(fn2)
			.then(fn3)
			...
			.catch(fn)
```

### fetch的基本用法

```
fetch('/abc').then(data => {
	return data.text();	//text内置方法
}).then(ret => {
	//注意这里得到的才是最终的数据
	console.log(ret);
})
```

### fetch请求参数

#### 1.常用配置选项

method(String):HTTP请求方法，默认为GET(GET、POST、PUT、DELETE)

body(String):HTTP的请求参数

headers(Object):HTTP的请求头，默认为{}

```
fetch('/abc', {
	method: 'get'
}).then(data => {
	return data.text()
}).then(ret => {
	//注意这里得到的才是最终的数据
	console.log(ret)
})
```

##### 2.GET请求方式的参数传递

传统方式:

```
fetch('/abc?id=123').then(data => {
	return data.text()
}).then(ret => {
	//注意这里得到的才是最终的数据
	console.log(ret)
})
```

Restful

```
fetch('/abc/123', {
	method: 'get'
}).then(data => {
	return data.text()
}).then(ret => {
	//注意这里得到的才是最终的数据
	console.log(ret)
})


```

##### 3.DELETE请求方式的参数传递

```
fetch('/abc/123', {
	method: 'delete'
}).then(data => {
	return data.text()
}).then(ret => {
	//注意这里得到的才是最终的数据
	console.log(ret)
})
```

##### 4.POST请求方式的参数传递

```
fetch('/books', {
	method: 'post',
	body: 'uname=lisi&pwd=123',
	headers: {
		'Content-Type': 'application/x-www-form-urlencoded',
	}
}).then(data => {
	return data.text()
}).then(ret => {
	console.log(ret)
})
```

```
fetch('/books', {
	method: 'post',
	body: JSON.stringify({
		uname: 'lisi',
		age: 12
	})
	headers: {
		'Content-Type': 'application/json',
	}
}).then(data => {
	return data.text()
}).then(ret => {
	console.log(ret)
})
```

##### 5.PUT请求方式的参数传递

```
fetch('/books123', {
	method: 'put',
	body: JSON.stringify({
		uname: 'lisi',
		age: '12'
	})
	headers: {
		'Content-Type': 'application/json',
	}
}).then(data => {
	return data.text()
}).then(ret => {
	console.log(ret)
})
```

### fetch响应结果

#### 响应数据格式

text():将返回体处理成字符串类型

json():返回结果和JSON.parse(responseText)一样

```
fetch('/abc').then(data => {
	//return data.text();
	return data.json()
}).then(ret => {
	console.log(ret)
});
```

## axios的基本特性

```
axios.get('/adata')
	.then(ret => {
		//data属性名称是固定的,用于获取后台响应的数据
		console.log(ret.data)
	})
```

### axios的常用API

get: 查询数据

post：添加数据

put：修改数据

delete：删除数据

#### 1.GET传递参数

通过URL传递参数

通过params选项传递参数

```
axios.get('/adata?id=123')
	.then(ret = {
		console.log(ret.data)
	})
```

```
axios.get('/adata/123')
​	.then(ret => {
	console.log(ret.data)
})
```



```
axios.get('/adata', {
	params: {
		id: 123
	}
})
.then(ret ={
	console.log(ret.data)
})
```

#### 2.DELETE传递参数

参数传递方式与GET类似

```
axios.delete('/adata?id=123').then(ret => {
	console.log(ret.data)
})
```

```
axios.delete('/adata/123').then(ret => {
	console.log(ret.data)
})
```

```
axios.delete('/adata', {
	params: {
		id: 123
	}
}).then(ret => {
	console.log(ret.data)
})
```

#### 3.POST传递参数

通过选项传递参数(默认传递的是json格式的数据)

```
axios.post('/adata', {
	uname: 'tom',
	pwd: 123
}).then(ret => {
	console.log(ret.data)
})
```

通过URLSearchParams传递参数(application/x-www-form-urlencoded)

```
const params = new URLSearcheParams()
params.append('params', 'value1')
params.append('params2', 'value2')
axios.post('/api/test', params).then(ret => {
	console.log(ret.data)
})
```

##### 4.PUT传递参数

参数传递方式与POST类似

```
axios.put('/adata/123', {
	uname: 'tom',
	pwd: 123
}).then(ret => {
	console.log(ret.data)
})
```

### axios的响应结果

data：实际响应回来的数据

headers：响应头信息

status：响应状态码

statusText：响应状态信息

```
axios.post('/axios-json').then(ret => {
	console.log(ret)
})
```

### axios的全局配置

axios.defaults.tiemout = 3000;	//超时时间

axios.defaults.baseURL = 'http://localhost:3000/app';	//默认地址

axios.defaults.headers['mytoken'] = '自定义内容 '	//设置请求头

### axios拦截器

#### 1.请求拦截器

```
//添加一个请求拦截器
axios.interceptors.request.use(function (confing) {
	//在请求发出之前进行一些信息设置
	return config;
}),function (err) {
	//处理响应的错误信息
}
```

#### 2.响应拦截器

```
//添加一个响应拦截器
axios.interceptors.response.use(function (res) {
	//在这里对返回的哦数据进行处理
	return res;
}), function (err) {
	//处理响应的错误信息
}
```

## async/await的基本用法

async/await是ES7引入的新语法，可以更加方便的进行异步操作

async关键字用于函数上(async函数的返回值是Promise实例对戏那个)

await关键字用于async函数当中(await可以得到异步的结果)

```
async function queryData(id) {
	const ret = await axios.get('/adata')
	return	ret.data
}
queryData.then(data => {
	console.log(data)
})
```

### async/await处理多个异步请求

多个异步请求的场景

```
async function queryData(id) {
	const info = await axios.get('/async1')
	const ret = await axios.get('async2?infi='+ info.data)
	return ret
}
queryData.then(ret => {
	console.log(ret)
})
```

