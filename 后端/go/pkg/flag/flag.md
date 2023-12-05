## package flag

软件包标志实现了命令行标志解析。

### Usage

使用 flag.String()、Bool()、Int() 等定义标志。

此处声明了一个整数标志-n，存储在指针 nFlag 中，类型为 *int：

```go
import "flag"
var nFlag = flag.Int("n", 1234, "help message for flag n")
```

如果您愿意，还可以使用 Var() 函数将标记与变量绑定。

```go
var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

或者，您也可以创建满足 Value 接口（带有指针接收器）的自定义标记，并通过以下方式将它们与标记解析结合起来

```go
flag.Var(&flagVal, "name", "help message for flagname")
```

对于此类标志，默认值只是变量的初始值。

定义完所有标记后，调用

```go
flag.Parse()
```

将命令行解析为已定义的标志。

然后就可以直接使用标志。如果使用标志本身，它们都是指针；如果绑定到变量，它们都是值。

```go
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```

解析后，标志后面的参数可作为片段 flag.Args() 或单独作为 flag.Arg(i)。参数的索引范围从 0 到 flag.NArg()-1。

### Command line flag syntax

允许使用以下形式：

```go
-flag
--flag   // 也允许使用双破折号
-flag=x
-flag x  // 仅限非布尔标志
```

可以使用一个或两个破折号，它们是等价的。布尔标志不允许使用最后一种形式，因为命令

```go
cmd -x *
```

其中 * 是一个 Unix shell 通配符，如果存在名为 0、false 等的文件，它就会改变。要关闭布尔标志，必须使用 -flag=false 形式。

标志解析会在第一个非标志参数（"-"为非标志参数）之前或结束符"--"之后停止。

整数标志可以是 1234、0664、0x1234，也可以是负数。布尔标志可以是:

```go
1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
```

持续时间标记接受任何对 time.ParseDuration.Duration.Duration 有效的输入。

默认的命令行标志集由顶级函数控制。FlagSet 类型允许定义独立的标志集，例如在命令行接口中实现子命令。FlagSet 的方法类似于命令行标志集的顶级函数。

## Index

### Variables

```go
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
```

CommandLine 是一组默认的命令行标志，由 os.Args 解析而来。BoolVar 和 Arg 等顶级函数是 CommandLine 方法的封装器。

```go
var ErrHelp = errors.New("flag: help requested")
```

ErrHelp 是在调用 -help 或 -h 标记但未定义此类标记的情况下返回的错误信息。

```go
var Usage = func() {
	fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	PrintDefaults()
}
```

Usage 会在 CommandLine 的输出（默认为 os.Stderr）中打印一条使用信息，记录所有已定义的命令行标志。如果在解析标志时发生错误，就会调用该函数。函数是一个变量，可以修改为指向自定义函数。默认情况下，它会打印一个简单的页眉并调用 PrintDefaults；有关输出格式和如何控制输出的详细信息，请参阅 PrintDefaults 文档。自定义使用函数可以选择退出程序；默认情况下，由于命令行的错误处理策略被设置为 ExitOnError，退出程序还是会发生。

### func Arg(i int) string

Arg 返回第 i 个命令行参数。Arg(0) 是处理完标志后剩余的第一个参数。如果请求的元素不存在，Arg 返回空字符串。

### func Args() []string

Args 返回非 flag 命令行参数。

### func Bool(name string, value bool, usage string) *bool

Bool 定义一个具有指定名称、默认值和使用字符串的 bool 标志。返回值是存储标志值的 bool 变量的地址。

### func BoolFunc(name, usage string, fn func(string) error) 添加于1.21.0

BoolFunc 用指定的名称和用法字符串定义一个标志，但不要求值。每次看到标志时，都会调用 fn 并返回标志的值。如果 fn 返回非零错误，则会被视为标志值解析错误。

### func BoolVar(p *bool, name string, value bool, usage string)

BoolVar 用指定的名称、默认值和使用字符串定义一个 bool 标志。参数 p 指向一个用于存储标志值的 bool 变量。

### func Duration(name string, value time.Duration, usage string) *time.Duration

Duration 定义具有指定名称、默认值和使用字符串的 time.Duration 标志。返回值是存储标志值的 time.Duration 变量的地址。该标志接受 time.ParseDuration.Duration 的可接受值。

### func DurationVar(p *time.Duration, name string, value time.Duration, usage string)

DurationVar 用指定的名称、默认值和使用字符串定义 time.Duration 标志。参数 p 指向用于存储标志值的 time.Duration 变量。该标志接受 time.ParseDuration.DurationVar 可以接受的值。

### func Float64(name string, value float64, usage string) *float64

Float64 定义了一个 float64 标志，并指定了名称、默认值和使用字符串。返回值是存储标志值的 float64 变量的地址。

### func Float64Var(p *float64, name string, value float64, usage string)

Float64Var 用指定的名称、默认值和使用字符串定义 float64 标志。参数 p 指向一个 float64 变量，用于存储标志的值。

### func Func(name string, usage int, fn func(string) error) 添加于1.16

Func 用指定的名称和用法字符串定义了一个标志。每次看到标志时，都会调用带有标志值的 fn。如果 fn 返回非零错误，则会被视为标志值解析错误。

### func Int(name string, value int, usage string) *int

Int 定义一个带有指定名称、默认值和使用字符串的 int 标志。返回值是存储标志值的 int 变量的地址。

### func int64(name string, value int64, usage string) *int64

Int64 定义了一个带有指定名称、默认值和使用字符串的 int64 标志。返回值是存储标志值的 int64 变量的地址。

### func Int64Var(p *int64, name string, value int64, usage string)

Int64Var 用指定的名称、默认值和使用字符串定义一个 int64 标志。参数 p 指向一个 int64 变量，用于存储标志的值。

### func IntVar(p *int, name string, value int, usage string)

IntVar 用指定的名称、默认值和使用字符串定义一个 int 标志。参数 p 指向一个 int 变量，用于存储标志的值。

### func NArg() int

NArg 是处理完标志后剩余的参数个数。

### func NFlag() int

NFlag 返回已设置的命令行标志数。

### func Parse()

解析 os.Args[1:] 中的命令行标志。必须在定义了所有标志之后、程序访问标志之前调用。

### func Parsed() bool

Parsed 报告是否已解析命令行标志。

### func PrintDefaults()

除非另有配置，否则 PrintDefaults 会在标准错误中打印使用信息，显示所有已定义的命令行标志的默认设置。对于整数值标志 x，默认输出格式为

```go
-x int
	usage-message-for-x (default 7)
