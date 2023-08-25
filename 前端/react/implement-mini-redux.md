# 动手实现一个redux

## 获取数据

现在定义了三个组件名为ComponentOne、ComponentTwo、ComponentThree
在组件一中展示数据、组件二用来修改数据

```jsx
const ComponentOne = () => <section>组件一</section>
```

```jsx
const ComponentTwo = () => <section>组件二</section>
```

__那么数据从哪里来的呢？__
使用的是react的上下文
先简单介绍一下context，context是react 16.3中新增特性。用来解决传递props时需要使用的组件比较深层次，那么要从父组件一直传递，期间无需用到props的组件也需要接收props。而context可以直接将数据传递到需要使用数据的组件不需要再层层传递。

```jsx
// 创建一个context，这里可以给它值或不给。
const MyContext = createContext(null)
// 如何提供数据？
// 直接像普通的组件一样使用，不过需要加 .Provider 和value
<MyContext.Provider value="dark">
  <Son>
</MyContext.Provider>
// 如何获取数据？
// 使用 useContext并传入创建的context，返回值即使需要使用的数据。
function Son() {
  const theme = useContext(MyContext)
  return <div>{theme}</div>
}
```

现在使用Context来获取数据

```jsx
function App () {
  const [appState, setAppState] = useState({
    user: { name: "fishworm", age: 18 }
  })
  const contextValue = { appState, setAppState }

  return (
    <appContext.Provider value={contextValue}>
      <ComponentOne />
      <ComponentTwo />
      <ComponentThree />
    </appContext.Provider>
  )
}
```

组件一使用`contextValue`来获取数据

```jsx
const ComponentOne = () => <section>组件一<User /></section>

const User = () => {
  const contextValue = useContext(appContext)
  return <div>User: {contextValue.appState.user.name}</div>
}
```

组件二使用 `contextValue` 的 `setAppState` 来修改数据

```jsx
const ComponentTwo = () => <section>组件二<UserModifier /></section>

const UserModifier = () => {
  const contextValue = useContext(appContext)
  const { appState, setAppState } = contextValue
  const onChange = e => {
    appState.user.name = e.target.value
    setAppState({ ...contextValue.appState })
  }

  return <div>
    <input value={contextValue.appState.user.name}
      onChange={onChange}
    />
  </div>
}
```

完整代码

```jsx
import { useState, useContext, createContext } from "react"

const appContext = createContext(null)

function App () {
  const [appState, setAppState] = useState({
    user: { name: "fishworm", age: 18 }
  })
  const contextValue = { appState, setAppState }

  return (
    <appContext.Provider value={contextValue}>
      <ComponentOne />
      <ComponentTwo />
      <ComponentThree />
    </appContext.Provider>
  )
}

const ComponentOne = () => <section>组件一<User /></section>
const ComponentTwo = () => <section>组件二<UserModifier /></section>
const ComponentThree = () => <section>组件三</section>

const User = () => {
  const contextValue = useContext(appContext)
  return <div>User: {contextValue.appState.user.name}</div>
}

const UserModifier = () => {
  const contextValue = useContext(appContext)
  const { appState, setAppState } = contextValue
  const onChange = e => {
    appState.user.name = e.target.value
    setAppState({ ...contextValue.appState })
  }

  return <div>
    <input value={contextValue.appState.user.name}
      onChange={onChange}
    />
  </div>
}

export default App
```

## 规范创建流程

上面的 `appState` 创建流程不规范，现在按照redux中的reducer的使用方法来规范一下创建数据。

```jsx
const reducer = (state, { type, payload }) => {
  if (type === 'updateUser') {
    return {
      ...state,
      user: {
        ...state.user,
        ...payload
      }
    }
  }

  return state
}
```

```diff
const UserModifier = () => {
  const contextValue = useContext(appContext)
  const { appState, setAppState } = contextValue
  const onChange = e => {
    appState.user.name = e.target.value
-   setAppState({ ...contextValue.appState })
+   setAppState(reducer(appState, { type: 'updateUser', payload: { name: e.target.value } }))
  }

  return <div>
    <input value={contextValue.appState.user.name}
      onChange={onChange}
    />
  </div>
}
```

