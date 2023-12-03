## package errors

包 errors 实现了处理错误的函数。

New 函数创建内容仅为文本信息的错误。

如果 e 的类型具有以下方法之一，则错误 e 会封装另一个错误

```go
Unwrap() error
Unwrap() []error
```

如果 e.Unwrap() 返回一个非零错误 w 或一个包含 w 的片段，那么我们说 e 封装了 w。如果一个 Unwrap 方法返回的[]error 包含零错误值，则该方法无效。

创建封装错误的简单方法是调用 fmt.Errorf，并对错误参数应用 %w verb：

```go
wrapsErr := fmt.Errorf("... %w ...", ..., err, ...)
```

错误的连续解包会创建一棵树。Is 和 As 函数在检查错误树时，首先检查错误本身，然后依次检查其每个子代的树（预排序、深度优先遍历）。

Is 会检查第一个参数的错误树，寻找与第二个参数匹配的错误。它会报告是否找到了匹配的错误。与简单的相等检查相比，Is 应优先使用：

```go
if errors.Is(err, fs.ErrExist)
```

优于

```go
if err == fs.ErrExist
```

因为如果 err wraps io/fs.ErrExist.ErrExist.ErrExist.ErrExist.ErrExist.ErrExist.ErrExist.ErrExist.ErrExist.ErrExist.ErrExist

As 会检查第一个参数的树，寻找可以赋值给第二个参数（必须是指针）的错误。如果成功，则执行赋值并返回 true。否则，返回 false。形式为

```go
var perr *fs.PathError
if errors.As(err, &perr) {
	fmt.Println(perr.Path)
}
```

优于

```go
if perr, ok := err.(*fs.PathError); ok {
	fmt.Println(perr.Path)
}
```

因为如果 err 封装了 *io/fs.PathError，前者就会成功。

## Index

### Variables

```go
var ErrUnsupported = New("unsupported operation")
```

ErrUnsupported 表示请求的操作因不支持而无法执行。例如，在使用不支持硬链接的文件系统时调用 os.Link。

函数和方法不应返回此错误，而应返回一个包含适当上下文的错误，该上下文应满足

```go
errors.Is(err, errors.ErrUnsupported)
```

要么直接封装 ErrUnsupported，要么实现 Is 方法。

函数和方法应记录在哪些情况下会返回封装 ErrUnsupported 的错误。

### func As(err error, target any) bool 添加于1.13

as 在 err 的树中查找第一个与 target 匹配的错误，如果找到，则将 target 设置为该错误值并返回 true。否则，返回 false。

错误树由 err 本身和反复调用 Unwrap 得到的错误组成。当 err 封装多个错误时，As 会先检查 err，然后深度遍历其子代。

如果错误的具体值可赋值给 target 指向的值，或者如果错误有 As(interface{}) bool 方法，且 As(target) 返回 true，则该错误与 target 匹配。在后一种情况下，As 方法负责设置 target。

一种错误类型可能会提供一个 As 方法，这样它就可以被当作另一种错误类型来处理。

如果 target 不是指向实现 error 的类型或任何接口类型的非空指针，As 就会慌乱。

### func Is(err, target error) bool 添加于1.13

报告错误树中的任何错误是否与目标匹配。

错误树由 err 本身和反复调用 Unwrap 得到的错误组成。当 err 封装了多个错误时，Is 会先检查 err，然后深度遍历其子代。

如果错误等于目标，或者实现了 Is(error) bool 方法，且 Is(target) 返回 true，则认为该错误与目标匹配。

错误类型可能会提供 Is 方法，这样它就可以被视为等同于现有的错误。例如，如果 MyError 定义了

```go
func (m MyError) Is(target error) bool { return target == fs.ErrExist }
```

则 Is(MyError{}, fs.ErrExist) 返回 true。标准库中的示例请参见 syscall.Errno.Is。Is 方法只能对 Err 和目标进行浅层比较，而不能调用 Unwrap。

### func Join(errs ...error) error 添加于1.20

Join 返回一个包含给定错误的错误。任何零错误值都会被丢弃。如果 Ers 中的每个值都为零，Join 返回零。错误格式为调用 Ers 中每个元素的 Error 方法得到的字符串的连接，每个字符串之间有换行符。

Join 返回的非零错误实现了 Unwrap() []error 方法。

### func New(text string) error

New 返回格式为给定文本的错误值。即使文本相同，每次调用 New 都会返回一个不同的错误值。

### func Unwrap(err error) error 添加于1.13

如果 Er 的类型包含返回错误的 Unwrap 方法，则 Unwrap 返回调用 Er 的 Unwrap 方法的结果。否则，Unwrap 返回 nil。

Unwrap 只调用形式为 "Unwrap() error "的方法。尤其是，Unwrap 不会解除 Join 返回的错误。