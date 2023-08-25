# VUE-Router

## 路由

### 1.路由的概念

路由的本质就是一种对应关系，比如说我们在url地址中输入我们要访问的url地址之后，浏览器要去请求这个url地址对应的资源。
那么url地址和真实的资源之间就有一种对应的关系，就是路由。

路由分为前端路由和后端路由
1).后端路由是由服务器端进行实现，并完成资源的分发
2).前端路由是依靠hash值(锚链接)的变化进行实现 

后端路由性能相对前端路由来说较低，所以，我们接下来主要学习的是前端路由
前端路由的基本概念：根据不同的事件来显示不同的页面内容，即事件与事件处理函数之间的对应关系
前端路由主要做的事情就是监听事件并分发执行事件处理函数

### 2.实现简易前端路由

基于URL中的hash实现(点击菜单的时候改变URL的hash，根据hash的变化控制组件的切换)

```
//监听window的onhashchange事件，根据获取到的最新的hash值，切换要显示的组件的名称window.onhashchenge = function () {
	//通过localhost.hash获取到最新的hash值
}
```

## 1.基本使用步骤

1.引入相关的库文件

2.添加路由链接

3.添加路由填充位

4.定义路由组件

5.配置路由规则并创建路由实例

6.把路由挂在到Vue根实例中

### 1.引入相关的路文件

```
<!-- 导入vue文件，为全局window对象挂在Vue构造函数 -->
<script src="./lib/vue_2.5.22.js"></script>

<!-- 导入Vue-router文件，为全局window对象挂在VueRouter构造函数 -->
<script	src="./lib/vue-router_3.0.2.js"></script>
```

### 2.添加路由连接

```
<!-- router-link是vue中提供的标签，默认会被渲染为a标签 -->
<!-- to属性默认会被渲染为href属性 -->
<!-- to属性的默认值会被渲染为#开头的hash地址 -->
<router-link to="/user">User</router-link>
<router-link to="/register">Register</router-link>
```

### 3.添加路由填充位

```
<!-- 路由填充位(也叫路由占位符) -->
<!-- 将来通过路由规则匹配到的组件，将会被渲染到router-view所在的位置 -->
<router-View></router-view>
```

### 4.定义路由组件

```
var User = {
	template: `<div>User</div>`
}
var Register = {
	template: `<div>Register</div>`
}
```

### 5.配置路由规则并创建路由实例

```
//创建路由实例对象
var touter = new VueRouter({
	//routes是路由规则数组
	routes: [
		//每个路由规则都是一个配置对象，其中至少包含path和component两个属相:
		//path表示当前路由规则匹配的hash地址
		//component表示当前路由规则对应要展示的组件
		{path: '/user',component: User},	//component只接收对象
		{path:'/register',component: Register}
	]
})
```

### 6.把路由挂在到Vue根实例中

```
new Vue({
	el: '#app',
	//为了能够让路由规则生效，必须把路由对象挂在到vue实例对象上
	//router: router
	router
})
```

## 2.路由重定向

路由重定向是指：用户在访问地址A的时候，强制用户跳转到c，从而展示特定的组件页面；

通过路由规则redirect属性，指定一个新的路由地址，可以很方便地设置路由的重定向

```
var touter = new VueRouter({
	routes: {
		//其中，path表示需要被重定向的原地址，redirect表示将要被重新定向到的新地址
		{path: '/', redirect: '/user'},
		{path: '/user', component: User},
		{path: '/register', component: Register}
	}
}) 
```

### 3.嵌套路由用法

#### 1.嵌套路由功能

点击父级路由链接显示模板内容

模板内容中又有子级路由链接

点击子级链接显示子级模板内容

#### 2.父路由组件模板

父级路由连接

父组件路由填充位

```
<p>
	<router-link to="/user">User</router-link>
	<router-link to="register">Register</router-link>
</p>
<div>
	<!-- 控制组件的现实位置 -->
	<router-view></router-view>
</div>
```

