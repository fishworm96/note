## package fmt

软件包 fmt 使用类似于 C 的 printf 和 scanf 的函数实现了格式化 I/O。格式 "动词 "源自 C 语言，但更为简单。

### Printing

动词

一般：

```go
%v	默认格式的值打印结构体时，加号标记 (%+v) 会添加字段名称
%#v	值的 Go 表示语法
%T	表示值类型的 Go 表示语法
%%	百分号
```

布尔值：

```go
%t 真或假
```

整数：

```go
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式：U+1234，等价于"U+%04X"
```

浮点和复数成分：

```go
%b 无分量科学记数法，指数为 2 的幂、
	采用 strconv.FormatFloat 的方式，格式为 "b"、
	例如：-123456p-78
%E 科学记数法，例如：-1.234456e+78
%E 科学记数法，例如：-1.234456E+78
%f 小数点但无指数，如 123.456
%F %f 的同义词
大指数时使用 %g %e，否则使用 %f。精确度将在下文讨论。
%G %E 表示大指数，否则表示 %F
%x 十六进制符号（十进制二幂指数），如 -0x1.23abcp+20
%X 大写十六进制符号，如 -0X1.23ABCP+20
```

字符串和字节片（与这些动词等同处理）：

```go
%s 未解释的字符串字节或片段
%q 用 Go 语法安全转义的双引号字符串
%x 十六进制，小写，每个字节两个字符
%X 十六进制，大写，每个字节两个字符
```

切片：

```go
%p 第 0 个元素的地址，以基数 16 表示，前导为 0x
```

指针:

```go
%p 基 16 符号，前导为 0x
%b、%d、%o、%x 和 %X 动词也适用于指针、
将指针值格式化为整数。
```

%v 的默认格式为:

```go
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

对于复合对象，使用这些规则递归打印元素，布局如下：

```go
struct:             {field0 field1 ...}
array, slice:       [elem0 elem1 ...]
maps:               map[key1:value1 key2:value2 ...]
pointer to above:   &{}, &[], &map[]
```

宽度由动词前的可选十进制数指定。如果没有，宽度就是表示数值所需的宽度。精度在（可选）宽度之后，用一个句号和一个小数来指定。如果没有句号，则使用默认精度。如果句号后面没有数字，则精度为零。示例:

```go
%f:    默认宽度，默认精度
%9f    宽度9，默认精度
%.2f   默认宽度，精度2
%9.2f  宽度9，精度2
%9.f   宽度9，精度0    
```

宽度和精度的单位是 Unicode 码点，即符点（这与 C 的 printf 不同，后者的单位始终是字节）。(这与 C 的 printf 不同，printf 的单位总是以字节为单位。）可以用字符 "*"替换任一个或两个标志位，使其值从下一个操作数（格式化操作数之前的操作数）中获取，而下一个操作数的类型必须是 int。

对于大多数值，宽度是要输出的最小符文数，必要时可在格式化后的表格中填充空格。

不过，对于字符串、字节片和字节数组，精度限制的是要格式化的输入长度（而不是输出的大小），必要时会截断。通常以符节为单位，但对于使用 %x 或 %X 格式化的这些类型，则以字节为单位。

对于浮点数值，宽度设置字段的最小宽度，精度设置小数点后的位数（如果合适），但对于%g/%G，精度设置最大有效位数（去掉尾数零）。例如，给定 12.345，格式 %6.3f 打印 12.345，而 %.3g 打印 12.3。对于 %e、%f 和 %#g，默认精度为 6；对于 %g，则是唯一标识数值所需的最小位数。

对于复数，宽度和精度分别适用于两个部分，结果用括号表示，因此 %f 应用于 1.2+3.4i 会产生 (1.200000+3.400000i) 的结果。

当使用 %q 格式化单个整数码位或符文字符串（[]符文类型）时，无效的 Unicode 码位将被转换为 Unicode 替换字符 U+FFFD，如 strconv.QuoteRune 中所示。

其他标志：

```go
'+'	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
' '	对数值，正数前加空格而负数前加负号；
'-'	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
'#'	切换格式：
  	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）；
 	对%q（%#q），如果strconv.CanBackquote返回真会输出反引号括起来的未转义字符串；
 	对%U（%#U），输出Unicode格式后，如字符可打印，还会输出空格和单引号括起来的go字面值；
  	对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格；