## 使用dispatch规范setState流程

现在我们每次更新数据都要使用 `setAppState(reducer())` 太繁琐了，有没有什么办法简化一下

```jsx
setAppState(reducer(appState, { type: 'updateUser', payload: { name: e.target.value } }))
setAppState(reducer(appState, { type: 'updateAge', payload: { name: e.target.value } }))
```

使用dispatch简化 `setState`

```jsx
const Wrapper = () => {
  const { appState, setAppState } = useContext(appContext)
  const dispatch = (action) => {
    setAppState(reducer(appState, action))
  }

  return <UserModifier dispatch={dispatch} state={appState} />
}

const UserModifier = ({ dispatch, state }) => {
  const onChange = e => {
    dispatch({ type: 'updateUser', payload: { name: e.target.value } })
  }

  return <div>
    <input value={state.user.name}
      onChange={onChange}
    />
  </div>
}
```

```diff
- const ComponentTwo = () => <section>组件二<UserModifier /></section>
+ const ComponentTwo = () => <section>组件二<Wrapper /></section>
```

## 将组件修改成可复用的高阶组件connect

上面代码虽然是优化了 `setAppState(reducer())` ，但是这个组件只能在一个地方使用。现在这样在其他组件中就没办法获取到 `dispatch` 和 `state`，如果每个组件都套一层 `Wrapper` 太冗余了。所以需要修改一下

这里在简单介绍一下什么是高阶组件：就是接收一个组件作为参数，在不修改样式只修改逻辑的情况下返回组件。
了解了什么是高阶组件，我们就按照这个逻辑进行修改。

```jsx
const connect = (Component) => {
  return (props) => {
    const {appState, setAppState} = useContext(appContext)
    const dispatch = (action) => {
      setAppState(reducer(appState, action))
    }
    return <Component {...props} dispatch={dispatch} state={appState} />
  }
}
```

使用高阶组件

```jsx
// 这是他原来的样子
const UserModifier = ({ dispatch, state }) => {}
// 修改后的样子
const UserModifier = connect(({ dispatch, state, children }) => {
    const onChange = e => {
    dispatch({ type: 'updateUser', payload: { name: e.target.value } })
  }

  return <div>
    {children}
    <input value={state.user.name}
      onChange={onChange}
    />
  </div>
})
```

```diff
- const ComponentTwo = () => <section>组件二<Wrapper /></section>
+ const ComponentTwo = () => <section>组件二<UserModifier /></section>
```

## 精准渲染组件

```jsx
// App.jsx
import { useState, useContext, createContext } from "react"

const appContext = createContext(null)

function App () {
  const [appState, setAppState] = useState({
    user: { name: "fishworm", age: 18 }
  })
  const contextValue = { appState, setAppState }

  return (
    <appContext.Provider value={contextValue}>
      <ComponentOne />
      <ComponentTwo />
      <ComponentThree />
    </appContext.Provider>
  )
}

const ComponentOne = () => {
  console.log('componentOne' + Math.random())
  return <section>组件一<User /></section>
}
const ComponentTwo = () => {
  console.log('componentTwo' + Math.random())
  return <section>组件二<UserModifier /></section>
}
const ComponentThree = () => {
  console.log('componentThree' + Math.random())
  return <section>组件三</section>
}

const User = () => {
  const contextValue = useContext(appContext)
  return <div>User: {contextValue.appState.user.name}</div>
}

const reducer = (state, { type, payload }) => {
  if (type === 'updateUser') {
    return {
      ...state,
      user: {
        ...state.user,
        ...payload
      }
    }
  }

  return state
}

const connect = (Component) => {
  return (props) => {
    const { appState, setAppState } = useContext(appContext)
    const dispatch = (action) => {
      setAppState(reducer(appState, action))
    }
    return <Component {...props} dispatch={dispatch} state={appState} />
  }
}

const UserModifier = connect(({ dispatch, state, children }) => {
  console.log('UserModifier执行了' + Math.random())
  const onChange = e => {
    dispatch({ type: 'updateUser', payload: { name: e.target.value } })
  }

  return <div>
    {children}
    <input value={state.user.name}
      onChange={onChange}
    />
  </div>
})

export default App
```