#### 3.子级路由模板

子级路由连接

子级路由填充为

```
const Register = {
	template: `<div>
		<h1>Register组件</h1>
		<hr/>
		<router-link to="/register/tab1">Tab1</router-link>
		<router-link to="/register/tab2">Tab2</touter-link>
		
		<!-- 子路由填充位置 -->
		<router-view/>
	`
}
```

#### 4.嵌套路由配置

父级路由通过children属性配置子级路由

```
const router = new VueRouter({
	routes: [
		{path: 'user', component: Usedr},
		{
			path: '/register',
			component: Register,
			//通过children属性，为/register添加子路由规则
			children: [
				{path: '/register/tab1', component: Tab1},
				{path: '/register/tab2', component: Tab2}
			]
		}
	]
})
```

### 4.动态路由

应用场景：通过动态路由参数的模式进行路由匹配

```
var router = new VueRouter({
	routes: [
		//动态路径参数 以冒号开头
		{path: '/user/:id', component: User}
	]
})
```

```
const User = {
	//路由组件中通过$router.params获取路由参数
	template: `<div>User {{$router.params.id}}</div>`
}
```

#### 路由组件传参数

$router与对应路由形成高度耦合，不够灵活，所以可以使用props将组件和路由解耦

##### 1.props的值为布尔类型

```
const router = new VueRouter({
	routes: [
		//如果props被设置为true，router.params将会被设置为组件属性
		{path: '/user/:id', component: User, props: true}
	]
})

const User = {
	props: ['id'], //使用props接收路由参数
	template: `<div>用户ID: {{id}}</div>`	//使用路由参数
}
```

##### 2.props的值为对象类型

```
const router = new VueRouter({
	routers: [
		//如果props是一个对象，它会被按照原样设置为组件属性
		{path: '/user/:id', component: User, props:{uname: 'lisi', age:12}}
	]
})
const User = {
	props: ['uname', 'age'],
	template: `<div>用户信息: {{uname + '---' + age}}</div>`
}
```

3.props的值为函数类型

```
const router = new VueRouter({
	routers: [
		//如果props是一个函数，则这个函数接收router对象为子级的形参
		{path: '/user/:id',
		component: User,
		props: route => ({uname: 'zs', age: 20, id: route.params.id})}
	]
})
const User = {
	props: ['uname', 'age', 'id'],
	template: `<div>用户信息: {{uname + '---' + age + '---' + id}}</div>`
}
```

## 3.命名路由的配置规则

为了更加方便的表示路由的路径，可以给路由规则起一个别名，即为"命名路由"。

```
const router = new VueRouter({
	routes: [
	{
		path: '/user/:id',
		name: 'user',
		comoponent: User
	}
	]
})
```

```
<router-link :to="{name: 'user',params: {id: 123}}">User</router-link>
router.push{{naem: 'user', params: {id:123}}}
```

## 4.编程式导航

1.声明式导航：通过点击链接实现导航的方式，叫做声明式导航

例如：普通网页中的<a></a>链接或vue中的<router-link></router-link>

2.编程式导航：通过调用Javasciript形式的API实现导航的方式，叫做编程式导航

例如：普通网页中的localtion.href

### 1.编程式导航基本用法

常用的编程式导航API如下：

this.$router.push('hash地址')

this.$router.go(n)

```
const User = {
	template: `<div><button @click="goRegister">跳转到注册页面</button></div>`,
	methods: {
		goRegister: function () {
			//用编程的方式控制路由跳转
		this.$router.push('/register')
		}
	}
}
```

### 2.编程式导航参数规则

router.push()方法的参数规则

```
//字符串(路径名称)
router.push('/home')
//对象
router.push({path: '/home'})
//命名的路由(传递参数)
router.push({name: '/user', params: {userId: 123}})
//带查询参数，编程/register?uname=lisi
router.push({path: '/register', query: {uname: 'lisi'}})
```

