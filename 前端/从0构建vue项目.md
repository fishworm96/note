---
title: 从0构建vue项目
date: 2021-06-02 17:26:13
tags: vue项目
---

项目目录：

<!-- more -->

————node_modules:
————public:放置公共文件
——————favicon.ico:标签栏看到的小图标
——————index.html:模板文件，webpack打包的时候会根据这个文件生成index.html文件
————src:项目主文件
——————assets：存放静态资源文件，图片、图标字体。
——————componets：组件，抽离出可复用的文件
——————views：视图，也就是页面
————————App.vue:基础组件
——————main.js:项目入口文件，开发运行和编译都会从这个文件作为起始点进入编译
——————router.js:路由文件
——————store.js：状态管理文件
————.eslintrc.js:配置eslint规则的文件
————.gitignore:git提交的忽略文件
————.postcssrc.js:css自动补充兼容性代码的配置
————babel.config.js:babel配置文件
————package-lock.json:
————package.json:项目描述
————vue.config.js:vue配置文件

```
vue.config.js目录下代码表示取消`eslint检测`
    module.exports = {
      lintOnSave: false
    }
```

```
//使用vscode可以添加一个`.editorconfig`文件，可以用来配置编辑器的使用习惯
//启用配置文件
root = true
//所有文件都有效
[*]
//utf-8编码
charset = utf-8
//使用tabs缩进，可以切换成space缩进
indent_style = tabs
//缩进2个字符
indent_size = 2
//还需要安装EditorConfig for VS Code插件
```

添加目录
————src/api:接口，项目json请求可以写在这里面，做一个统一的管理
————src/assets/img：放置图片
————src/assets/font：放置字体
————src/config/index.js：项目配置，使用`export default{}`导出配置对像。引入方式`import config from './config/index.js'`可以简写去掉`js`或去掉`index.js`。
————src/directive/index.js:用来放置`vue`自定义指令
————src/lib/util.js:与业务结合的工具方法
————src/lib/tools.js：纯粹工具方法
————src/router/router.js:抽离路由文件,只做一些路由的列表配置，放置路由模块。
————src/router/index.js:抽离路由文件
————src/store:vuex文件夹
————src/store/mutations.js:
————src/store/store.js
————src/store/actions.js
————src/store/module/user.js:存放用户名，用户信息
————src/mock/index.js:请求模拟;`import Mock from 'mockjs`，要安装`npm install mockjs -D`，最后把`mock`导出去`export default Mock`

```
//router/index配置
import Vue from 'vue'
import VueRouter from 'vue-router'
import routes from './router'

Vue.use(VueRouter)

export default new VueRouter({
  routes
})

//router/router配置
import Home from '../views/Home.vue'

export default [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]
```

```
//store/index配置
import Vue from 'vue'
import Vuex from 'vuex'
import state from './state'
import mutations from './mutations'
import actions from './actions'
import user from './module/user'

Vue.use(Vuex)

export default new Vuex.Store({
  state,
  mutations,
  actions,
  modules: {
    user
  }
})

//router/module/user.js配置
const state = {

}
const mutations = {

}
const actions = {

}

export default {
  state,
  mutations,
  actions
}
```

```
//vue.config.js配置
const path = require('path')

const resolve = dir => path.join(__dirname, dir)

//判断开发环境还是生产环境
// '/'一个斜线意思是指定在域名的根目录下
//如果是开发环境只用写一个斜杠
const BASE_URL = process.env.NODE_ENV === 'procution' ? '/iview-admin' : '/'

module.exports = {
  lintOnSave: false,
  //项目基本路径
  //baseUrl: Base_URL,
  chainWebpack: config => {
    config.resolve.alias
    // 代表当前路径拼接上src
    //@代表src路径，比如要使用api直接@/api
    .set('@', resolve('src'))
    //_c即代表src/components路径
    .set('_c', resolve('src/components'))
  },
  //打包时不生成.map文件，减少打包体积，加快打包速度
  productionSourceMap: false,
  // 配置跨域，告诉开发服务器将任何未知请求，就是没有匹配到静态文件的请求都代理到这个url来满足跨域的请求
  devServer: {
    proxy: 'http://localhost:4000'
  }
}

```

## 路由配置

内容：
  1.router-link和router-view组件
  2.路由配置
  —1.动态路由
  —2.嵌套路由
  —3.命名路由
  —4.命名视图
  3.JS操作路由
  4.重定向和别名
  在`router.js`中配置路由列表，在`index.js`中创建路由实例。

### 组件

```
<!-- router-link是vue中提供的标签，默认会被渲染为a标签 -->
<!-- to属性默认会被渲染为href属性 -->
<!-- to属性的默认值会被渲染为#开头的hash地址 -->
<router-link to="/user">User</router-link>
<router-link to="/register">Register</router-link>
```

```
<!-- 路由填充位(也叫路由占位符) -->
<!-- 将来通过路由规则匹配到的组件，将会被渲染到router-view所在的位置 -->
<router-View></router-view>
```

### 路由配置

  path: 即要拼接的路径
  component: 即组件，访问的路径的文件。
  component: () => import(/* webpackChunkName: "about" */ '../views/About.vue') 懒加载，只有访问到这个页面的时候才会加载。`webpackChunkName: "about"`代表打包后备注的名称

#### 动态路由匹配

```
    // name为动态路由参数
    path: '/argu/:name',
    component: () => import('@/views/argu.vue')
```

```
//在argu文件中的内容
<template>
  <div>
    {{ $route.params.name}}
  </div>
</template>
```

`$route`代表当前路由对象`params`参数`name`值这里的`name`就是路由中的`:name`,当改变参数的时候，不管`:name`的参数是什么他匹配到的都是这个路由对象。同一个页面处理不同的逻辑。

#### 嵌套路由

```
path: '/parent',
    component: () => import('@/views/parent.vue'),
    children: [
      {
        //子嵌套路由不用写斜线
        path: 'child',
        component: () => import('@/views/child.vue')
      }
    ]
```

#### 命名路由

```
//就是给路由对象设置一个name属性
    path: '/',
    name: 'Home',
    component: Home
```

```
//在App.vue文件下
//to传入一个对象，就要使用动态绑定，后面写一个对象。
      <router-link :to="{{ naem: 'home' }}">Home</router-link> 
      <router-link :to="{{ name: 'about' }}">About</router-link>
```

#### 命名视图

```
//怎么在同一个页面上显示多个视图，而且每一个视图显示在指定的位置
    <router-view name="email"/>
    <router-view name="tel"/>
//router/router.js文件下
    path: '/named_view',
    //这里要多加个s，代表多个组件
    components: {
      //如果没有命名router-view就加载这个default组件
      default: () => import('@/views/child.vue'),
      email: () => import('@/views/email.vue'),
      tel: () => import('@/views/tel.vue')
    }
```

### 重定向

```
    path: '/main',
    //当我们访问/路径时就重定向到main路径
    /* redirect: '/'
    //也可以使用命名路由
    redirect: {
      name: 'home'
    } */
    //还可以使用一个方法
    redirect: to => {
      //这里必须return一个对象或字符串路径
      return {
        name: 'home'
      }
      //还可以直接简写成 to => '/'
    }
```

### 别名

```
    path: '/',
    //当访问的是首页
    alias: '/home_page',
    name: 'Home',
    component: Home
```

### 编程式导航

```
//Home.vue文件下
<button @click="handleClick">返回上一页</button>

  methods: {
    handleClick(type) {
      // this.$router.go(-1)
      this.$router.back()
    }
  }
```

```
<button @click="handleClick('back')">返回上一页</button>
<button @click="handleClick('push')">返回上一页</button>
<button @click="handleClick('parent')">替换到parent</button>

methods: {
    handleClick(type) {
      if(type === 'back') this.$router.back()
      //else if (type === 'push') this.$router('/parent')
      //这种方法也是可以的
      else if (type === 'push') {
        this.$router({
        name: 'parent',
        //此时跳转到parent页面会添加一个参数lison
         query: {
            name: 'lison'
          }
      })else if (type === 'replace') {
        // replace替换，把当前的浏览历史替换成parent这个页面，之后再做回退会回到到parent
        this.$router.replace({
          name: 'parent'
        })
      }
      }
    }
  }
  if(type === 'push') {
        this.$router.push({
          //使用es6方法简写
          //const name: 'lison
          //name: '/argu/${name}'
          name: 'argu',
          params: {
            name: 'lison'
          }
          }
        )
      }

      //错误方法
      this.$router.push({
        //想要这样写可以改成name: 'argu'
        path: '/argu',
        params: {
          name: 'lison'
        }
      })
```

## 路由进阶

  1.路由组件传参
  2.HTML5 History模式
  3.导航守卫
  4.路由元信息
  5.过度效果

### 路由组件

  路由组件传参有三种形式。
  第一种是布尔模式，适用于动态路由匹配中，有动态路由参数的动态配置中。比如之前`argu`这个动态路由，里面有动态参数`:name`

```
views/argu.vue页面下
  <div>
    <!-- {{ $route.params.name}} -->
    <!-- 这样这里可以只写name -->
    {{name}}
  </div>

//作为属性传入
  props: {
    //这个属性是name
    name: {
      //可以设定类型，如果只想要是名字就只设定String类型
      type: String,
      //没有设定值，默认值。
      default: 'lison'
    }
  }

router/router.js页面下argu路由中
    //传入一个对象，页面上显示的就是这个值。
    props: {
      food: 'banana'
    }
```

  第二种是对象模式，在普通的页面，不是动态匹配的页面。比如在`About`页面。在`About`页面定义属性。

```
views/About.vue页面下
  <div class="about">
    <b>{{food}}</b>
  </div>

  props: {
    food: {
      type: String,
      // 路由没有传值页面上显示的就是这个值
      default: 'apple'
    }
  }

router/router.js页面下About路由中
    //传入一个对象，页面上显示的就是这个值。
    props: {
      food: 'banana'
    }
```

  第三种是函数模式，它适合于在传入的属性中能够根据当前的路由来做一些处理逻辑，从而设置一些我们设置的属性值。比如在`Home`页

```
views/Home.vue页面下
  <div class="home">
    <!-- 如果没有传入参数，那么就是默认的apple -->
    <b>{{ food }}</b>
    <router-view></router-view>
  </div>

  props: {
    food: {
      type: String,
      default: apple
    }
  }

router/router.js页面下Home路由中
// 代表当前这个参数就是路由对象，如果想返回一个对象使用一个括号抱住，或者直接使用return {}
    props: route => ({
      //想根据query里的food来传入这个属性
      food: route.query.food
    })
```

### HTML5 History模式

```
router/index.js页面下
export default new VueRouter({
  //mode默认值为hash。在路径后面携带一个井号用来完成无刷新的模拟路由跳转。
  //不想看到井号，使用history模式，它是使用history的API来完成无刷新的页面跳转，需要后端支持。没有匹配到静态资源的时候都会返回一个index.html文件，如果使用history模式，会产生一个问题匹配不到静态资源，而且路由匹配不到组件的话就会有问题了。
  mode: 'hash',
  routes
})

router/router.js页面下
  // 所以需要在这里添加一个配置，一定要放置在最后。因为他是从上自下执行的，如果放到上面会影响其他路由。
  {
    path: '*',
    component: () => import('@/views/error_404.vue')
  }
```

### 导航守卫

导航守卫功能，用来判断用户有没有登录，没有登录的话跳转到登录页面，如果已经登录就跳转到允许登录的页面。还有在做权限控制的时候，如果这个页面没有权限浏览，就做一些相应的处理。

#### 全局守卫：在全局设置一个守卫。

```
//router/index.js页下
const router = new VueRouter({
  routes
})

// 模拟一个登录，实际是通过接口来判断的。
const HAS_LOGINED = true

router.beforeEach((to, from, next) => {
  //to代表即将跳转的路由对象
  //form代表从哪里跳转过来
  //next代表放行
  if(to.name !== 'login') {
    // 如果已经登录就放行
    if(HAS_LOGINED) next()
    //如果没有登录就跳转到登录页面
    else next({ name: 'login' })
  } else {
    // 如果跳转到的是登录页面，并且已经登录过了就跳转到首页
    if(HAS_LOGINED) next({ name: 'home'})
    //没有登录就放行，跳转套登录页面
    else next()
  }
})

export default routes
```

#### 后置钩子：

```
//router/router.js
//导航钩子不能阻止之后跳转的页面进行操作，只能执行一些简单的逻辑。
router.afterEach((to, from) => {
  //在路由跳转之后做一些操作
  //logining = false
})

//全局守卫，它是在导航被确认之前和所有组件导航守卫，异步路由组件解析之后被调用。
//确认的意思是所有的导航钩子都结束。
// router.beforeResolve((to, from, next) => {})
```

#### 路由独享守卫

```
//router/router.js页下
//一定要添加next()方法，不然不会跳转
beforeEnter: (to, from, next) => {
      next()
    }
```

#### 组件内的守卫

```
  beforeRouterEnter(to, from, next) {
    //跳转的页面此时this还没有加载出来，是不能用this的
    next(vm => {
      //这个vm就是组件的实例，这样就能在里面使用this了
    })
  },
  beforeRouterLeave(to, from, next) {
    //这个时候组件已经被渲染了，可以使用this
    //将要离开页面时调用钩子方法
    const leave = confirm('您确定要离开吗？')
    if(leave) next()
    else next(false)
  },
    // 路由发生变化，组件被复用的时候调用
  beforeRouterUpdate(to, from, next) {
    //这个时候组件已经被渲染了，可以使用this
  }
}
```

### 元信息

  存放一些我们要定义的信息，比如说页面是否需要权限，需要权限可以在路由前置守卫里可以做一些处理。

```
//router/router.js页面下，需要添加的路由中。
    meta: {
      //想让每个跳转的页面title都不一样
      title: '关于'
    }

router/index.js页面下
//结构赋值导入
imort { setTitle} from '@/lib/util.js'
//全局前置路由里添加判断，是否有meta属性，如果没有就不执行。因为有些路由里没有添加meta属性就会出错。
  to.meta && setTitle(to.meta.title)

lib/util.js页面下定义方法，如果有就添加修改的title，没有就是用默认的admin。并暴露出去
export const setTitle = (title) => {
  window.document.title = title || 'admin'
}
```

### 过度效果

添加过度的样式

```
//App.vue页面下
//如果只要包住router-view 可以只写transition组件就可以了，需要给每个组件添加一个key值给transition-group添加一个name
    <transition-group name="router">
      <router-view key="default" />
      <router-view key="email" name="email" />
      <router-view key="tel" name="tel" />
    </transition-group>

// 这里是name的值。这里是页面进入的效果
.router-enter{
  opacity: 0;
}
.router-enter-active {
  transition: opacity 1s ease;
}
.router-enter-to {
  opacity: 1;
}
//页面注销离开的效果
.router-leave{
  opacity: 1;
}
.router-leave-active {
  transition: opacity 1s ease;
}
.router-leave-to {
  opacity: 0;
}

//另一种方法
需要把transition的name名称改为routerTransition
  data () {
    return {
      routerTransition: ''
    }
  },
  watch: {
    //to代表路由对象
    '$router' (to) {
      // 在to的query里面找有没有transiName，首先要判断有没有query这个字段。都满足了才改名字。
      //想为某个页面展示特定特效可以用这个方式
      to.query && to.query.transitionName && (this.routerTransition = to.query.transitionName)
    }
  }
}
当页面访问到特定的路由会添加特效
```

## 状态管理

  1.bus
  2.Vuex-基础-state&getter
  3.Vuex-基础-mutation&action/module
  4.Vuex-进阶

### 父子、兄弟组件

```
父子组件和兄弟组件通信
在views下创建一个store.vue。父组件
<template>
  <div>
    <a-input @input="handleInput"></a-input>
    <p>{{ inputValue }}</p>
    <a-show :content="inputValue" ></a-show>
  </div>
</template>

<script>
import AInput from '_c/AInput.vue'
import AShow from '_c/AShow.vue'
export default {
  name: 'store',
  components: {
    AInput,
    AShow
  },
  data() {
    return {
      inputValue: ''
    }
  },
  methods: {
    handleInput(val) {
      this.inputValue = val
    }
  }
}
</script>

//components下创建一个AInput.vue。子组件
<template>
  <input @input="handleInput" :value="value">
</template>

<script>
export default {
  name: 'AInput',
  props: {
    value: {
      type: [String, Number],
      default: ''
    }
  },
  methods: {
    handleInput(event) {
      const value = event.target.value
      this.$emit('input', value)
    }
  },
}
</script>

//components下创建AShow.vue。兄弟组件
<template>
  <div>
    <p>AShow: {{ content }}</p>
  </div>
</template>

<script>
export default {
  props: {
    content: {
      type: [String, Number],
      default: ''
    }
  }
}
</script>

//路由下添加
  {
    path: '/store',
    component: () =>  import('@/views/store.vue')
  },
```

### bus

在App.vue里使用了命名视图，name为email和name为tel这两组件之间的简单场景下通信就是用到了bus。
bus就是创建一个空的实例，作为交互的中介。

```
//在lib文件夹下创建bus.js
import Vue from 'vue'
<!-- 创建vue实例 -->
const Bus = new Vue()
<!-- 暴露出去 -->
export default Bus

//在main.js中导入bus，import Bus from './lib/bus.js'。
// bus注册到根实例
Vue.prototype.$bus = Bus

//views/email.vue页面下
<template>
  <div>
    <!-- 使用方法 -->
    <button @click="handleClick" >按钮</button>
  </div>
</template>
<script>
export default {
  methods() {
    // 把事件提交给bus的方法
    this.$bus.$emit('on-click', 'hello')
    // $emit是触发当前实例上的一些事件
  }
}
</script>

//views/tel.vue页面下
<template>
  <div>
    <!-- 接收事件 -->
    <p>{{message}}</p>
  </div>
</template>

<script>
export default {
  data () {
    return {
      message: ''
    }
  },
  mounted () {
    this.$bus.$on('on-click', mes => {
      // $on是给当前事件绑定一个事件监听
      this.message = mes
    })
  }
}
</script>
```

### state和getter

#### state的使用

```
//store/index.js首先要在实例中引入。
import Vue from 'vue'
import Vuex from 'vuex'
import state from './state'
import mutations from './mutations'
import actions from './actions'
import user from './module/user'

Vue.use(Vuex)

// 一定要Vuex.store方法创建实例
export default new Vuex.Store({
  state,
  mutations,
  actions,
  modules: {
    user
  }
})
```

```
//store/state.js页面下
const state = {
  // 这里定义的值可以在各个组件中使用
  appName: 'admin'
}

export default state

在//views/store.vue页面下就可以直接使用
<template>
  <div>
   {{ appName }}
  </div>
</template>
appName () {
  return this.$store.state.appName
}
```

```
//store/module/user.js模块下使用方法，先在模块中定义好
const state = {
// 在模块中使用
  userName: 'Lison'
}
export default {
  state,
  mutations,
  actions
}
//views/state.vue页面下
<template>
  <div>
    {{userName}}
  </div>
</template>

userName () {
    // 这里要写模块名
  return this.$store.state.user.userName
}
```

最后就是命名空间方法

```
//store/module/user.js页面下，这里加入了namespaced: true启用了命名空间
export default {
  // 使用命名空间
  namespaced: true,
  state,
  mutations,
  actions
}
//views/store.vue页面下
<template>
  <div>
    {{ appName }}
    {{userName}}
  </div>
</template>

<script>
// 对象解构方法引入
// import {mapState} from 'vuex'
// 命名空间方法引入
import {createNamespacedHelpers} from 'vuex'
const {mapState} = createNamespacedHelpers('user')
export default {
  computed: {
    // 命名空间方法
    ...mapState({
      userName: state => state.user.userName
    })
    //使用对象解构方法引入也可以用这种方法，第一个参数为模块名。
    ...mapState('user',{
      userName: state => state.userName
    })

    // ...mapState({
    //   // 对象方法
    //   appName: state => state.appName,
    //   // 模块要加模块名
    //   userName: state => state.user.userName
    // })

    // 数组方法
    // ...mapState(['appName'])
}
</script>
```

#### getter使用

getter相当于vue组件里的一个计算属性，getter的方法和state相似。

```
//router/geter页面下
const getters = {
  // 现在计算一个值，这个值是根据state中的appname来计算的
  // 这里的state就是当前vue实例里同级的state
  appNameWithVersion: (state) => {
    return `${state.appName}v2.0`
  }
}
export default getters

//routoer/index.js页面下引入
import Vue from 'vue'
import Vuex from 'vuex'
import state from './state'
import getters from './getter'

Vue.use(Vuex)

// 一定要Vuex.store方法创建实例
export default new Vuex.Store({
  getters,
  modules: {
    user
  }
})

//views/store页面下
<template>
  <div>
    <p>{{ inputValue }}->{{inputValueLastLetter}}</p>
    <p>{{appNameWithVersion}}</p>
  </div>
</template>

<script>
const {mapGetters} = createNamespacedHelpers('user')

  computed: {
    ...mapGetters([
      'appNameWithVersion'
    ]),
    inputValueLastLetter() {
    //始终返回最后1个字符串
      return this.inputValue.substr(-1, 1)
    },
    appNameWithVersion() {
      return this.$store.getters.appNameWithVersion
    }
  }
</script>

```

