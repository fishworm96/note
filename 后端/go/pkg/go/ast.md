## package ast

Package ast 声明了用于表示 Go 包语法树的类型。

## Index

### func FileExports(src *File) bool

FileExports 会就地修剪 Go 源文件的 AST，只保留导出节点：删除所有未导出的顶级标识符及其相关信息（如类型、初始值或函数体）。未导出的字段和导出类型的方法也会被删除。文件注释列表不变。

FileExports 会报告是否有导出声明。

### func FilterDecl(decl Decl, f Filter) bool

FilterDecl 通过删除所有未通过过滤器 f 的名称（包括结构字段和接口方法名称，但不包括参数列表中的名称），对 Go 声明的 AST 进行修剪。

FilterDecl 会报告过滤后是否还有任何声明名称。

### func FilterFile(src *File, f Filter) bool

FilterFile 通过删除顶层声明中所有未通过过滤器 f 的名称（包括结构字段和接口方法名称，但不包括参数列表），对 Go 文件的 AST 进行就地修剪。导入声明始终会被删除。File.Comments 列表不会更改。

FilterFile 会报告过滤后是否还有任何顶层声明。

### func FilterPackage(pkg *Package, f Filter) bool

FilterPackage 通过删除顶层声明（包括结构字段和接口方法名称，但不包括参数列表）中所有未通过过滤器 f 的名称，对 Go 软件包的 AST 进行修剪。pkg.Files 列表不会改变，因此文件名和顶层软件包注释不会丢失。

FilterPackage 会报告过滤后是否还有任何顶级声明。

### func Fprint(w io.Writer, fset *token.FileSet, x any, f FieldFilter) error

如果 fset != nil，位置信息将相对于该文件集进行解释。否则，位置信息将以整数值打印（文件集特定偏移）。

可以提供一个非零的字段过滤器 f 来控制输出：f(fieldname, fieldvalue) 为 true 的结构体字段将被打印；所有其他字段将从输出中过滤掉。未导出的 struct 字段将永不打印。

### func Inspect(node Node, f func(Node) bool)

Inspect 以深度优先的顺序遍历 AST：它首先调用 f(node)；node 必须不是 nil。如果 f 返回 true，Inspect 会对 node 的每个非 nil 子节点递归调用 f，然后再调用 f(nil)。

### func IsExported(name string) bool

IsExported 报告名称是否以大写字母开头。

### func IsGenerated(file *File) bool 添加于1.21.0

IsGenerated 通过检测 https://go.dev/s/generatedcode 上描述的特殊注释，报告文件是否由程序生成，而非手写。

语法树必须使用 ParseComments 标志进行过解析。示例：

```go
f, err := parser.ParseFile(fset, filename, src, parser.ParseComments|parser.PackageClauseOnly)
if err != nil { ... }
gen := ast.IsGenerated(f)
```

### func NotNilFilter(_ string, v reflect.Value) bool

如果字段值不是 nil，NotNilFilter 返回 true；否则返回 false。

### func PackageExports(pkg *Package) bool

PackageExports 会对 Go 软件包的 AST 进行修剪，只保留导出的节点。pkg.Files 列表不会改变，因此文件名和顶层软件包注释不会丢失。

PackageExports 会报告是否有导出声明；否则返回 false。

### func Print(fset *token.FileSet, x any) error

Print 将 x 打印到标准输出，跳过 nil 字段。Print(fset, x) 与 Fprint(os.Stdout, fset, x, NotNilFilter) 相同。

### func SortImports(fset *token.FileSet, f *File)

SortImports 会对 f 中导入块中的连续导入行进行排序，并在不丢失数据的情况下删除重复导入。

### func Walk(v Visitor, node Node)

Walk 以深度优先的顺序遍历 AST：开始时，它调用 v.Visit(node)；节点必须不是 nil。如果 v.Visit(node)返回的访问者 w 不是 nil，Walk 就会对 node 的每个非 nil 子节点递归调用访问者 w，然后再调用 w.Visit(nil)。

### type ArrayType

```go
type ArrayType struct {
	Lbrack token.Pos // "["的位置
	Len Expr // [...]T 数组类型的椭圆节点，片段类型为 nil
	Elt Expr // 元素类型
}
```

ArrayType 节点表示数组或片段类型。

#### func (x *ArrayType) End() token.Pos

#### func (x *ArrayType) Pos() token.Pos

### type AssignStmt

```go
type AssignStmt struct {
	Lhs []Expr
	TokPos token.Pos // Tok 的位置
	Tok token.Token // 赋值标记，DEFINE
	Rhs []Expr
}
```

