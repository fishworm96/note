## package printer

包打印机实现了 AST 节点的打印。

## Index

### func Fprint(output io.Writer, fset *token.FileSet, node any) error

Fprint 将 AST 节点 "漂亮地打印 "出来。它使用默认设置调用 Config.Fprint。请注意，gofmt 使用制表符缩进，但使用空格对齐；使用 format.Node （包 go/format）可获得与 gofmt 一致的输出。

### type CommentedNode

```go
type CommentedNode struct {
	Node any // *ast.File, or ast.Expr, ast.Decl, ast.Spec, or ast.Stmt
	Comments []*ast.CommentGroup
}
```

一个注释节点（CommentedNode）捆绑了一个 AST 节点和相应的注释。它可以作为参数提供给任何一个 Fprint 函数。

### type Config

```go
type Config struct {
	Mode Mode // 默认值: 0
	Tabwidth int // 默认值: 8
	Indent int // 默认值：0（所有代码至少缩进这么多）
}
```

Config 节点控制 Fprint 的输出。

#### func (cfg *Config) Fprint(output io.Writer, fset *token.FileSet, node any) error

Fprint 针对给定的配置 cfg，"漂亮地打印 "一个 AST 节点输出。位置信息是相对于文件集 fset 来解释的。节点类型必须是 *ast.File、*CommentedNode、[]ast.Decl、[]ast.Stmt，或与 ast.Expr、ast.Decl、ast.Spec 或 ast.Stmt 赋值兼容。

### type Mode

```go
type Mode uint
```

模式值是一组标志（或 0）。它们控制打印。

```go
const (
	RawFormat Mode = 1 << iota // 不使用制表符；如果设置，则忽略 UseSpaces
	TabIndent // 使用制表符缩进，与 UseSpaces 无关
	UseSpaces // 使用空格而不是制表符对齐
	SourcePos // 发送 // 行指令以保留原始源位置
)
```