### mutation&action/module

#### mutations

在组件中是不能直接通过直接赋值的方式来修改App.vue，要通过mutation来修改App.vue。
比如要修改store中appName

```
//store/mutations.js页面下定义
// 引入vue
import vue from 'vue'

const mutations = {
  // 这个方法有两个参数，第一个参数state，state是指同级的state对象。第二个参数params在Vue文档中叫载荷，这个参数有两种形式，如果只需要一个值那么它就是一个值直接使用。如果是多个值，那么就是一个对象。
  SET_APP_NAME (state, params) {
    // state.appName = params
    state.appName = params.appName
  },
  //现在没有版本号，但以后想给它添加一个应该这么做
  // 先定义一个函数，这里要传一个参数。
  SET_APP_VERSION(state) {
    // 使用vue的set方法，第一个参数是要给谁设置值。第二个参数是要设置的名字。第三个参数是要设置的值。
    // 如果不是用vue.set方法将不会添加，这就是vue响应式的原则。
    vue.set(state, 'appVersion', '1.0')
  }
}

export default mutations

//views/store.vue页面下
<button @click="handleChangeAppName">修改appName</button>

appName: {
  // 调用appName的时候使用getter方法，设置一个值的时候使用setter方法
  set: function (newValue) {
    this.inputValue = newValue + 'sd'
  },
  get: function () {
    return this.inputValue + 'asdf'
  }
},
```

如果想修改appName

```
//views/store
<template>
  <div>
    <button @click="handleChangeAppName">修改appName</button>
    {{appVersion}}
  </div>
</template>

import { mapState} from 'vuex'

computed: {
  // 命名空间方法
  ...mapState({
    appVersion: state => state.appVersion
  })
}

methods: {
    handleChangeAppName () {
      // 要修改appName不能直接使用复制的方式
      // this.appName = 'newAppName'
      // 要修改的话就要commit方法，这里第一个参数就是要提交的名称，第二个参数是要赋的新的名称
      //修改单个参数的方法
      // this.$store.commit('SET_APP_NAME', 'newAppName')
      //修改对象的方法
      // this.$store.commit('SET_APP_NAME', {
      //   appName: 'newAppName'
      // })
      // 直接一个参数，对象的写法
      this.$store.commit({
        // 这个type就是要提交方法的名称
        type: 'SET_APP_NAME',
        appName: 'newAppName'
      })
        // 这里再使用commit方法
        this.$store.commit('SET_APP_VERSION')
    }
  }
}
```

mapMutations工具方法

```
//views/sotre页面下
import { mapMutations} from 'vuex'

  methods: {
    ...mapMutations([
      'SET_APP_NAME'
    ]),
    handleInput (val) {
      this.inputValue = val
    },
    handleChangeAppName () {
      //这里如果只有一个值就直接传进来
      //this.SET_APP_NAME('newAppName')
      //也可以用对象的方式
      this.SET_APP_NAME({
          appName: 'newAppName'
        })
    }

  }
```

在user模块中定义的mutation。

```
//sotre/module/user.js页面下
const state = {
// 在模块中使用
  userName: 'Lison'
}
const mutations = {
  //这里修改userName
  // 第一个参数就是要作用的名称，第二个参数是传过来的值
  SET_USER_NAME (state, params) {
    state.userName = params
  }
}

//views/store页面下
<template>
  <div>
    <p>Username: {{userName}}</p>
    <button @click="changeUerName">修改用户名</button>
  </div>
</template>

  methods: {
    //这里vuex把getters、mutations、actions统统注册在全局中，如果想单独使用可以使用命名空间。
    ...mapMutations([
      'SET_USER_NAME'
    ]),
    changeUerName() {
      this.SET_USER_NAME('vue-cource')
    }
  }
}
```

#### actions 

比如现在需要从接口中调用数据来完成异步操作来修改state中的appName

```
//sotre/cations页面下
// 假如是一个api接口请求
import { getAppName } from '@/api/app.js'

const actions = {
  //这是第一种方法
  // 第一个参数是一个对象，这是一个方法，调用它去提交mutation。这里用的是es6的对象结构，相当于是传入paramsObj在函数体内使用const commit = paramsObj.commit
  // updateAppName ({ commit }) {
  //   // 模拟异步操作
  //   getAppName().then(res => {
  //     // 提交方法名
  //     // 基本用法
  //     // commit('SET_APP_NAME', res.info.appName)
  //     //使用es6结构方法
  //     const { info: { appName } } = res
  //     commit('SET_APP_NAME', res.info.appName)
  //   }).catch((res) => {
  //     // 处理错误
  //     console.log(err)
  //   })
  // }

  //第二种方法
  //使用es8的异步函数来处理问题
  async updateAppName ({ commit }) {
    try {
      const { info: { appName } } = await getAppName()
      commit('SET_APP_NAME', appName)
    } catch (err) {
      // 处理异常信息
      console.log(err)
    }
  }
}
export default actions

//views/store页面下ia
<template>
  <div>
    {{ appName }}
    <button @click="handleChangeAppName">修改appName</button>
    <button @click="changeUerName">修改用户名</button>
  </div>
</template>

import { mapActions} from 'vuex'

  methods: {
    ...mapActions([
      'updateAppName'
    ]),
     handleChangeAppName () {
       //直接调用方法
        this.updateAppName()
    },
    changeUerName() {
      //使用实例的方法
      // this.$store.dispatch('updateAppName', '123')
    }
  }
}
```

#### module

当项目非常庞大的时候store会非常的臃肿，把它拆成各个模块会变得比较清晰，模块里面还可以嵌套模块。

```
//store/module/user.js页面下
const actions = {
  // 这里的第一个参数是提交，第二个参数是这里的state实例，第三个参数是store下面的state实例，可以直接操作。第四个参数是action实例的提交方法。
  updateUserName ({ commit, state, rootState, despatch}) {
    // 操作state下的appName
    rootState.appName
    // 这里可以调用dispatch来触发aaa这个action。
    dispatch('aaa', '')
  },
  aaa() {
  }
}

```

使用store实例，动态注册模块

```
//vews/store页面下
<template>
  <div>
    <button @click="registerModule">动态注册模块</button>
    <p v-for="(li, index) in todoList" :key="index">{{li}}</p>
  </div>
</template>

computed: {
  // 命名空间方法
  ...mapState({
    // 使用方法，再进行判断，是否存在todoList。
    todoList: state => state.todo ? state.todo.todoList : []
  })
},
methods: {
      registerModule() {
       第一个参数是要动态添加的模块名称todo，
       this.$store.registerModule('todo', {
         state: {
           todoList: [
             '学习mutations',
             '学习actions'
           ]
         }
       })
}
```

在模块下嵌套子模块

```
//vews/store页面下
<template>
  <div>
    <button @click="registerModule">动态注册模块</button>
    <p v-for="(li, index) in todoList" :key="index">{{li}}</p>
  </div>
</template>

computed: {
  // 命名空间方法
  ...mapState({
    // 使用方法，再进行判断，是否存在todoList。
      todoList: state => state.user.todo ? state.user.todo.todoList : []
  })
},
methods: {
      // 给user模块添加一个子模块
      this.$store.registerModule(['user', 'todo'], {
        state: {
          todoList: [
            '学习mutations',
            '学习actions'
          ]
        }
      })
    }
```

## Vuex进阶

  1.插件
  2.严格模式
  3.vuex+双向绑定

### 插件

现在往/store下新建plugin文件夹并创建saveInLocal.js文件。持久化存储插件，因为store是存在内存中的，不是存在本地的。用户刷新后便会消失，有时候需要将一些东西存在本地，这样用户刷新后不会丢失。

```
//store/index.js页面下
import saveInLocal from './plugin/saveInLocal.js'
  // 持久化插件
 export default new Vuex.Store({
  plugins: [ saveInLocal ]
})

//store/plugin/saveInLocal.js页面下
//在每次实例初始化的时候调用
export default  store => {
  // 每次浏览器刷新后第一次做的操作
  // 每次提交mutaions后就会提交回调函数，第一个参数mutation，每次提交的一些信息。第二个参数state，就是当前模块的state
  // 添加实例： 先判断是否存在，存在就替换掉。这里不能直接赋值要使用store里的replaceState方法再转为对象替换掉。
  if(localStorage.state) store.replaceState(JSON.parse(localStorage.state))
  store.subscribe((mutation, state) => {
    // 每次提交都保存在本地，state就是当前模块的state，是一个对象。存储了以后要添加到state实例里。
    localStorage.state = JSON.stringify(state)
  })
}
```

### 严格模式

store里的state必须要提交mutation来修改它的值，不能直接在组件里通过赋值的方式来修改它。严格模式就是在创建store这个文件里设置一个`strict: true`。如果现在使用赋值就会报错。如果不设置或false就不会报错。这里的模式应该不要写死，而是通过判断来决定。`process.env.NODE_ENV === 'development'`如果是开发环境就会报错，生产环境不会。

### Vuex双向绑定

如果当前的值不是在`data`里面定义的，而是在全局里面定义的。这种情况使用v-model就会有问题。
有两种方法解决这个问题
第一种

```
//store/state.js页面下
const state = {
  //先定义好stateValue
  stateValue: 'asd'
}

//store/mutation.js页面下
const mutations = {
  SET_STATE_VALUE (state, value) {
    state.stateValue = value
  }
}

//views/store页面下
<template>
  <div>
  //这里要使用:value来绑定stateValue，通过input来触发
    <a-input :value="stateValue"
             @input="handleStateValue"></a-input>
    <p>{{ stateValue }}</p>
  </div>
</template>

computed: {
  ...mapState({
      stateValue: state => state.stateValue
  })
}
methods: {
    ...mapMutations([
      'SET_STATE_VALUE'
    ]),
    handleStateValue (VAL) {
    this.SET_STATE_VALUE(VAL)
  }
}

```

第二种使用get和set方法。

```
//store/state.js页面下
const state = {
  //先定义好stateValue
  stateValue: 'asd'
}

//store/mutation.js页面下
const mutations = {
  SET_STATE_VALUE (state, value) {
    state.stateValue = value
  }
}

//views/store页面下
<template>
  <div>
  //这里还是使用v-model
    <a-input v-model="stateValue"></a-input>
    <p>{{ stateValue }}</p>
  </div>
</template>

computed: {
      stateValue: {
      get() {
        // 读取stateValue的时候就会调用这个方法，调用的stateValue就是return的stateValue
        return this.$store.state.stateValue
      },
      set(value) {
        // 调用stateValue时候就会改变set方法，set有一个参数就是新赋的值是什么。这里用commit通过mutation来修改。
        this.SET_STATE_VALUE(value)
      }
    }
}
methods: {
    ...mapMutations([
      'SET_STATE_VALUE'
  ]),
    handleStateValue (VAL) {
    this.SET_STATE_VALUE(VAL)
  }
}
```

## Ajax请求

  1.解决跨域问题
  2.封装axios
  ——1.请求拦截
  ——2.响应拦截
  3.请求实战

### 跨域

这里配置好了代理，能直接发送异步请求。
如果没有设置代理，而是填写了域名的完整路径。需要在服务器端配置

```
app.all('*', (req, res, next) => {
  res.header('Access-Control-Allow-Origin', '*')
  res.header('Access-Control-Allow-Header', 'X-Requested-With,Content-Type')
  res.header('Access-Control-Allwo-Methods', 'PUT, POST, GET, DELETE, OPTIONS')
})
```

### 封装axios

在lib文件夹下创建axios.js

```
import axios from 'axios'
import { baseURL } from '@/config.js'
// 用类的形式封装
class HttpRequest {
  // constructor方法是每个类必须的方法，如果没有添加它会默认添加这个方法。
  // 这里可以接收一些参数，如options参数对象。创建new HttpRequest()实例时可以传入一个参数
  constructor(baseUrl = baseURL) {
    // 这里es6写法传入一个baseUrl，这个写法相当于的baseUrl = baseUrl || baseURL
    // 这里this指代要创建的实例
    this.baseUrl = baseUrl
    this.queue = {}
  }
  // 在内部设置一个全局的配置，使用方法的形式
  getInsideConfig() {
    const config = {
      baseURL: this.baseUrl,
      header: {

      }
    }
    return config
  }
  // 请求拦截器
  interceptors(instance) {
    instance.interceptors.request.use(config => {
      //添加全局的loading
      //Spin组件，添加遮罩层，覆盖遮罩层后就无法点击。Spin.show()
      return config
    }, error => {
      return Promise.reject(error)
    })
    // 响应拦截器
    instance.interceptors.response.use(res => {
      console.log(res)
      return res
    }, error => {
      return Promise.reject(error)
    })
  }
  request(options) {
    // 利用axios.create方法创建一个实例
    const instance = axios.create()
    // options需要一个合并，利用es6的assign方法进行合并。这个方法会把这两个对象合并成一个对象。如果有相同的key值的话，后面对象的key值会覆盖前面对象key的值。所以创建请求的配置要放在后面。
    options = Object.assign(this.getInsideConfig(), options)
    // 传入实例
    this.interceptors(instance)
    // 返回实例
    return instance(options)
  }
}
export default HttpRequest
```

在config页面下创建index.js文件

```
//如果这里是生产环境，判断里就写一个生产环境的实际接口。如果是本地开发环境就写一个空的字符串，在配置代理里填写域名。如果没有在vue.config.js里写代理就需要在这里配置端口。
export const baseURL = process.env.NODE_ENV === 'PRODUCTION' 
  ? 'http://production.com'
  //这里的地址就是axios里的baseURL
  : 'http://localhost:3000'
```

在api下创建一个index.js

```
// 所有请求都会放到这里管理
import HttpRequest from '@/lib/axios'
const axios = new HttpRequest()
export default axios
```

在api下创建一个user.js

```
import axios from './index.js'

export const getUserInfo = ({ userId }) => {
  return axios.request({
    url: '/getUserInfo',
    method: 'post',
    data: {
      userId
    }
  })
}
```

## Mock模拟Ajax响应

  1.响应模拟
  2.Mock用法

### 响应模拟

在mock下创建response文件夹存放请求，在response下创建user.js

```
模拟请求并暴露
export const getUserInfo = (options) => {
  return {
    namr: 'lison'
  }
}
```

在mock下处理

```
//mock/index.js文件下
import Mock from 'mockjs'
//引入请求
import { getUserInfo } from './response/user.js' 

//第一个参数可以是字符串或者正则表达式。第二个参数是请求方式写post或者get等等，这里可以省略。第三个参数是一个模板或者方法。
// Mock.mock('http://localhost:3000/getUserInfo', getUserInfo)
//使用正则的方式
Mock.mock(/\/getUserInfo/, getUserInfo)

export default {
  Mock
}
```

## 从数字渐变组件到第三方JS库的使用

  1.组件封装基础
  2.组件中使用id值
  3.组件中获取DOM

需要安装countup依赖，使用npm install countup。
  在views下创建count-to.vue

  ```
<template>
  <div>
    <count-to ref="countTo" :end-val="endVal" @on-animation-end="handleEnd">
      <span slot="left">总金额: </span>
      <span slot="right">元</span>
    </count-to>
    <button @click="getNumber">获取数值</button>
    <button @click="up">更新值</button>
  </div>
</template>

 <script>
import CountTo from '@/components/count-to'
export default {
  name: 'count_to',
  components: {
    CountTo
  },
  data() {
    return {
      endVal: 100
    }
  },
  methods: {
    getNumber() {
      // ref在组件上使用就是调用实例，在html上使用就是调用dom
      this.$refs.countTo.getCount()
    },
    up() {
      this.endVal += Math.random() * 100
    },
    handleEnd(endVal) {
      console.log('end ->', endVal)
    }
  }
}
</script>
  ```

  在router.js下添加路由

  ```
{
  // 封装count-to组件
  path: '/count-to',
  name: 'count_to',
  component: () => import('@/views/count-to.vue'),
}
  ```

  在components下创建count-to文件夹再创建count-to.vue

  ```
  <template>
  <div>
    <slot name="left"></slot><span ref="number"
          :class="countClass"
          :id='eleId'></span>
    <slot name="right"></slot>
  </div>
</template>

<script>
import CountUp from 'countup'
//第一种引入样式
// import './count-to.less'
export default {
  name: 'CountTo',
  computed: {
    eleId () {
      //如果countUp组件在全局使用时就需要一个id。但是id重复会出错。
      // 每一个vue实例this都有一个_uid，这个_uid都不相同。把这个_uid和id值拼接。
      return `count_up_${this._uid}`
    },
    countClass () {
      //第一个参数是样式的名称，第二个参数是样式值，这里类型设置为字符串，默认为空。
      return [
        'count-to-number',
        this.className
      ]
    }
  },
  data () {
    return {
      counter: {}
    }
  },
  props: {
    //@description 起始值
    startVal: {
      type: Number,
      default: 0
    },
    //@description 最终值
    endVal: {
      type: Number,
      require: true
    },
    //@description 小数点后保留几位小数
    decimals: {
      type: Number,
      default: 0
    },
    //@description 动画延迟开始时间
    delay: {
      type: Number,
      default: 0
    },
    //@description 渐变时常
    duration: {
      type: Number,
      default: 1
    },
    //@description 是否使用变速效果
    useEasing: {
      type: Boolean,
      defautl: false
    },
    //@description 是否使用分组
    useGrouping: {
      type: Boolean,
      default: true
    },
    //@description 分组符号
    separator: {
      type: String,
      default: ','
    },
    //@description 整数和小数分割符号
    decimal: {
      type: String,
      default: '.'
    },
    className: {
      type: String,
      default: ''
    }
  },
  methods: {
    getCount () {
      return this.$refs.number.innerText
    },
    emitEndEvent () {
      setTimeout(() => {
        this.$emit('on-animation-end', Number(this.getCount()))
      }, this.duration * 1000)
    }
  },
  watch: {
    // 更新值
    endVal (newVal, oldVal) {
      this.counter.update(newVal)
      this.emitEndEvent()
    }
  },
  mounted () {
    //countUp组件是在dom上修改显示的，所以需要挂载到dom上。
    //因为mounted函数只是挂在dom，不确定是否已经渲染好，使用$nextTick方法，方法调用一个回调函数，会在渲染完毕后调用这个方法。在这里创建实例，第一个参数是id值。
    this.$nextTick(() => {
      //因为这个id是在全局使用的，同一个id值会发生冲突。所以需要计算属性创建不一样的id。
      //这里创建一个countUp实例，这里传入的参数是在全局需要使用的参数。创建完实例需要在data中定义counter，这样就能使用countUp的内置方法了。
      this.counter = new CountUp(this.eleId, this.startVal, this.endVal, this.decimals, this.duration, {
        useEasing: this.useEasing,
        useGrouping: this.useGrouping,
        separator: 'this.separator',
        decimal: 'this.decimal'
      })
      setTimeout(() => {
        this.counter.start()
        this.emitEndEvent()
      }, this.delay)
    })
  }
}
</script>
<style lang="less">
// 第二种引入样式
  .count-to-number {
    color: pink
  }
  // 第三种引入样式
  // @import './count-to.less'
</style>
  ```

  在conponents下创建index.js文件

  ```
import CountTo from './count-to.vue'
// 导出组件，使用时只用@/components/count-to即可，它会自动找到index.js。这里已经导出了，所以可以简练的写。
export default CountTo
  ```

  添加样式

  ```
//第一种引入样式
// import './count-to.less'
<style lang="less">
// 第二种引入样式
  .count-to-number {
    color: pink
  }
  // 第三种引入样式
  // @import './count-to.less'
</style>
  ```

## 从SplitPane组件谈Vue中如何“操作”DOM

  1.简单两列布局
  2.如何让两个div改变宽度
  3.鼠标拖动效果
  4.v-model和.sync的用法
  在views下创建split-pane文件

### 制作一个能自由控制平均宽度高度的盒子。

  ```
<template>
  <div class="split-pane-con">
    <!-- 这里传递value可以直接用v-mode="offset"这种写法 -->
    <!-- :value="offset" @input="handleInput"第二种写法 -->
    <!-- .sync，:value首先绑定offset值，并绑定事件用来更新offset这个值 -->
    <split-pane :value.sync="offset">
      <div slot="left">left</div>
      <div slot="right">right</div>
    </split-pane>
  </div>
</template>

<script>
import SplitPane from '_c/split-pane'
export default {
  components: {
    SplitPane
  },
  data() {
    return {
      offset: 0.8
    }
  },
  methods: {
    // handleInput(value) {
    //   this.offset = value
    // }
  }
}
</script>

<style lang="less">
  .split-pane-con {
    width: 400px;
    height: 200px;
    background-color: pink;
  }  
</style>
  ```

  在components文件夹下创建split-pane文件夹
  在split-pane下创建index.js文件和split-pane.vue文件