```

除了名称为单字节的 bool 标志外，其他任何标志的使用信息都将显示在单独一行中。对于 bool 标志，类型会被省略，如果标志名称是一个字节，则使用信息会显示在同一行。如果默认值为零，则省略括号中的默认值。列出的类型（此处为 int）可以通过在标志的使用字符串中插入一个带反引的名称来更改；信息中的第一个此类项目将被视为在信息中显示的参数名称，并且在显示时会去掉信息中的反引。例如

```go
flag.String("I", "", "search `directory` for include files")
```

输出将是

```go
-I directory
	search directory for include files.
```

要更改标记信息的目的地，请调用 CommandLine.SetOutput 命令。

### func Set(name, value string) error

Set 设置指定命令行标志的值。

### func String(name string, value string, usage string) *string

String 用指定的名称、默认值和使用字符串定义一个字符串标志。返回值是存储标志值的字符串变量的地址。

### func StringVar(p *string, name string, value string, usage string)

StringVar 用指定的名称、默认值和使用字符串定义字符串标志。参数 p 指向一个字符串变量，用于存储标志的值。

### func TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, ...) 添加于1.19

TextVar 定义了一个具有指定名称、默认值和使用字符串的标志。参数 p 必须是一个指向变量的指针，该变量将保存标志的值，并且 p 必须实现 encoding.TextUnmarshaler。如果使用了标记，标记值将传递给 p 的 UnmarshalText 方法。默认值的类型必须与 p 的类型相同。

### func Uint(name string, value uint64, usage string) *uint64

Uint 用指定的名称、默认值和使用字符串定义一个 uint 标志。返回值是存储标志值的 uint 变量的地址。

### func Uint64(name string, value uint64, usage string) *uint64

Uint64 定义一个 uint64 标志，并指定名称、默认值和使用字符串。返回值是存储标志值的 uint64 变量的地址。

### func Uint64Var(p *uint64, name string, value uint64, usage string)

Uint64Var 用指定的名称、默认值和使用字符串定义一个 uint64 标志。参数 p 指向存储标志值的 uint64 变量。

### func UintVar(p *uint, name string, value uint, usage string)

UintVar 用指定的名称、默认值和使用字符串定义一个 uint 标志。参数 p 指向一个 uint 变量，用于存储标志的值。

### func UnquoteUsage(flag *Flag) (name string, usage string) 添加于1.5

UnquoteUsage 从标志的用法字符串中提取一个后引号名称，并返回该名称和未加引号的用法。如果给定 "一个要显示的名称"，则返回（"名称"、"一个要显示的名称"）。如果没有反引号，则名称是对标志值类型的推测，如果标志是布尔型，则返回空字符串。

### func Var(value Value, name string, usage string)

Var 用指定的名称和用法字符串定义一个标志。标志的类型和值由 Value 类型的第一个参数表示，该参数通常包含用户定义的 Value 实现。例如，调用者可以创建一个将逗号分隔字符串转化为字符串片段的标记，方法是赋予片段 Value 的方法；特别是，Set 会将逗号分隔字符串分解为片段。

### func Visit(fn func(*Flag))

Visit 按词典顺序访问命令行标志，对每个标志调用 fn。它只访问已设置的标志。

### func VisitAll(fn func(*Flag))

VisitAll 按词典顺序访问命令行标志，对每个标志调用 fn。它会访问所有标记，即使是未设置的标记。

### type ErrorHandling

```go
type ErrorHandling int
```

ErrorHandling 定义了 FlagSet.Parse 在解析失败时的处理方式。

```go
const (
	ContinueOnError ErrorHandling = iota // 返回描述性错误。
	ExitOnError                          // 调用 os.Exit(2) 或 -h/-help Exit(0)。
	PanicOnError                         // 使用描述性错误调用 panic。
)
```

如果解析失败，这些常量会导致 FlagSet.Parse 按照描述的方式运行。

### type Flag

```go
type Flag struct {
	Name     string // 命令行中的名称
	Usage    string // 帮助信息
	Value    Value  // 设定值
	DefValue string // 默认值（文本）；使用信息
}
```

旗帜代表旗帜的状态。

#### func Lookup(name string) *Flag

Lookup 返回指定命令行标志的 Flag 结构，如果不存在则返回 nil。

### type FlagSet

```go
type FlagSet struct {
	// Usage 是在解析标记时发生错误时调用的函数。该字段是一个函数（而不是方法），可以修改为指向自定义错误处理程序。调用 Usage 后会发生什么取决于 ErrorHandling 设置；对于命令行，默认为 ExitOnError，即调用 Usage 后退出程序。
	Usage func()
	// 包含已过滤或未导出的字段
}
```

FlagSet 表示一组已定义的标志。FlagSet 的零值没有名称，并具有 ContinueOnError 错误处理功能。

在一个 FlagSet 中，标志名称必须是唯一的。如果试图定义一个名称已被使用的标志，将导致程序崩溃。

#### func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet

NewFlagSet 返回一个新的空标志集，并带有指定的名称和错误处理属性。如果名称不为空，则会打印在默认使用信息和错误信息中。

#### func (f *FlagSet) Arg(i int) string

Arg 返回第 i 个参数。Arg(0) 是处理完标志后剩余的第一个参数。如果请求的元素不存在，Arg 返回空字符串。

#### func (f *FlagSet) Args() []string

Args 返回非标志参数。

#### func (f *FlagSet) Bool(name string, value bool, usage string) *bool

Bool 定义一个具有指定名称、默认值和使用字符串的 bool 标志。返回值是存储标志值的 bool 变量的地址。

#### func (f *FlagSet) BoolFunc(p *bool, name string, value bool, usage string) 添加于1.21.0

BoolFunc 用指定的名称和用法字符串定义了一个标志，但不要求值。每次看到标志时，都会调用 fn 并返回标志的值。如果 fn 返回非零错误，则会被视为标志值解析错误。

#### func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)

BoolVar 用指定的名称、默认值和使用字符串定义一个 bool 标志。参数 p 指向一个用于存储标志值的 bool 变量。

#### func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration

Duration 定义具有指定名称、默认值和使用字符串的 time.Duration 标志。返回值是存储标志值的 time.Duration 变量的地址。该标志接受 time.ParseDuration.Duration 的可接受值。

#### func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)

DurationVar 用指定的名称、默认值和使用字符串定义 time.Duration 标志。参数 p 指向用于存储标志值的 time.Duration 变量。该标志接受 time.ParseDuration.DurationVar 可以接受的值。

#### func (f *FlagSet) ErrorHandling() ErrorHandling 添加于1.10

ErrorHandling 返回所设置标志的错误处理行为。

#### func (f *FlagSet) Float64(name string, value float64, usage string) *float64

Float64 定义了一个 float64 标志，并指定了名称、默认值和使用字符串。返回值是存储标志值的 float64 变量的地址。

#### func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)

Float64Var 用指定的名称、默认值和使用字符串定义 float64 标志。参数 p 指向一个 float64 变量，用于存储标志的值。

#### func (f *FlagSet) Func(name, usage string, fn func(string) error) 添加于1.16

Func 用指定的名称和用法字符串定义了一个标志。每次看到标志时，都会调用带有标志值的 fn。如果 fn 返回非零错误，则会被视为标志值解析错误。

#### func (f *FlagSet) Init(name string, errorHandling ErrorHandling)

Init 为一个标志集设置名称和错误处理属性。默认情况下，零 FlagSet 使用空名称和 ContinueOnError 错误处理策略。

#### func (f *FlagSet) Int(name string, value int, usage string) *int

Int 定义一个带有指定名称、默认值和使用字符串的 int 标志。返回值是存储标志值的 int 变量的地址。

#### func (f *FlagSet) Int64(name string, value int64, usage string) *int64

Int64 定义了一个带有指定名称、默认值和使用字符串的 int64 标志。返回值是存储标志值的 int64 变量的地址。

#### func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)

Int64Var 用指定的名称、默认值和使用字符串定义一个 int64 标志。参数 p 指向一个 int64 变量，用于存储标志的值。

#### func (f *FlagSet) IntVar(p *int, name string, value int, usage string)

IntVar 用指定的名称、默认值和使用字符串定义一个 int 标志。参数 p 指向一个 int 变量，用于存储标志的值。

#### func (f *FlagSet) Lookup(name string) *Flag

Lookup 返回指定标志的 Flag 结构，如果不存在则返回 nil。

#### func (f *FlagSet) NArg() int

NArg 是处理完标志后剩余的参数个数。

#### func (f *FlagSet) NFlag() int

NFlag 返回已设置标志的个数。

#### func (f *FlagSet) Name() string 添加于1.10

Name 返回所设置标志的名称。

#### func (f *FlagSet) Output() io.Writer 添加于1.10

如果输出未设置或设置为 nil，则返回 os.Stderr。

#### func (f *FlagSet) Parse(arguments []string) error

Parse 从参数列表中解析标志定义，参数列表不应包括命令名称。必须在定义了 FlagSet 中的所有标志后、程序访问标志前调用。如果设置了 -help 或 -h 但未定义，返回值将是 ErrHelp。

#### func (f *FlagSet) Parsed() bool

Parsed 报告是否调用了 f.Parse。

#### func (f *FlagSet) PrintDefaults()

除非另有配置，否则 PrintDefaults 会将设置中所有已定义的命令行标志的默认值打印到标准错误中。更多信息，请参阅全局函数 PrintDefaults 的文档。

#### func (f *FlagSet) Set(name, value string) error

Set 设置指定标志的值。

#### func (f *FlagSet) SetOutput(output io.Writer)

SetOutput 设置使用和错误信息的目的地。如果输出为 nil，则使用 os.Stderr。

#### func (f *FlagSet) String(name string, value string, usage string) *string

String 用指定的名称、默认值和使用字符串定义一个字符串标志。返回值是存储标志值的字符串变量的地址。

#### func (f *FlagSet) StringVar(p *string, name string, value string, usage string)

StringVar 用指定的名称、默认值和使用字符串定义字符串标志。参数 p 指向一个字符串变量，用于存储标志的值。

#### func (f *FlagSet) TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, ...) 添加于1.19

TextVar 定义了一个具有指定名称、默认值和使用字符串的标志。参数 p 必须是一个指向变量的指针，该变量将保存标志的值，并且 p 必须实现 encoding.TextUnmarshaler。如果使用了标记，标记值将传递给 p 的 UnmarshalText 方法。默认值的类型必须与 p 的类型相同。

#### func (f *FlagSet) Uint(name string, value uint, usage string) *uint

Uint 用指定的名称、默认值和使用字符串定义一个 uint 标志。返回值是存储标志值的 uint 变量的地址。

#### func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64

Uint64 定义一个 uint64 标志，并指定名称、默认值和使用字符串。返回值是存储标志值的 uint64 变量的地址。

#### func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)

Uint64Var 用指定的名称、默认值和使用字符串定义一个 uint64 标志。参数 p 指向存储标志值的 uint64 变量。

#### func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)

UintVar 用指定的名称、默认值和使用字符串定义一个 uint 标志。参数 p 指向一个 uint 变量，用于存储标志的值。

#### func (f *FlagSet) Var(value Value, name string, usage string)

Var 用指定的名称和用法字符串定义一个标志。标志的类型和值由 Value 类型的第一个参数表示，该参数通常包含用户定义的 Value 实现。例如，调用者可以创建一个将逗号分隔字符串转化为字符串片段的标记，方法是赋予片段 Value 的方法；特别是，Set 会将逗号分隔字符串分解为片段。

#### func (f *FlagSet) Visit(fn func(*Flag))

Visit 按词典顺序访问标记，每个标记都调用 fn。它只访问已设置的标志。

#### func (f *FlagSet) VisitAll(fn func(*Flag))

VisitAll 按词典顺序访问标记，对每个标记调用 fn。它会访问所有标记，即使是未设置的标记。

### type Getter 添加于1.2

```go
type Getter interface {
	Value
	Get() any
}
```

Getter 是一个允许检索 Value 内容的接口。它封装了 Value 接口，而不是 Value 接口的一部分，因为它出现在 Go 1 及其兼容性规则之后。除了 Func 使用的类型外，本软件包提供的所有 Value 类型都满足 Getter 接口的要求。

### type Value

```go
type Value interface {
	String() string
	Set(string) error
}
```

Value 是存储在标志中的动态值的接口。(默认值用字符串表示）。

如果 Value 的 IsBoolFlag() bool 方法返回值为 true，命令行解析器会将 -name 等同于 -name=true 而不是使用下一个命令行参数。

每出现一个标志，Set 都会按命令行顺序调用一次。标志包可以使用零值接收器（如 nil 指针）调用 String 方法。