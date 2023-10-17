## package context

import "context"

包上下文定义了上下文类型，该类型可跨 API 边界和进程间传输截止日期、取消信号和其他请求范围值。

向服务器发送的请求应创建一个 Context，向服务器发出的调用应接受一个 Context。它们之间的函数调用链必须传播 Context，可选择用使用 WithCancel、WithDeadline、WithTimeout 或 WithValue 创建的派生 Context 将其替换。当一个 Context 被取消时，从它派生出来的所有 Context 也会被取消。

WithCancel、WithDeadline 和 WithTimeout 函数接收一个 Context（父节点），并返回一个派生 Context（子节点）和一个 CancelFunc。调用 CancelFunc 会取消子代及其子代，删除父代对子代的引用，并停止任何相关的计时器。不调用 CancelFunc 会泄漏子代及其子代，直到父代被取消或计时器启动。go vet 工具会检查所有控制流路径是否使用了 CancelFunc。

WithCancelCause 函数会返回一个 CancelCauseFunc，它会接收一个错误并将其记录为取消原因。在已取消的上下文或其任何子上下文上调用 Cause 可获取原因。如果没有指定原因，Cause(ctx) 返回与 ctx.Err() 相同的值。

使用上下文的程序应遵循这些规则，以保持各软件包接口的一致性，并使静态分析工具能够检查上下文传播：

```go
func DoSomething(ctx context.Context, arg Arg) error {
  // ... use ctx ...
}
```

即使函数允许，也不要传递 nil Context。如果不确定使用哪个 Context，请传递 context.TODO。

context Values 仅用于传输进程和 API 的请求范围数据，不用于向函数传递可选参数。

同一 Context 可传递给在不同程序中运行的函数；多个程序可安全地同时使用 Context。

使用 Contexts 的服务器示例代码请参见 https://blog.golang.org/context。

## Index

### Variables

```go
var Canceled = errors.New("content canceled")
```

Canceled 是取消上下文时 [Context.Err] 返回的错误。

```go
var DeadlineExceeded error = deadlineExceededError{}
```

DeadlineExceeded 是上下文的截止日期过去时 [Context.Err] 返回的错误。

### func AfterFunc(ctx Context, f func()) (stop func() bool) 添加于1.21

AfterFunc 安排在 ctx 完成（取消或超时）后在自己的 goroutine 中调用 f。如果 ctx 已经完成，AfterFunc 会立即在自己的例行程序中调用 f。

在一个上下文中对 AfterFunc 的多次调用是独立运行的，一个调用不会取代另一个调用。

如果调用停止了 f 的运行，则返回 true。如果 stop 返回 false，则要么上下文已完成，f 已在其自身的 goroutine 中启动；要么 f 已经停止。stop 函数在返回之前不会等待 f 完成。如果调用者需要知道 f 是否已完成，必须明确地与 f 协调。

如果 ctx 有 "AfterFunc(func()) func() bool "方法，AfterFunc 将使用该方法调度调用。

### func Cause(c Context) error 添加于1.2

Cause 返回一个非零错误，解释 c 被取消的原因。c 或其父节点的第一次取消会设置原因。如果取消是通过调用 CancelCauseFunc(err) 发生的，则 Cause 返回 err。否则，Cause(c) 返回与 c.Err() 相同的值。如果 c 尚未被取消，则 Cause 返回 nil。

### func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

WithCancel 返回父上下文的一个副本，并带有一个新的 Done 通道。当调用返回的取消函数或父上下文的 Done 通道关闭时（以先发生者为准），返回上下文的 Done 通道将被关闭。

取消此上下文会释放与其相关的资源，因此代码应在此上下文中运行的操作完成后立即调用 cancel。

### func WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc) 添加于1.2

WithCancelCause 的行为与 WithCancel 类似，但返回的是 CancelCauseFunc 而不是 CancelFunc。调用 cancel 时，如果错误（"原因"）为非零，则会在 ctx 中记录该错误；然后可以使用 Cause(ctx) 检索该错误。如果调用 cancel 时错误为空，则会将原因设置为 "已取消"。