```
//components/index.js页面下
import SplitPane from './split-pane.vue'
export default SplitPane

//components/split-pane.vue页面下
<template>
  <div class="split-pane-wrapper" ref="outer">
   <div class="pane pane-left" :style="{ width: leftOffsetPercent}">
     <slot name="left"></slot>
   </div>
   <div class="pane-trigger-con" @mousedown="handleMousedown" :style="{ left: triggerLeft, width: `${triggerWidth}px`}"></div>
   <div class="pane pane-right" :style="{ left: leftOffsetPercent}">
     <slot name="right"></slot>
   </div>
  </div>
</template>

<script>
export default {
  name: 'SplitPane',
  props: {
    // 初始偏移值
    value: {
      type: Number,
      default: 0.5
    },
    triggerWidth: {
      type: Number,
      default: 8
    },
      // 拖动条能拖动的最小值
  min: {
    type: Number,
    default: 0.1
  },
  // 拖动条能拖动的最大值
  max: {
    type: Number,
    default: 0.9
  }
  },
  data () {
    return {
      // 初始布局位置
      // leftOffset: 0.3,
      // 控制能否移动
      canMove: false,
      // 移动时控制拖动条初始值
      initOffset: 0
    }
  },
  computed: {
    leftOffsetPercent () {
      return `${this.value * 100}%`
    },
    triggerLeft() {
      // calc是css3的计算方法
      return `calc(${this.value * 100}% - ${this.triggerWidth / 2}px)`
    }
  },
  methods: {
    handleClick () {
      this.value -= 0.02
    },
    handleMousedown(event) {
      // 点击时控制拖动条
      document.addEventListener('mousemove', this.handleMousemove)
      // 松开拖动条撤销控制
      document.addEventListener('mouseup', this.handleMouseup)
      // 添加鼠标点击位置的距离
      this.initOffset = event.pageX - event.srcElement.getBoundingClientRect().left
      this.canMove = true
    },
    // event是事件对象
    handleMousemove(event) {
      if(!this.canMove) return
      const outerRect = this.$refs.outer.getBoundingClientRect()
      // event.pageX距离页面左侧距离
      // offset偏移的像素数
      // 这里this.initOffset + this.triggerWidth计算结果是鼠标点击拖动条的宽度。不会在点击时鼠标点击时移动到一个中心的位置。
      let offsetPercent = (event.pageX - this.initOffset + this.triggerWidth - outerRect.left) / outerRect.width
      // 控制拖动条，拖动最小值的时候使它不能再拖动。
      if(offsetPercent < this.min) offsetPercent = this.min
      if(offsetPercent > this.max) offsetPercent = this.max
      // 子向父传值
      // this.$emit('input', offsetPercent)
      // 使用$emit触发一个事件，这里固定update:，这里value就是前面传递的属性名。
      this.$emit('update:value', offsetPercent)
      // 计算出偏移的百分比
      // this.value = offsetPercent
    },
    handleMouseup() {
      this.canMove = false
    }
  }
}
</script>
<style lang="less">
.split-pane-wrapper {
  height: 100%;
  width: 100%;
  position: relative;
  .pane {
    height: 100%;
    top: 0;
    position: absolute;
    //&代表父级选择器
    &-left {
      // width: 30%;
      background: red;
    }
    &-right {
      right: 0;
      bottom: 0;
      // left: 30%;
      background: palegoldenrod;
    }
    &-trigger-con {
      height: 100%;
      background: blue;
      position: absolute;
      top: 0;
      z-index: 10;
      user-select: none;
      // 鼠标样式
      cursor: col-resize;
    }
  }
}
</style>
```

  在router下配置路由

```
//router/router.js页面下
  {
    path: '/split-pane',
    name: 'split_pane',
    component: () => import('@/views/split-pane.vue'),
  }
```

## 渲染函数和JSX快速掌握

1.render函数
2.函数式组件
3.JSX
4.作用域插槽

### rander函数

在views下创建render-page.vue，并添加相应的路由。

```
//在main页面下
new Vue({
  router,
  store,
  render: h => {
  //在vue中我们使用模板HTML语法组建页面的，使用render函数我们可以用js语言来构建DOM因为vue是虚DOM，所以在拿到template模板时也要转译成VNode的函数，而用render函数构建DOM，vue就免去了转译过程。当使用render函数描述虚拟DOM时，vue提供一个函数，这个函数是就构建虚拟DOM所需要的工具。官方给他起了个名字叫createElement。还有约定的简写叫h,vm中有一个方法_c,也是这个函数的别名
  // h方法有三个参数，第一个是必选参数，就是要渲染的组件、标签、字符串或函数。第二个参数是配置对象。第三个参数是字符串或者数组这两种值。
    return h('div', {
      attrs: {
        id: 'box'
      },
      style: {
        color: 'red'
      }
    },'lison')
  }
}).$mount('#app')
```

render的使用方法

```
//在main页面下
  render: h => {
    return h(CountTO, {
      // 在render下使用class的方法，可以是字符串、数组和对象。
      'class': 'count-to',
      attrs: {},
      style: {},
      // 在render下使用dom的方法
      // domProps: {
      //   innerHTML: '123'
      // },
      props: {
        endVal: 100
      },
      on: {
        'on-animation-end': (val) => {
          console.log('on-animation-end')
        }
      },
      // 在render下使用click的方法
      nativeOn: {
        'click': () => {
          console.log('click')
        }
      },
      // 这里可以定义自定义的指令
      directives: [],
      // 作用域插槽
      scopedSlots: {},
      slot: '',
      key: '',
      ref: ''
    })
}
```

render下使用map遍历出所有dom节点

```
//在main页面下
const handleClick = event => {
  console.log(event)
  // 在render下阻止冒泡
  event.stopPropagation()
}

let list = [{name: 'lison'}, {name: 'li'}]
const getListEleArr = (h) => {
  return list.map((item, index) => h('li', {
    on: {
      'click': handleClick
    },
    key: `list_item_${index}`
  }, item.name))
}

render: h => h('div', [
  // 这里只能写数组或字符串，不能去掉[]直接写h('span', '111')
  // h('span', '111')
  h('ul', {
    on: {
      'click': handleClick
    }
  }, getListEleArr(h))
])
```

### 函数式组件

在components文件夹下创建list文件夹，在list文件夹下创建index.js和list.vue，这里index.js做引入用。

```
//components/list/list.vue文件下
<template>
  <div>
    <li v-for="(item, index) in list" :key="`list_item_${index}`">
      <span v-if="!render">{{ item.name }}</span>
      <render-dom v-else :render-func="render" :name="item.name"></render-dom>
    </li>
  </div>
</template>

<script>
import RenderDom from '_c/render-dom.js'
export default {
  name: 'List',
  components: {
    RenderDom
  },
  props: {
    list: {
      type: Array,
      default: () => []
    },
    // 如果想在span标签里显示这里的render就需要用到render组件。
    render: {
      type: Function,
      default: () => {}
    }
  }
}
</script>
```

在components文件夹下创建render-dom.js

```
//components/render-dom.js文件下
// 函数式组件，这个组件没有生命周期。它只是一个接收参数的函数，当functional为true的时候，意味着它就是没有状态，没有实力的对象。但是把它引进但做一个组件去用的时候，vue会把它做一个处理。在render这里传入render函数，vue会使用render函数，把里面的逻辑、虚拟节点做一个渲染。所以
export default {
  functional: true,
  // render属性
  porps: {
    name: String,
    renderfunc: Function
  },
  // 这里render就是用户传进来的render。第一个参数是h函数，第二个参数是实例。因为这里是没有实例的，所以用ctx代表这这个文嘉你的对象
  render: (h, ctx) => {
    return ctx.props.renderFunc(h, ctx.props.name)
  }
}
```

```
//views/render-page.vue文件下
<template>
  <div>
    <list :list="list" :render="renderFunc"></list>
  </div>
</template>

<script>
import List from '_c/list'
export default {
  data() {
    return {
      list: [
        { name: 'lison' },
        { name: 'li' }
      ]
    }
  },
  components: {
    List
  },
  methods: {
    renderFunc(h, name) {
      return h('i', {
        style: {
          color: 'pink'
        }
      }, name)
    }
  }
}
</script>
```

### JSX

JSX是react最先提出的，在js里写html标签和特定的语法，最后把字符串转义成js用render函数去渲染

```
//components/render-dom.js文件下
export default {
  functional: true,
  // render属性
  porps: {
    number: Number,
    renderfunc: Function
  },
  // 这里render就是用户传进来的render。第一个参数是h函数，第二个参数是实例。因为这里是没有实例的，所以用ctx代表这这个文嘉你的对象
  render: (h, ctx) => {
    return ctx.props.renderFunc(h, ctx.props.number)
  }
}
```

```
//components/list/list.vue文件下
<template>
  <div>
    <li v-for="(item, index) in list" :key="`list_item_${index}`">
      <span v-if="!render">{{ item.number }}</span>
      <render-dom v-else :render-func="render" :number="item.number"></render-dom>
    </li>
  </div>
</template>

<script>
import RenderDom from '_c/render-dom.js'
export default {
  name: 'List',
  components: {
    RenderDom
  },
  props: {
    list: {
      type: Array,
      default: () => []
    },
    // 如果想在span标签里显示这里的render就需要用到render组件。
    render: {
      type: Function,
      default: () => {}
    }
  }
}
</script>
```

```
//views/rander-page.vue文件下
<template>
  <div>
    <list :list="list" :style="{color: 'red'}" :render="renderFunc"></list>
  </div>
</template>

<script>
import List from '_c/list'
import CountTo from '_c/count-to'
export default {
  data() {
    return {
      list: [
        { number: 100},
        { number: 45}
      ]
    }
  },
  components: {
    List
  },
  methods: {
    // 使用JSX语法时这里一定要使用h
    renderFunc(h, number) {
      return (
        // notiveOn-click为原生绑定事件，on-'on-animation-end'为自定义事件。
        <CountTo nativeOn-click={this.handleClick} on-on-animation-end={this.handleEnd} endVal={number} style={{color: 'pink'}}></CountTo>
      )
    },
    handleClick(event) {
      console.log(event)
    },
    // 自定义事件
    handleEnd() {
      console.log('end')
    }
  }
}
</script>
```

### 作用域插槽

```
//components/list/list.vue文件下
<template>
  <div>
    <li v-for="(item, index) in list" :key="`list_item_${index}`">
      <slot :number="item.number"></slot>
    </li>
  </div>
</template>

<script>
import RenderDom from '_c/render-dom.js'
export default {
  name: 'List',
  components: {
    RenderDom
  },
  props: {
    list: {
      type: Array,
      default: () => []
    },
    // 如果想在span标签里显示这里的render就需要用到render组件。
    render: {
      type: Function,
      default: () => {}
    }
  }
}
</script>
```

```
//views/render-page.vue文件下
<template>
  <div>
    <list :list="list" :style="{color: 'red'}">
      <count-to slot-scope="count" :end-val="count.number"></count-to>
    </list>
  </div>
</template>

<script>
import List from '_c/list'
import CountTo from '_c/count-to'
export default {
  data() {
    return {
      list: [
        { number: 100},
        { number: 45}
      ]
    }
  },
  components: {
    List,
    CountTo
  },
  methods: {
    // 使用JSX语法时这里一定要使用h
    renderFunc(h, number) {
      return (
        // notiveOn-click为原生绑定事件，on-'on-animation-end'为自定义事件。
        <CountTo nativeOn-click={this.handleClick} on-on-animation-end={this.handleEnd} endVal={number} style={{color: 'pink'}}></CountTo>
      )
    },
    handleClick(event) {
      console.log(event)
    },
    // 自定义事件
    handleEnd() {
      console.log('end')
    }
  }
}
</script>
```

## 递归组件

1.封装简单Menu组件
2.递归组件

### 封装简单Menu组件

在components目录下创建menu文件夹，在menu文件夹下创建a-menu-item.vue、a-menu.vue、a-submenu.vue和index.js。index.js用来简写路径，a-menu.vue作为容器，a-menu-item.vue存放一级菜单，a-submenu.vue存放子级菜单。

```
在a-meun.vue文件下
<template>
  <div class="a-menu">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: 'AMenu'
}
</script>

<style lang="less">
  .a-menu {
    & * {
    list-style: none;
    }
    ul {
      padding: 0;
      margin: 0;
    }
  }
</style>
```

```
在a-menu-item文件下
<template>
  <div>
    <li class="a-menu-item">
      <slot></slot>
    </li>
  </div>
</template>

<script>
export default {
  name: 'AMenuItem'
}
</script>

<style lang="less">
.a-menu-item {
  background: rgb(90, 92, 104);
  color: #fff;
}
</style>
```

```
在a-submenu文件下
<template>
  <div>
    <ul class="a-submenu">
    //菜单栏绑定事件
      <div class="a-submenu-titel" @click="handleClick">
      <slot name="title"></slot>
      //判断菜单是否点击，点击后 将图标旋转180度
      <span class="shrink-icon" :style="{ transform: `rotateZ(${showChild ? 0 : 180}deg)` }">^</span>
      </div>
      <div v-show="showChild" class="a-submenu-child-box">
        <slot></slot>
      </div>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'ASubmenu',
  data() {
    return {
      //控制是否展开菜单栏
      showChild: false
    }
  },
  methods: {
    handleClick() {
      this.showChild = !this.showChild
    }
  }
}
</script>

<style lang="less">
  .a-submenu {
    background: rgb(33, 35, 39);
    &-titel {
    color: #fff;
    position: relative;
    .shrink-icon {
      position: absolute;
      top: 4px;
      right: 10px;
    }
    }
    &-child-box {
      overflow: hidden;
      padding-left: 20px;
    }
    li {
      background: rgb(33, 35, 39);
    }
}
</style>
```

### 递归组件

```
//views/menu-page.vue文件下
<template>
  <div class="menu-box">
    <!-- <a-menu>
      <a-menu-item>1111</a-menu-item>
      <a-menu-item>2222</a-menu-item>
      <a-submenu>
        <div slot="title">3333</div>
        <a-menu-item>3333-11</a-menu-item>
        <a-submenu>
          <div slot="title">3333-22</div>
          <a-menu-item>3333-22-11</a-menu-item>
          <a-menu-item>3333-22-22</a-menu-item>
        </a-submenu>
      </a-submenu>
    </a-menu> -->
    <a-menu>
      <!-- 用来循环一级菜单 -->
    <template v-for="(item, index) in list">
      <!-- 判断是否还有子菜单，没有子菜单就直接渲染。 -->
      <a-menu-item v-if="!item.children"
                   :key="`menu_item_${index}`">{{ item.title }}</a-menu-item>
                   <!-- 有子菜单，父向子传递数据 -->
      <re-submenu v-else
                  :key="`meun_item_${index}`"
                  :parent="item"
                  :index="index"></re-submenu>
    </template>
  </a-menu>
  </div>
</template>

<script>
import menuComponents from '_c/menu'
import ReSubmenu from './re-submenu.vue'
export default {
  name: 'menu_page',
  components: {
    ...menuComponents,
    ReSubmenu
  },
  data () {
    return {
      list: [
        {
          title: '111'
        }, {
          title: '222'
        }, {
          title: '333',
          children: [
            {
              title: '333-1'
            }, {
              title: '333-2',
              children: [
                {
                  title: '333-1-1'
                },
                {
                  title: '333-1-2'
                }
              ]
            },
          ]
        }
      ]
    }
  }
}
</script>

<style lang="less">
.menu-box {
  width: 300px;
  height: 400px;
}
</style>
```

```
//views/re-submenu.vue文件下
<template>
  <div>
    <a-submenu>
      <div slot="title">{{ parent.title }}</div>
      <template v-for="(item, i) in parent.children">
      <!-- 没有子菜单，直接渲染 -->
      <a-menu-item v-if="!item.children" :key="`meun_item_${index}_${i}`">{{ item.title }}</a-menu-item>
      <!-- 有就显示这里 -->
      <re-submenu v-else :key="`meun_item_${index}_${i}`" :parent="item"></re-submenu>
      </template>
    </a-submenu>
  </div>
</template>

<script>
import menuComponents from '_c/menu'
export default {
  name: 'ReSubmenu',
  components: {
    ...menuComponents
  },
  props: {
    parent: {
      type: Object,
      // 对象需要返回一个回调函数，({})简写
      default: () => ({})
    },
    // 传递的索引号
    index: Number
  }
}
</script>
```

## 登录/登出以及JWT认证

1.后端代码概览
2.登录以及Token处理
3.Token过期处理
4.退出登录

### 1后端代码概览

index.js配置

```
var express = require('express');
var router = express.Router();
const jwt = require('jsonwebtoken')

const getPasswordByName = (name) => {
  return { password: '123' }
}

router.post('/getUserInfo', function(req, res, next) {
  res.status(200).send({
    code: 200,
    data: {
      name: 'Lison'
    }
  })
});

router.post('/login', function(req, res, next) {
  //获取到前端传递过来的userName和password
  const { userName, password } = req.body
  if (userName) {
    //判断是否有用户名，没有用户名就返回空，有用户名就去数据库里查询这个用户信息，返回信息。
    const userInfo = password ? getPasswordByName(userName) : ''
    //如果没有查到用户信息，返回null。返回前端401。
    //在有用户信息和密码的时候再用传过来的密码和数据库的密码比对，如果不一样还是报401.
    if (!userInfo || !password || userInfo.password !== password) {
      res.status(401).send({
        code: 401,
        mes: 'user name or password is wrong',
        data: {}
      })
    } else {
      //如果用户名也对，密码也对就返回前端一个token。这里token是jwt生成的。
      res.send({
        code: 200,
        mes: 'success',
        data: {
          //第一个字段是自定义字段，第二个字段是加密字段，第三个字段可以设置信息，这里是过期时间。
          token: jwt.sign({ name: userName }, 'abcd', {
            expiresIn: '1d'
          })
        }
      })
    }
  } else {
    //如果前端没有传用户名，也是401.
    res.status(401).send({
      code: 401,
      mes: 'user name is empty',
      data: {}
    })
  }
});

module.exports = router;
```

app.js配置

```
var createError = require('http-errors');
var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');
const jwt = require('jsonwebtoken')

var indexRouter = require('./routes/index');
var usersRouter = require('./routes/users');
var dataRouter = require('./routes/data');

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

//白名单
const whiteListUrl = {
  get: [
  ],
  post: [
    '/index/login',
  ]
}

const hasOneOf = (str, arr) => {
  //这里只要有一个函数返回为true，这个结果就是true。
  return arr.some(item => item.includes(str))
}

app.all('*', (req, res, next) => {
  //从请求中获取到所有请求方式，再转为小写。请求的方式
  let method = req.method.toLowerCase()
  //获取到当前请求的路径
  let path = req.path
  //这里第一个判断的是白名单，如果在白名单对象里面，在判断定义的方法。第二个判断是一个函数。这里第一个参数是当前请求它的路径就是上面的str，第二个参数是传进来的数组，比如是post请求，就会获取白名单post请求里的post数组，再传递给hasOneOf函数，用来判断。如果有一个包含的请求路径就是true。就直接往下走不进行token校验。
  if(whiteListUrl[method] && hasOneOf(path, whiteListUrl[method])) next()
  else {
    //如果没有通过校验就走到这里，这里获取到headers里的authorization字段。
    const token = req.headers.authorization
    //判断是否有token，没有token就返回前端401代码，同时发送错误信息。
    if (!token) res.status(401).send('there is no token, please login')
    else {
      //如果有token就是用jwt进行校验token，第一个参数是传过来的token，第二个参数是前面加密的信息。第三个参数是一个函数，函数里第一个参数是一个错误信息。第二个参数是从token里解码出来的信息，就是在生成token时生成的自定义对象。
      jwt.verify(token, 'abcd', (error, decode) => {
        //这里error是一个对象不是一个null的说明有错误，比如token过期或token错误。返回错误信息。
        if (error) res.send({
          code: 401,
          mes: 'token error',
          data: {}
        })
        else {
          //都成功就从解码出来的decode里的name字段赋值userName，调用next()往下走，走到调用接口的地方。
          req.userName = decode.name
          next()
        }
      })
    }
  }
})

app.use('/index', indexRouter);
app.use('/users', usersRouter);
app.use('/data', dataRouter);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

module.exports = app;

```

user.js配置

```
var express = require('express');
var router = express.Router();
const jwt = require('jsonwebtoken')

/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('respond with a resource');
});

router.get('/getUserInfo', function(req, res, next) {
  res.send('success')
})
//这里req能从app配置中获得到已经校验过的token，校验过后再加密重新返回给前端。
router.get('/authorization', (req, res, next) => {
  const userName = req.userName
  res.send({
    code: 200,
    mes: 'success',
    data: {
      token: jwt.sign({ name: userName }, 'abcd', {
        expiresIn: '1d'
      })
    }
  })
})

module.exports = router;
```

### 登录以及token处理

需要安装依赖js-cookie@2.2.0、md5@2.2.1

```
//views/login.vue文件下，配置基本的请求页面。
<template>
  <div>
    <input v-model="userName">
    <input type="password" v-model="password">
    <button @click="handleSubmit">登录</button>
  </div>
</template>

<script>
// 使用工具函数引入
import { mapActions } from 'vuex'
export default {
  name: 'login_page',
  data() {
    return {
      userName: '',
      password: ''
    }
  },
  methods: {
    // 使用拆分操作符展开，这里的login就是store/model/user.js下定义的login
    ...mapActions([
      'login'
    ]),
    handleSubmit() {
      this.login({
        userName: this.userName,
        password: this.password
      }).then(() => {
        console.log('success')
        // 成功就跳转到home页
        this.$router.push({
          name: 'home'
        })
        // 失败就返回错误信息
      }).catch(error => {
        console.log(error)
      })
    }
  }
}
</script>
```

配置请求的信息，模块化请求信息，利于修改。

