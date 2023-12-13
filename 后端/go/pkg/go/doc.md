## package doc

Package doc 可从 Go AST 中提取源代码文档。

## Index

### Variables

```go
var IllegalPrefixes = []string{
	"copyright",
	"all rights",
	"author",
}
```

IllegalPrefixes 是一个小写前缀列表，用于识别不属于文档注释的注释。这有助于避免将紧接在软件包声明之前的版权声明误解为文档注释的常见错误。

### func IsPredeclared(s string) bool 添加于1.8

IsPredeclared 报告 s 是否是预先声明的标识符。

### func Synopsis(text string) string 废除

### func ToHTML(w io.Writer, text string, words map[string]string) 废除

### func ToText(w io.Writer, text string, prefix, codePrefix string, width int) 废除

### type Example

```go
type Example struct {
	Name string // 示例项目的名称（包括可选后缀）
	Suffix string // 示例后缀，不含前导 '_'（仅由 NewFromFiles 填充）
	Doc string // 示例函数 doc string
	Code ast.Node
	Play *ast.File // 示例的整个程序版本
	Comments []*ast.CommentGroup
	Output string // 预期输出
	Unordered bool
	EmptyOutput bool // 预期空输出
	Order int // 原始源代码顺序
}
```

示例表示测试源文件中的一个示例函数。

#### func Examples(testFiles ...*ast.File) []*Example

Examples 返回在 testFiles 中找到的示例，按名称字段排序。顺序字段记录了遇到示例的顺序。直接调用 Examples 时不会填充后缀字段，只有 NewFromFiles 在 _test.go 文件中找到示例时才会填充后缀字段。

可播放示例必须位于名称以"_test "结尾的软件包中。在上述任何一种情况下，示例都是 "可播放的"（"播放 "字段非零）：

- 示例函数是自包含的：函数仅引用其他包中的标识符（或预先声明的标识符，如 "int"），且测试文件不包含点导入。
- 整个测试文件都是示例：文件中正好包含一个示例函数、零测试、模糊测试或基准函数，以及除示例函数外的至少一个顶级函数、类型、变量或常量声明。

### type Filter

```go
type Filter func(string) bool
```

### type Func

```go
type Func struct {
	Doc string
	Name string
	Decl *ast.FuncDecl

	// 方法
	// （对于函数，这些字段各自的值为零）
	Recv 字符串 // 实际接收器 "T "或"*T"，后面可能跟类型参数[P1, ..., Pn]
	Orig string // 原始接收器 "T "或"*T
	Level int // 嵌入级别；0 表示未嵌入

	// Examples 是与该函数或方法相关的示例的排序列表。
	// 函数或方法相关的示例排序列表。示例是从提供给 NewFromFiles 的 _test.go 文件中提取的。
	// 提供给 NewFromFiles。
	Examples []* Example
}
```

Func 是 func 声明的文档。

### type Mode

```go
type Mode int
```

模式值控制新建和从文件新建的操作。

```go
const (
	// AllDecls 表示提取所有软件包级 声明的文档，而不仅仅是导出声明。
	AllDecls Mode = 1 << iota

	// AllMethods 表示显示所有嵌入的方法，而不仅仅是 不可见（未导出）的匿名字段的方法。
	AllMethods

	// PreserveAST 规定不修改 AST。最初，为了在 godoc 中节省内存，AST 的部分内容（如函数体）会被 nil 掉，但并非所有程序都希望这样做。
	PreserveAST
)
```

### type Note 添加于1.1

```go
type Note struct {
	Pos, End token.Pos // 包含标记的注释的位置范围
	UID string // 与标记一起找到的 uid
	Body string // 注释正文文本
}
```

注释代表以 "MARKER(uid): note body "开头的标记注释。任何带有 2 个或更多大写[A-Z]字母标记和至少一个字符 uid 的注释都能被识别。uid后面的": "是可选项。备注会被收集到以备注标记为索引的 Package.Notes 地图中。

### type Package