用例：

```go
ctx, cancel := context.WithCancelCause(parent)
cancel(myError)
ctx.Err() // 返回 context.Canceled
context.Cause(ctx) // 返回 myError
```

### func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

如果父上下文的截止日期已经早于 d，那么 WithDeadline(parent, d) 在语义上等同于父上下文。当截止日期到期、调用返回的取消函数或父上下文的 Done 通道关闭时（以先发生者为准），返回的 [Context.Done] 通道将被关闭。

取消此上下文会释放与其相关的资源，因此代码应在此上下文中运行的操作完成后立即调用 cancel。

### func WithDeadlineCause(parent Context, d time.Time, cause error) (Context, CancelFunc)

WithDeadlineCause 的行为与 WithDeadline 类似，但也会在超过截止日期时设置返回 Context 的原因。返回的 CancelFunc 不会设置原因。

### func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

WithTimeout 返回 WithDeadline(parent, time.Now().Add(timeout)).

取消此上下文会释放与其相关的资源，因此代码应在此上下文中运行的操作完成后立即调用取消：

```go
func slowOperationWithTimeout(ctx context.Context) (Result, error) {
  ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
  defer cancel() // 如果 slowOperation 在超时前完成，则释放资源
  return slowOperation(ctx)
}
```

### func WithCancelCause(parent Context, timeout time.Duration, cause error) (Context, CancelFunc)

WithTimeoutCause 的行为与 WithTimeout 类似，但也会在超时时设置返回 Context 的原因。返回的 CancelFunc 不会设置原因。

### type CancelCauseFunc

```go
type CancelCauseFunc func(cause error)
```

CancelCauseFunc 的行为与 CancelFunc 类似，但会额外设置取消原因。可以通过在被取消的上下文或其任何派生上下文上调用 Cause 来获取该原因。

如果上下文已被取消，CancelCauseFunc 不会设置原因。例如，如果 childContext 从 parentContext 派生：

- 如果在使用 cause2 取消子上下文之前，使用 cause1 取消了父上下文，那么 Cause(parentContext) == Cause(childContext) == cause1
- 如果在用 cause1 取消 parentContext 之前用 cause2 取消了 childContext，则 Cause(parentContext) == cause1 和 Cause(childContext) == cause2

### type CancelFunc

```go
type CancelFunc func()
```

CancelFunc 命令操作放弃工作。CancelFunc 不会等待工作停止。CancelFunc 可以同时被多个程序调用。在第一次调用后，对 CancelFunc 的后续调用将不起作用。

### type Context