```
//api/user.js文件下
import axios from './index.js'

export const getUserInfo = ({ userId }) => {
  return axios.request({
    url: '/getUserInfo',
    method: 'post',
    data: {
      userId
    }
  })
}

export const login = ({ userName, password}) => {
  return axios.request({
    url: '/index/login',
    method: 'post',
    data: {
      userName,
      password
    }
  })
}
//这里传递token到服务器端进行校验。
export const authorization = () => {
  return axios.request({
    url: '/users/authorization',
    method: 'get'
  })
}
```

在vuex里保存发起的异步请求。

```
//state/model/user.js文件下
// 引入配置好的请求方式
import { login } from '@/api/user'
import { setToken } from '@/lib/util'

const actions = {
  // 这里的第一个参数是提交，第二个参数是这里的state实例，第三个参数是store下面的state实例，可以直接操作。第四个参数是action实例的提交方法。
  updateUserName ({ commit, state, rootState, despatch}) {
    // 操作state下的appName
    // rootState.appName
  },
  //通过载荷的形式传递userName和pssword。
  login({ commit }, { userName, password}) {
    login({ userName, password}).then(res => {
      console.log(res)
    }).catch(error => {
      console.log(error)
    })
  }
}

export default {
  state,
  mutations,
  actions,
  module: {

  }
}
```

发起请求前需要保证的环境

```
//main.js文件下，在真实的环境下发起请求，mock会对发起的请求进行拦截。如果配置了mock，需要判断是否在生产环境，如果在生产环境需要注释掉mock。
// if(process.env.NODE_ENV !== 'production') require('./mock')
//vue.config.js文件下，如果服务器端配置跨域，就需要前端配置代理。
  devServer: {
    proxy: 'http://localhost:3000'
  }
//config/index.js文件下，如果配置了代理，这里baseURL就要设置为。
export const baseURL = process.env.NODE_ENV === 'PRODUCTION' 
  ? 'http://production.com'
  : ''
```

获取到token之后需要把token存起来，每次接口调用的时候都把token放到header里传给服务器做一个验证。结合业务的方法，封装到util里面。

```
//lib/util.js文件下
import Cookies from 'js-cookie'

export const setTitle = (title) => {
  window.document.title = title || 'admin'
}
//这里表示存在Cookie里面名字叫什么，给tokenName设一个默认值，如果没有设置就叫token。
export const setToken = (token, tokenName = 'token') => {
  Cookies.set(tokenName, token)
}
//获取token，这里也有tokenName，名字就叫token
export const getToken = (tokenName = 'token') => {
  return Cookies.get(tokenName)
}
```

因为前面全局守卫是写死的，这里需要判断是否携带token

```
//router/index.js页面下
import Vue from 'vue'
import VueRouter from 'vue-router'
import routes from './router'
import store from '@/store'
//这里引入保存token的页面
import {setTitle, setToken, getToken} from '@/lib/util.js'

Vue.use(VueRouter)

const router = new VueRouter({
  routes
})

// 模拟一个登录，实际是通过接口来判断的。
const HAS_LOGINED = false

router.beforeEach((to, from, next) => {
  //获取token判断是否登录过
  const token = getToken()
  if(token) {
    //携带token时，这里判断token是否是有效的。
    store.dispatch('authorization', token).then(() => {
      if(to.name === 'login') next({ name: 'home'})
      else next()
    }).catch(() => {
      //如果token失效，但是又携带token。会一直重复执行这个函数，所以这里把token设为空。
      setToken('')
      next({ name: 'login' })
    })
  }else {
    //判断没有携带token，判断是否去登录页，如果不是跳转到登录也。
    if(to.name === 'login') next()
    else next({ name: 'login'})
  }
})

export default router
```

```
//lib/axios.js页面下，需要在请求拦截器中添加config.headers['Authorization'] = getToken()
//这里引入获取到的token
import { getToken } from '@/lib/util'

 interceptors(instance, url) {
    instance.interceptors.request.use(config => {
      //添加全局的loading
      //Spin组件，添加遮罩层，覆盖遮罩层后就无法点击。Spin.show()
      if(!Object.keys(this.queue).length) {}/* Spin.show() */
      this.queue[url] = true
      // 传递token，每次调用都会获取token，并把token放到header这个字段里面。
      config.headers['Authorization'] = getToken()
      return config
    }, error => {
      return Promise.reject(error)
    })
```

在user.js中保存发起校验token的请求

```
//store/model/user.js页面下
import { login, authorization } from '@/api/user'

//在login请求下面添加token校验请求，发送登录请求时同时发送校验token请求。
  authorization ({ commit }, token) {
    return new Promise((resolve, reject) => {
      authorization().then(res => {
        if (parseInt(res.code) == 401) {
          reject(new Error('token error'))
        } else {
          resolve()
        }
      }).catch(error => {
        reject(error)
      })
    })
  }
```

### token过期

```
//store/model/user.js页面下
import { login, authorization } from '@/api/user'

//在login请求下面添加token校验请求，发送登录请求时同时发送校验token请求。
  authorization ({ commit }, token) {
    return new Promise((resolve, reject) => {
      authorization().then(res => {
        if (parseInt(res.code) == 401) {
          reject(new Error('token error'))
        } else {
          //因为token是有一个过期时间，如果用户还在浏览页面，但是token过期了就会退出。这里保存token，让用户在浏览页面时会添加token的时间不会过期。
          setToken(res.data.token)
          resolve()
        }
      }).catch(error => {
        reject(error)
      })
    })
  }
```

### 退出登录

```
//viwes/home.vue页面下
<template>
  <div class="home">
    <button @click="handleLogout">退出登录</button>
    <router-view></router-view>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  methods: {
    ...mapActions([
      'logout'
    ]),
    handleLogout() {
      this.logout()
      this.$router.push({
        name: 'login'
      })
    }
  }
}
</script>

//store/model/user.js文件下
//在登录的请求下面添加退出登录，让token为空并保存在vuex中。
  logout () {
    setToken('')
  }
```

## 响应式布局

1.vue-cli3中使用iview
2.布局组件的使用
3.栅格组件实现响应式布局

### vue-cli3中使用iview

引入 iView #
一般在 webpack 入口页面 main.js 中如下配置：

```
import Vue from 'vue';
import VueRouter from 'vue-router';
import App from 'components/app.vue';    // 路由挂载
import Routers from './router.js';       // 路由列表
import iView from 'iview';
import 'iview/dist/styles/iview.css';    // 使用 CSS

Vue.use(VueRouter);
Vue.use(iView);

// 路由配置
const RouterConfig = {
    routes: Routers
};
const router = new VueRouter(RouterConfig);

new Vue({
    el: '#app',
    router: router,
    render: h => h(App)
});
```

按需引用 #
如果您想在 webpack 中按需使用组件，减少文件体积，可以这样写：

```
import Checkbox from 'iview/src/components/checkbox';
```

具体参考iview官网

### 布局组件的使用

安装iview，使用npm install iview --save
在views下创建layout文件

```
//router/router.js文件下添加路由
import Layout from '@/views/layout.vue'
  {
    path: '/',
    //当访问的页是首页
    alias: '/home_page',
    name: 'home',
    component: Layout,
    //使用嵌套路由
    children: [
      {
        path: 'home',
        component: Home
      }
    ]
  }
  在App.vue文件下格式化样式
html,body {
height: 100%;
}
body {
  margin: 0;
}
```

layout布局

```
//views/laytou.vue页面下
<template>
  <div class="layout-wrapper">
    <Layout class="layout-outer">
    //collapsible为侧边栏是否可收起
      <Sider collapsible
      //breakpoint控制浏览器缩小的大小来实现响应式，sm为控制多少像素才缩放。
             breakpoint="sm"
             //控制侧边栏收起与展开
             v-model="collapsed"></Sider>
      <Layout>
        <Header class="header-wrapper">
        //绑定css动画
          <Icon :class="triggerClasses"
          //Icon是没有click事件的使用native调用最外侧html标签来调用原生click事件
                @click.native="handleCollapsed"
                //头部菜单栏三条横线
                type="md-menu"
                :size="32" />
        </Header>
        //栅格区域
        <Content class="content-card">
          <Card shadow
                class="page-card">
            <router-view></router-view>
          </Card>
        </Content>
      </Layout>
    </Layout>
  </div>
</template>

<script>
export default {
  data () {
    return {
      collapsed: false
    }
  },
  computed: {
    triggerClasses () {
      return [
        'trigger-icon',
        this.collapsed ? 'rotate' : ''
      ]
    }
  },
  methods: {
    handleCollapsed () {
      this.collapsed = !this.collapsed
    }
  }
}
</script>

<style lang="less">
.layout-wrapper,
.layout-outer {
  height: 100%;
  .header-wrapper {
    background-color: #fff;
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);
    padding: 0 23px;
    .trigger-icon {
      cursor: pointer;
      transition: transform 0.3s ease;
      &.rotate {
        transform: rotateZ(-90deg);
        transition: transform 0.3s ease;
      }
    }
  }
  .content-card {
    padding: 10px;
  }
  .page-card {
    //less实现计算属性
    min-height: ~"calc(100vh - 84px)";
  }
}
</style>
```

### 栅格组件实现响应式布局

在home文件下栅格组件

```
<template>
  <div class="home">
    <Row>
      <i-col></i-col>
    </Row>
    //gutter每一列间隔为10像素
    <Row :gutter="10">
    //span为字符串或者数值，总栅格数为24，span=“12”为对半分
      <i-col span="12"></i-col>
      <i-col span="12"></i-col>
    </Row>
        <Row :gutter="10" class="blue">
        //offset="1"所有栅格会往右移动1 push="1"只有这个栅格会往右移动
      <i-col :md="6" :sm="12" :xs="24" offset="1" push="1"></i-col>
        //md为响应式栅格所占的份数，当页面为特定像素比如md代表页面大于等于992像素。比如md="6"即24/6份 在一行上分为4列展示
      <i-col :md="6" :sm="12" :xs="24"></i-col>
      //sm为页面像素大于等于768像素，这里:sm="12"就是一行上展示2列
      <i-col :md="6" :sm="12" :xs="24"></i-col>
      //xs为像素小鱼768像素时，这里:xs="24"就是一行上展示一列
      <i-col :md="6" :sm="12" :xs="24"></i-col>
    </Row>
  </div>
</template>

<style lang="less">
  .home {
    .ivu-col {
      height: 50px;
      margin-top: 10px;
      background-color: pink;
      //只想让内容区域填充padding，多余的区域不会有颜色 。
      background-clip: content-box;
    }
    .blue {
      .ivu-col {
        background: blue;
      background-clip: content-box;
      }
    }
  }
</style>
```

## 可收缩多级菜单的实现

1.递归组件实战
2.v-if和v-show对比

### 递归组件实战

在compones下创建side-menu文件夹，在文件夹下创建side-menu.vue、re-dropdown.vue、re-submenu.vue和index.js文件。index.js文件引入side-menu.vue文件用来展示组件。side-menu.vue用来展示展开与收缩是的组件，re-submenu.vue为展开时的组件，re-dropdown.vue为收缩时的组件。

```
//components/side-menu/side-menu.vue文件下
<template>
  <div class="side-menu-wrapper">
  //顶上还有一个可以自定义的图片区域，这里添加一个插槽用来自定义。
    <slot></slot>
    //这里展示menu组件，主题theme设置为黑色。v-show控制是收缩哪个菜单。on-select控制点击的是哪个菜单
    <Menu v-show="!collapsed" width="auto" @on-select="handleSelect"
          theme="dark">
          //递归组件不能用menu-item直接来做，要用template。template里不能有key。
      <template v-for="item in list">
      //现在有两种情况。一种情况是有子级
        <re-submenu v-if="item.children"
        //用name值当作key
                    :key="`menu_${item.name}`"
                    //给子组件re-submenu.vue传递数据源。
                    :parent="item"
                    //传递name值
                    :name="item.name">
        </re-submenu>
        //每一级menu-item就是一个一级菜单
        //一种情况是没有子级
        <menu-item v-else
                   :key="`menu_${item.name}`"
                    //传递name值
                   :name="item.name">
          <Icon :type="item.icon" />
          {{ item.title }}
        </menu-item>
      </template>
    </Menu>
    //这里展示dropdown组件
    <div v-show="collapsed" class="drop-wrapper">
      <template v-for="item in list">
      //有子级菜单显示这里，递归组件传递属性这里icon-color让第一级菜单显示白色，:show-title为是否显示一级菜单名称。on-select是子级菜单传递过来的自己菜单点击事件
        <re-dropdown @on-select="handleSelect" v-if="item.children" :show-title="false" icon-color="#fff" :key="`drop_${item.name}`" :parent="item"></re-dropdown>
        //在没有子菜单的时候，收缩显示这个菜单的title值。
        <Tooltip v-else transfer
                 :content="item.title"
                 //控制鼠标放在菜单上文字靠右边排列显示
                 placement="right"
                 :key="`drop_${item.name}`">
              //click在收缩状态下的点击事件
          <span @click="handleClick(item.name)" class="drop-menu-span">
            <Icon :type="item.icon"
            //让一级菜单图标为白色
            color="#fff"
                  :size="20" />
          </span>
        </Tooltip>
      </template>
    </div>
  </div>
</template>

<script>
import ReSubmenu from './re-submenu.vue'
import ReDropdown from './re-dropdown.vue'
export default {
  name: 'SideMenu',
  components: {
    ReSubmenu,
    ReDropdown
  },
  props: {
    collapsed: {
      type: Boolean,
      //默认展开，即false
      default: false
    },
    list: {
      type: Array,
      //子组件的类型如果是数组、对象、函数，不能直接写，要写一个函数，这个函数return的值就是这个空数组、对象、函数。
      default: () => []
    }
  },
  methods: {
    handleSelect(name) {
      console.log(name)
    },
    handleClick(name) {
      console.log(name)
    }
  }
}
</script>

<style lang="less">
.side-menu-wrapper {
  width: 100%;
  .ivu-tooltip,
  .drop-menu-span {
    display: block;
    width: 100%;
    text-align: center;
    padding: 5px 0;
  }
  //因为子级菜单的宽度属性都是和父级菜单一样的，这里就只选择类名来修改居中
  .drop-wrapper > .ivu-dropdown {
    display: block;
    padding: 5px;
    margin: 0 auto;
  }
}
</style>
```

```
//views/layout.vue文件下
<template>
  <div class="layout-wrapper">
    <Layout class="layout-outer">
      <Sider collapsible
             :width="300"
             breakpoint="sm"
             v-model="collapsed">
             //在这里展示收缩和展开的组件，用collapsed来控制。给子组件side-menu传递collpased。
        <side-menu :collapsed="collapsed"
        //向子组件传递获取的数据
                   :list="menuList"></side-menu>
      </Sider>
      <Layout>
        <Header class="header-wrapper">
          <Icon :class="triggerClasses"
                @click.native="handleCollapsed"
                type="md-menu"
                :size="32" />
        </Header>
        <Content class="content-card">
          <Card shadow
                class="page-card">
            <router-view></router-view>
          </Card>
        </Content>
      </Layout>
    </Layout>
  </div>
</template>

<script>
//引入收缩展开组件
import SideMenu from '_c/side-menu'
export default {
  components: {
    SideMenu
  },
  data () {
    return {
      collapsed: true,
      //循环的假数据
      menuList: [
        {
          title: '1',
          name: 'menu1',
          icon: 'ios-alarm'
        },
        {
          title: '2',
          name: 'menu2',
          icon: 'ios-alarm'
        },
        {
          title: '3',
          name: 'menu3',
          icon: 'ios-alarm',
          children: [{
            title: '3-1',
            name: 'menu11',
            icon: 'ios-alarm'
          },
          {
            title: '3-2',
            name: 'menu12',
            icon: 'ios-alarm',
            children: [
              {
                title: '3-2-1',
                name: 'menu12-1',
                icon: 'ios-alarm'
              },
              {
                title: '3-2-2',
                name: 'menu12-2',
                icon: 'ios-alarm'
              },
            ]
          },
          ]
        },
      ]
    }
  },
  computed: {
    triggerClasses () {
      return [
        'trigger-icon',
        this.collapsed ? 'rotate' : ''
      ]
    }
  },
  methods: {
    handleCollapsed () {
      this.collapsed = !this.collapsed
    }
  }
}
</script>

<style lang="less">
.layout-wrapper,
.layout-outer {
  height: 100%;
  .header-wrapper {
    background-color: #fff;
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);
    padding: 0 23px;
    .trigger-icon {
      cursor: pointer;
      transition: transform 0.3s ease;
      &.rotate {
        transform: rotateZ(-90deg);
        transition: transform 0.3s ease;
      }
    }
  }
  .content-card {
    padding: 10px;
  }
  .page-card {
    min-height: ~"calc(100vh - 84px)";
  }
}
</style>
```

递归组件，哪一块是会重复递归的就需要把那一块拆成组件。

```
//components/side-menu/re-submenu.vue文件下
<template>
  <div>
    <Submenu :name="parent.name">
    //这里插槽用来展示可以展开的菜单的title值
      <template slot="title">
        <Icon :type="parent.icon" />
        {{ parent.title }}
      </template>
      <template v-for="item in parent.children">
        <re-submenu v-if="item.children"
                    :key="`menu_${item.name}`"
                    :parent="item"
                    :name="item.name">
        </re-submenu>
        <menu-item v-else
                   :key="`menu_${item.name}`"
                   :name="item.name"><Icon :type="item.icon" />{{ item.title }}</menu-item>
      </template>
    </Submenu>
  </div>
</template>

<script>
export default {
  name: 'ReSubmenu',
  props: {
    parent: {
      type: Object,
      default: () => ({})
    }
  }
}
</script>
```

收缩组件

```
//components/side-meun/re-dropdown.vue文件下
<template>
    //这里dropdown标签和一级标签是同级的。placement控制子级菜单显示的位置。on-click控制子级菜单点击
    <Dropdown @on-click="handleClick" placement="right-start">
    //控制鼠标放上去显示子级菜单，:style让有子级菜单的菜单居中
      <span class="drop-menu-span" :style="titleStyle">
        <Icon :type="parent.icon"
              //接收递父级传递的颜色，一级标签颜色
              :color="iconColor"
              :size="20" />
            //控制是否展示一级菜单title
        <span v-if="showTitle">{{ parent.title }}</span>
      </span>
      //下拉菜单，作为插槽来定义。
      <DropdownMenu slot="list">
        <template v-for="item in parent.children">
          <re-dropdown v-if="item.children"
                       :key="`drop_${item.name}`"
                       :parent="item"></re-dropdown>
            //鼠标放上去显示子级菜单
          <DropdownItem v-else
                    :name="item.name"
                    :key="`drop_${item.name}`">
            <Icon :type="item.icon"
                  color="515a6e"
                  :size="20" />
            {{ item.title }}
          </DropdownItem>
        </template>
      </DropdownMenu>
    </Dropdown>
</template>

<script>
export default {
  name: 'ReDropdown',
  props: {
    parent: {
      type: Object,
      default: () => ({})
    },
    iconColor: {
      type: String,
      default: '#515a6e'
    },
    //让第一级不显示title
    showTitle: {
      type: Boolean,
      default: true
    }
  },
  computed: {
    titleStyle() {
      return {
        //让有子级菜单的菜单居中
         textAlign: this.showTitle ? 'left' : 'center',
        paddingLeft: this.showTitle ? '16px' : ''
      }
    }
  },
  methods: {
    handleClick(name) {
      //因为点击会触发2次，这里判断是否是一级菜单
      if(!this.showTitle) this.$emit('on-select', name)
    }
  }
}
</script>
```

## 可编辑表格的实现

1.JSX进阶
2.单个单元格编辑表格
3.多单元格编辑表格

### 准备工作

在api文件夹下创建data.js用来发起data数据请求

```
//api/data.js文件下
import axios from './index'

export const getTableData = () => {
  return axios.request({
    url: '/getTableData',
    method: 'get'
  })
}
```

在mock/response文件夹下创建data.js用来模拟服务器返回data数据

```
//mock/response/data.js
import { doCustomTimes } from '@/lib/tools'
import Mock from 'mockjs'

export const getTableData = () => {
  const template = {
    name: '@name',
    'age|18-25': 0,
    email: '@email'
  }
  let arr = []
  //发起请求时返回5条数据
  doCustomTimes(5, () => {
    arr.push(Mock.mock(template))
  })
  return arr
}
```

工具函数，表示返回data函数的条数

```
//lib/tools.js
  //调用这个方法，第一个参数是第二个参数执行几次
export const doCustomTimes = (times, callback) => {
  let i = -1
  while (++i < times) {
    callback()
  }
}
```

### 可编辑表格组件

在views下创建table.vue文件，用来展示table数据，添加路由。