给每个组件添加了 `log` 。现在有一个问题，使用 `setAppState` 修改内容后每个组件都会重新渲染。但是现在组件三并没有修改，我也不希望他重新渲染。只有几个组件时不会有什么问题，如果组件很多就会有性能问题。
__这是什么原因造成的呢？__
原因就是使用 `useState` 重新渲染后父组件会重新渲染，父组件重新渲染后子组件也会重新渲染。
__那应该怎么办？__
使用 `useMemo` 来优化

```jsx
const x = useMemo = (() => {
  return <ComponentThree>
}, [appState.user.age])

    <appContext.Provider value={contextValue}>
      <ComponentOne />
      <ComponentTwo />
      {x}
    </appContext.Provider>
```

这样修改虽然可以，但是每个组件都要使用 `useMemo` 还是太麻烦了。有没有其他办法，用到 `user` 时才会重新执行。
现在重新修改一下 `state`

```jsx
const appContext = createContext(null)

const store = {
  state: {
    user: { name: 'fishworm', age: 18 }
  },
  setState: (newState) {
    store.state = newState
  },
}

function App () {
  return (
    <appContext.Provider value={store}>
      <ComponentOne />
      <ComponentTwo />
      <ComponentThree />
    </appContext.Provider>
  )
}

const User = () => {
  const state = useContext(appContext)
  return <div>User: {state.user.name}</div>
}

const connect = (Component) => {
  return (props) => {
    const { state, setState } = useContext(appContext)
    const dispatch = (action) => {
      setState(reducer(appState, action))
    }
    return <Component {...props} dispatch={dispatch} state={state} />
  }
}

```

修改后就不使用 `useState` 重新渲染了，虽然修改代码后不会重新渲染多余的组件了，但是需要重新渲染的组件也不渲染了。
这个问题简单。

```jsx
const connect = (Component) => {
  return (props) => {
    const { state, setState } = useContext(appContext)
    // 使用 useState重新渲染当前组件
    const [, update] = useState({})
    const dispatch = (action) => {
      setState(reducer(appState, action))
      // 使用 useState重新渲染当前组件
      update({})
    }
    return <Component {...props} dispatch={dispatch} state={state} />
  }
}
```

因为每个对象的引用是不一样的，在 `connect` 中使用 `useState` 就会重新渲染当前组件。
虽然当前修改的组件会重新渲染了，但是展示的组件 `组件一` 没有重新渲染。
这里我们可以用到发布订阅模式来进行修改。

```jsx
export const store = {
  state: {
    user: { name: 'fishworm', age: 18 }
  },
  setState (newState) {
    store.state = newState
    // 需要重新渲染的组件执行渲染
    // 这里放不放store.state都行，放了方便取出新的state
    store.listeners.map(fn => fn(store.state))
  },
  // 订阅中心
  listeners: [],
  subscribe (fn) {
    store.listeners.push(fn)
    // 取消订阅
    return () => {
      const index = store.listeners.indexOf(fn)
      store.listeners.splice(index, 1)
    }
  }
}

export const connect = (Component) => {
  return (props) => {
    const { state, setState } = useContext(appContext)
    const [, update] = useState({})
    useEffect(() => {
      // 重新渲染
      store.subscribe(() => {
        update({})
      })
    }, [])
    const dispatch = (action) => {
      setState(reducer(state, action))
    }
    return <Component {...props} dispatch={dispatch} state={state} />
  }
}
```
抽离 `redux` 逻辑
完整代码

