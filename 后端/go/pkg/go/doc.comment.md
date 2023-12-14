## package doc/comment

包注释实现了对 Go doc 注释（文档注释）的解析和重新格式化，文档注释是紧接在包、const、func、type 或 var 的顶层声明之前的注释。

Go doc 注释语法是 Markdown 的简化子集，支持链接、标题、段落、列表（无嵌套）和预格式化文本块。语法的详细信息请访问 https://go.dev/doc/comment。

要解析与 doc 注释相关的文本（移除注释标记后），请使用解析器：

```go
var p comment.Parser
doc := p.Parse(text)
```

结果是一个 *Doc。要将其重新格式化为文档注释、HTML、Markdown 或纯文本，请使用打印机：

```go
var pr comment.Printer
os.Stdout.Write(pr.Text(doc))
```

解析器和打印机类型都是结构体，其字段可以修改，以自定义操作。有关详细信息，请参阅这些类型的文档。

需要对重新格式化进行额外控制的用例可以通过检查解析的语法本身来实现自己的逻辑。有关其他类型的概述和链接，请参阅 Doc、Block 和 Text 文档。

## Index

### func DefaultLookupPackage(name string) (importPath string, ok bool)

DefaultLookupPackage 是默认的软件包查找函数，在 Parser.LookupPackage 为 nil 时使用。它能识别标准库中带有单元素导入路径（如数学）的软件包名称，否则将无法命名这些软件包。

请注意，go/doc 软件包会根据当前软件包中使用的导入提供更复杂的查找。

### type Block

```go
type Block interface {
  // 包含已过滤或未导出字段
}
```

块是文档注释中的块级内容，是*代码、*标题、*列表或*段落中的一种。

### type Code

```go
type Code struct {
  // 文本是预先格式化的文本，以换行符结束。
	// 文本可以是多行，每行都以换行符结束。
	// 文本永远不会为空，也不会以空行开始或结束。
  Text string
}
```

代码是预先格式化的代码块。

### type Doc

```go
type Doc struct {
	// Content 是注释中内容块的序列。
	Content []Block

	// Links 是评论中的链接定义。
	Links []*LinkDef
}
```

Doc 是经过解析的 Go 文档注释。

### type DocLink

```go
type DocLink struct {
  Text []Text // 链接文本

  // ImportPath、Recv 和 Name 可识别作为链接目标的 Go 软件包或符号。非空字段的可能组合是
  // - ImportPath：指向另一个软件包的链接
  // - ImportPath、Name：指向另一个软件包中 const、func、type 或 var 的链接
  // - ImportPath、Recv、Name：指向另一个软件包中方法的链接
  // - Name：指向此软件包中 const、func、类型或 var 的链接
  // - Recv, 名称：指向此软件包中方法的链接
  ImportPath string // 为方法 const、func、type、var 或方法名导入路径接收器类型，不含任何指针星型
  Recv string
  Name string
}
```

DocLink 是 Go 软件包或符号的文档链接。

#### func (l *DocLink) DefaultURL(baseURL string) string

DefaultURL 构建并返回 l 的文档 URL，使用 baseURL 作为指向其他软件包链接的前缀。

DefaultURL 返回的可能形式有

baseURL/ImportPath，用于指向其他软件包的链接
baseURL/ImportPath#Name，用于链接到其他软件包中的 const、func、type 或 var
baseURL/ImportPath#Recv.Name，用于链接到另一个软件包中的方法
#Name，用于链接到本软件包中的 const、func、类型或 var
#Recv.Name，用于指向本软件包中方法的链接
如果 baseURL 以斜线结尾，则 DefaultURL 会在锚定形式中的 ImportPath 和 # 之间插入斜线。例如，以下是一些 baseURL 值及其可生成的 URL：

```go
"/pkg/" → "/pkg/math/#Sqrt"
"/pkg"  → "/pkg/math#Sqrt"
"/"     → "/math/#Sqrt"
""      → "/math#Sqrt"
```

### type Heading