```
//viwes/table.vue文件下
<template>
  <div>
  //单个单元格可编辑表格
  //向子组件传递数据
    <!-- <edit-table :columns="columns" v-model="tableData" @on-edit="handleEdit"></edit-table> -->
    //多个单元格可编辑表格
    <edit-table-mul :columns="columns" v-model="tableData"></edit-table-mul>
  </div>
</template>

<script>
import { getTableData } from '@/api/data'
import EditTable from '_c/edit-table'
import EditTableMul from '_c/edit-table-mul'
export default {
  components: {
    EditTable,
    EditTableMul
  },
  data () {
    return {
      tableData: [],
      //定义表格的名称、索引和是否可编辑。
      columns: [
        { key: 'name', title: '姓名' },
        { key: 'age', title: '年龄', editable: true },
        { key: 'email', title: '邮箱', editable: true }
      ]
    }
  },
  methods: {
    //子组件向父组件传递修改后数据，再向后端传递数据，后端就知道前端修改了什么数据了。
    handleEdit({row, index, column, newValue}) {
      console.log({row, index, column, newValue})
    }
  },
  mounted () {
    getTableData().then(res => {
      this.tableData = res
    })
  }
}
</script>
```

在components文件夹下创建edit-table文件夹，在文件夹下创建index.js和edit-table.vue文件。index.js作为快捷路径

```
<template>
  <div>
  //引入父组件数据，这里的数据使用编辑后的数据。
    <Table :columns="insideColumns" :data="value"></Table>
  </div>
</template>

<script>
import clonedeep from 'clonedeep'
export default {
  name: 'EditTable',
  data() {
    return {
      //处理好的数据
      insideColumns: [],
      //全局的标记
      edittingId: '',
      //每次输入的内容
      edittingContent: ''
    }
  },
  props: {
    columns: {
      type: Array,
      default: () => []
    },
    //父组件使用的是v-model，子组件固定只能使用value来接收。
    value: {
      type: Array,
      default: () => []
    }
  },
  watch: {
    //如果表头改变了数据就会改变，所以需要监听columns的更新
    columns() {
      this.handleColumns()
    }
  },
  methods: {
    //接收参数
    handleClick({row, index, column}) {
      //判断是在保存还是编辑状态，如果为true就是编辑状态，下面处理的就是要保存的逻辑。
      if(this.edittingId === `${column.key}_${index}`) {
        //这里不能直接赋值修改的新的值，需要深拷贝后再赋值。
        let tableData = clonedeep(this.value)
        //
        tableData[index][column.key] = this.edittingContent
        //使用emit触发input，v-model一定要触发input事件才能让子组件修改父组件的值，再把修改后的数据传给父组件进行修改。
        this.$emit('input', tableData)
        //向父组件传递一个事件，告诉父组件每次更新后的值是什么。
        this.$emit('on-edit', { row, index, column, newValue: this.edittingContent})
        //每次点击保存后让这个全局的标记为空，这样保存就会变会编辑。
        this.edittingId = ''
        //保存后输入的值为空。
        this.edittingContent = ''
      }else {
        //对标记进行拼接，这样我们就能知道是哪一列哪一行。如果不在编辑状态就让它变为编辑状态。
      this.edittingId = `${column.key}_${index}`
      }
    },
    handleInput(newValue) {
      //把每次输入的内容保存到edittingContent里面
      this.edittingContent = newValue
    },
    handleColumns() {
      //使用map方法映射，item代表一列的对象。
    const insideColumns = this.columns.map(item => {
      //先判断是否有render，如果已经定义了，render函数就直接返回。再次判断是否需要编辑，不需要编辑的也直接返回。
      if(!item.render && item.editable) {
        //render必须要携带h，这里第二个参数是一个对象，这里使用数据解构。行、列和索引。
        item.render = (h, {row, index, column}) => {
          //全局的标记，判断edittingId等于点击的单元格
          const isEditting = this.edittingId === `${column.key}_${index}`
          //这里使用JSX语法来写，这个括号里最外层只能有一个html元素。
          return (
            <div>
            //在JSX里逻辑只能在{}里写。在JSX不能使用vue语法v-if和v-show使用js逻辑来完成。这里完成的功能是当点击一个单元格这个单元格会显示一个输入框和一个保存字样，当点击保存后数据修改，保存字样变为编辑。同时只能点击一个单元个。在点击一个单元格的时候这个单元格唯一的标记被激活，全局的值变为当前点击的单元格的值。通过标记来控制哪个输入框的显示。
            //这里定义edittingId为标记，每一行有一个行号，每行都有一个key值，通过行号和key值来确定一个单元格。这个标记就是行号和key值拼接起来。
            //这里row是行对象column是列对象,row[column.key]就能取到每行的数据。value={row[column.key]}这里就是点击时输入框的数据就是当前的数据。现在点击的单元格判断成功就显示输入框，不成功就显示文字。使用on-input触发事件
            {isEditting ? <i-input value={row[column.key]} style="width: 50px" on-input={this.handleInput}></i-input> : <span>{row[column.key]}</span>}
            //iview的组件，在JSX里面标签名和原生html标签名重复的要在前面加i。因为只在age和email里添加了editable，所以在这两个前面添加，iview定义的on-click事件不是原生的事件，如果事件名叫on-click就要在on-click前面再加一个on-前缀。JSX里语句里使用的{}代替""。如果想在里面传参数要在后面添加.bind，在this上去执行。参数可以是对象结构赋值或者逗号隔开。
              <i-button on-click={this.handleClick.bind(this, {row, index, column})}>{isEditting ? '保存' : '编辑'}</i-button>
            </div>
          )
        }
        return item
        //不需要判断的直接返回
      }else return item
    })
    //把处理好的数据重新赋值使用。
    this.insideColumns = insideColumns
    }
  },
  mounted () {
    //如果columns更新了，这里就再更新。
    this.handleColumns()
  }
}
</script>
```

### 多单元格编辑

table表格不变，在components下创建edit-table-mul文件夹，文件夹下创建edit-table-mul.vue和index.js文件index.js做引入文件。
思路：把传进来的数据拷贝一份，在这个数据上面做一个处理，每一行都维护一个编辑状态，比如第一行是一个对象，3个字段把进入编辑状态的key传到一个数组里，判断这个数组里存不存在key来提供显示状态。

```
<template>
  <Table :columns="insideColumns"
         :data="value"></Table>
</template>

<script>
import clonedeep from 'clonedeep'
export default {
  name: 'EditTable',
  data () {
    return {
      insideData: [],
      insideColumns: [],
      edittingId: '',
      edittingContent: ''
    }
  },
  props: {
    columns: {
      type: Array,
      default: () => []
    },
    value: {
      type: Array,
      default: () => []
    }
  },
  watch: {
    columns () {
      this.handleColumns()
    },
    value () {
      //输入的值改变的时候也执行
      this.handleColumns()
    }
  },
  methods: {
    handleClick ({ row, index, column }) {
      //判断是保存状态还是编辑状态，还在编辑状态就添加一个key值，编辑完成就给-1状态，展示编辑字样。
      let keyIndex = this.insideData[index].edittingKeyArr ? this.insideData[index].edittingKeyArr.indexOf(column.key) : -1
      //把要判断的行号赋值给rowObj
      let rowObj = this.insideData[index]
      //判断是在保存状态是编辑状态
      if (keyIndex > -1) {
        //把正在编辑数组里的key值删掉
        rowObj.edittingKeyArr.splice(keyIndex, 1)
        //把修改过后的值赋值给现在点击的这个行
        this.insideData.splice(index, 1, rowObj)
        //通过$emit方法提交一个input，提交修改的值传递给父组件
        this.$emit('input', this.insideData)
        //通过$emit方法提交on-edit事件，在内部获得的行号传递给父组件修改后端。
        this.$emit('on-edit', { row, index, column, newValue: this.insideData[index][column.key] })
      } else {
        //当前非编辑状态，有可编辑的key值，但是不在编辑状态就添加可编辑的key值并把这行添加到数组中，让这行处于编辑状态。
        rowObj.edittingKeyArr = (rowObj.edittingKeyArr) ? [...rowObj.edittingKeyArr, column.key] : [column.key]
        //index就是行号，再替换掉这行。
        this.insideData.splice(index, 1, rowObj)
      }
    },
    handleInput (row, index, column, newValue) {
      //把修改后的新值赋值给所修改的行，用来给父组件传递。不能直接在输入框中向父组件传递数据，如果直接传递数据会造成输入一次页面渲染一次，不能连续输入。需要输入完成，点击保存后在向父组件传递数据，完成视图渲染。
      this.insideData[index][column.key] = newValue
    },
    handleColumns () {
      //深拷贝输入的值
      this.insideData = clonedeep(this.value)
      //遍历columns
      const insideColumns = this.columns.map(item => {
        //判断是否能编辑
        if (!item.render && item.editable) {
          item.render = (h, { row, index, column }) => {
            //给父组件函数添加一个新的数据字段KeyArr。insideData[index]代表行，在每一行数据对象添加edittingKeyArr，代表可编辑。
            const keyArr = this.insideData[index] ? this.insideData[index].edittingKeyArr : []
            return (
              <div>
              //判断行数据字段有edittingKeyArr数据字段代表可编辑，通过indexOf判断column.key有没有key值，大于-1就是有，就是在数组里面。
                {keyArr && keyArr.indexOf(column.key) > -1
                  ? <i-input value={row[column.key]} style="width: 50px;" on-input={this.handleInput.bind(this, row, index, column)}></i-input>
                  : <span>{row[column.key]}</span>}
                <i-button on-click={() => { this.handleClick({ row, index, column }) }}>{keyArr && keyArr.indexOf(column.key) > -1 ? '保存' : '编辑'}</i-button>
              </div>
            )
          }
          return item
        } else return item
      })
      this.insideColumns = insideColumns
    }
  },
  mounted () {
    this.handleColumns()
  }
}
</script>

<style>
</style>

```

## Tree组件实现文件目录-基础实现

1.Tree组件使用
2.扁平数据树状化
3.自定义组件结构

### Tree组件使用

在views下folder-tree文件夹，在文件夹下创建folder-tree.vue文件，创建添加相应的路由。
data.js中添加发起的请求，在开发中经常需要获取独立的文件列表，如果都要就发2次请求。

```
//api/data.js
export const getFolderList = () => {
  return axios.request({
    url: '/getFolderList',
    method: 'get'
  })
}

export const getFileList = () => {
  return axios.request({
    url: '/getFileList',
    method: 'get'
  })
}
```

在mock里做请求拦截，返回数据。

```
//mock/response/data.js

import { doCustomTimes } from '@/lib/tools'
import Mock from 'mockjs'

export const getFileList = () => {
  const template = {
    //中文词
    'name|5': '@cword',
    //日期加时间
    'creat_tiem': '@datatiem',
    //属于哪个文件夹的文件夹id
    'folder_id|1-5': 0,
    //文件id，每次循环+1
    'id|+1': 10000
  }
  let arr = []
  //返回10条信息
  doCustomTimes(10, () => {
    arr.push(Mock.mock(template))
  })
  return arr
}

export const getFolderList = () => {
  const template1 = {
    'name|1': '@word',
    'creat_tiem': '@datatime',
    //文件夹可以嵌套，这就是嵌套的id
    'folder_id': 0,
    //自身的id
    'id|+1': 1
  }
  const template2 = {
    'name|1': '@word',
    'creat_time': '@datatime',
    'folder_id|+1': 1,
    //id从4开始每次加1
    'id|+1': 4
  }
  let arr = []
  doCustomTimes(3, () => {
    arr.push(Mock.mock(template1))
  })
  doCustomTimes(2, () => {
    arr.push(Mock.mock(template2))
  })
  return arr
}
```

在index.js中使用

```
//mock/index.js文件下
import Mock from 'mockjs'
import { getFileList, getFolderList } from './response/data'

//第一个参数可以是字符串或者正则表达式。第二个参数是请求方式写post或者get等等，这里可以省略。第三个参数是一个模板或者方法。
Mock.mock(/\/getFileList/, 'get', getFileList)
Mock.mock(/\/getFolderList/, 'get', getFolderList)

export default Mock
```

### 扁平数据树状化

```
<template>
  <div class="folder-wrapper">
    <Tree :data="folderTree" :render="renderFunc"></Tree>
  </div>
</template>

<script>
import { getFolderList, getFileList } from '@/api/data'
import { putFileInFolder,transferFolderToTree } from '@/lib/util'
export default {
  data () {
    return {
      folderList: [],
      fileList: [],
      folderTree: []
    }
  },
  mounted () {
    Promise.all([getFolderList(), getFileList()]).then(res => {
      this.folderTree = transferFolderToTree(putFileInFolder(res[0],res[1]))
    })
  }
}
</script>

<style lang="less">
.folder-wrapper{
  width: 300px;
}
</style>
```

使用两个方法，把文件放到对应的文件夹里，再把文件按照层级关系，一层一层的拼接起来。

```
把文件放到文件里，这个方法只用遍历一遍就可以
//lib/util.js文件下
//避免修改传过来的数据，使用深拷贝。
import clonedeep from 'clonedeep'

export const putFileInFolder = (folderList, fileList) => {
  //首先深拷贝一份
  const folderListCloned = clonedeep(folderList)
  const fileListCloned = clonedeep(fileList)
  //遍历文件夹列表，遍历文件夹的同时遍历文件，这样文件就放到文件夹里面了。使用return返回处理好的数据
  return folderListCloned.map(folderItem => {
    //每遍历一次folderItem判断文件列表里有哪个folderId是当前的folderItem的id
    //文件夹id
    const folderId = folderItem.id
    //遍历完folderItem，这个文件就不属于其他文件夹，找到文件归属，再次遍历就不会遍历这个文件了。
    //从后面往前面删东西，这样就不会改变索引号。
    let index = fileListCloned.length
    while (--index >= 0) {
      //获取到当前遍历到的文件对象
      const fileItem = fileListCloned[index]
      //文件归属id和文件夹的id匹配对了
      if (fileItem.folder_id === folderId) {
        //匹配对了就把这个文件移除，并获得移除的文件对象。
        const file = fileListCloned.splice(index, 1)[0]
        //后端返回的名称各种各样，iview格式是title，把name赋值给title方便修改。
        file.title = file.name
        //把移除的文件对象放到对应的文件夹里面，使用children做嵌套
        //因为一开始还没有children属性，这里判断一下，有children属性后在放到文件夹里面
        if (folderItem.children) folderItem.children.push(file)
        //如果没有文件夹属性就添加一个文件夹属性
        else folderItem.children = [file]
      }
    }
    //添加文件夹类型，用来判断是文件夹还是文件。
    folderItem.type = 'folder'
    //这里返回加工后的数组
    return folderItem
  })
}

export const transferFolderToTree = folderList => {
  //如果文件夹列表是空就直接返回空数组
  if (!folderList.length) return []
  //深拷贝文件夹列表，以免操作原数组
  const folderListCloned = clonedeep(folderList)
  //文件夹是有层级的，这里要用到递归。
  //定义一个函数，参数是id，就是文件夹的id
  const handle = id => {
    //用来存放文件夹的空数组
    let arr = []
    //遍历文件夹对象
    folderListCloned.forEach(folder => {
      //如果遍历到的文件夹的id和传进来的id相等，遍历到的文件夹就是传进来文件夹的子级。这里遍历所有文件夹父级都是一级文件夹。
      if (folder.folder_id === id) {
        //当前的文件夹也有子级，把它的id传进去，再去找它的id还有哪个children。
        const children = handle(folder.id)
        //如果当前的文件夹有children，把已有的当前文件夹的children和children合并复制给当前文件夹的children。
        if (folder.children) folder.children = [].concat(folder.children, children)
        //如果没有children就直接等于children
        else folder.children = children
        //这里也是靠title来完成渲染的，这里把name改成title。
        folder.title = folder.name
        //当前遍历的文件夹是传进来的文件夹的子级，push到文件夹中。
        arr.push(folder)
      }
    })
    //这里把处理后的数据返回给arr数组
    return arr
  }
  //所有第一级文件夹folder_id都是0，这里传入0。这里得到的就是所有一级文件夹下的子级
  return handle(0)
}
```

### 自定义组件结构

```
<template>
  <span class="folder-wrapper">
    <Tree :data="folderTree"
          :render="renderFunc"></Tree>
  </span>
</template>

<script>
import { getFolderList, getFileList } from '@/api/data'
import { putFileInFolder, transferFolderToTree } from '@/lib/util'
export default {
  data () {
    return {
      folderList: [],
      fileList: [],
      folderTree: [],
      //使用render自定义组件
      renderFunc: (h, { root, node, data }) => {
        return (
          <div class="tree-item">
            { data.type === 'folder' ? <icon type="ios-folder" color="#2d8cf0" style="margin-right: 10px" /> : ""}
            { data.title}
          </div>
        )
      }
    }
  },
  mounted () {
    Promise.all([getFolderList(), getFileList()]).then(res => {
      this.folderTree = transferFolderToTree(putFileInFolder(res[0], res[1]))
    })
  }
}
</script>

<style lang="less">
.folder-wrapper {
  width: 300px;
  .tree-item {
    display: inline-block;
    width: ~"calc(100% - 50px)";
    height: 30px;
    line-height: 30px;
  }
}
</style>
```

## Tree组件实现文件目录-高级实现

1.封装文件目录组件
2.操作目录
3.多个属性v-model替代方案
4.增加钩子函数

封装的组件作用就是尽量减少操作，增加复用。

```
//views/folder-tree/folder-tree.vue文件下
<template>
  <div class="folder-wrapper">
    <folder-tree
      //传入文件夹列表
      :folder-list.sync="folderList"
      //传入文件列表
      :file-list.sync="fileList"
      //向子组件传递下拉菜单的选项
      :folder-drop="folderDrop"
      :file-drop="fileDrop"
      //增加一个钩子，用来执行删除前的接口调用确定。
      :beforeDelete="beforeDelete"
    />
  </div>
</template>

<script>
import { getFolderList, getFileList } from '@/api/data'
import FolderTree from '_c/folder-tree'
export default {
  components: {
    FolderTree
  },
  data () {
    return {
      folderList: [],
      fileList: [],
      //下拉菜单的选项
      folderDrop: [
        {
          name: 'rename',
          title: '重命名'
        },
        {
          name: 'delete',
          title: '删除文件夹'
        }
      ],
      fileDrop: [
        {
          name: 'rename',
          title: '重命名'
        },
        {
          name: 'delete',
          title: '删除文件'
        }
      ]
    }
  },
  methods: {
    beforeDelete () {
      //实际要用接口，这里用Promise代替。
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          let error = new Error('error')
          //如果没有出错
          if (!error) {
            //执行成功
            resolve()
            //执行错误
          } else reject(error)
        }, 2000)
      })
    }
  },
  mounted () {
    Promise.all([getFolderList(), getFileList()]).then(res => {
      //文件夹数组
      this.folderList = res[0]
      //文件数组
      this.fileList = res[1]
    })
  }
}
</script>

<style lang="less">
.folder-wrapper{
  width: 300px;
}
</style>
```

在components文件夹下创建folder-tree.vue和index.js文件，index.js作为快捷路径文件。