```jsx
// App.jsx
import React from 'react'
import { appContext, store, connect } from './redux.jsx'

function App () {
  return (
    <appContext.Provider value={store}>
      <ComponentOne />
      <ComponentTwo />
      <ComponentThree />
    </appContext.Provider>
  )
}

const ComponentOne = () => {
  console.log('componentOne' + Math.random())
  return <section>组件一<User /></section>
}
const ComponentTwo = () => {
  console.log('componentTwo' + Math.random())
  return <section>组件二<UserModifier /></section>
}
const ComponentThree = () => {
  console.log('componentThree' + Math.random())
  return <section>组件三</section>
}

const User = connect(({ state, dispatch }) => {
  console.log('User执行了' + Math.random())
  return <div>User: {state.user.name}</div>
})



const UserModifier = connect(({ dispatch, state, children }) => {
  console.log('UserModifier执行了' + Math.random())
  const onChange = e => {
    dispatch({ type: 'updateUser', payload: { name: e.target.value } })
  }

  return <div>
    {children}
    <input value={state.user.name}
      onChange={onChange}
    />
  </div>
})

export default App
```

```jsx
// redux.jsx
import React, { useContext, useEffect, useState } from 'react'

export const store = {
  state: {
    user: { name: 'fishworm', age: 18 }
  },
  setState (newState) {
    store.state = newState
    store.listeners.map(fn => fn(store.state))
  },
  listeners: [],
  subscribe (fn) {
    store.listeners.push(fn)
    return () => {
      const index = store.listeners.indexOf(fn)
      store.listeners.splice(index, 1)
    }
  }
}

export const reducer = (state, { type, payload }) => {
  if (type === 'updateUser') {
    return {
      ...state,
      user: {
        ...state.user,
        ...payload
      }
    }
  }

  return state
}

export const connect = (Component) => {
  return (props) => {
    const { state, setState } = useContext(appContext)
    const [, update] = useState({})
    useEffect(() => {
      store.subscribe(() => {
        update({})
      })
    }, [])
    const dispatch = (action) => {
      setState(reducer(state, action))
    }
    return <Component {...props} dispatch={dispatch} state={state} />
  }
}

export const appContext = React.createContext(null)
```

## connect 接受 selector

先看下需要实现什么功能，在 `redux` 中的 `connect` 执行的第一个括号可以接收一个 `selector` 参数，这个参数就是在 `state` 中的变量，在第二个括号中就能获取到第一个括号中返回的值。
__这解决了什么问题呢？__
如果要获取值得名称很长，那么用这个方法就可以简化。

```jsx
// 比如
const User = connect(({state, dispatch}) => {
  return <div>{{state.xxx.xxx.xxx.name}}</div>
})
// 解决
const User = connect((state) => {
  return {name: state.xxx.xxx.xxx.name}
})(({name}) => {
  return <div>{{name}}</div>
})
```

现在刷要执行两次函数，并且第一个函数需要获取一个selector参数，第二个函数获取Component参数

```jsx
const connect = (selector) => (Component) => {
    return <Component {...props} {...data} dispatch={dispatch} />
}
```

在这个基础上第一个函数的返回值是第二个函数的参数

```jsx
const connect = (selector) => (Component) => {
  // 获取到参数，如果有参数的返回处理后的参数，没有参数就返回原始值
  const data = selector ? selector(state) : { state }
  return <Component {...props} {...data} dispatch={dispatch} />
}
```

看下完整代码

修改后的 `connect`

```jsx
export const connect = (selector) => (Component) => {
  return (props) => {
    const { state, setState } = useContext(appContext)
    const [, update] = useState({})
    const data = selector ? selector(state) : { state }
    useEffect(() => {
      store.subscribe(() => {
        update({})
      })
    })
    const dispatch = (action) => {
      setState(reducer(state, action))
    }
    return <Component {...props} {...data} dispatch={dispatch} />
  }
}
```

使用 `connect`