```go
type Context interface {
  // 截止日期返回代表此上下文完成的工作应取消的时间。
	// 应该取消的时间。如果没有设置截止日期，Deadline 将返回 ok==false
	// 设置。连续调用 Deadline 会返回相同的结果。
  Deadline() (deadline time.Time, ok bool)

	// Done 返回一个通道，当代表本上下文完成的工作应该取消时，该通道将被关闭。
	// 上下文应被取消。如果此上下文可以
	// 永远不会被取消。连续调用 Done 会返回相同的值。
	// Done 通道的关闭可以异步进行、
	// 在 cancel 函数返回后。
	//
	// WithCancel 安排在调用 cancel 时关闭 Done；
	// WithDeadline 安排在截止日期结束时关闭 Done； 
  // WithTimeout 安排在截止日期结束时关闭 Done。
	// 在超时时关闭 Done。
	// 结束时关闭 Done。
	//
	// Done 用于选择语句：
	//
	// // 流使用 DoSomething 生成值并将其发送到 out
	// // 直到 DoSomething 返回错误或 ctx.Done 关闭。
	//  func Stream(ctx context.Context, out chan<- Value) error {
	//  	for {
	//  		v, err := DoSomething(ctx)
	//  		if err != nil {
	//  			return err
	//  		}
	//  		select {
	//  		case <-ctx.Done():
	//  			return ctx.Err()
	//  		case out <- v:
	//  		}
	//  	}
	//  }
	//
	// 有关如何使用
	// 取消的通道。
	// 请参阅 https://blog.golang.org/pipelines，了解如何使用
	// 取消的通道。
  Done <-chan struct{}

	// 如果 Done 尚未关闭，Err 返回 nil。
	// 如果 Done 已关闭，Err 返回一个非零错误，解释原因：
	// 如果上下文被取消，则返回 Canceled
	// 如果上下文的截止日期已过，则返回 DeadlineExceeded。
	// Err 返回非零错误后，连续调用 Err 将返回相同的错误。
  Err() error

	// Value 返回 key 在此上下文中的相关值，如果 key 没有相关值，则返回 nil。
	// 如果 key 没有关联值。连续调用 Value
	// 返回相同的结果。
	//
	// 使用上下文值只能用于传输请求范围内的数据，而不能用于传输请求范围内的数据。
	// 进程和 API 边界的请求作用域数据，而不是用于将可选参数传递给
	// 函数的可选参数。
	//
	// 键标识上下文中的特定值。希望
	// 要在 Context 中存储值的函数通常会在一个全局变量中分配一个键，然后使用该键作为 context.WithVal 值的参数。
	然后使用该键作为 context.WithValue 和 // Context.Value 的参数。
	// Context.Value的参数。键可以是任何支持相等的类型；
	软件包应将 key 定义为未导出类型，以避免 // 碰撞。
	// 碰撞。
	//
	// 定义 Context 键的软件包应为使用该键存储的值提供类型安全的访问器
	// 为使用该键存储的值提供类型安全的访问器：
	//
	// // 包 user 定义了一个存储在 Context 中的 User 类型。
	// 包用户
	//
	// 导入 "上下文"
	//
	// // User 是存储在 Contexts 中的值类型。
	// type User struct {...}
	//
	// key 是本软件包中定义的键的未导出类型。
	// // 这样可以防止与其他软件包中定义的键发生冲突。
	// type key int
	//
	// // userKey 是 Contexts 中 user.User 值的键。它是
	// 用户使用 user.NewContext 和 user.FromContext
	// 而不是直接使用该键。
	// var userKey key
	//
	// // NewContext 返回一个携带值 u 的新 Context。
	// func NewContext(ctx context.Context, u *User) context.Context { // return context.WithValue.
	// return context.WithValue(ctx, userKey, u)
	// }
	//
	// // FromContext 返回存储在 ctx 中的用户值（如果有）。
	// func FromContext(ctx context.Context) (*User, bool) { // u, ok := ctx.
	// u, ok := ctx.Value(userKey).
	// return u, ok
	// }
  Value(key any) any
}
```

情境可跨 API 界携带截止日期、取消信号和其他值。

Context 的方法可同时被多个程序调用。

#### func Background() Context

Background 返回一个非零、空的 Context。它不会被取消，没有值，也没有截止时间。它通常用于主函数、初始化和测试，并作为传入请求的顶级 Context。

#### func TODO() Context

TODO 返回一个非空的 Context。如果不清楚要使用哪个 Context，或者还没有可用的 Context（因为周围的函数还没有扩展到接受 Context 参数），代码就应该使用 context.TODO。

#### func WithValue(parent Context, key, val any) Context

WithValue 返回父类的副本，其中与 key 关联的值为 val。

上下文值只能用于传输进程和 API 的请求范围数据，而不能用于向函数传递可选参数。

提供的 key 必须具有可比性，且不应是字符串或任何其他内置类型，以避免使用 context 的软件包之间发生冲突。WithValue 的用户应自行定义键的类型。为避免在赋值给接口{}时进行分配，上下文键通常采用具体的 struct{} 类型。或者，导出上下文键变量的静态类型应为指针或接口。

#### func WIthoutCancel(parent Context) Context

WithoutCancel 返回父节点的副本，当父节点被取消时，该副本不会被取消。返回的上下文不返回 Deadline 或 Err，其 Done 通道为 0。对返回的上下文调用 Cause 会返回 nil。