```
//components/folder-tree/folder-tree.vue文件下
<template>
  <Tree :data="folderTree" :render="renderFunc"></Tree>
</template>

<script>
//引入方法
import { putFileInFolder, transferFolderToTree, expandSpecifiedFolder } from '@/lib/util'
import clonedeep from 'clonedeep'
export default {
  name: 'FolderTree',
  data () {
    return {
      //接收处理后的数据
      folderTree: [],
      //当前正在重命名的id的标识。
      currentRenamingId: '',
      currentRenamingContent: '',
      renderFunc: (h, { root, node, data }) => {
        //判断是文件夹对象还是文件对象
        const dropList = data.type === 'folder' ? this.folderDrop : this.fileDrop
        //循环dropList，给每条文件或文件夹对象添加选择名，name={item.name}相当于每条数据的id，就能知道点击的是哪条数据。
        const dropdownRender = dropList.map(item => {
          return (<dropdownItem name={item.name}>{ item.title }</dropdownItem>)
        })
        //获取到当前是文件夹对象+id还是文件对象+id，用来判断是什么对象在重命名。这里就是全局的命名标识。
        const isRenaming = this.currentRenamingId === `${data.type || 'file'}_${data.id}`
        return (
          <div class="tree-item">
            { data.type === 'folder' ? <icon type="ios-folder" color="#2d8cf0" style="margin-right: 10px;"/> : ''}
            {
              //如果当前正在重命名，就显示输入框。
              isRenaming
                ? <span>
                //这里输入框里的值就是当前的值，这里input事件就是重命名事件。
                  <i-input value={data.title} on-input={this.handleInput} class="tree-rename-input"></i-input>
                  //添加两个图标，一个是保存图标一个是取消图标。保存图标就把修改后保存到空间里的值取出来再复制。
                  <i-button size="small" type="text" on-click={this.saveRename.bind(this, data)}><icon type="md-checkmark" /></i-button>
                  <i-button size="small" type="text"><icon type="md-close" /></i-button>
                </span>
                //如果没有重命名就是显示名称
                : <span>{ data.title }</span>
            }
              //操作目录，给每个render添加dropdown下拉菜单。判断这里如果不为undefined就显示dropdown组件，如果是undefined就显示空字符串。
            {
              //如果正在重命名是不能有其他操作，这里下拉菜单就不显示。
              dropList && !isRenaming ? <dropdown placement="right-start" on-on-click={this.handleDropdownClick.bind(this, data)}>
              //触发显示下拉列表的元素，在JSX中标签名和原生html一样要加i-
                <i-button size="small" type="text" class="tree-item-button">
                  <icon type="md-more" size={12}/>
                </i-button>
                //slot添加list插槽，dropdownMenu提供下拉菜单显示。
                <dropdownMenu slot="list">
                //这里面是活的，根据父组件folderDrop来渲染出来。
                  { dropdownRender }
                </dropdownMenu>
              </dropdown> : ''
            }
          </div>
        )
      }
    }
  },
  props: {
    //接收父级传递的文件夹对象
    folderList: {
      type: Array,
      default: () => []
    },
    //接收父级传递的文件对象
    fileList: {
      type: Array,
      default: () => []
    },
    //文件夹操作列表，可以为空。判断时只用判断是否为Undefined
    folderDrop: Array,
    //文件操作列表
    fileDrop: Array,
    //接收钩子
    beforeDelete: Function
  },
  watch: {
    //处理异步的任务，如果传进来的是异步的任务就需要在这里做处理。
    folderList () {
      this.transData()
    },
    fileList () {
      this.transData()
    }
  },
  methods: {
    //父组件传给子组件数据，让子组件来做数据处理
    transData () {
      this.folderTree = transferFolderToTree(putFileInFolder(this.folderList, this.fileList))
    },
    //判断是文件夹还是文件
    isFolder (type) {
      return type === 'folder'
    },
     //删除的逻辑
    handleDelete (data) {
      //当前传入的数据
      const folderId = data.folder_id
      //判断是否是文件夹
      const isFolder = this.isFolder(data.type)
      //判断是文件夹还是文件在决定更新
      let updateListName = isFolder ? 'folderList' : 'fileList'
      //获取列表，要删除的是哪个列表。如果是文件夹就克隆文件夹，如果是文件就克隆文件
      let list = isFolder ? clonedeep(this.folderList) : clonedeep(this.fileList)
      //返回不删除的东西，删除完剩下的。
      list = list.filter(item => item.id !== data.id)
      this.$emit(`update:${updateListName}`, list)
      //视图还没有渲染完，这里需要调用$nextTick方法。$nextTick方法会在视图渲染完再执行函数。
      this.$nextTick(() => {
        //这里删除以后只是前端的删除，后端可能有各种原因没有删除，这样数据就会不一致，这里需要先调用接口在执行删除。
        expandSpecifiedFolder(this.folderTree, folderId)
      })
    },
    //原生传递的时候name会拼接在data后面
    handleDropdownClick (data, name) {
      if (name === 'rename') {
        //文件夹名称可能和文件名称相同，这里做一个处理，拼接起来让名字不一样。这里赋值的结果是判断类型是folder就folder+id如果好似file就file+id。
        this.currentRenamingId = `${data.type || 'file'}_${data.id}`
        //删除操作
      } else if (name === 'delete') {
        //弹出提示信息
        this.$Modal.confirm({
          title: '提示',
          content: `您确定要删除${this.isFolder(data.type) ? '文件夹' : '文件'}《${data.title}》吗？`,
          onOk: () => {
            //如果传进来钩子函数，钩子函数调用成功，执行删除的操作。
            this.beforeDelete ? this.beforeDelete().then(() => {
              //执行删除操作
              this.handleDelete(data)
            }).catch(() => {
              //如果调用钩子报错，就返回删除失败。
              this.$Message.error('删除失败')
              //删除的逻辑
            }) : this.handleDelete(data)
          }
        })
      }
    },
    //重命名事件
    handleInput (value) {
      //输入的值保存到一个空间里。
      this.currentRenamingContent = value
    },
    //
    updateList (list, id) {
      let i = -1
      let len = list.length
      while (++i < len) {
        let folderItem = list[i]
        //如果当前文件夹等于传进来的id值
        if (folderItem.id === id) {
          //修改name属性
          folderItem.name = this.currentRenamingContent
          //用splice把原数组删掉
          list.splice(i, 1, folderItem)
          break
        }
      }
      //返回修改后的数据
      return list
    },
    //重命名事件
    saveRename (data) {
      //当前点击的是哪个文件
      const id = data.id
      //当前点击的是什么类型
      const type = data.type
      //如果修改的是文件夹目录的话，需要遍历出所有的文件夹
      if (type === 'folder') {
        //不能直接修改原始数据，要拷贝一份。
        const list = this.updateList(clonedeep(this.folderList), id)
        //把修改后的数据传递给父组件
        this.$emit('update:folderList', list)
      } else {
        //如果是文件列表，获取到的是fileList的length
        const list = this.updateList(this.fileList, id)
        //把修改后的数据传递给父组件
        this.$emit('update:fileList', list)
      }
      //点保存后让全局标识为空，图标消失。
      this.currentRenamingId = ''
    }
  },
  mounted () {
    //使用钩子函数，让传进来就有数据
    this.transData()
  }
}
</script>

<style lang="less">
.tree-item{
  display: inline-block;
  width: ~"calc(100% - 50px)";
  height: 30px;
  line-height: 30px;
  & > .ivu-dropdown{
    float: right;
  }
  ul.ivu-dropdown-menu{
    padding-left: 0;
  }
  li.ivu-dropdown-item{
    margin: 0;
    padding: 7px 16px;
  }
  .tree-rename-input{
    width: ~"calc(100% - 80px)";
  }
}
</style>
```

功能介绍：如果删除文件夹但是文件夹下面还有其他文件，这个时候不想文件夹收缩起来。

```
//components/lib/util.js文件下
//这里传入两个参数，第一个参数是处理好的树状列表，第二个参数是哪个id。
export const expandSpecifiedFolder = (folderTree, id) => {
  return folderTree.map(item => {
    //如果遍历的是文件夹
    if (item.type === 'folder') {
      //当前的id是传进来的id
      if (item.id === id) {
        //就展开
        item.expand = true
      } else {
        //当前遍历的文件夹不是要展开的文件夹，就遍历它的children。当前有children并children有子节点
        if (item.children && item.children.length) {
          //做一个递归，遍历当前这个子节点
          item.children = expandSpecifiedFolder(item.children, id)
          //判断当前的节点，如果其中一个子节点有expand属性等于true的话就返回
          if (item.children.some(child => {
            return child.expand === true
          })) {
            //有一个子节点的expand为true，那么这个节点也为true。这个节点就展开。
            item.expand = true
          } else {
            //如果都没有为true就不展开。
            item.expand = false
          }
        }
      }
    }
    //要返回这个数组才能生效。
    return item
  })
}
```

## 文件上传前后端(Node.js)实现

1.Node.js服务
2.前端上传、下载
3.自行控制文件上传

### Node.js服务

配置需要，安装mysql数据库，在mysql中添加数据库名为fss的数据库。安装redis。
在根目录下创建file_storage文件夹，用来存放上传的文件。
服务器配置在fss-server/.env文件下。
在fss-server文件下使用npm run local启动服务。

### 前端上传、下载

在views下创建upload.vue文件

```
//views/upload.vue文件下
<template>
  <div>
  //添加upload组件，这是触发上传选择文件。upload必须必须添加一个:action，action的值就是文件上传的url。这个action是基础路径拼接上后端的请求路径。multiple为设置批量上传。upload组件有几个钩子函数，这里:before-upload在上传之前触发。:on-success钩子函数，当上传成功将执行的函数。
    <Upload ref="upload" :action="`${baseURL}/upload_file`" multiple :before-upload="beforeUpload" :on-success="handleSuccess" :show-upload-list="false">
      <Button icon="ios-cloud-upload-outline">Upload Files</Button>
    </Upload>
    //### 自行控制文件上传，有些时候文件要选择，如果能上传才让上传。
    <Button @click="handleUpload">上传吧</Button>
    //用来显示上传文件的列表，columns是自己定义的，filiList是后端获取的。
    <Table :columns="columns" :data="fileList"></Table>
    //展示text内容，这里通过showModel控制
    <Modal v-model="showModal">
      <div style="height: 300px; overflow: auto">
        {{ content }}
      </div>
    </Modal>
  </div>
</template>

<script>
//引入baseURL
import { baseURL } from '@/config'
import { getFilesList, getFile, deleteFile } from '@/api/data'
//引入配置好的模拟form表单验证
import { downloadFile } from '@/lib/util'
export default {
  data () {
    return {
      //挂在baseURL实例，用来拼接上传路径。
      baseURL,
      showModal: false,
      content: '',
      //要上传的文件对象
      file: {},
      columns: [
        //key相当于文件的id
        { title: '文件key', key: 'key' },
        { title: '文件名', key: 'file_name' },
        { title: '文件类型', key: 'file_type' },
        { title: '文件大小', key: 'file_size' },
        { title: '上传时间', key: 'createdAt' },
        {
          //如果在后端不设置有效期就是永久
          title: '存储有效期',
          key: 'storage_time',
          render: (h, { row }) => {
            return (
              //判断有没有有效期，根据有效期显示。
              <span>{ row.storage_time ? row.storage_time : '永久' }</span>
            )
          }
        },
        {
          title: '操作',
          key: 'handle',
          render: (h, { row }) => {
            return (
              <span>
                <i-button on-click={this.download.bind(this, row.key)}>下载</i-button>
                //如果文件类型包括text就能显示点击，不包括就禁用。
                <i-button disabled={!row.file_type.includes('text')} on-click={this.showFileContent.bind(this, row.key)}>显示内容</i-button>
                <i-button on-click={this.deleteFile.bind(this, row.key)}>删除</i-button>
              </span>
            )
          }
        }
      ],
      fileList: []
    }
  },
  methods: {
    //上传钩子函数，如果在这里返回一个false或者返回一个Promise将终止上传。
    beforeUpload (file) {
      //要上传的文件
      this.file = file
      //阻止上传
      return false
    },
    handleUpload () {
      this.$refs.upload.post(this.file)
    },
    //上传成功钩子函数，将执行此函数。
    handleSuccess () {
      //显示上传成功的信息
      this.$Message.success('文件上传成功')
      this.updateFileList()
      //上传成功，上传制空。
      this.file = null
    },
    //提供用户下载的函数
    download (key) {
      //调用模拟form表单方法，拼接基础路径，传递值。
      downloadFile({
        //发起下载请求
        url: `${baseURL}/get_file`,
        params: {
          key,
          type: 'download'
        }
      })
    },
    //控制点击显示text内容，这里发起请求，获取数据展示。
    showFileContent (key) {
      getFile({
        key,
        type: 'text'
      }).then(res => {
        this.content = res
        this.showModal = true
      })
    },
    //删除文件，发起请求删除文件。
    deleteFile (key) {
      deleteFile(key).then(res => {
        this.updateFileList()
      })
    },
    //发起获取文件列表请求
    updateFileList () {
      getFilesList().then(res => {
        //把数据赋值给filiList用来显示数据。
        this.fileList = res
      })
    }
  },
  mounted () {
    this.updateFileList()
  }
}
</script>

<style>

</style>
```

定义发起的各种请求

```
//api/data.js文件下
//获取文件列表请求
export const getFilesList = () => {
  return axios.request({
    url: 'get_file_list',
    params: {
      userId: 1
    },
    method: 'get'
  })
}
//下载文件请求
export const getFile = ({ key, type }) => {
  return axios.request({
    url: 'get_file',
    data: {
      key,
      type
    },
    method: 'post'
  })
}
//删除文件的请求
export const deleteFile = key => {
  return axios.request({
    url: 'delete_file',
    data: {
      key
    },
    method: 'delete'
  })
}
```

下载处理
后端做了download请求，前端也要做处理，前端如果要发起post请求，接口响应不会触发下载，为了安全它只会触发form表单的方式才能下载。这里封装一个方法模拟提交表单

```
//lib/uitl.js文件下
//第一个参数是提交的url，第二个参数是提交的残害苏
export const downloadFile = ({ url, params }) => {
  //创建一个表单
  const form = document.createElement('form')
  //设置一个action，要往哪里提交
  form.setAttribute('action', url)
  //要以什么方式提交
  form.setAttribute('method', 'post')
  //遍历传进来的参数
  for (const key in params) {
    //如果要提交一个值就要创建一个input标签
    const input = document.createElement('input')
    //设置input属性，在页面上不会看到
    input.setAttribute('type', 'hidden')
    //每一个表单元素都要有一个name元素，这样服务端form才能通过name取到，name就是参数里的key
    input.setAttribute('name', key)
    //value就是当前params里key对应的值
    input.setAttribute('value', params[key])
    //把input的节点放到form里
    form.appendChild(input)
  }
  //把表单放到body里里面
  document.body.appendChild(form)
  //触发form提交
  form.submit()
  //再把表单移除
  form.remove()
}
```

## Form表单

1.基础表单
2.动态组件
3.动态表单

### 基础表单

在views下创建form.vue文件

```
//iviews/form.vue文件下
<template>
  <div class="form-wrapper">
    <Form ref="form"
          :label-width="80"
          :model="formData"
          :rules="rules">
      <FormItem label="姓名"
      //prop使用规则，在label前面添加红星。
                prop="name">
        <Input v-model="formData['name']"></Input>
      </FormItem>
      <FormItem label="年龄">
        <Input v-model="formData['age']"></Input>
      </FormItem>
      <FormItem>
        <Button @click="handleSubmit"
                type="primary">提交</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
import { sentFormData } from '@/api/data'
//自定义表单验证，一定要调用callback()不然不会执行。
const validateName = (rule, value, callback ) => {
  if (value !== 'lison') {
    callback(new Error('Name error'))
  }else {
    callback()
  }
}
export default {
  data () {
    return {
      formData: {
        name: 'lison',
        age: 18
      },
      rules: {
        name: [
          { required: true, message: 'this name cannot be empty', trigger: 'blur' },
          //自定义表单验证
          { validator: validateName, trigger: 'blur'}
        ]
      }
    }
  },
  methods: {
    handleSubmit () {
      this.$refs.form.validate(valid => {
        if (valid) {
          sentFormData(this.formData).then(res => {
            console.log(res)
          })
        }
      })
    }
  }
}
</script>

<style lang="less">
.form-wrapper {
  padding: 20px;
}
</style>
```

添加请求

```
export const sentFormData = (data) => {
  return axios.request({
    url,
    data,
    method: 'post'
  })
}
```

### 动态组件

有时候需要分发表单页，表单页又不是我们来制作的，需要给给其他非技术人员通过拖拽来实现表单页。
在components下创建form-group文件夹，文件夹下创建fomr-group.vue和index.js。index.js用来做路径引导。

```
//views/form.vue文件下
<template>
  <div class="form-wrapper">
    <form-group :list="formList" :url="url"></form-group>
  </div>
</template>

<script>
import FormGroup from '_c/form-group'
export default {
  components: {
    FormGroup
  },
  data () {
    return {
      url: '/data/formData',
      //这里实际是从后端获取的数据
      formList: [
        {
          //姓名输入框
          name: 'name',
          type: 'i-input',
          value: '',
          label: '姓名',
          rule: [
            { required: true, message: 'not', trigger: 'blur'}
          ]
        },
        {
          //选择范围
          name: 'range',
          type: 'slider',
          //值的范围
          value: [10, 40],
          //这里填true，就能自由控制最小与最大。
          range: true,
          label: '范围'
        },
        {
          //级联选择器
          name: 'sex',
          type: 'i-select',
          value: '',
          label: '性别',
          children: {
            type: 'i-option',
            list: [
              { value: 'man', title: '男' },
              { value: 'woman', title: '女' }
            ]
          }
        },
        {
          //单选
          name: 'educatioin',
          type: 'radio-group',
          value: 1,
          label: '学历',
          children: {
            type: 'radio',
            list: [
              { label: 1, title: '本科' },
              { label: 2, title: '硕士' },
              { label: 3, title: '博士' }
            ]
          }
        },
        {
          //多选
          name: 'skill',
          type: 'checkbox-group',
          value: [],
          label: '技能',
          children: {
            type: 'checkbox',
            list: [
              { label: 1, title: 'Vue' },
              { label: 2, title: 'Node.js' },
              { label: 3, title: 'JS' }
            ]
          }
        },
        {
          //开关按钮
          name: 'inWork',
          type: 'i-switch',
          value: true,
          label: '在职'
        }
      ]
    }
  },
  methods: {

  }
}
</script>

<style lang="less">
.form-wrapper {
  padding: 20px;
}
</style>
```

```
//components/form-group.vue文件下
<template>
  <Form ref="form"
        //验证表单是否有值，修复有值还是显示错误信息的bug
        v-if="Object.keys(valueList).length"
        //设置label宽度
        :label-width="labelWidth"
        //传入当前的数据
        :model="valueList"
        //当前的规则
        :rules="rules">
        //所有的表单项都要包裹在FormItem里面，所以在这里循环。
    <FormItem v-for="(item, index) in list"
              :label="item.label"
              //把label的值放在左边与input同一行。
              label-position="left"
              //绑定校验规则
              :prop="item.name"
              //绑定获取到的后端错误信息会展示为字符串的信息
              :error="errorStore[item.name]"
              //使用.native绑定原生事件，可能在其他组件里绑定事件函数不一样就会报错。
              @click.native="handleFocus(item.name)"
              //一个页面可能不止一个表单，如果key重复就会报错，这里用唯一的_uid来拼接组成唯一的索引。
              :key="`${_uid}_${index}`">
              //如果有很多的表单元素的话这里就要写很多的input表单用很多v-if来判断,这样组件会很大。这里使用vue的内置组件component动态组件，它有一个必填值:is，代表它是什么组件。你要什么标签在这里面填入什么标签就行。
      <component :is="item.type"
              //:range用来控制范围
                 :range="item.range"
                  //提交到服务器的数据
                 v-model="valueList[item.name]">
                 //判断级联选择器有没有子级再渲染。
        <template v-if="item.children">
        //渲染级联选择器
          <component v-for="(child, i) in item.children.list"
                      //使用i-select标签
                     :is="item.children.type"
                     :key="`${_uid}_${index}_${i}}`"
                     //子级使用单选多选
                     :label="child.label"
                     :value="child.value">{{child.title}}</component>
        </template>
      </component>
    </FormItem>
    <FormItem>
      <Button @click="handleSubmit"
              type="primary">提交</Button>
      <Button @click="handleReset">重置</Button>
    </FormItem>
  </Form>
</template>

<script>
import clonedeep from 'clonedeep'
import { sentFormData } from '@/api/data'
export default {
  name: 'FormGroup',
  data () {
    return {
      initValueList: [],
      rules: {},
      valueList: {},
      //获取的后端错误验证信息
      errorStore: {}
    }
  },
  props: {
    //父组件传递过来的数组
    list: {
      type: Array,
      default: () => []
    },
    labelWidth: {
      type: Number,
      default: 100
    },
    url: {
      type: String,
      require: true
    }
  },
  watch: {
    //监听list变化
    list () {
      this.setInitValue()
    }
  },
  methods: {
    //处理表单数据
    setInitValue () {
      //存放遍历出来的规则
      let rules = {}
      //所以值的对象
      let valueList = {}
      //获取到初始传进来的value值，方便做重置操作。
      let initValueList = {}
      let errorStore = {}
      //当前表单对象rule写到rules对象上
      this.list.forEach(item => {
        //item.name作为属性名，rule作为属性值
        rules[item.name] = item.rule
        valueList[item.name] = item.value
        initValueList[item.name] = item.value
        //为空不报错
        errorStore[item.name] = ''
      })
      //
      this.rules = rules
      this.valueList = valueList
      this.initValueList = initValueList
      this.errorStore = errorStore
    },
    //重置
    handleReset () {
      //深克隆一遍传过来的值再赋值，变为初始值。
      this.valueList = clonedeep(this.initValueList)
    },
    //提交
    handleSubmit () {
      //验证表单规则
      this.$refs.form.validate(valid => {
        if (valid) {
          sentFormData({
            //提交的地址
            url: this.url,
            //提交的值
            data: this.valueList
          }).then(res => {
            this.$emit('on-submit-success', res)
          }).catch(err => {
            this.$emit('on-submit-error', err)
            //修改密码时需要和后端先比对密码，如果错误就显示错误信息，提高使用者体验。
            for (let key in err) {
              this.errorStore[key] = err[key]
            }
          })
        }
      })
    },
    handleFocus(name) {
      this.errorStore[name] = ''
    }
  },
  mounted () {
    //表单一开始传进来就要做一个处理
    this.setInitValue()
  }
}
</script>
```

修改请求时携带url

```
//lib/data.js
export const sentFormData = ({ url, data}) => {
  return axios.request({
    url,
    data,
    method: 'post'
  })
}
```

## 权限控制

1.简单权限控制
2.页面级别
3.组件级别

### 简单权限控制

在iviewadmin中路由列表里给每个需要权限控制的路由列表，在meta里配置一个access这个权限字段，它是一个数组。表示哪个用户组可以浏览，里面可以是字符串或者数值。如果这个用户不是super_admin，那么这个用户就看不到这个页面，左侧菜单也不会渲染，如果这级访问不了那么它的子级也访问不了。如果不需要访问权把access这个字段删除即可。

```  
{
    path: '/component',
    name: 'component',
    component: Layout,
    meta: {
      access: ['super_admin'],
      icon: 'md-funnel',
      showAlways: true,
      title: '二级-2'
    }，
    children: [
      {
        path: 'table',
        name: 'table_page',
        meta: {
          title: '表格'
        },
        component: () => import('@/views/table.vue')
      }
    ]
  }
```