```go
type Heading struct {
  Text []Text // 标题文本
}
```

标题是文档注释的标题。

#### func (h *Heading) DefaultID() string

DefaultID 返回标题 h 的默认锚点 ID。

默认锚点 ID 是通过将所有非 ASCII 字母数字的符号转换为下划线，然后添加前缀 "hdr-"而生成的。例如，如果标题文本为 "Go Doc Comments"，则默认 ID 为 "hdr-Go_Doc_Comments"。

### type Italic

```go
type Italic string
```

斜体是指呈现为斜体文本的字符串。

### type Link

```go
type Link struct struct {
	Auto bool // 这是一个字面 URL 的自动（隐式）链接吗？
	Text []Text // 链接文本
	URL string // 链接的目标 URL
}
```

链接是指向特定 URL 的链接。

### type LinkDef

```go
type LinkDef struct {
	Text string // 链接文本
	URL string // 链接 URL
	Used bool // 注释是否使用该定义
}
```

LinkDef 是一个链接定义。

### type List

```go
type List struct {
  // Items 是列表项目。
  Items []*ListItem

  // ForceBlankBefore 表示在重新格式化注释时，列表前必须有一个空行，它优先于通常的条件。请参阅 BlankBefore 方法。
  // 
  // 注释解析器会为任何前面有空行的列表设置 ForceBlankBefore，以确保打印时保留空行。
  ForceBlankBefore bool

  // ForceBlankBetween 表示在重新格式化注释时，列表项之间必须用空行隔开，这优先于通常的条件。请参阅 BlankBetween 方法。
  // 
  // 注释解析器会为任何两个项目之间有空行的列表设置 ForceBlankBetween，以确保打印时保留空行。
  ForceBlankBetween bool
}
```

列表是编号列表或项目列表。列表总是非空的：len(Items) > 0。在编号列表中，每个 Items[i].Number 都是非空字符串。在项目列表中，每个 Items[i].Number 都是空字符串。

#### func (l *List) BlankBefore() bool

BlankBefore 报告注释的重新格式化是否应在列表前加入空行。默认规则与 [BlankBetween] 相同：如果列表项内容包含任何空行（即至少有一个项目包含多个段落），则列表本身之前必须有空行。可以通过设置 List.ForceBlankBefore 来强制前面的空行。

#### func (l *List) BlankBetween() bool

BlankBetween 报告注释的重新格式化是否应在每对列表项之间包含空行。默认规则是，如果列表项内容包含任何空行（即至少一个列表项包含多个段落），则列表项本身必须用空行分隔。可以通过设置 List.ForceBlankBetween 来强制使用空行分隔符。

### type ListItem

```go
type ListItem struct {
  // 在编号列表中，"编号 "是一个十进制字符串；在项目列表中，"编号 "是一个空字符串。
  Number string // "1"、"2"、......；""表示项目列表

  // Content 是列表内容。目前，由于解析器和打印机的限制，Content 的每个元素都必须是 *Paragraph。
  Content []Block // 该项目的内容。
}
```

ListItem 是编号列表或项目列表中的单个项目。

### type Paragraph

```go
type Paragraph struct {
  Text []Text
}
```

段落是一段文字。

### type Parser