'0'	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
```

不需要标记的动词会忽略标记。例如，没有其他的十进制格式，因此 %#d 和 %d 的行为完全相同。

每个类似 Printf 的函数都有一个 Print 函数，该函数不使用任何格式，相当于对每个操作数说 %v。另一个变体 Println 会在操作数之间插入空格并加上换行符。

无论使用哪种动词，如果操作数是接口值，则使用内部的具体值，而不是接口本身。因此:

```go
var i interface{} = 23
fmt.Printf("%v\n", i)
```

将打印 23。

除使用动词 %T 和 %p 打印外，对于实现某些接口的操作数，还需考虑特殊的格式问题。按应用顺序排列：

1.如果操作数是 reflect.Value，则操作数将被其持有的具体值替换，然后继续按下一条规则打印。

2.2. 如果操作数实现了 Formatter 接口，则将调用该接口。在这种情况下，动词和标志的解释由该实现控制。

3.如果 %v 动词与 # 标志（%#v）一起使用，且操作数实现了 GoStringer 接口，则将调用该接口。

如果格式（Println 等隐含 %v）对字符串有效（%s %q %v %x %X），则适用以下两条规则：

4.如果操作数实现了错误接口，将调用 Error 方法将对象转换为字符串，然后按照动词（如果有）的要求进行格式化。

5.如果操作数实现了 String() string 方法，则将调用该方法将对象转换为字符串，然后按照动词（如果有）的要求进行格式化。

对于片段和结构体等复合操作数，格式适用于每个操作数的递归元素，而不是整个操作数。因此，%q 将引用字符串片段的每个元素，而 %6.2f 将控制浮点数组每个元素的格式。

但是，当打印带有类似字符串的动词（%s %q %x %X）的字节片段时，它的处理方式与字符串相同，都是单个项目。

为了避免诸如

```go
type X string
func (x X) String() string { return Sprintf("<%s>", x) }
```

在循环使用前转换数值：

```go
func (x X) String() string { return Sprintf("<%s>", string(x)) }
```

自引用数据结构也可能触发无限递归，例如，如果片段类型具有 String 方法，那么片段本身就包含一个元素。不过，这种情况并不多见，软件包也不会对其进行保护。

打印结构体时，fft 无法调用格式化方法（如 Error 或 String），因此也不会调用未导出字段的格式化方法。

### 显式参数索引

在 Printf、Sprintf 和 Fprintf 中，默认行为是每个格式化动词对调用中传递的连续参数进行格式化。但是，紧接动词前的符号 [n] 表示将格式化第 n 个单引号参数。在宽度或精度的 "*"前使用同样的符号，可以选择持有该值的参数索引。在处理括号中的表达式 [n] 之后，除非另有指示，否则后续动词将使用参数 n+1、n+2 等。

例如

```go
fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
```

将产生 "22 11"，而

```go
fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
```

相当于

```go
fmt.Sprintf("%6.2f", 12.0)
```

将产生 " 12.00"。由于显式索引会影响后面的动词，因此可以通过重置第一个要重复的参数的索引，多次打印相同的值：

```go
fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
```

将产生 "16 17 0x10 0x11"。

### 格式错误

如果为动词提供的参数无效，例如向 %d 提供字符串，生成的字符串将包含问题描述，如以下示例所示：

```go
类型错误或未知动词：%!verb(type=value)
	Printf("%d", "hi")：        %!d(string=hi)
参数过多：参数过多： %!
	Printf("hi", "guys")：hi%!(EXTRA string=guys)
参数太少：参数太少：%!verb(缺失)
	Printf("hi%d"): hi%!d(MISSING)
宽度或精度为非整数：%!(BADWIDTH) 或 %!
	Printf("%*s", 4.5, "hi")：  %!(BADWIDTH)hi
	Printf("%.*s", 4.5, "hi")：%!(BADPREC)hi
参数索引无效或使用无效：%!(BADINDEX)
	Printf("%*[2]d", 7)：       %!d(BADINDEX)
	Printf("%.[2]d", 7)：       %!d(BADINDEX)