```jsx
const User = connect((state) => {
  return {user: state.user}
})(({user}) => {
  console.log('User执行了' + Math.random())
  return <div>User: {user.name}</div>
})
```

## selector的精准渲染

__selector的精准渲染是什么意思？__
组件只在自己的数据变化时渲染

__要做什么呢？__
只用判断一下数据有没有修改过

```jsx
const new Data = selector ? selector(store.state) : { state: store.state }
// 主要是这里
if (changed(data, newData)) {
  update({})
}
```

__`changed`做了什么？__
  主要把两个对象遍历出来挨个判断，如果相等说明没有修改，不相等说明修改了就重新执行渲染。

```jsx
const changed = (oldState, newState) => {
  let changed = false
  for (let key in oldState) {
    if (oldState[key] !== newState[key]) {
      changed = true
    }
  }
  return changed
}
```
 
```jsx
useEffect(() => {
  store.subscribe(() => {
  const newData = selector
    ? selector(state)
    : { state: state }
  if (changed(data, newData)) {
    update({})
  }
})
}, [selector])
```

__结束了吗？__
还没有，这里还需要取消订阅。如果不取消订阅，给`useEffect`添加第二个依赖`state`每次修改都会多次渲染。当然我们不需要给它添加第二个依赖

`store.subscribe`的返回值就是取消订阅的函数，执行这个函数即可

```jsx
useEffect(() => {
  const unsubscribe = store.subscribe(() => {
  const newData = selector
    ? selector(state)
    : { state: state }
  if (changed(data, newData)) {
    update({})
  }
})
  return unsubscribe
}, [selector])

// 这里可以将它简写
useEffect(() => {
  return store.subscribe(() => {
  const newData = selector
    ? selector(state)
    : { state: state }
  if (changed(data, newData)) {
    update({})
    }
  })
}, [selector])

// 再进一步简写
useEffect(() => store.subscribe(() => {
  const newData = selector
    ? selector(state)
    : { state: state }
  if (changed(data, newData)) {
    update({})
    }
}), [selector])
```

现在以组件三为例使用

```jsx
// App.jsx
const ComponentThree = connect(state => {
  return { group: state.group }
})(({ group }) => {
  console.log('componentThree' + Math.random())
  return <section>组件三<div>Group: {group.name}</div></section>
})
```

## 添加mapDispatchToProps

__`mapDispatchToProps`有什么用呢？__

```jsx
// 修改前
const UserModifier = connect()(({ dispatch, state, children }) => {
  console.log('UserModifier执行了' + Math.random())
  const onChange = e => {
    dispatch({ type: 'updateUser', payload: { name: e.target.value } })
  }
  ...
}

// 修改后
const UserModifier = connect(null, (dispatch) => {
  return {
    updateUser: (attrs) => dispatch({type: 'updateUser', payload: attrs})
  }
})(({ updateUser, state, children }) => {
  console.log('UserModifier执行了' + Math.random())
  const onChange = e => {
    updateUser({ name: e.target.value })
  }
  ...
})
```

__这里做了什么？__
就是向`connect`的第二个函数传入一个函数，函数返回一个修改数据的的函数。
__为什么要这么做？解决了什么？__
这么做主要是修改更新数据过程，看起来更好看了些。

__主要做了什么？__
主要就是需要判断`dispatchSelector`是否存在，如果存在就往`dispatchSelector`中传入需要更新数据的函数。

```jsx
const connect = (selector, dispatchSelector) => (Component) => {
  const dispatch = (action) => {
    setState(reducer(state, action))
  }
  const dispatchers = dispatchSelector ? dispatchSelector(dispatch) : { dispatch }
}
```

来看下完整的实现

