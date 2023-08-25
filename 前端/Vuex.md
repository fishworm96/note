# Vuex

## 1.Vuex是什么

Vuex是实现组件全局状态(数据)管理的一种机制，可以方便的实现组件之间数据的共享

## 2.安装vuex依赖包

### 1.安装vuex依赖包

```
npm install vuex --save
```

### 2.导入vuex包

```
import Vuex from 'vuex'
```

### 3.创建store对象

```
const store = new Vuex.Store({
	//state中存放的就是全局共享的数据
	state: { count: 0 }
})
```

### 4.将store对象挂在到vue实例中

```
new Vue({
	el: '#app',
	render: h => h(app),
	router,
	//将创建的共享数据对象，挂在到Vue实例中
	//所有的组件，就可以直接从store中获取全局的数据了
	store
})
```

## 3.核心概念概述

Vuex中的主要核心概念如下

State

Mutation

Action

Getter

### 1.State

State提供唯一的公共数据源，所有共享的数据都要统一放到Store的State中进行存储

```
//创建store数据源，提供唯一公共数据
const store = new Vuex.Store({
	state: { count: 0 }
})
```

组件访问State中数据的第一种方式：

```
this.$store.state.全局数据名称
```

组件访问State中数据的第二种方式：

```
//1.从vuex中按需导入mapState函数
import { mapState } from 'vuex'
```

通过刚才导入的mapState函数，将当前组件按需要的全局数据，映射为当前组件的computed计算属性：

```
//2.将全局数据，映射为当前组件的计算属性
computed: {
	...mapState(['count'])
}
```

### 2.Mutation

Mutation用于变更Store中的数据

1.只能通过mutation变更Store数据，不可以直接操作Store中的数据

2.通过这种方式虽然操作起来稍微繁琐一些，但是可以集中监控所有数据局的变化

```
//定义Mutation
const store = new Vuex.Store({
	state: {
		ocunt: 0
	},
	mutations: {
		add(state) {
			state.count++
		}
	}
})
```

```
//触发mutation
methods: {
	handle() {
	//触发mutations的第一种方式
	this.$store.commit('add')
	}
}
```

可以在触发mutations时传递参数：

```
//定义Mutation
const store = new Vuex.store({
	state: {
		count: 0
	},
	mutations: {
		addN(state, step) {
			//变更状态
			state.count += step
		}
	}
})
```

```
//触发mutation
methods: {
	handle2() {
		//在调用commit函数，
		//触发mutation时携带参数
		this.$store.commit('addN', 3)
	}
}
```

触发mutations的第二种方式：

```
//1.从vuex中按需导入mapMutations函数
import {mapMutations} from 'vuex'
```

通过刚才导入的mapMutations函数，将需要mutations函数，映射为当前组件的methods方法：

```
//2.将指定的mutations函数，映射为当前组件的methods函数
methods: {
    ...mapMutations(['sub', 'subN']),
    btnHandle1 () {
      this.sub()
    },
    btnHandle2 () {
      this.subN(3)
    }
  }
```

### 3.Action

Action用于处理异步任务。

如果通过异步操作变更数据，必须通过Action，而不能使用Muattion，但是Action中还是要通过触发Mutation的方式间接变更数据

```
//定义Action
const store = new Vuex.Store({
	mutations: {
	add(state) {
		state.count++
	}
	},
	actions: {
		addAsync(context) {
		setTimeout (() => {
			context.commit('add')
		}, 1000)
		}
	}
})
```

```
//触发Action
methods: {
	handle() {
		//触发actions的第一种方式
		this.$store.dispatch('addAsync')
	}
}
```

触发actions异步任务时携带参数：

```
//定义Action
const store = new Vuex,store({
	mutations: {
		addN(state, step) {
			state.count += step
		}
	},
	actions: {
		addNAync(context, step) {
			setTimeout(() => {
				context.commit('addN', step)
			})
		}
	}
})
```

```
//触发Action
methods: {
	handle() {
		//在调用dispatch函数，
		//触发actions时携带参数
		this.$store.dispatch('addNAsync', 5)
	}
}
```

第二种方式：

```
//1.从vuex中按需导入mapActions函数
import { mapActions } from 'vuex'
```

通过刚才导入的mapActions函数，将需要的actions函数，映射为当前组件的method方法：

```
//2.将指定的actions函数，映射为当前组件的methods函数
methods: {
	...mapActions(['subAsync', 'subNAsync']),
	btnHandle3() {
		this.subAsync()
	}
}
或
直接在插槽中使用
<button @click="subAsync()">setTimeout-N</button>
```

### 4.Getter

Getter用于对Store中的数据进行加工处理形成新的数据。//不会修改Vue中原数据，只会包装数据

1.Getter可以对Store中已有的数据加工处理之后形成新的数据，类似Vue的计算属性。

2.Store中数据发生变化，Getter的数据也会跟着变化

```
//定义Getter
const store = new Vuex.Store({
	state: {
		count: 0
	},
	getters: {
		showNum: state => {
			return '当前最新的数量是【'+ store.count +'】'
		}
	}
})
```

使用getter的第一种方式：

```
this.$store.getter.名称
```

使用getter的第二种方式：

```
import { mapGetters} from 'vuex'

<h3>{{showNum}}</h3>

computed: {
	...mapGetters(['showNum'])
}
```