```

所有错误都以字符串"%!"开头，有时后面跟一个字符（动词），并以括号内的描述结束。

如果 Error 或 String 方法在被打印例程调用时触发了 panic，则 fmt 包会重新格式化 panic 中的错误信息，并在信息上标明它是通过 fmt 包产生的。例如，如果字符串方法调用 panic("bad")，生成的格式化信息将如下所示

```go
%!s(PANIC=bad)
```

%！s 只显示故障发生时使用的打印动词。然而，如果恐慌是由 Error 或字符串方法的 nil 接收器引起的，输出则是未装饰的字符串"<nil>"。

### 扫描

一组类似的函数会扫描格式化文本以生成数值。Scan、Scanf 和 Scanln 从 os.Stdin 读取；Fscan、Fscanf 和 Fscanln 从指定的 io.Reader 读取；Sscan、Sscanf 和 Sscanln 从参数字符串读取。

Scan、Fscan 和 Sscan 将输入中的换行符视为空格。

Scanln、Fscanln 和 Sscanln 会在换行处停止扫描，并要求项目后跟换行或 EOF。

Scanf、Fscanf 和 Sscanf 根据格式字符串解析参数，类似于 Printf 的格式字符串。在下文中，"空格 "指除换行符外的任何 Unicode 空格符。

在格式字符串中，由 % 字符引入的动词会消耗和解析输入；下文将详细介绍这些动词。格式中除 %、空格或换行符之外的其他字符都会准确消耗该输入字符，且该字符必须存在。格式字符串中一个换行符前有零个或多个空格，则输入中会出现零个或多个空格，随后出现一个换行符或输入结束符。格式字符串中换行后的空格会在输入中消耗零个或多个空格。否则，格式字符串中任何一个或多个空格都会在输入中消耗尽可能多的空格。除非格式字符串中的空格与换行符相邻，否则该空格必须至少占用输入中的一个空格或找到输入的末尾。

对空格和换行符的处理与 C 的 scanf 系列不同：在 C 中，换行符与其他空格一样处理，在格式字符串中运行空格时，如果发现输入中没有空格，则不会出错。

动词的行为与 Printf 类似。例如，%x 会将整数扫描为十六进制数，而 %v 则会扫描数值的默认表示格式。Printf 动词 %p 和 %T 以及标志 # 和 + 未执行。对于浮点数和复数值，所有有效的格式化动词（%b %e %E %f %F %g %G %x %X 和 %v）都是等价的，并接受十进制和十六进制符号（例如："2.3e+7"、"0x4.5p-8"）和数字分隔下划线（例如："3.14159_26535_89793"）。

动词处理的输入是隐式空格限制的：除了 %c 以外，每个动词的执行都会从剩余输入开始丢弃前导空格，而 %s 动词（以及 %v 读入字符串）会在第一个空格或换行符时停止消耗输入。

在扫描无格式的整数或使用 %v verb 时，可以使用我们熟悉的基数设置前缀 0b（二进制）、0o 和 0（八进制）以及 0x（十六进制），也可以使用数字分隔下划线。

宽度在输入文本中解释，但没有精确扫描的语法（没有 %5.2f，只有 %5f）。如果提供了宽度，则会在修剪前导空格后使用，并指定为满足动词要求而要读取的最大符文数。例如

```go
Sscanf(" 1234567 ", "%5s%d", &s, &i)
```

将 s 设置为 "12345"，将 i 设置为 67，而

```go
Sscanf(" 12 34 567 ", "%5s%d", &s, &i)
```

将 s 设置为 "12"，将 i 设置为 34。

在所有扫描功能中，紧接着换行符的回车符被视为普通换行符（\r\n 的意思与 \n 相同）。

在所有扫描函数中，如果操作数实现了 Scan 方法（即实现了 Scanner 接口），该方法将用于扫描该操作数的文本。此外，如果扫描的参数数少于提供的参数数，则会返回错误信息。

所有要扫描的参数必须是基本类型的指针或 Scanner 接口的实现。

与 Scanf 和 Fscanf 类似，Sscanf 不需要消耗整个输入。无法恢复 Sscanf 使用了多少输入字符串。

注意：Fscan 等可以读取超出其返回输入的一个字符（符文），这意味着循环调用扫描例程可能会跳过部分输入。通常只有在输入值之间没有空格时才会出现这种问题。如果提供给 Fscan 的阅读器实现了 ReadRune，那么将使用该方法读取字符。如果读取器也实现了 UnreadRune，则该方法将用于保存字符，连续调用不会丢失数据。要为不具备 ReadRune 和 UnreadRune 功能的阅读器附加 ReadRune 和 UnreadRune 方法，请使用 bufio.NewReader.ReadRune 方法。

## Index

### func Append(b []byte, a ...any) []byte 添加于1.19

附加格式使用操作数的默认格式，将结果附加到字节片段，并返回更新后的片段。

### func Appendf(b []byte, format string, a ...any) []byte 添加于1.19

Appendf 根据格式指定符进行格式化，将结果追加到字节片段，并返回更新后的片段。

### func Appendln(b []byte, a ...any) []byte 添加于1.19

Appendln 使用操作数的默认格式进行格式化，将结果追加到字节片段，并返回更新后的片段。操作数之间总是添加空格，并附加换行符。

### func Errorf(format string, a ...any) error

Errorf 根据格式指定符进行格式化，并将字符串作为满足错误的值返回。

如果格式指定符包含一个带有 error 操作数的 %w verb，则返回的 error 将执行一个 Unwrap 方法，返回该操作数。如果有多个 %w verb，返回的 error 将执行 Unwrap 方法，返回一个包含所有 %w 操作数的[]error，顺序与参数中出现的顺序一致。向 %w verb 提供未实现错误接口的操作数是无效的。另外，%w 动词是 %v 的同义词。

### func FormatString(sate State, varb rune) string 添加于1.20

FormatString 返回一个字符串，表示 State 捕捉到的完全合格的格式化指令，后面跟参数动词（State 本身不包含动词）。缺失的标志、宽度和精度将被省略。该函数允许 Formatter 重构触发 Format 调用的原始指令。

### func Fprint(w io.Writer, a ...any) (n int, err error)

当操作数都不是字符串时，会在操作数之间添加空格。它会返回写入的字节数和遇到的任何写入错误。

### func Fprintf(w io.Writer, format string, a ...any) (n int, err error)

它返回写入的字节数和遇到的任何写入错误。

### func Fprintln(w io.Writer, a ...any) (n int, err error)

Fprintln 使用默认格式对操作数进行格式化，并写入 w。它会返回写入的字节数和遇到的任何写入错误。

### func Fscan(r io.Reader, a ...any) (n int, err error)

Fscan 扫描从 r 读取的文本，将连续的空格分隔值存储到连续的参数中。换行算作空格。它会返回成功扫描的条目数。如果少于参数个数，err 将报告原因。

### func Fscanf(r io.Reader, a ...any) (n int, err error)

Fscanf 扫描从 r 读取的文本，根据格式将连续的空格分隔值存储到连续的参数中。它返回成功解析的条目数。输入中的换行符必须与格式中的换行符一致。

### func Fscanln(r io.Reader, a ...any) (n int, err error)

Fscanln 与 Fscan 类似，但在换行时停止扫描，最后一个项目后必须有换行或 EOF。

### func Print(a ...any) (n int, err error)

打印格式使用操作数的默认格式，并写入标准输出。当操作数都不是字符串时，会在操作数之间添加空格。它会返回写入的字节数和遇到的任何写入错误。

### func Printf(format string, a ...any) (n int, err error)\

Printf 根据格式指定符进行格式化，并写入标准输出。它会返回写入的字节数和遇到的任何写入错误。

### func Println(a ...any) (n int, err error)

Println 使用操作数的默认格式进行格式化，并写入标准输出。操作数之间总是加上空格，并附加换行符。它会返回写入的字节数和遇到的任何写入错误。

### func Scan(a ...any) (n int, err error)

Scan 扫描从标准输入端读取的文本，将连续的空格分隔值存储到连续的参数中。换行算作空格。它会返回成功扫描的条目数。如果少于参数数，err 将报告原因。

### func Scanf(format string, a ...any) (n int, err error)

Scanf 扫描从标准输入端读取的文本，根据格式将连续的空格分隔值存储到连续的参数中。它会返回成功扫描的条目数。如果少于参数个数，err 将报告原因。输入中的换行符必须与格式中的换行符一致。但有一个例外：动词 %c 总是扫描输入中的下一个符，即使是空格（或制表符等）或换行符。

### func Scanln(a ...any) (n int, err error)

Scanln 与 Scan 类似，但在换行符处停止扫描，并且在最后一个项目后必须有换行符或 EOF。

### func Sprint(format string, a ...any) string

Sprint 格式使用操作数的默认格式，并返回结果字符串。当操作数都不是字符串时，会在操作数之间添加空格。

### func Sprintf(format string, a ...any) string

Sprintf 根据格式指定符进行格式化，并返回结果字符串。

### func Sprintln(a ...any) string

Sprintln 使用操作数的默认格式进行格式化，并返回结果字符串。操作数之间总是加上空格，并附加换行符。

### func Sscan(str string, a ...any) (n int, err error)

Sscan 扫描参数字符串，将连续的空格分隔值存储到连续的参数中。换行算作空格。它会返回成功扫描的条目数。如果少于参数数，err 将报告原因。

### func Sscanf(str string, format string, a ...any) (n int, err error)

Sscanf 扫描参数字符串，根据格式将连续的空格分隔值存储到连续的参数中。它会返回成功解析的条目数。输入中的换行符必须与格式中的换行符一致。

### func Sscanln(str string, a ...any) (n int, err error)

Sscanln 与 Sscan 类似，但在换行时停止扫描，最后一个项目后必须有换行或 EOF。

### type Formatter

```go
type Formatter interface {
  Format(f State, verb rune)
}
```

Formatter 由任何具有 Format 方法的值实现。该实现控制如何解释状态和符文，并可调用 Sprint() 或 Fprint(f) 等生成输出。

### type GoStringer

```go
type GoStringer interface {
  GoString() string
}
```

GoStringer 由任何具有 GoString 方法的值实现，该方法定义了该值的 Go 语法。GoString 方法用于将作为操作数传递的值打印为 %#v 格式。

### type ScanState

```go
type ScanState interface {
	// ReadRune 从输入中读取下一个符节（Unicode 代码点）。如果在 Scanln、Fscanln 或 Sscanln 期间调用，ReadRune() 将在返回第一个"\n "后或在读取超出指定宽度时返回 EOF。
	ReadRune() (r rune, size int, err error)
	// UnreadRune 会导致下一次调用 ReadRune 时返回相同的符文。
	UnreadRune() error
	// SkipSpace 跳过输入中的空格。换行符会根据所执行的操作进行适当处理；更多信息请参阅软件包文档。
	SkipSpace()
	// 如果 skipSpace 为 true，令牌会跳过输入中的空格，然后返回满足 f(c) 的 Unicode 代码点 c 的运行。 如果 f 为空，则使用 !unicode.IsSpace(c) ；也就是说，令牌将保留非空格字符。换行符将根据所执行的操作进行适当处理；更多信息请参阅软件包文档。返回的片段指向共享数据，可能会被下一次调用 Token、调用使用 ScanState 作为输入的 Scan 函数或调用的 Scan 方法返回时覆盖。
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
	// Width 返回宽度选项的值及其是否已设置。单位为 Unicode 代码点。
	Width() (wid int, ok bool)
	// 由于 ReadRune 是由接口实现的，因此扫描例程不应调用 Read，而 ScanState 的有效实现可以选择总是从 Read 返回错误。
	Read(buf []byte) (n int, err error)
}
```

ScanState 表示传递给自定义扫描器的扫描器状态。扫描程序可以进行运行时扫描，也可以要求 ScanState 发现下一个空格分隔标记。

### type Scanner

```go
type Scanner interface {
  Scan(state ScanState, verb rune) error
}
```

Scanner 由任何具有 Scan 方法的值实现，该方法扫描输入以查找值的表示形式，并将结果存储在接收器中，接收器必须是指针才有用。Scan、Scanf 或 Scanln 的任何参数都会调用 Scan 方法。

### type State

```go
type State interface {
	// Write 是调用的函数，用于输出格式化的打印输出。
	Write(b []byte) (n int, err error)
	// Width 返回宽度选项的值以及是否已设置。
	Width() （wid int, ok bool）
	// Precision 返回精度选项的值以及是否已设置。
	Precision() （prec int, ok bool）

	// 标志报告标志 c（一个字符）是否已设置。
	Flag(c int) bool
}
```

State 表示传递给自定义格式器的打印机状态。它提供了对 io.Writer 接口的访问，以及操作数格式指定器的标志和选项信息。

### type Stringer

```go
type Stringer interface {
	String() string
}
```

Stringer 由任何具有 String 方法的值实现，该方法定义了该值的 "本地 "格式。String 方法用于将作为操作数传递的值打印到任何可接受字符串的格式或未格式化的打印机（如 Print）。