```jsx
export const connect = (selector, dispatchSelector) => (Component) => {
  return (props) => {
    const dispatch = (action) => {
      setState(reducer(state, action))
    }
    const { state, setState } = useContext(appContext)
    const [, update] = useState({})
    const data = selector ? selector(state) : { state }
    const dispatchers = dispatchSelector
      ? dispatchSelector(dispatch)
      : { dispatch }
    useEffect(() => store.subscribe(() => {
        const newData = selector
          ? selector(store.state)
          : { state: store.state }
        if (changed(data, newData)) {
          update({})
        }
      }), [selector])
    return <Component {...props} {...data} {...dispatchers} />
  }
}
```

```jsx
// 使用
const UserModifier = connect(null, (dispatch) => {
  // 返回 updateUser
  return {
    updateUser: (attrs) => dispatch({type: 'updateUser', payload: attrs})
  }
})(({ updateUser, state, children }) => {
  // 接收到 updateUser 这个函数
  console.log('UserModifier执行了' + Math.random())
  const onChange = e => {
    updateUser({ name: e.target.value })
  }

  ...
})
```

## 添加connectToUser

经过上面一顿操作，实现了一个简单的redux，但是现在发现有重复的代码出现。

```jsx
// 比如都需要使用`connect`然后执行两个函数
const User = connect((state) => {
  return { user: state.user }
})(({ user }) => {
  ...
}

const UserModifier = connect(null, (dispatch) => {
  return {
    updateUser: (attrs) => dispatch({type: 'updateUser', payload: attrs})
  }
})(({ updateUser, state, children }) => {
  ...
}
```

现在就需要把代码抽离出来复用

创建一个`connecters`的文件用来处理`User`的逻辑

```jsx
// connectToUser.jsx
import { connect } from '../redux'
// 读数据
const userSelector = state => {
  return { user: state.user }
}
// 写数据
const userDispatcher = (dispatch) => {
  return {
    updateUser: (attrs) => dispatch({type: 'updateUser', payload: attrs})
  }
}

export const connectToUser = connect(userSelector, userDispatcher)
```

```jsx
// 抽离前的代码
// 直接使用 connect
const User = connect((state) => {
  return { user: state.user }
})(({ user }) => {
  console.log('User执行了' + Math.random())
  return <div>User: {user.name}</div>
})

const UserModifier = connect(null, (dispatch) => {
  return {
    updateUser: (attrs) => dispatch({type: 'updateUser', payload: attrs})
  }
})(({ updateUser, state, children }) => {
  console.log('UserModifier执行了' + Math.random())
  const onChange = e => {
    updateUser({ name: e.target.value })
  }

  return <div>
    {children}
    <input value={state.user.name}
      onChange={onChange}
    />
  </div>
})

// 抽离后的代码
// 使用抽离的 connectToUser
const User = connectToUser(({ user }) => {
  console.log('User执行了' + Math.random())
  return <div>User: {user.name}</div>
})

const UserModifier = connectToUser(({ updateUser, user, children }) => {
  console.log('UserModifier执行了' + Math.random())
  const onChange = e => {
    updateUser({ name: e.target.value })
  }

  return <div>
    {children}
    <input value={user.name}
      onChange={onChange}
    />
  </div>
})
```

## 封装Provider和createStore

现在发现了一个问题，`state`的数据和`reducer`的逻辑都是写死的，需要修改成动态的。
把`state`和`reducer`都置空
`state: undefined, reducer: undefined`
使用函数来修改

```jsx
export const store = {
  state: undefined,
  reducer: undefined,
  setState (newState) {
    store.state = newState
    store.listeners.map(fn => fn(store.state))
  },
  listeners: [],
  subscribe (fn) {
    store.listeners.push(fn)
    return () => {
      const index = store.listeners.indexOf(fn)
      store.listeners.splice(index, 1)
    }
  }
}
```

### 创建`store`

```jsx
export const createStore = (reducer, initState) => {
  store.state = initState
  store.reducer = reducer
  return store
}
```

拆解`reducer`与初始数据

```jsx
// 通过传入2个参数第一个是 reducer，第二个是初始数据
const store = createStore((state, { type, payload }) => {
  if (type === 'updateUser') {
    return {
      ...state,
      user: {
        ...state.user,
        ...payload
      }
    }
  }
  return state
}, {
  user: { name: 'fishworm', age: 18 },
  group: { name: '前端组' }
})
```