```go
type Package struct {
  Doc string
  Name string
  importPath string
  Imports []string
  Filenames []string
  Notes map[string][]*Note

  // 已废弃： 为了向后兼容，Bugs 仍在使用，但所有新代码都应使用 Notes。
  Bugs []string

  // 宣言
  Consts []*Value
  Types []*Type
  Vars []*Value
  Funcs []*Func

  // Examples 是与软件包相关的示例的排序列表。示例是从提供给 NewFromFiles 的 _test.go 文件中提取的。
  Examples []*Example
  // 包含已过滤或未导出的字段
}
```

软件包是整个软件包的文档。

#### func New(pkg *ast.Package, importPath string, mode Mode) *Package

New 为给定的软件包 AST 计算软件包文档。New 拥有 AST pkg 的所有权，可以编辑或覆盖它。要填充 Examples 字段，请使用 NewFromFiles 并包含软件包的 _test.go 文件。

#### func NewFromFiles(fset *token.FileSet, files []*ast.File, importPath string, opts ...any) (*Package, error) 添加于1.14

NewFromFiles 计算软件包的文档。

软件包由 *ast.Files 列表和相应的文件集指定，文件集不能为空。在计算文档时，NewFromFiles 会使用所有提供的文件，因此调用者有责任只提供与所需构建上下文匹配的文件。"go/build".Context.MatchFile 可用于确定文件是否与具有所需 GOOS 和 GOARCH 值以及其他构建约束的构建上下文相匹配。软件包的导入路径由 importPath 指定。

在 _test.go 文件中找到的示例会根据其名称与相应的类型、函数、方法或软件包相关联。如果示例名称中有后缀，则在 Example.Suffix 字段中进行设置。名称畸形的示例将被跳过。

可以选择提供一个 Mode 类型的额外参数，以控制文档提取行为的底层方面。

除非 PreserveAST 模式位处于开启状态，否则 NewFromFiles 将获得 AST 文件的所有权并可对其进行编辑。

#### func (p *Package) Filter(f Filter)

f. TODO(gri)：将 "Type.Method "识别为名称。

#### func (p *Package) HTML(text string) []byte 添加于1.19

HTML 返回文档注释文本的 HTML 格式。

要自定义 HTML 的详细信息，请使用 Package.Printer 获取 comment.Printer，并在调用其 HTML 方法前对其进行配置。

#### func (p *Package) Markdown(text string) []byte 添加于1.19

Markdown 返回文档注释文本的格式化 Markdown。

要自定义 Markdown 的细节，请使用 Package.Printer 获取 comment.Printer，并在调用其 Markdown 方法前对其进行配置。

#### func (p *Package) Parser() *comment.Parser 添加于1.19

每次调用都会返回一个新的解析器，以便调用者在使用前对其进行自定义。

#### func (p *Package) Printer() *comment.Printer 添加于1.19

每次调用都会返回一个新打印机，以便调用者在使用前对其进行自定义。

#### func (p *Package) Synopsis(text string) string 添加于1.19

Synopsis 返回文本中第一句话的简洁版本。该句子在第一个句号之后结束，句号后面是空格，且前面没有一个大写字母，或者在第一个段落分隔符处结束。结果字符串没有 \n、\r 或 \t 字符，单词之间只使用单空格。如果文本以任何 IllegalPrefixes 开头，则结果字符串为空字符串。

#### func (p *Package) Text(text string) []byte 添加于1.19

文本返回文档注释文本的格式化文本，以 80 个 Unicode 代码点包装，并使用制表符表示代码块缩进。

要自定义格式化细节，请使用 Package.Printer 获取 comment.Printer，并在调用其 Text 方法前对其进行配置。

### type Type

```go
type Type struct {
  Doc string
  Name string
  Decl *ast.GenDecl

	// 相关声明
	Consts []*Value // 此类型常量（大部分）的排序列表
	Vars []*Value // （大部分）此类型变量的排序列表
	Funcs []*Func // 返回本类型的函数的排序列表
	Methods []*Func // 此类型方法（包括嵌入式方法）的排序列表

  // Examples 是与该类型相关的示例的排序列表。示例摘自提供给 NewFromFiles 的 _test.go 文件。
  Examples []*Example
}
```

Type 是类型声明的文档。

### type Value

```go
类型 Value struct {
	Doc string
	Names []string // 按声明顺序排列的 var 或 const 名称
	Decl *ast.GenDecl
	// 包含已过滤或未导出字段
}
```

Value 是 var 或 const 声明（可能已分组）的文档。