```go
type Parser struct {
  // Words 是围棋标识符词的映射，这些标识符词应该被斜体化，并可能被链接。如果 Words[w] 是空字符串，那么 w 只被斜体化。否则，将使用 Words[w] 作为链接目标，对其进行链接。Words 相当于 [go/doc.ToHTML] words 参数。
  Words map[string]string

  // LookupPackage 可将软件包名称解析为导入路径。
  // 
  // 如果 LookupPackage(name) 返回 ok == true，则 [name]（或 [name.Sym] 或 [name.Sym.Method]）被视为指向 importPath 软件包文档的文档链接。返回 ""、true 也是有效的，在这种情况下，[name] 将被视为指向当前软件包。
  // 
  // 如果 LookupPackage(name) 返回 ok == false，那么 [name]（或 [name.Sym] 或 [name.Sym.Method]）将不被视为文档链接，除非 name 是标准库中某个包的完整（但单元素）导入路径，例如 [math] 或 [io.Reader]。为了允许引用具有相同包名的其他包的导入，仍会调用 LookupPackage。
  // 
  // 将 LookupPackage 设置为 "nil "相当于将其设置为一个总是返回""、false 的函数。
  LookupPackage func(name string) (importPath string, ok bool)

  // LookupSym 报告当前软件包中是否存在符号名或方法名。
  // 
  // 如果 LookupSym("", "Name")返回 true，则 [Name] 被视为 const、func、type 或 var 的文档链接。
	//
	// 同样，如果 LookupSym("Recv", "Name") 返回 true，那么 [Recv.Name] 将被视为类型 Recv 的方法 Name 的文档链接。
	//
	// 将 LookupSym 设置为 nil 相当于将其设置为一个总是返回 false 的函数。
  LookupSym func(recv, name string) (ok bool)
}
```

解析器是一个文档注释解析器。结构体中的字段可以在调用 Parse 之前填写，以便自定义解析过程的细节。

#### func (p *Parser) Parse(text string) *Doc

Parse 会解析文档注释文本，并返回 *Doc 表单。文本中的注释标记（/* // 和 */）必须已被删除。

### type Plain

```go
type Plain string
```

Plain 是以纯文本（非斜体）呈现的字符串。

### type Printer

```go
type Printer struct {
  // HeadingLevel 是 HTML 和 Markdown 标题使用的嵌套级别。如果 HeadingLevel 为零，则默认为第 3 级，即使用 <h3> 和 ###。
  HeadingLevel int

  // HeadingID 是一个函数，用于计算生成 HTML 和 Markdown 时标题 h 所使用的标题 ID（锚标签）。如果 HeadingID 返回空字符串，则省略标题 ID。如果 HeadingID 为空，则使用 h.DefaultID。
  HeadingID func(h *Heading) string

  // DocLinkURL 是一个计算给定 DocLink 的 URL 的函数。如果 DocLinkURL 为 nil，则使用 link.DefaultURL(p.DocLinkBaseURL)。
  DocLinkURL func(link *DocLink) string

  // DocLinkBaseURL is used when DocLinkURL is nil, passed to [DocLink.DefaultURL] to construct a DocLink's URL. See that method's documentation for details.
	DocLinkBaseURL string

	// TextPrefix is a prefix to print at the start of every line when generating text output using the Text method.
	TextPrefix string

	// TextCodePrefix is the prefix to print at the start of each preformatted (code block) line when generating text output, instead of (not in addition to) TextPrefix. If TextCodePrefix is the empty string, it defaults to TextPrefix+"\t".
	TextCodePrefix string

	// TextWidth is the maximum width text line to generate, measured in Unicode code points, excluding TextPrefix and the newline character. If TextWidth is zero, it defaults to 80 minus the number of code points in TextPrefix. If TextWidth is negative, there is no limit.
	TextWidth int
}
```

打印机是文档注释打印机。可以在调用任何打印方法之前填写结构体中的字段，以便自定义打印过程的细节。

#### func (p *Printer) Comment(d *Doc) []byte

注释返回文档的标准 Go 格式，不带任何注释标记。

#### func (p *Printer) HTML(d *Doc) []byte

HTML 返回文档的 HTML 格式。有关自定义 HTML 输出的方法，请参阅打印机文档。

#### func (p *Printer) Markdown(d *Doc) []byte

Markdown 返回文档的 Markdown 格式。有关自定义 Markdown 输出的方法，请参阅打印机文档。

#### func (p *Printer) Text(d *Doc) []byte

Text 返回文档的文本格式。有关自定义文本输出的方法，请参阅打印机文档。

### type Text

```go
type Text interface {
  // 包含已过滤或未导出的字段
}
```

文本是文档注释中的文本级内容，可选普通、斜体、*链接或*DocLink。