这样看是不是有点不清晰，再把他们拆开。

```jsx
const reducer = (state, { type, payload }) => {
  if (type === 'updateUser') {
    return {
      ...state,
      user: {
        ...state.user,
        ...payload
      }
    }
  }
  return state
}

const initState = {
  user: { name: 'fishworm', age: 18 },
  group: { name: '前端组' }
}

const store = createStore(reducer, initState)
```

### 创建Provider
我们发现`redux`不是使用`appContext.Provider`这样使用的有没有办法修改成`redux`的样子？
可以的，只用修改一下

```jsx
export const Provider = ({store, children}) => {
  return (
    <appContext.Provider value={store}>
      {children}
    </appContext.Provider>
  )
}
```

```jsx
function App () {
  return (
    <Provider store={store}>
      <ComponentOne />
      <ComponentTwo />
      <ComponentThree />
    </Provider>
  )
}
```

## 总结
首先我们有一个`App`它里面会有其他组件，我们需要做的就是让它访问到一个全局的`state`

### 第一个概念`state`
`state`就是存储数据用的，我们发现`redux`中的`state`是放在`store`中。
__那我们如何将`state`与组件连接起来呢？__
在`react-redux`中使用`connect`将组件与`state`连接起来，它提供了读与写功能。如果要读就从组件的属性中获取`state`，如果要写就在组件中获取`dispatch`来修改数据。如果要精确的读就传入`selector`，如果要精确的写就使用`mapDispatchToProps`函数返回的`updateUser`这个自己封装的`api`。
`react-redux`所要解决的问题就是将组件与`state`连接，并提供两个读写`api`
### 第二个概念connect
`connect`做的是对组件做了一次封装，封装成一个`Wrapper`并返回出去。
主要做了三件事情
- 1、从上下文中获取到`state`和`setState`。不过不用上下文用`store`也可以。
- 2、得到具体的数据和具体的`dispatch`。
- 3、在恰当的时间进行更新，对`store`进行订阅。一旦发现`store`变化，就在数据更的情况下使用`update`进行更新。

然后`dispatch`又拆分为`reducer`、`initState`、`action`
reducer：在本篇中是用于规范创建`state`的过程，为了不修改原来的`state`。
initState: 初始化`state`。
action: 对一次变动的描述。如：type、payload

## Api重构

`redux`中暴露的`Api`为`getState()、dispatch(action)、subscribe(listener)`
现在重构成和`redux`相同，本质还是隐藏起来不暴露出去所有的属性

```jsx
// getState
// 原来
export const store = {
    state: {
    user: undefined,
    group: undefined
  }
}

// 重构
let state = undefined
export const store = {
  getState() {
    return state
  }
}
```

```jsx
// dispatch
// 原来
export const store = {
    setState (newState) {
    store.state = newState
    store.listeners.map(fn => fn(store.state))
  }
}
export const createStore = (reducer, initState) => {
  store.state = initState
  store.reducer = reducer
  return store
}
const dispatch = (action) => {
  setState(store.reducer(state, action))
}

// 重构
let reducer = undefined
const setState = (newState) => {
  state = newState
  listeners.map(fn => fn(state))
}
export const store = {
  dispatch: (action) => {
    setState(reducer(state, action))
  }
}
const dispatch = store.dispatch
export const createStore = (_reducer, initState) => {
  state = initState
  reducer = _reducer
  return store
}
```

```jsx
// subscribe
// 原来
export const store = {
  listeners: [],
  subscribe (fn) {
    store.listeners.push(fn)
    return () => {
      const index = store.listeners.indexOf(fn)
      store.listeners.splice(index, 1)
    }
  }
}

// 重构
let listeners = []
export const store = {
  subscribe (fn) {
    listeners.push(fn)
    return () => {
      const index = listeners.indexOf(fn)
      listeners.splice(index, 1)
    }
  }
}
```