配置完路由列表还需要配置路由守卫，在路由守卫里面进行拦截。通过一个方法canTurnTo，这个方法有三个参数。第一个参数是要访问的name传进去，第二个参数是把access用户的权限字段传进去，权限字段是登录之后通过接口获得的权限字段，它是一个列表或数组。第三个参数是routes就是路由列表。这个方法会把传递的权限字段和路由进行匹配，如果当前用户有权限

```
import Vue from 'vue'
import Router from 'vue-router'
import routes from './routers'
import store from '@/store'
import iView from 'iview'
import { setToken, getToken, canTurnTo, setTitle } from '@/libs/util'
import config from '@/config'
const { homeName } = config

Vue.use(Router)
const router = new Router({
  routes,
  mode: 'history'
})
const LOGIN_PAGE_NAME = 'login'

const turnTo = (to, access, next) => {
  if (canTurnTo(to.name, access, routes)) next() // 有权限，可访问
  else next({ replace: true, name: 'error_401' }) // 无权限，重定向到401页面
}

router.beforeEach((to, from, next) => {
  iView.LoadingBar.start()
  const token = getToken()
  if (!token && to.name !== LOGIN_PAGE_NAME) {
    // 未登录且要跳转的页面不是登录页
    next({
      name: LOGIN_PAGE_NAME // 跳转到登录页
    })
  } else if (!token && to.name === LOGIN_PAGE_NAME) {
    // 未登陆且要跳转的页面是登录页
    next() // 跳转
  } else if (token && to.name === LOGIN_PAGE_NAME) {
    // 已登录且要跳转的页面是登录页
    next({
      name: homeName // 跳转到homeName页
    })
  } else {
    if (store.state.user.hasGetInfo) {
      turnTo(to, store.state.user.access, next)
    } else {
      store.dispatch('getUserInfo').then(user => {
        // 拉取用户信息，通过用户权限和跳转的页面的name来判断是否有权限访问;access必须是一个数组，如：['super_admin'] ['super_admin', 'admin']
        if (canTurnTo(to.name, user.access, routes)) next() //有权限，可访问
        else next({ replace: true, name: 'error_401}) //无权限，重定向到401页面
      })
    }
  }
})

router.afterEach(to => {
  setTitle(to, router.app)
  iView.LoadingBar.finish()
  window.scrollTo(0, 0)
})

export default router
```

### 页面级别

这种方法好处就是后端只用返回一个用户组列表其他就交给前端在路由列表里配置就行了。但是如果有上百种用户权限字段的话就不适用了。这种方式是没有权限跳转到404页面没有页面也跳转到404页面，可以设置一个401页面，判断没有权限跳转到401页面，没有页面跳转到404页面
第二种方式可以解决这个问题，通过服务端返回一个可访问的路由列表，然后做一个过滤，动态的过滤挂在到路由实例上。有一个弊端，就是路由列表都需要一个name属性。
在`router/router`文件下，把login路由和404路由单独暴露出来，数组名为`routes`。其他路由在`routerMap`下。

```
// 后端代码 user.js文件下

var express = require('express');
var router = express.Router();
const jwt = require('jsonwebtoken')

/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('respond with a resource');
});

router.get('/getUserInfo', function(req, res, next) {
  res.send('success')
})

router.get('/authorization', (req, res, next) => {
  const userName = req.userName
  res.send({
    code: 200,
    mes: 'success',
    data: {
      token: jwt.sign({ name: userName }, 'abcd', {
        expiresIn: '1d'
      }),
      rules: {
        page: {
          home: true,
          home_index: true,
          about: true,
          argu: true,
          count_to: true,
          menu_page: true,
          upload: true,
          form: true,
          folder_tree: true,
          table_page: true,
          render_page: true,
          split_pane: true,
          parent: true,
          child: true,
          named_view: true,
          store: true,
          main: true
        },
        component: {
          edit_button: true,
          publish_button: false
        }
      }
    }
  })
})

module.exports = router;

```

前端

```
//router/index.js文件下
import Vue from 'vue'
import VueRouter from 'vue-router'
import { routes } from './router'
import store from '@/store'
import { setTitle, setToken, getToken } from '@/lib/util.js'

Vue.use(VueRouter)

const router = new VueRouter({
  routes
})

// 模拟一个登录，实际是通过接口来判断的。
const HAS_LOGINED = false

router.beforeEach((to, from, next) => {
  to.meta && setTitle(to.meta.title)
  //权限控制
  //获取token
  const token = getToken()
  if (token) {
    //这里说明这里已经登录，判断是否获取权限列表
    if (!store.state.router.hasGetRules) {
      //调用接口获取权限列表，这里使用authorization方法
      store.dispatch('authorization').then(rules => {
        //使用dispatch方法传递合并方法携带rules参数
        store.dispatch('concatRoutes', rules).then(routers => {
          //router是路由实例，调用addRoutes方法动态挂在。
          //这里修改store里的routers可能报错，只用关闭vuex的严格模式即可。
          router.addRoutes(routers)
          //调用next，展开to。因为路由列表可能没有挂在完，直接使用next可能会有错误，拆分到对象里面，使用replace设为true。就是要访问的路径使用替换的形式。
          next({ ...to, replace: true })
        }).catch(() => {
          //出现异常直接跳转到登录页
          next({ name: 'login' })
        })
      })
    } else {
      //如果已经获取过权限列表直接放行
      next()
    }
  } else {
    if (to.name === 'login') next()
    else next({ name: 'login' })
  }
})

export default router
```

在store下的module文件夹下创建router.js，这个文件的作用是将后端传递的权限字段进行过滤匹配。
在store/index.js文件下引入router并挂载到modules里

```
  modules: {
    user,
    router
  }
```

在store里的module文件夹下里的user.js中authorization方法，它会调用接口。只用在resolve中使用`res.data.rules.page`就能调用后端权限列表。

```
//store/module/router.js文件下
import { routes, routerMap } from '@/router/router'
const state = {
  //routers就是最后挂载到路由实例上的路由
  //引入所有人都能访问的路由页
  routers: routes,
  //存储权限状态，默认为false。获取了用户权限列表后设为true
  hasGetRules: false
}
//合并路由列表的方法
const mutations = {
  CONCAT_ROUTES (state, routerList) {
    //合并方法，权限列表能访问的放前面，所有用户能访问的放后面合并出用户能访问的所有列表。这里如果把所有用户能访问的放前面，因为404页面在routes里面，所有页面都会变成404。
    state.routers = routerList.concat(routes)
    //返回状态，说明已经获取了用户权限。
    state.hasGetRules = true
  }
}

//筛选路由列表的递归方法
const getAccessRoutreList = (routes, rules) => {
  return routes.filter(item => {
    //在item上访问rules的name如果return是true就会添加进去，如果return的是false就不能访问。
    if (rules[item.name]) {
      //路由列表是有嵌套的，调用自身，如果有children再次调用递归，把自身传递过去。
      if (item.children) item.children = getAccessRoutreList(item.children, rules)
      return true
    } else return false
  })
}

const actions = {
  //定义合并函数，第一个参数是一个对象这里获取到commit，接收rules。
  concatRoutes ({ commit }, rules) {
    //返回Promise
    return new Promise((resolve, reject) => {
      try {
        //定义一个空的路由列表
        let routerList = []
        //entries方法可以把对象转为二位数组，返回属性名和属性值，再用every来遍历，选择返回数组的第二个元素就是属性值。如果返回的是true说明这个数组里都是true就不需要使用递归去筛选。
        if (Object.entries(rules).every(item => item[1])) {
          //这个路由列表都能访问
          routerList = routerMap
        } else {
          //如果没有返回true，这里需要筛选出能够访问的路由列表。传入路由列表和后端的权限列表
          routerList = getAccessRoutreList(routerMap, rules)
        }
        //使用commit对state进行修改，提交一个方法CONCAT_ROUTES传递routerList值
        commit('CONCAT_ROUTES', routerList)
        //调用resolve返回，这样就不用再去state里面去取了
        resolve(state.routers)
      }catch (err) {
        //接收报错
        reject(err)
      }
    })
  }
}
//导出模块
export default {
  state,
  mutations,
  actions
}
```

### 组件级别

```
//store/module/user.js页面下
import { login, authorization } from '@/api/user'
import { setToken } from '@/lib/util'

const state = {
  // 在模块中使用
  userName: 'Lison',
  rules: {}
}
const mutations = {
  // 第一个参数就是要作用的名称，第二个参数是传过来的值
  SET_USER_NAME (state, params) {
    state.userName = params
  },
  //定义mutation的组件
  SET_RULES(state, rules) {
    state.rules = rules
  }
}
const actions = {
  // 这里的第一个参数是提交，第二个参数是这里的state实例，第三个参数是store下面的state实例，可以直接操作。第四个参数是action实例的提交方法。
  updateUserName ({ commit, state, rootState, despatch }) {
    // 操作state下的appName
    // rootState.appName
  },
  //通过载荷的形式传递userName和pssword。
  login ({ commit }, { userName, password }) {
    return new Promise((resolve, reject) => {
      //这里返回一个promise，通过then的方式接收
      login({ userName, password }).then(res => {
        // 判断返回的状态码是否是200和token是否为空。
        if (res.code === 200 && res.data.token) {
          // 成功返回token就保存token
          setToken(res.data.token)
          resolve()
        } else {
          //如果判断失败就就是显示一个错误信息
          reject(new Error('错误'))
        }
        //返回的错误信息
      }).catch(error => {
        reject('error')
      })
    })
  },
  authorization ({ commit }, token) {
    return new Promise((resolve, reject) => {
      authorization().then(res => {
        if (parseInt(res.code) == 401) {
          reject(new Error('token error'))
        } else {
          setToken(res.data.token)
          resolve(res.data.rules.page)
          //提交SET_RULES方法，传递component
          commit('SET_RULES', res.data.rules.component)
        }
      }).catch(error => {
        reject(error)
      })
    })
  },
  logout () {
    setToken('')
  }
}

export default {
  state,
  mutations,
  actions,
  module: {

  }
}
```

在home也中展示使用。有些权限只能编辑，发布需要更高的权限。

```
<template>
  <div class="home">
    <Row :gutter="10"
         class="blue">
      <i-col :md="6" :sm="12" :xs="24"></i-col>
      <i-col :md="6" :sm="12" :xs="24"></i-col>
      <i-col :md="6" :sm="12" :xs="24"></i-col>
      <i-col :md="6" :sm="12" :xs="24">{{ rules }}</i-col>
    </Row>
    //这里通过修改后端传递的布尔值来判断是否渲染按钮
    <Button v-if="rules.edit_button">编辑</Button>
    <Button v-if="rules.publish_button">发布</Button>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import { getUserInfo } from '@/api/user.js'
//引入权限列表组件mapState
import { mapState, mapActions } from 'vuex'

export default {
  name: 'Home',
  components: {
    HelloWorld
  },
  props: {
    food: {
      type: String,
      default: 'apple'
    }
  },
  beforeRouterEnter (to, from, next) {
    //跳转的页面此时this还没有加载出来，是不能用this的
    next(vm => {
      //这个vm就是组件的实例，这样就能在里面使用this了
    })
  },
  beforeRouterLeave (to, from, next) {
    //将要离开页面时调用钩子方法
    const leave = confirm('您确定要离开吗？')
    if (leave) next()
    else next(false)
  },
  methods: {
    ...mapActions([
      'logout'
    ]),
    handleClick (type) {
      // this.$router.go(-1)
      if (type === 'back') this.$router.back()
      else if (type === 'push') {
        this.$router.push({
          name: 'parent',
          query: {
            name: 'lison'
          }
        }
        )
      } else if (type === 'replace') {
        // replace替换，把当前的浏览历史替换成parent这个页面，之后再做回退会回到到parent
        this.$router.replace({
          name: 'parent'
        })
      }
    },
    getInfo () {
      getUserInfo({ userId: 21 }).then(res => {
        console.log('res: ', res)
      })
    },
    handleLogout () {
      this.logout()
      this.$router.push({
        name: 'login'
      })
    }
  },
  computed: {
    //使用权限列表组件
    ...mapState ({
      rules: state => state.user.rules
    })
  }
}
</script>

<style lang="less">
.home {
  .ivu-col {
    height: 50px;
    margin-top: 10px;
    background-color: pink;
    background-clip: content-box;
  }
  .blue {
    .ivu-col {
      background: blue;
      background-clip: content-box;
    }
  }
}
</style>
```

## icon组件

1.Unicode&Symbol
2.font-class
3.封装单色和多色icon组件

### Unicode&Symbol

使用阿里的iconfont，这里利用下载方法使用，下载后放到assets/font目录下
在App.vue里使用，引入的样式是全局的。

```
//App.vue文件下
<style lang="less">
html,
body {
  height: 100%;
}
body {
  margin: 0;
}
//定义Unicode全局样式
@font-face {
  font-family: "./assets/font/iconfont";
  src: url("./assets/font/iconfont.eot");
  src: url("./assets/font/iconfont.eot?#iefix") format("embedded-opentype"),
    url("./assets/font/iconfont.woff") format("woff"),
    url("./assets/font/iconfont.svg#iconfont") format("svg");
}
.iconfont {
  font-family: "iconfont" !important;
  font-size: 16px;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -webkit-text-stroke-width: 0.2px;
  -moz-osx-font-smoothing: grayscale;
}
//需要在main里面引入import '@/assets/font/iconfont.js'，定义Symbol全局样式。
.icon-svg {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>

```

在viwes下创建icon_page.vue文件

```
<template>
  <div>
  //Unicode使用
  //md-menu为菜单的字体，size为fontsize，color颜色属性。
    <Icon type="md-menu" :size="50" color="pink"/>
  //使用字体图标
    <i class="iconfont smile">&#xe619;</i>
    <i class="iconfont">&#xe618;</i>
    //使用svg
    <svg class="iconfont-svg" aria-hidden="true" style="font-size: 70px;">
      <use xlink:href="#icon-zhaoxiangji-"></use>
    </svg>
    //font-class样式，在main里引入import '@/assets/font/iconfont.css'。
    <!-- <i class="iconfont icon-shouye4"></i> -->
    <Icon custom="iconfont icon-shouye4"/>
    //单色组件
    <icon-font icon="shouye3" :size="50" color="blue"></icon-font>
    //多色组件
    <icon-svg icon="zhuye-" :size="100"></icon-svg>
  </div>
</template>

<script>
export default {
  //
}
</script>

<style>
.smile{
  color: red;
}
</style>

```

### 封装单色和多色icon组件

单色组件
在components下创建icon-font文件夹，文嘉下创建icon-font.vue和index.js，index做引导文件。
可以在全局使用，在main.js里引入`import IconFont from '_c/icon-font'`，并且挂在到实例上`Vue.component('icon-font', IconFont)`

```
//components/icon-font/icon-font.vue文件下
<template>
  <i :class="classes" :style="styles"></i>
</template>

<script>
export default {
  props: {
    icon: {
      type: String,
      default: ''
    },
    size: {
      type: Number,
      default: 12
    },
    color: {
      type: String,
      default: '#515a6e'
    }
  },
  computed: {
    classes () {
      return [
        'iconfont',
        `icon-${this.icon}`
      ]
    },
    styles () {
      return {
        color: this.color,
        fontSize: `${this.size}px`
      }
    }
  }
}
</script>

<style>

</style>

```

多色组件
在components下创建icon-svg文件夹，文嘉下创建icon-svg.vue和index.js，index做引导文件。
在全局使用需要在main里引入`import IconSvg from '_c/icon-svg'`，挂在实例`Vue.component('icon-svg', IconSvg)`

```
//components/icon-svg/icon-svg.vue文件下
<template>
  <svg class="iconfont-svg" aria-hidden="true" :style="style">
    <use :xlink:href="iconName"></use>
  </svg>
</template>

<script>
export default {
  name: 'IconSvg',
  props: {
    icon: {
      type: String,
      default: ''
    },
    size: {
      type: Number,
      default: 20
    }
  },
  computed: {
    iconName () {
      return `#icon-${this.icon}`
    },
    style () {
      return {
        fontSize: `${this.size}px`
      }
    }
  }
}
</script>
```

## 大数据量性能优化

1.列表优化
2.大型表单优化
3.表格优化
在vue中修改视图，就会发生相应的变化。这个便捷的功能也会带来些问题，我们无法控制vue渲染视图的时机，当数据量非常的时候这个问题尤为明显。因为我们的数据放到vue中它会做一个遍历遍历，给每一个属性添加一个getter和setter，这样非常冗余。每一次一个数据变动会牵扯到很多的数据和视图的变化。

### 列表优化

select下拉菜单如果数据量超过几百会有延迟卡顿，现在解决这一问题。这里用到`vue-virtual-scroll-list`插件，这是一个组件。
这个组件的实现原理：当滚动下拉菜单的时候，他会计算上下的尺寸，比如列表全量渲染出来，现在要渲染1000条，它会计算出1000条数据的高度，算出总高度。现在只渲染11条，计算出11条的高度。再用总的条数高度减去11条高度就是未渲染的上面的高度和下面的高度。它是用padding-ton和padding-bottom把未渲染的做一个填充。那么就和全量渲染的条数一样。它内部还会做一个优化，这样渲染速度就会快很多。
在views下创建optimize.vue文件

```
//views/optimize.vue文件
<template>
  <div class="box">
    <!-- <Select v-model="selectData"
            style="width:200px">
            //插件使用方法，使用这个组件把要循环的数据包住。在组件上要给它设两个值，第一个是:size代表每行的高度。第二个是:remain代表需要渲染多少条。
      <virtual-list :size="30"
                    :remain="6">
        <Option v-for="item in list"
                :value="item.value"
                :key="item.value">{{ item.label }}</Option>
      </virtual-list>
    </Select>
    <Select v-model="selectData"
            style="width:200px">
      <virtual-list :size="30"
                    :remain="6">
        <Option v-for="item in list"
                :value="item.value"
                :key="item.value">{{ item.label }}</Option>
      </virtual-list>
    </Select> -->
    //在checkbox中使用
    <!-- <CheckboxGroup v-model="checkedArr">
      <virtual-list :size="30"
                    :remain="6">
        <p v-for="item in list"
           :key="`check${item.value}`"
           style="height: 30px">
          <Checkbox :label="item.value">
            <Icon type="logo-twitter"></Icon>
            <span>{{item.label}}</span>
          </Checkbox>
        </p>
      </virtual-list>
    </CheckboxGroup> -->
  </div>
</template>

<script>
import { doCustomTimes } from '@/lib/tools'
//引入插件
import VirtualList from 'vue-virtual-scroll-list'
export default {
  components: {
    VirtualList
  },
  data () {
    return {
      list: [],
      selectData: 0,
      checkedArr: []
    }
  },
  mounted () {
    //定义这个作用域接收名
    let list = []
    //使用之前在lib/tools.js文件下定义了一个doCustomTimes()能够将一个指定参数调用指定次数。
    doCustomTimes(1000, (index) => {
      list.push({
        value: index,
        label: `select${index}`
      })
    })
    //循环出的数据赋值给全局list
    this.list = list
  }
}
</script>

<style>
</style>

```

### 大型表单的优化

由于数据量太大会影响视图渲染，所以不要把所有的数据都绑到一个组件上，在外面做一个循环，一个组件只传入一个表单元素的数据。
在components下创建form-single文件夹，在文件夹下创建form-single.vue和index.js。index文件作为引导路径文件。

```
//components/form-single/fomr-single.vue文件下
<template>
  //把form-group.vue里的代码复制过来，这里把Object.key(valueList).length改为config，如果有config才能显示表单。现在只做一条表单验证，就把Form放在formSingle里，如果放到外面就没法做了。
  <Form ref="form" v-if="config" :label-width="100" :model="valueData" :rules="ruleData">
  //一个form组件就渲染一个formItem就不需要循环了。
    <FormItem
    //这里所有的item都换成config，因为这里是通过config传进来的
      :label="config.label"
      label-position="left"
      :prop="config.name"
      :error="errorStore[config.name]"
      :key="`${_uid}`"
      @click.native="handleFocus(config.name)"
    >
      <component :is="config.type" :range="config.range" v-model="valueData[config.name]">
        <template v-if="config.children">
          <component
            v-for="(child, i) in config.children.list"
            :is="config.children.type"
            :key="`${_uid}_${i}`"
            :label="child.label"
            :value="child.value">{{ child.title }}</component>
        </template>
      </component>
    </FormItem>
  </Form>
</template>

<script>
export default {
  name: 'FormSingle',
  props: {
    //接收传递的值
    valueData: {
      type: Object,
      default: () => ({})
    },
    ruleData: {
      type: Object,
      default: () => ({})
    },
    errorStore: {
      type: Object,
      default: () => ({})
    },
    //当前表单元素的配置
    config: Object
  },
  methods: {
    handleFocus (name) {
      this.errorStore[name] = ''
    },
    validate (callback) {
    //调用组件实例的validate方法
      //在外面传入一个回调函数valid，如果验证通过就会传一个true
      this.$refs.form.validate(valid => {
        //执行传进来的回调函数
        callback(valid)
      })
    }
  }
}
</script>
```

这里把form.vue下formList里的数据抽离出来，在mock/response下创建form-data.js。放到form-data.js里。在form页里引入。

```
//viwe/form文件下
<template>
  <div class="form-wrapper">
    <!-- <form-group :list="formList" :url="url"></form-group> -->
    <Button @click="handleSubmit" type="primary">提交</Button>
    <Button @click="handleReset">重置</Button>
    //循环数据
    <form-single
    //在v-for循环出来的组件上，使用ref获取到的是一个数组。
      ref="formSingle"
      v-for="(item, index) in formList"
      :key="`form_${index}`"
      //用来传当前配置的数据
      :config="item"
      //用来传递数据，在form-single.vue中使用，作为数据源。
      :value-data="valueData"
      //用来传递数据，在form-single.vue中使用，作为规则。
      :rule-data="ruleData"
      //用来传递数据，在form-single.vue中使用，作为错误提示。
      :error-store="errorStore"
    ></form-single>
  </div>
