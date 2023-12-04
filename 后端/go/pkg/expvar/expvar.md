## package expvar

包 expvar 为公共变量（如服务器中的操作计数器）提供了一个标准化接口。它通过 HTTP 在 /debug/vars 以 JSON 格式公开这些变量。

设置或修改这些公共变量的操作都是原子操作。

除添加 HTTP 处理程序外，该软件包还注册了以下变量：

```go
cmdline   os.Args
memstats  runtime.Memstats
```

有时导入该软件包只是为了注册 HTTP 处理程序和上述变量。要以这种方式使用它，请将此软件包链接到您的程序中：

```go
import _ "expvar"
```

## Index

### func Do(f func(KeyValue))

为每个导出变量调用 f。在迭代过程中，全局变量映射被锁定，但现有条目可以同时更新。

### func Handler() http.Handler 添加于1.8

Handler 返回 expvar HTTP 处理程序。

只有在将处理程序安装到非标准位置时才需要这样做。

### func Publish(name string, v aVar)

Publish 声明一个命名的导出变量。当软件包创建 Vars 时，应在其 init 函数中调用该变量。如果变量名已被注册，则会出现 log.Panic.Publish。

### type Float

```go
type Float struct {
  // 包含已过滤或未导出的字段
}
```

Float 是一个满足 Var 接口的 64 位浮点变量。

#### func NewFloat(name string) *Float

#### func (v *Float) Add(delta float64)

在 v 中添加 delta。

#### func (v *Float) Set(value float64)

Set 将 v 设为值。

#### func (v *Float) String() string

#### func (v *Float) Value() float64 添加于1.8

### type Func

```go
type Func func() any
```

Func 通过调用函数和使用 JSON 格式化返回值来实现 Var。

#### func (f Func) String() string

#### func (f Func) Value() any 添加于1.8

### type Int

```go
type Int struct {
  // 包含已过滤或未导出的字段
}
```

Int 是一个满足 Var 接口的 64 位整数变量。

#### func NewInt(name string) *Int

#### func (v *Int) Add(delta int64)

#### func (v *Int) Set(value int64)

#### func (v *Int) String() string

#### func (v *Int) Value() int64 添加于1.8

### type KeyValue

```go
type KeyValue struct {
  Key string
  Value Var
}
```

KeyValue 表示 Map 中的一个条目。

### type Map

```go
type Map struct {
  // 包含已过滤或未导出的字段
}
```

Map 是满足 Var 接口的字符串到 Var 的映射变量。

#### func NewMap(name string) *Map

#### func (v *Map) Add(key string, delta int64)

将 delta 添加到存储在给定映射键下的 *Int 值中。

#### func (v *Map) AddFloat(key string, delta float64)

AddFloat 将 delta 添加到存储在给定的映射键下的 *Float 值中。

#### func (v *Map) Delete(key string) 添加于1.12

Delete 从地图中删除给定的键。

#### func (v *Map) Do(f func(KeyValue))

为地图中的每个条目调用 f。在迭代过程中，地图会被锁定，但现有条目可能会被同时更新。

#### func (v *Map) Get(key string) Var

#### func (v *Map) Init() *Map

初始化会删除地图上的所有键。

#### func (v *Map) Set(key string, av Var)

#### func (v *Map) String() string

### type String

```go
type String struct {
  // 包含已过滤或未导出的字段
}
```

String 是字符串变量，符合 Var 接口。

#### func NewString(name string) *String

#### func (v *String) Set(value string)

#### func (v *String) String() string

String 实现了 Var 接口。要获取未加引号的字符串，请使用 Value。

#### func (v *String) Value() string 添加于1.8

### type Var

```go
type Var interface {
  // String 返回变量的有效 JSON 值。不返回有效 JSON 的 String 方法类型（如 time.Time）不得用作 Var。
  String() string
}
```

Var 是所有导出变量的抽象类型。

#### func Get(name string) Var

Get 获取已命名的导出变量。如果变量名尚未注册，则返回 nil。