AssignStmt 节点表示赋值或简短的变量声明。

#### func (s *AssignStmt) End() token.Pos

#### func (s *AssignStmt) Pos() token.Pos

### type BadDecl


```go
type BadDecl struct {
	From, To token.Pos // 坏声明的位置范围
}
```

BadDecl 节点是包含语法错误的声明的占位符，无法为其创建正确的声明节点。

#### func (d *BadDecl) End() token.Pos

#### func (d *BadDecl) Pos() token.Pos

### type BadExpr

```go
type BadExpr struct {
	From, To token.Pos // 坏表达式的位置范围
}
```

BadExpr 节点是表达式的占位符，其中包含语法错误，无法创建正确的表达式节点。

#### func (x *BadExpr) End() token.Pos

#### func (x *BadExpr) Pos() token.Pos

### type BadStmt

```go
type BadStmt struct {
	From, To token.Pos // 坏语句的位置范围
}
```

BadStmt 节点是语法错误语句的占位符，这些语句无法创建正确的语句节点。

#### func (x *BadStmt) End() token.Pos

#### func (x *BadStmt) Pos() token.Pos

### type BasicLit

```go
type BasicLit struct {
	ValuePos token.Pos // 字面位置
	Kind token.Token // token.INT、token.FLOAT、token.IMAG、token.CHAR 或 token.STRING
	Value string // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`.
}
```

BasicLit 节点表示基本类型的文字。
#### func (x *BasicLit) End() token.Pos

#### func (x *BasicLit) Pos() token.Pos

### type BinaryExpr

```go
type BinaryExpr struct {
	X Expr // 左操作数
	OpPos token.Pos // 运算符的位置
	Op token.Token // 运算符
	Y Expr // 右操作数
}
```

BinaryExpr 节点表示二进制表达式。

#### func (x *BinaryExpr) End() token.Pos

#### func (x *BinaryExpr) Pos() token.Pos

### type BlockStmt

```go
type BlockStmt struct {
	Lbrace token.Pos // "{" 的位置
	List []Stmt
	Rbrace token.Pos // "}"的位置（如果有的话）（可能因语法错误而不存在
}
```

BlockStmt 节点表示一个括号语句列表。

#### func (s *BlockStmt) End() token.Pos

#### func (s *BlockStmt) Pos() token.Pos

### type BranchStmt

```go
type BranchStmt struct {
	TokPos token.Pos // Tok 的位置
	Tok token.Token // 关键字标记（BREAK、CONTINUE、GOTO、FALLTHROUGH）
	Label *Ident // 标签名称；或 nil
}
```

BranchStmt 节点表示中断、继续、GOTO 或突破语句。

#### func (s *BranchStmt) End() token.Pos

#### func (s *BranchStmt) Pos() token.Pos

### type CallExpr

```go
type CallExpr struct {
	Fun Expr // 函数表达式
	Lparen token.Pos // "("的位置
	Args []Expr // 函数参数；或 nil
	Ellipsis token.Pos // "...... "的位置（如果没有"......"，则 token.NoPos）
	Rparen token.Pos // ") "的位置
}
```

CallExpr 节点表示一个表达式，后面跟一个参数列表。

#### func (x *CallExpr) End() token.Pos

#### func (x *CallExpr) Pos() token.Pos

### type CaseClause

```go
type CaseClause struct {
	Case token.Pos // "情况 "或 "默认 "关键字的位置
	List []Expr // 表达式或类型列表；nil 表示默认情况
	Colon token.Pos // ": "的位置
	Body []Stmt // 语句列表；或 nil
}
```

CaseClause 表示表达式或类型转换语句的一种情况。

#### func (x *CaseClause) End() token.Pos

#### func (x *CaseClause) Pos() token.Pos

### type ChanDir

```go
type ChanDir int
```

通道类型的方向由位掩码指示，该位掩码包括以下一个或两个常量。

```go
const (
	SEND ChanDir = 1 << iota
	RECV
)
```

### type ChanType

```go
type ChanType struct {
	Begin token.Pos // "chan "关键字或"<-"的位置（以先到者为准）
	Arrow token.Pos // "<-"的位置（如果没有"<-"，则为 token.NoPos）
	Dir ChanDir // 通道方向
	值 Expr // 值类型
}
```

ChanType 节点表示通道类型。

#### func (x *ChanType) End() token.Pos

#### func (x *ChanType) Pos() token.Pos

### type CommClause

```go
type CommClause struct {
	Case token.Pos // "情况 "或 "默认 "关键字的位置
	Comm Stmt // 发送或接收语句；nil 表示默认情况
	Colon token.Pos // ": "的位置
	Body []Stmt // 语句列表；或 nil
}
```

CommClause 节点表示选择语句的一种情况。

#### func (x *CommClause) End() token.Pos

#### func (x *CommClause) Pos() token.Pos

### type Comment

```go
type Comment struct {
	Slash token.Pos // 注释开头"/"的位置
	Text string // 注释文本（对于 // 样式的注释，"//n "不包括在内）
}
```

一个 Comment 节点代表一条 //-style 或 /*-style 注释。

文本字段包含源代码中可能存在的不含回车（\r）的注释文本。由于注释的结束位置是使用 len(Text) 计算的，因此 End() 报告的位置与包含回车的注释的真正源结束位置不一致。

#### func (c *Comment) End() token.Pos

#### func (c *Comment) Pos() token.Pos

### type CommentGroup

```go
type CommentGroup struct {
	List []*Comment // len(List) > 0
}
```

一个 CommentGroup 表示一个没有其他标记、中间没有空行的注释序列。

#### func (g *CommentGroup) End() token.Pos

#### func (g *CommentGroup) Pos() token.Pos

#### func (g *CommentGroup) Text() string

文本返回注释文本。注释标记（//、/* 和 */）、行注释的第一个空格以及前导空行和尾部空行都会被删除。注释指令（如"//line "和"//go:noinline"）也会被删除。多个空行会被缩减为一个，行尾空格也会被修剪。除非结果为空，否则将以换行结束。

### type CommentMap 添加于1.1

```go
type CommentMap map[Node][]*CommentGroup
```

CommentMap 将 AST 节点映射到与其关联的注释组列表。有关关联的描述，请参阅 NewCommentMap。

#### func NewCommentMap(fset *token.FileSet, node Node, comments []*CommentGroup) CommentMap 添加于1.1

NewCommentMap 将注释列表中的注释组与 node 指定的 AST 节点关联起来，创建新的注释映射。

在以下情况下，评论组 g 与节点 n 关联：

g 开始于 n 结束的同一行
g 开始于紧接 n 之后的一行，并且在 g 之后和下一个节点之前至少有一个空行
g 开始于 n 之前，且未通过前面的规则与 n 之前的节点关联
NewCommentMap 会尽量将评论组与 "最大 "的节点关联：例如，如果注释是一个赋值后面的行注释，那么该注释将与整个赋值相关联，而不是只与赋值中的最后一个操作数相关联。

#### func (cmap CommentMap) Comments() []*CommentGroup 添加于1.1

Comments 返回评论地图中的评论组列表。结果按来源顺序排序。

#### func (cmap CommentMap) Filter(node Node) CommentMap 添加于1.1

筛选器会返回一个新的注释映射，该映射只包含 cmap 中那些在 node 指定的 AST 中存在相应节点的条目。

#### func (cmap CommentMap) String() string 添加于1.1

#### func (cmap CommentMap) Update(old, new Node) Node 添加于1.1

更新会用新节点替换注释图中的旧节点，并返回新节点。与旧节点关联的注释将与新节点关联。

### type CompositeLit

```go
type CompositeLit struct {
	Type Expr // 字面类型；或 nil
	Lbrace token.Pos // "{" 的位置
	Elts []Expr // 复合元素列表；或 nil
	Rbrace token.Pos // "}" 的位置
	Incomplete bool // 如果 Elts 列表中缺少（源）表达式，则为 true
}
```

CompositeLit 节点表示一个复合字面。

#### func (x *CompositeLit) End() token.Pos

#### func (x *CompositeLit) Pos() token.Pos

### type Decl

```go
type Decl interface {
	Node
	// 包含已过滤或未导出的方法
}
```

所有声明节点都实现 Decl 接口。

### type DeclStmt

```go
type DeclStmt struct {
	Decl Decl // *GenDecl，带 CONST、TYPE 或 VAR 标记
}
```

DeclStmt 节点表示语句列表中的一个声明。

#### func (s *DeclStmt) End() token.Pos

#### func (s *DeclStmt) Pos() token.Pos

### type DeferStmt

```go
type DeferStmt struct {
	Defer token.Pos // "推迟 "关键字的位置
	Call *CallExpr
}
```

DeferStmt 节点表示一条延迟语句。

#### func (x *Ellipsis) End() token.Pos

#### func (x *Ellipsis) Pos() token.Pos

### type Ellipsis

```go
type Ellipsis struct {
	Ellipsis token.Pos // "... "的位置
	Elt Expr // 省略号元素类型（仅限参数列表）；或 nil
}
```

Ellipsis 节点代表参数列表中的"... "类型或数组类型中的"... "长度。

#### func (x *Ellipsis) End() token.Pos

#### func (x *Ellipsis) Pos() token.Pos

### type EmptyStmt

```go
type EmptyStmt struct {
	Semicolon token.Pos // 后面"; "的位置
	Implicit bool // 如果设置，源代码中省略了";"
}
```

EmptyStmt 节点表示空语句。空语句的 "位置 "是紧跟其后的（显式或隐式）分号的位置。

### type Expr

```go
type Expr interface {
	Node
	// 包含已过滤或未导出的方法
}
```

所有表达式节点都实现 Expr 接口。

### type ExprStmt

```go
type ExprStmt struct {
	X Expr // 表达式
}
```

ExprStmt 节点表示语句列表中的一个（独立）表达式。

#### func (s *ExprStmt) End() token.Pos

#### func (s *ExprStmt) Pos() token.Pos

### type Field

```go
type Field struct {
	Doc *CommentGroup // 相关文档；或 nil
	Names []*Ident // 字段/方法/（类型）参数名称；或 nil
	Type Expr // 字段/方法/参数类型；或 nil
	Tag *BasicLit // 字段标签；或 nil
	Comment *CommentGroup // 行注释；或 nil
}
```

字段表示结构类型中的字段声明列表、接口类型中的方法列表或签名中的参数/结果声明。对于未命名参数（仅包含类型的参数列表）和嵌入式结构字段，Field.Names 为空。在后一种情况下，字段名称就是类型名称。

#### func (f *Field) End() token.Pos

#### func (f *Field) Pos() token.Pos

### type FieldFilter

```go
type FieldFilter func(name string, value reflect.Value) bool
```

可向 Fprint 提供 FieldFilter（字段过滤器），以控制输出。

### type FieldList

```go
type FieldList struct {
	Opening token.Pos // 开头括号/小括号/方括号的位置（如果有的话
	List []*Field // 字段列表；或 nil
	Closing token.Pos // 结束括号/小括号/方括号的位置（如果有的话
}
```

FieldList 表示一个字段列表，由小括号、大括号或方括号括起来。

#### func (f *FieldList) End() token.Pos

#### func (f *FieldList) NumFields() int

NumFields 返回 FieldList 所代表的参数或结构字段的数量。

#### func (f *FieldList) Pos() token.Pos

### type File

```go
type File struct {
	Doc *CommentGroup // 相关文件；或 nil
	Package token.Pos // "包 "关键字的位置
	Name *Ident // 包名
	Decls []Decl // 顶层声明；或 nil

	FileStart, FileEnd token.Pos // 整个文件的起点和终点
	Scope *Scope // 包范围（仅此文件）
	Imports []*ImportSpec // 本文件中的导入
	Unresolved []*Ident // 本文件中未解决的标识符
	Comments []*CommentGroup // 源文件中所有注释的列表
	GoVersion string // //go:build 或 // +build 指令要求的最小 Go 版本
}
```

文件节点表示 Go 源文件。

Comments 列表包含源文件中按出现顺序排列的所有注释，包括通过 Doc 和 Comment 字段从其他节点指向的注释。

要正确打印包含注释的源代码（使用 go/format 和 go/printer 软件包），必须特别注意在修改文件的语法树时更新注释：在打印时，注释根据标记的位置穿插在标记之间。如果语法树节点被移除或移动，其附近的相关注释也必须相应移除（从 File.Comments 列表中移除）或移动（更新其位置）。可以使用注释图来简化其中的一些操作。

注释是否以及如何与节点关联取决于操作程序对语法树的解释：除了与节点直接关联的 Doc 和 Comment 注释外，其余的注释都是 "自由浮动 "的（另见问题 #18593、#20744）。

#### func MergePackageFiles(pkg *Package, mode MergeMode) *File

MergePackageFiles 通过合并软件包所属文件的 AST 来创建文件 AST。模式标志控制合并行为。

#### func (f *File) End() token.Pos

End 返回文件中最后一条声明的结束。(使用 FileEnd 返回整个文件的结束）。

#### func (f *File) Pos() token.Pos

Pos 返回软件包声明的位置。(使用 FileStart 获取整个文件的起始位置）。

### type Filter

```go
type Filter func(string) bool
```

### type ForStmt

```go
type ForStmt struct {
	For token.Pos // "for "关键字的位置
	Init Stmt // 初始化语句；或 nil
	Cond Expr // 条件；或 nil
	Post Stmt // 后置迭代语句；或 nil
	Body *BlockStmt
}
```

ForStmt 表示 for 语句。

#### func (s *ForStmt) End() token.Pos

#### func (s *ForStmt) Pos() token.Pos

### type FuncDecl

```go
type FuncDecl struct {
	Doc *CommentGroup // 相关文档；或 nil
	Recv *FieldList // 接收者（方法）；或 nil（函数）
	Name *Ident // 函数/方法名称
	Type *FuncType // 函数签名：参数类型和值、结果以及 "func "关键字的位置
	Body *BlockStmt // 函数体；或 nil（外部（非 Go）函数
}
```

FuncDecl 节点表示函数声明。

#### func (d *FuncDecl) End() token.Pos

#### func (d *FuncDecl) Pos() token.Pos

### type FuncLit

```go
type FuncLit struct {
	Type *FuncType // 函数类型
	Body *BlockStmt // 函数正文
}
```

FuncLit 节点表示函数文字。

#### func (x *FuncLit) End() token.Pos

#### func (x *FuncLit) Pos() token.Pos

### type FuncType

```go
type FuncType struct {
	Func token.Pos // "func "关键字的位置（如果没有 "func"，则 token.NoPos）
	TypeParams *FieldList // 类型参数；或 nil
	Params *FieldList // （输入）参数；非零
	Results *FieldList // （传出）结果；或 nil
}
```

FuncType 节点表示函数类型。

#### func (x *FuncType) End() token.Pos

#### func (x *FuncType) Pos() token.Pos

### type GenDecl

```go
type GenDecl struct {
	Doc *CommentGroup // 相关文档；或 nil
	TokPos token.Pos // Tok 的位置
	Tok token.Token // IMPORT、CONST、TYPE 或 VAR
	Lparen token.Pos // '(' 的位置（如果有的话
	Specs []Spec
	Rparen token.Pos // ')' 的位置，如果有的话
}
```

GenDecl 节点（通用声明节点）表示导入、常量、类型或变量声明。一个有效的 Lparen 位置（Lparen.IsValid()）表示一个带括号的声明。

Tok 值与 Specs 元素类型之间的关系：

```go
token.IMPORT  *ImportSpec
token.CONST   *ValueSpec
token.TYPE    *TypeSpec
token.VAR     *ValueSpec
```

#### func (d *GenDecl) End() token.Pos

#### func (d *GenDecl) Pos() token.Pos

### type GoStmt

```go
type GoStmt struct {
	Go token.Pos // "go "关键字的位置
	Call *CallExpr
}
```

GoStmt 节点表示 go 语句。

#### func (s *GoStmt) End() token.Pos

#### func (s *GoStmt) Pos() token.Pos

### type Ident

```go
type Ident struct {
	NamePos token.Pos // 标识符位置
	Name string // 标识符名称
	Obj *Object // 表示对象；或 nil
}
```

Ident 节点表示一个标识符。

#### func NewIdent(name string) *Ident

NewIdent 创建一个没有位置的新 Ident。这对 Go 解析器以外的代码生成的 AST 很有用。

#### func (x *Ident) End() token.Pos

#### func (id *Ident) IsExported() bool

IsExported 报告 id 是否以大写字母开头。

#### func (x *Ident) Pos() token.Pos

#### func (id *Ident) String() string

### type IfStmt

```go
type IfStmt struct {
	If token.Pos // "if "关键字的位置
	Init Stmt // 初始化语句；或 nil
	Cond Expr // 条件
	Body *BlockStmt
	Else Stmt // 其他分支；或 nil
}
```

IfStmt 节点表示 if 语句。

#### func (s *IfStmt) End() token.Pos

#### func (s *IfStmt) Pos() token.Pos

### type ImportSpec

```go
type ImportSpec struct {
	Doc *CommentGroup // 相关文档；或 nil
	Name *Ident // 本地软件包名称（包括"."）；或 nil
	Path *BasicLit // 导入路径
	Comment *CommentGroup // 行注释；或 nil
	EndPos token.Pos // 规范结束（如果不为零，则覆盖 Path.Pos）
}
```

一个 ImportSpec 节点代表一个单一的软件包导入。

#### func (s *ImportSpec) End() token.Pos

#### func (s *ImportSpec) Pos() token.Pos

### type Importer

```go
type Importer func(imports map[string]*Object, path string) (pkg *Object, err error)
```

导入器将导入路径解析为软件包对象。导入映射记录了已导入的软件包，并以软件包 ID（规范导入路径）为索引。导入器必须确定规范导入路径，并检查该路径是否已出现在导入映射中。如果是，导入器可以返回映射条目。否则，导入器应将给定路径的软件包数据加载到一个新的 *Object (pkg)，在导入映射中记录 pkg，然后返回 pkg。

### type IncDecStmt

```go
type IncDecStmt struct {
	X Expr
	TokPos token.Pos // Tok 的位置
	Tok token.Token // INC 或 DEC
}
```

IncDecStmt 节点表示递增或递减语句。

#### func (s *IncDecStmt) End() token.Pos

#### func (s *IncDecStmt) Pos() token.Pos

### type IndexExpr

```go
type IndexExpr struct {
	X Expr // 表达式
	Lbrack token.Pos // "["的位置
	Index Expr // 索引表达式
	Rbrack token.Pos // "]"的位置
}
```

IndexExpr 节点表示一个表达式，后面跟一个索引。

#### func (x *IndexExpr) End() token.Pos

#### func (x *IndexExpr) Pos() token.Pos

### type IndexListExpr 添加于1.18

```go
type IndexListExpr struct {
	X Expr // 表达式
	Lbrack token.Pos // "["的位置
	Indices []Expr // 索引表达式
	Rbrack token.Pos // "]"的位置
}
```

IndexListExpr 节点表示一个表达式，后面跟多个索引。

#### func (x *IndexListExpr) End() token.Pos 添加于1.18

#### func (x *IndexListExpr) Pos() token.Pos 添加于1.18

### type InterfaceType

```go
type InterfaceType struct {
	Interface token.Pos // "接口 "关键字的位置
	Methods *FieldList // 嵌入接口、方法或类型的列表
	Incomplete bool // 如果方法列表中缺少（源）方法或类型，则为 true
}
```

InterfaceType 节点表示接口类型。

#### func (x *InterfaceType) End() token.Pos

#### func (x *InterfaceType) Pos() token.Pos

### type KeyValueExpr

```go
type KeyValueExpr struct {
	键 Expr
	Colon token.Pos // ": "的位置
	值 Expr
}
```

KeyValueExpr 节点表示复合文字中的（键 : 值）对。

#### func (x *KeyValueExpr) End() token.Pos

#### func (x *KeyValueExpr) Pos() token.Pos

### type LabeledStmt

```go
type LabeledStmt struct {
	Label *Ident
	Colon token.Pos // ": "的位置
	Stmt Stmt
}
```

一个 LabeledStmt 节点表示一个带标签的语句。

#### func (s *LabeledStmt) End() token.Pos

#### func (s *LabeledStmt) Pos() token.Pos

### type MapType

```go
type MapType struct {
	Map token.Pos // "map "关键字的位置
	Key Expr
	Value Expr
}
```

MapType 节点表示地图类型。

#### func (x *MapType) End() token.Pos

#### func (x *MapType) Pos() token.Pos

### type MergeMode

```go
类型 MergeMode uint
```

MergeMode 标志控制 MergePackageFiles 的行为。

```go
const (
	// 如果设置，则排除重复的函数声明。
	FilterFuncDuplicates MergeMode = 1 << iota
	// 如果设置了，则与特定 AST 节点（作为 Doc 或 Comment）无关的注释将被排除。
	// AST 节点（作为文档或注释）无关的注释将被排除。
	FilterUnassociatedComments
	// 如果设置，将排除重复的导入声明。
	FilterImportDuplicates
)
```

### type Node

```go
type Node interface {
	Pos() token.Pos // 属于节点的第一个字符的位置
	End() token.Pos // 紧随节点之后的第一个字符的位置
}
```

所有节点类型都实现了 Node 接口。

### type ObjKind

```go
type ObjKind int
```

ObjKind 描述了一个对象代表什么。

```go
const (
	Bad ObjKind = iota // 用于错误处理
	Pkg // 软件包
	Con // 常量
	Typ // 类型
	Var // 变量
	Fun // 函数或方法
	Lbl // 标签
)
```

可能的对象类型列表。

#### func (kind ObjKind) String() string

### type Object

```go
type Object struct {
	Kind ObjKind
	Name // 声明名称
	Decl any // 相应的 Field, XxxSpec, FuncDecl, LabeledStmt, AssignStmt, Scope; or nil
	Data // 特定于对象的数据；或 nil
	Type any // 类型信息的占位符；可能为 nil
}
```

对象描述一个已命名的语言实体，如包、常量、类型、变量、函数（包括方法）或标签。

数据字段包含特定于对象的数据：

#### func NewObj(kind ObjKind, name string) *Object

NewObj 按给定的种类和名称创建一个新对象。

#### func (obj *Object) Pos() token.Pos

Pos 计算对象名称声明的源位置。如果无法计算，结果可能是无效位置（obj.Decl 可能为 nil 或不正确）。

### type Package

```go
type Package struct {
	Name string // 软件包名称
	Scope *Scope // 软件包在所有文件中的范围
	Imports map[string]*Object // 包 ID -> 包对象的映射
	Files map[string]*File // 按文件名查找源文件
}
```

包节点表示一组源文件，它们共同构建一个 Go 包。

#### func NewPackage(fset *token.FileSet, files map[string]*File, importer 
Importer, ...) (*Package, error)

NewPackage 从一组文件节点创建一个新的包节点。它会解析各文件中未解析的标识符，并相应地更新每个文件的未解析列表。如果提供了非零的导入器和宇宙范围，它们将被用于解决未在任何包文件中声明的标识符。剩余的未解决标识符将被报告为未声明。如果文件属于不同的软件包，则会选择一个软件包名称，报告并忽略不同软件包名称的文件。结果是一个软件包节点和一个 scanner.ErrorList（如果存在错误）。

#### func (p *Package) End() token.Pos

#### func (p *Package) Pos() token.Pos

### type ParenExpr

```go
type ParenExpr struct {
	Lparen token.Pos // "("的位置
	X Expr // 括号表达式
	Rparen token.Pos // ") "的位置
}
```

ParenExpr 节点表示括号表达式。

#### func (x *ParenExpr) End() token.Pos

#### func (x *ParenExpr) Pos() token.Pos

### type RangeStmt

```go
type RangeStmt struct {
	For token.Pos // "for "关键字的位置
	Key, Value Expr // Key, Value 可以为 nil
	TokPos token.Pos // Tok 的位置；如果 Key == nil 则无效
	Tok token.Token // 如果 Key == nil、ASSIGN、DEFINE，则无效
	Range token.Pos // "范围 "关键字的位置
	X Expr // 范围的值
	Body *BlockStmt
}
```

RangeStmt 表示带有范围子句的 for 语句。

#### func (s *RangeStmt) End() token.Pos

#### func (s *RangeStmt) Pos() token.Pos

### type ReturnStmt

```go
type ReturnStmt struct {
	Return token.Pos // "return "关键字的位置
	Results []Expr // 结果表达式；或 nil
}
```

ReturnStmt 节点表示返回语句。

#### func (s *ReturnStmt) End() token.Pos

#### func (s *ReturnStmt) Pos() token.Pos

### type Scope

```go
type Scope struct {
	Outer *Scope
	Objects map[string]*Object
}
```

一个作用域维护着在该作用域中声明的命名语言实体集，以及与紧邻的（外层）作用域的链接。

#### func NewScope(outer *Scope) *Scope

NewScope 创建嵌套在外作用域中的新作用域。

#### func (s *Scope) Insert(obj *Object) (alt *Object)

如果作用域已包含一个同名的对象 alt，Insert 将保持作用域不变并返回 alt。否则，它会插入 obj 并返回 nil。

#### func (s *Scope) Lookup(name string) *Object

如果在作用域 s 中找到具有给定名称的对象，Lookup 将返回该对象，否则返回 nil。外部作用域将被忽略。

#### func (s *Scope) String() string

调试支持

### type SelectStmt

```go
type SelectStmt struct {
	Select token.Pos // "select "关键字的位置
	Body *BlockStmt // 仅 CommClauses
}
```

SelectStmt 节点表示选择语句。

#### func (s *SelectStmt) End() token.Pos

#### func (s *SelectStmt) Pos() token.Pos

### type SelectorExpr

```go
type SelectorExpr struct {
	X Expr // 表达式
	Sel *Ident // 字段选择器
}
```

SelectorExpr 节点表示一个表达式，后面跟一个选择器。

#### func (x *SelectorExpr) End() token.Pos

#### func (x *SelectorExpr) Pos() token.Pos

### type SendStmt

```go
type SendStmt struct {
	Chan Expr
	Arrow token.Pos // "<-" 的位置
	Value Expr
}
```

SendStmt 节点表示一条发送语句。

#### func (s *SendStmt) End() token.Pos

#### func (s *SendStmt) Pos() token.Pos

### type SliceExpr

```go
type SliceExpr struct {
	X Expr // 表达式
	Lbrack token.Pos // "["的位置
	Low Expr // 切分范围的起始点；或 nil
	High Expr // 切片范围的终点；或 nil
	Max Expr // 片的最大容量；或 nil
	Slice3 bool // 如果是 3 索引分片则为 true（存在 2 个冒号）
	Rbrack token.Pos // "]"的位置
}
```

SliceExpr 节点表示一个表达式，后面跟着切片索引。

#### func (x *SliceExpr) End() token.Pos

#### func (x *SliceExpr) Pos() token.Pos

### type Spec

```go
type Spec interface {
	Node
	// 包含已过滤或未导出的方法
}
```

Spec 类型代表 *ImportSpec、*ValueSpec 和 *TypeSpec 中的任意一种。

### type StarExpr

```go
type StarExpr struct {
	Star token.Pos // "*" 的位置
	X Expr // 操作数
}
```

StarExpr 节点表示 "*" Expression 形式的表达式。从语义上讲，它可以是一个一元 "*"表达式，也可以是一个指针类型。

#### func (x *StarExpr) End() token.Pos

#### func (x *StarExpr) Pos() token.Pos

### type Stmt

```go
type Stmt interface {
	Node
	// 包含已过滤或未导出的方法
}
```

所有语句节点都实现了 Stmt 接口。

### type StructType

```go
type StructType struct {
	Struct token.Pos // "struct "关键字的位置
	Fields *FieldList // 字段声明列表
	Incomplete bool // 如果字段列表中缺少（源）字段，则为 true
}
```

StructType 节点表示结构类型。

#### func (x *StructType) End() token.Pos

#### func (x *StructType) Pos() token.Pos

### type SwitchStmt

```go
type SwitchStmt struct {
	Switch token.Pos // "switch "关键字的位置
	Init Stmt // 初始化语句；或 nil
	Tag Expr // 标签表达式；或 nil
	Body *BlockStmt // 仅 CaseClauses
}
```

SwitchStmt 节点表示表达式开关语句。

#### func (s *SwitchStmt) End() token.Pos

#### func (s *SwitchStmt) Pos() token.Pos

### type TypeAssertExpr

```go
type TypeAssertExpr struct {
	X Expr // 表达式
	Lparen token.Pos // "（"的位置
	Type Expr // 断言类型；nil 表示类型 switch X.(type)
	Rparen token.Pos // ") "的位置
}
```

TypeAssertExpr 节点表示一个表达式，后面跟一个类型断言。

#### func (x *TypeAssertExpr) End() token.Pos

#### func (x *TypeAssertExpr) Pos() token.Pos

### type TypeSpec

```go
type TypeSpec struct {
	Doc *CommentGroup // 相关文档；或 nil
	Name *Ident // 类型名
	TypeParams *FieldList // 类型参数；或 nil
	Assign token.Pos // '='的位置，如果有的话
	Type Expr // *Ident, *ParenExpr, *SelectorExpr, *StarExpr, 或任何 *XxxTypes
	Comment *CommentGroup // 行注释；或 nil
}
```

TypeSpec 节点表示类型声明（TypeSpec 生产）。

#### func (s *TypeSpec) End() token.Pos

#### func (s *TypeSpec) Pos() token.Pos

### type TypeSwitchStmt

```go
type TypeSwitchStmt struct {
	Switch token.Pos // "switch "关键字的位置
	Init Stmt // 初始化语句；或 nil
	Assign Stmt // x := y.(type) 或 y.(type)
	Body *BlockStmt // 仅 CaseClauses
}
```

TypeSwitchStmt 节点表示类型切换语句。

#### func (s *TypeSwitchStmt) End() token.Pos

#### func (s *TypeSwitchStmt) Pos() token.Pos

### type UnaryExpr

```go
type UnaryExpr struct {
	OpPos token.Pos // 运算符的位置
	Op token.Token // 运算符
	X Expr // 操作数
}
```

UnaryExpr 节点表示一元表达式。一元 "*"表达式通过 StarExpr 节点表示。

#### func (x *UnaryExpr) End() token.Pos

#### func (x *UnaryExpr) Pos() token.Pos

### type ValueSpec

```go
type ValueSpec struct {
	Doc *CommentGroup // 相关文档；或 nil
	Names []*Ident // 值名称 (len(Names) > 0)
	Type Expr // 值类型；或 nil
	Values []Expr // 初始值；或 nil
	Comment *CommentGroup // 行注释；或 nil
}
```

ValueSpec 节点代表常量或变量声明（ConstSpec 或 VarSpec 生产）。

#### func (s *ValueSpec) End() token.Pos

#### func (s *ValueSpec) Pos() token.Pos

### type Visitor

```go
type Visitor interface {
	Visit（node Node） (w Visitor)
}
```

Walk 遇到的每个节点都会调用 Visitor 的 Visit 方法。如果结果访问者 w 不是 nil，Walk 会用访问者 w 访问节点的每个子节点，然后调用 w.Visit(nil)。