</template>

<script>
import FormGroup from '_c/form-group'
//引入formSingle组件
import FormSingle from '_c/form-single'
//引入抽离的数据
import formData from '@/mock/response/form-data'
import clonedeep from 'clonedeep'
import { sentFormData } from '@/api/data'
export default {
  components: {
    FormGroup,
    FormSingle
  },
  data () {
    return {
      //请求数据抽离的数据
      url: '/data/formData',
      formList: formData,
      //需要传递的值
      valueData: {},
      ruleData: {},
      errorStore: {},
      initValueData: {}
    }
  },
  methods: {
    //form-group里的重置和校验
    handleSubmit () {
      let isValid = true
      //这里是一个数组，这里每一个元素是循环的组件实例。
      this.$refs.formSingle.forEach(item => {
        //这里使用form-single里的validate方法，这里的值就是form-single里传递的值。
        item.validate(valid => {
          if (!valid) isValid = false
        })
      })
      //通过了验证再向服务器发起请求
      if (isValid) {
        sentFormData({
          url: this.url,
          data: this.valueData
        }).then(res => {
          this.$emit('on-submit-success', res)
        }).catch(err => {
          this.$emit('on-submit-error', err)
          for (let key in err) {
            this.errorStore[key] = err[key]
          }
        })
      }
    },
    handleReset () {
      this.valueData = clonedeep(this.initValueData)
    }
  },
  mounted () {
    let valueData = {}
    let ruleData = {}
    let errorStore = {}
    let initValueData = {}
  //遍历formData，如果数据是异步的可以在watch中操作。
    formData.forEach(item => {
      //遍历出需要传递的值，赋值给全局变量。
      valueData[item.name] = item.value
      ruleData[item.name] = item.rule
      errorStore[item.name] = ''
      //循环出来的初始值，用name做键。
      initValueData[item.name] = item.value
    })
    this.valueData = valueData
    this.ruleData = ruleData
    this.errorStore = errorStore
    //初始化的值
    this.initValueData = initValueData
  }
}
</script>

<style lang="less">
.form-wrapper{
  padding: 20px;
}
</style>
```

## 多Tab页开发

1.根据路由列表生成菜单
2.多标签实现
3.菜单、URL和标签联动

```
//views/layout.vue文件下
<template>
  <div class="layout-wrapper">
  //左侧菜单栏过多给超出的菜单添加滚动
    <Layout class="layout-outer">
      <Sider :width="200" collapsible hide-trigger reverse-arrow v-model="collapsed" class="sider-outer">
      //这里把数据换成过滤出来的数据，并且需要在side-menu.vue和re-rebmenu.vue里更改，因为是从路由元信息里获取的，要把展示的数据从item.title更改为item.meta.title。
        <side-menu :collapsed="collapsed" :list="routers"></side-menu>
      </Sider>
      <Layout>
      //去除多余的padding。
        <Header class="header-wrapper">
          <Icon :class="triggerClasses" @click.native="handleCollapsed" type="md-menu" :size="32"/>
        </Header>
        <Content class="content-con">
          <div>
          //:value为当前点击的路由对象，会在url里显示。这里的type是和name绑定的，这里定义一个方法，把params和query里的信息统统包含进去，这里使用在util.js中定义的方法拼接一个字符串。这里实现点击左侧菜单栏的路由在导航栏高亮。
          //handleClickTab事件为点击tab发生相应的跳转
            <Tabs type="card" @on-click="handleClickTab" :value="getTabNameByRoute($route)">
              //:label即在标签上显示的文字,这里添加一个rander函数用来做关闭标签页。:name即能知道选中的是哪一个
              <TabPane :label="labelRender(item)" :name="getTabNameByRoute(item)" 
              //这里的:key不应该简单的存一个name，如果要做动态路由还要带参数，这里路由改成params，区别在于后面动态的参数，name值相同就没办法区分。需要把name，params，query都存进去到一个对象里。
              v-for="(item, index) in tabList" :key="`tabNav${index}`"></TabPane>
            </Tabs>
          </div>
          <div class="view-box">
            <Card shadow class="page-card">
              <router-view/>
            </Card>
          </div>
        </Content>
      </Layout>
    </Layout>
  </div>
</template>

<script>
import SideMenu from '_c/side-menu'
import { mapState, mapMutations, mapActions } from 'vuex'
import { getTabNameByRoute, getRouteById } from '@/lib/util'
export default {
  components: {
    SideMenu
  },
  data () {
    return {
      collapsed: false,
      //如果想在template里面用需要挂载到data上面，因为如$router这些变量都是挂载到vue实例上的，这里必须挂载到vue实例上才能使用
      getTabNameByRoute
    }
  },
  computed: {
    triggerClasses () {
      return [
        'trigger-icon',
        this.collapsed ? 'rotate' : ''
      ]
    },
    //计算属性，
    ...mapState({
      //用来渲染的列表数组
      tabList: state => state.tabNav.tabList,
      //过滤出的路由列表用来做左侧菜单栏，这里再次过滤不想让登录和404页面出现在左侧菜单栏。
      routers: state => state.router.routers.filter(item => {
        return item.path !== '*' && item.name !== 'login'
      })
    })
  },
  methods: {
    ...mapActions([
      'handleRemove'
    ]),
    handleCollapsed () {
      this.collapsed = !this.collapsed
    },
    handleClickTab (id) {
      //这里跳转不能用name跳转，因为动态参数的name值都相同
      //这里根据id获得路由对象，这个方法是在util.js中定义的拼接函数。里面包含了parmas和query参数。
      let route = getRouteById(id)
      this.$router.push(route)
    },
    //关闭标签的点击事件，触发点击事件的时候会冒泡，所以这里传入全局事件。
    handleTabRemove (id, event) {
      //组织冒泡
      event.stopPropagation()
      //使用定义tabNav.js中的在移除方法
      this.handleRemove({
        //传入要删除的id
        id,
        //传入当前的路由
        $route: this.$route
      }).then(nextRoute => {
        //关闭成功获取到要跳转的路由的信息并跳转。
        this.$router.push(nextRoute)
      })
    },
    //关闭标签页函数
    labelRender (item) {
      //使用一个闭包
      return h => {
        return (
          <div>
          //展示tab标签名字
            <span>{item.meta.title}</span>
            //关闭按钮，icon组件是没有点击事件的，这里用native绑定点击事件。
            //要删除一个路由对象就要找到name、params、query，所以这里使用util.sj中定义的getTabNameByRoute()方法
            <icon nativeOn-click={this.handleTabRemove.bind(this, getTabNameByRoute(item))} type="md-close-circle" style="line-height:10px;"></icon>
          </div>
        )
      }
    }
  }
}
</script>

<style lang="less">
.layout-wrapper, .layout-outer{
  height: 100%;
  .header-wrapper{
    background: #fff;
    box-shadow: 0 1px 1px 1px rgba(0, 0, 0, .1);
    padding: 0 23px;
    .trigger-icon{
      cursor: pointer;
      transition: transform .3s ease;
      &.rotate{
        transform: rotateZ(-90deg);
        transition: transform .3s ease;
      }
    }
  }
  //超出的菜单栏添加滚动条
  .sider-outer{
    height: 100%;
    overflow: hidden;
    .ivu-layout-sider-children{
      //不想显示滚动条，隐藏滚动条。
      margin-right: -20px;
      //垂直方向有滚动条
      overflow-y: scroll;
      //水平方向隐藏滚动条
      overflow-x: hidden;
    }
  }
  .content-con{
    padding: 0;
    .ivu-tabs-bar{
      margin-bottom: 0;
    }
    .view-box{
      padding: 0;
    }
  }
  .page-card{
    min-height: ~"calc(100vh - 84px)";
  }
}
</style>
```

在store/module下创建tabNav.js用来存放标签的数组。在store/index.js下引入tabNav.js并注册`import tabNav from './module/tabNav'`

```
//store/module/tabNav.js文件下
import { routeHasExist, getRouteById, routeEqual, localSave, localRead } from '@/lib/util'

const state = {
  //保存打开的页面的列表，所有的列表是通过这个来渲染的。
  //每次执行先从本地存储中获取当前的tabList
  tabList: JSON.parse(localRead('tabList') || '[]')
}

//定义要往localstorage存储的值
const getTabListToLocal = tabList => {
  return tabList.map(item => {
    return {
      name: item.name,
      path: item.path,
      meta: item.meta,
      params: item.params,
      query: item.query
    }
  })
}

const mutations = {
  //把打开的标签页路由保存到state里面，点击显示已经点击过的标签页。
  UPDATE_ROUTER (state, route) {
    //实现功能如果已经有标签页，再次点击不会再次添加标签页。
    //判断有没有添加过标签页，没有就添加，有就不添加。
    //这里使用在util中定义的routeHasExist函数，传入第一个参数是保存的列表，第二个参数是跳转的路由。这里取反说明不存在，再往里面添加。这里把登录的tab页去掉，实际应该添加一个标识，哪些在第一次渲染时不要激活tab页。
    if (!routeHasExist(state.tabList, route) && route.name !== 'login') state.tabList.push(route)
    //往localstorage里存储数据
    localSave('tabList', JSON.stringify(getTabListToLocal(state.tabList)))
  },
  //删除匹配一致的值
  REMOVE_TAB (state, index) {
    state.tabList.splice(index, 1)
    //往localstorage里存储数据
    localSave('tabList', JSON.stringify(getTabListToLocal(state.tabList)))
  }
}

const actions = {
  //移除路由对象事件方法，第一个参数是一个对象这里获取到commit。第二个参数是传进来的对象。
  handleRemove ({ commit }, { id, $route }) {
    //返回一个Promise，因为还要做后续的跳转
    return new Promise((resolve) => {
      //使用util.js中定义的getRouteById()方法获取到路由对象
      let route = getRouteById(id)
      //通过路由对象找索引
      let index = state.tabList.findIndex(item => {
        //这里使用在util中定义的routeEqual方法判断点击传入的值与储存在列表中一致的值
        return routeEqual(route, item)
      })
      let len = state.tabList.length
      let nextRoute = {}
      //这里使用在util中定义的routeEqual方法判断当前激活的路由对象和要关闭的路由对象是否一致。
      if (routeEqual($route, state.tabList[index])) {
        //如果当前关闭的tab页右边还有tab页时，就激活下一个tab页
        if (index < len - 1) nextRoute = state.tabList[index + 1]
        //如果当前关闭的tab页是最后一个，就激活前一个tab页
        else nextRoute = state.tabList[index - 1]
      }
      //获取到name、params和query，因为路由跳转是依据这三个值。如果nextRoute没有值就说明是最后一个tab页就跳转到首页。
      const { name, params, query } = nextRoute || { name: 'home_index' }
      //提交在mutation中定义删除state中的值的方法，传入要删除的值。
      commit('REMOVE_TAB', index)
      resolve({
        name, params, query
      })
    })
  }
}

export default {
  state,
  mutations,
  actions
}
```

路由跳转与侧边菜单栏联动。

```
//components/side-menu/side-menu.vue文件下
<template>
  <div class="side-menu-wrapper">
    <slot></slot>
    //菜单栏联动，menu有一个属性:active-name直接绑定。因为左侧菜单栏不用添加动态属性。:open-names即为点击时展开父级菜单。
    <Menu ref="menu" :active-name="$route.name" :open-names="openNames" v-show="!collapsed" width="auto" theme="dark" @on-select="handleSelect">
      <template v-for="item in list">
        <re-submenu
          v-if="item.children"
          :key="`menu_${item.name}`"
          :name="item.name"
          :parent="item"
        >
        </re-submenu>
        <menu-item v-else :key="`menu_${item.name}`" :name="item.name">
          <Icon :type="item.icon" />
          {{ item.meta.title }}
        </menu-item>
      </template>
    </Menu>
    <div v-show="collapsed" class="drop-wrapper">
      <template v-for="item in list">
        <re-dropdown @on-select="handleSelect" v-if="item.children" :show-title="false" icon-color="#fff" :key="`drop_${item.name}`" :parent="item"></re-dropdown>
        <Tooltip v-else transfer :content="item.title" placement="right" :key="`drop_${item.name}`">
          <span @click="handleClick(item.name)" class="drop-menu-span">
            <Icon :type="item.icon" color="#fff" :size="20"></Icon>
          </span>
        </Tooltip>
      </template>
    </div>
  </div>
</template>

<script>
import ReSubmenu from './re-submenu.vue'
import ReDropdown from './re-dropdown.vue'
import { mapState } from 'vuex'
import { getOpenArrByName } from '@/lib/util'
export default {
  name: 'SideMenu',
  components: {
    ReSubmenu,
    ReDropdown
  },
  props: {
    collapsed: {
      type: Boolean,
      default: false
    },
    list: {
      type: Array,
      default: () => []
    }
  },
  computed: {
    ...mapState({
      //获取到保存的路由
      routers: state => state.router.routers
    }),
    //点击tab页时展开父级菜单
    openNames () {
      //使用定义在util.js里的函数
      return getOpenArrByName(this.$route.name, this.routers)
    }
  },
  watch: {
    //iview组件需要展开父级的时候要做一个手动的更新。
    openNames () {
      this.$nextTick(() => {
        this.$refs.menu.updateOpened()
      })
    }
  },
  methods: {
    handleSelect (name) {
      //路由跳转
      this.$router.push({
        name
      })
    },
    handleClick (name) {
      console.log(name)
    }
  }
}
</script>

<style lang="less">
.side-menu-wrapper{
  width: 100%;
  .ivu-tooltip, .drop-menu-span{
    display: block;
    width: 100%;
    text-align: center;
    padding: 5px 0;
  }
  .drop-wrapper > .ivu-dropdown{
    display: block;
    padding: 5px;
    margin: 0 auto;
  }
}
</style>
```

模拟打开有携带动态id的参数页。如果页面是动态路由，url根据动态参数变动，或者携带query。不同参数打开不同标签页。

```
//views/table.vue文件下
<template>
  <div>
    <!-- <edit-table :columns="columns" v-model="tableData" @on-edit="handleEdit"></edit-table> -->
    <edit-table-mul :columns="columns" v-model="tableData"></edit-table-mul>
    //生成参数
    <Button @click="turnTo">打开参数页</Button>
  </div>
</template>

<script>
import { getTableData } from '@/api/data'
import EditTable from '_c/edit-table'
import EditTableMul from '_c/edit-table-mul'
export default {
  components: {
    EditTable,
    EditTableMul
  },
  data () {
    return {
      tableData: [],
      columns: [
        { key: 'name', title: '姓名' },
        { key: 'age', title: '年龄', editable: true },
        { key: 'email', title: '邮箱', editable: true }
      ]
    }
  },
  methods: {
    handleEdit ({ row, index, column, newValue }) {
      console.log(row, index, column, newValue)
    },
    turnTo () {
      //这里获得一个随机的id
      let id = 'params' + (Math.random() * 100).toFixed(0)
      //路由跳转携带id
      this.$router.push({
        name: 'params',
        params: {
          id
        }
      })
    }
  },
  mounted () {
    getTableData().then(res => {
      this.tableData = res
    })
  }
}
</script>

<style>

</style>
```

方法

```
//lib/util.js文件下
//函数接受2个路由对象
export const routeEqual = (route1, route2) => {
  //判断params有params就为params，没有就为空对象
  const params1 = route1.params || {}
  const params2 = route2.params || {}
  //判断query
  const query1 = route1.query || {}
  const query2 = route2.query || {}
  //如果路由列表的name和要点击的路由对象相同，再判断判断属性名，属性值是否一一相等。这里在工具函数当中定义objEqual函数。先判断params在判断queyr，如果都相等就返回true。
  return route1.name === route2.name && objEqual(params1, params2) && objEqual(query1, query2)
}
//如果name相同参数页相同才能算一样的方法，用来处理点击的路由携带动态参数或者queyr的字段。这里第一个参数是路由列表，第二个参数是路由对象。
export const routeHasExist = (tabList, routeItem) => {
  //获取长度
  let len = tabList.length
  //默认不存在
  let res = false
  //这里使用之前定义的回调函数，这个方法能执行指定次数。
  doCustomTimes(len, (index) => {
    //当前遍历到的项，如果tabList和routeItem他们的name、params，query内容统统相等，那么就是true。
    //routeEqual判断相等的函数
    if (routeEqual(tabList[index], routeItem)) res = true
  })
  return res
}

//取出键值对方法，取出的结果如[['id', '123'],[]]
const getKeyValueArr = obj => {
  //保存的数组
  let arr = []
  //使用Object.entries方法处理键值对，params对象里读出的属性顺序是不一定的，里面属性相同，但是键值对不相同，这里用sort方法进行排序。
  Object.entries(obj).sort((a, b) => {
    return a[0] - b[0]
    //排序后遍历，这里参数是遍历的数组，每一个数组都有一个key和value，这里用数组结构赋值。
  }).forEach(([ _key, _val ]) => {
    arr.push(_key, _val)
  })
  //返回去除的键值对
  return arr
}

//params和query拼接函数
export const getTabNameByRoute = route => {
  //取出要拼接的三个字段
  const { name, params, query } = route
  把name赋给结果，最少传进来的对象都有name属性
  let res = name
  //判断有params的情况，并且存在属性。Object.keys()方法传入一个对象，它会把这个对象的所有属性名取出来放到一个数组里。拼接结果展示'argu:id_111_'，这里定义一个方法取出params里的键值对。
  if (params && Object.keys(params).length) res += ':' + getKeyValueArr(params).join('_')
  //拼接结果展示'argu:id_111&a_111&_b_222'
  if (query && Object.keys(query).length) res += '&' + getKeyValueArr(query).join('_')
  return res
}

const getObjBySplitStr = (id, splitStr) => {
  //将id按照传入的字符串进行分割，返回数组。
  let splitArr = id.split(splitStr)
  //取数组最后一个元素
  let str = splitArr[splitArr.length - 1]
  //再次切割数组，用_分割。
  let keyValArr = str.split('_')
  let res = {}
  let i = 0
  let len = keyValArr.length
  while (i < len) {
    //第一次循环时，[keyValArr[0]]即为属性名，keyValArer[i + 1]即为属性值。
    res[keyValArr[i]] = keyValArr[i + 1]
    //一对一对遍历这里所以加2
    i += 2
  }
  return res
}

//给一个id生成一个对象，用来做路由跳转。
export const getRouteById = id => {
  let res = {}
  //判断id包含是否包含&字符串，包含&说明包含query
  if (id.includes('&')) {
    //调用方法
    res.query = getObjBySplitStr(id, '&')
    //取&的左边的值
    id = id.split('&')[0]
  }
  if (id.includes(':')) {
    res.params = getObjBySplitStr(id, ':')
    //这边切割后只剩name
    id = id.split(':')[0]
  }
  res.name = id
  return res
}

export const getOpenArrByName = (name, routerList) => {
  //要得结果是数组
  let arr = []
  //使用some遍历，使用some不使用forEach更省性能
  routerList.some(item => {
    //如果当前点击的路由对象和激活的路由对象是相同的
    if (item.name === name) {
      //name值push到arr里面
      arr.push(item.name)
      return true
    }
    //遍历到路由对象的时候，有children，并且length不为0
    if (item.children && item.children.length) {
      //使用递归
      let childArr = getOpenArrByName(name, item.children)
      //不为0，有一个激活的路由对象
      if (childArr.length) {
        //做一个合并，把返回的itme.name和childArr这里要的childArr
        arr = arr.concat(item.name, childArr)
        return true
      }
    }
  })
  return arr
}

//每次tabList刷新就没了，因为是存储在store里面的。
//本地存储功能
export const localSave = (name, value) => {
  //存储路由跳转
  localStorage.setItem(name, value)
}

export const localRead = (name) => {
  //获取路由跳转
  return localStorage.getItem(name)
}
```

用来判断两个属性名和属性值是否一一相等的函数。

```
//lib/tools.js文件下
//传入两个对象
export const objEqual = (obj1, obj2) => {
  //使用Object的静态方法Object.keys方法获取到obj1的所有属性名，它是一个数组
  const keysArr1 = Object.keys(obj1)
  const keysArr2 = Object.keys(obj2)
  //判断长度是否相同，属性个数不相同的换直接返回false
  if (keysArr1.length !== keysArr2.length) return false
  //长度相等，如果都等于0就是两个空对象就返回true
  else if (keysArr1.length === 0 && keysArr2.length === 0) return true
  //使用some方法遍历数组，传入回调函数，some第一个参数是遍历到的一个元素。第二个参数是遍历的索引。第三个参数是当前的数组。这里只用第一个参数，就是当前的key。
  else return !keysArr1.some(key => obj1[key] !== obj2[key])
}
```