## 改写dispatch以支持函数action

先看下要实现什么功能

```jsx
const ajax = () => {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve({data: {name: '3秒后的fishworm'}})
    }, 3000)
  })
}

const fetchUser = (dispatch) => {
  ajax('/user').then(response => {
    dispatch({type: 'updateUser', payload: response.data})
  })
}

const UserModifier = connect(null, null)(({ state, dispatch }) => {
  console.log('UserModifier执行了' + Math.random())
  const onClick = e => {
    // updateUser({ name: e.target.value })
    dispatch(fetchUser)
  }

  return <div>
    <div>User: {state.user.name}</div>
    <button onClick={onClick}>异步获取</button>
  </div>
})
```

就是简化异步请求，让看起来更优雅。

__怎么做呢？__

```jsx
let dispatch = store.dispatch
// 保存老的dispatch
const prevDispatch = dispatch

dispatch = (action) => {
  // 如果传入的是一个函数就传一个新的dispatch
  if (action instanceof Function) {
    action(dispatch)
  } else {
    // 如果不是函数，直接传给老的dispatch修改
    prevDispatch(action)
  }
}
```

## 改写dispatch以支持payload为promise的action

有时候`payload`传入的可能是一个`promise`

```jsx
const fetchUserPromise = () => {
  return ajax('/user').then(response => response.data)
}

const fetchUser = (dispatch) => {
  return ajax('/user').then(response => dispatch({type: 'updateUser', payload: response.data}))
}

const UserModifier = connect(null, null)(({ state, dispatch }) => {
  console.log('UserModifier执行了' + Math.random())
  const onClick = e => {
    dispatch({
      type: 'updateUser',
      payload: fetchUserPromise()
    })
    // dispatch(fetchUser)
  }
  ...
}
```

__应该如何修改呢？__

```jsx
const prevDispatch2 = dispatch

dispatch = (action) => {
  // 判断传入的是否为promise
  if (action.payload instanceof Promise) {
    // 递归传入异步参数
    action.payload.then(data => {
      dispatch({...action, payload: data})
    })
  } else {
    prevDispatch2(action)
  }
}
```

总结：
其实就是改变了组织形式，没有对代码做任何优化。

## 阅读 redux-thunk 和 redux-promise 源码

redux是不支持函数和`payload`为`promise`，所以它就让我们自己添加一个中间件`applyMiddleware`
`const store = createStore(reducer, initState, applyMiddleware(reduxThunk, reduxPromise))`


```typescript
// redux-thunk
function createThunkMiddleware<
  State = any,
  BasicAction extends Action = AnyAction,
  ExtraThunkArg = undefined
>(extraArgument?: ExtraThunkArg) {
  const middleware: ThunkMiddleware<State, BasicAction, ExtraThunkArg> =
    ({ dispatch, getState }) =>
    next =>
    action => {
      if (typeof action === 'function') {
        return action(dispatch, getState, extraArgument)
      }
      return next(action)
    }
  return middleware
}

export const thunk = createThunkMiddleware()
```

是不是和我们写的很像，它使用`typeof`来判断函数，我们用`instanceof`来判断。如果不是函数，就调用`next`函数。可以是下一个中间件或`dispatch`

```jsx
// redux-promise
export default function promiseMiddleware({ dispatch }) {
  return next => action => {
    if (!isFSA(action)) {
      return isPromise(action) ? action.then(dispatch) : next(action);
    }
    // 判断是否为promise
    return isPromise(action.payload)
      ? action.payload
          .then(result => dispatch({ ...action, payload: result }))
          .catch(error => {
            dispatch({ ...action, payload: error, error: true });
            return Promise.reject(error);
          })
      : next(action);
  };
}
```

如果是`promise`就把`payload`放到`then`里面然后把结果放到`payload`的里面吗，不是`promise`就执行下一个中间件。

源码地址：https://github.com/fishworm96/implement-mini-redux