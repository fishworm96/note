## package parser

包 parser 实现了 Go 源文件的解析器。输入可以多种形式提供（参见各种 Parse* 函数）；输出是表示 Go 源文件的抽象语法树（AST）。解析器通过 Parse* 函数之一调用。

解析器接受的语言比 Go 规范允许的语法更大，这是为了简化，也是为了在出现语法错误时提高鲁棒性。例如，在方法声明中，接收器被视为普通的参数列表，因此可能包含多个条目，而规范只允许一个条目。因此，AST（ast.FuncDecl.Recv）字段中的相应字段不限于一个条目。

## Index

### func ParseDir(fset token.FileSet, path string, filter func(fs.FileInfo) bool, mode Mode) (pkgs map[string]*ast.Package, first error)

包 parser 实现了 Go 源文件的解析器。输入可以多种形式提供（参见各种 Parse* 函数）；输出是表示 Go 源文件的抽象语法树（AST）。解析器通过 Parse* 函数之一调用。

解析器接受的语言比 Go 规范允许的语法更大，这是为了简化，也是为了在出现语法错误时提高鲁棒性。例如，在方法声明中，接收器被视为普通的参数列表，因此可能包含多个条目，而规范只允许一个条目。因此，AST（ast.FuncDecl.Recv）字段中的相应字段不限于一个条目。

### func ParseExpr(x string) (ast.Expr, error)

ParseExpr 是一个方便函数，用于获取表达式 x 的 AST。错误信息中使用的文件名是空字符串。

如果发现语法错误，结果将是部分 AST（ast.Bad* 节点代表错误源代码片段）。多个错误会通过一个按源代码位置排序的 scanner.ErrorList 返回。

### func ParseExprFrom(fset *token.FileSet, filename string, src any, mode Mode) (expr ast.Expr, err error) 添加于1.5

ParseExprFrom 是一个用于解析表达式的方便函数。参数的含义与 ParseFile 相同，但来源必须是有效的 Go（类型或值）表达式。具体来说，fset 不能为 nil。

如果无法读取源代码，返回的 AST 将为 nil，错误信息将指明具体的失败原因。如果源代码已读取，但发现了语法错误，则返回部分 AST（ast.Bad* 节点代表错误的源代码片段）。多个错误会通过一个按源代码位置排序的 scanner.ErrorList 返回。

### func ParseFile(fset *token.FileSet, filename string, src any, mode Mode) (f *ast.File, err error)

ParseFile 会解析单个 Go 源文件的源代码，并返回相应的 ast.File 节点。源代码可以通过源文件的文件名或 src 参数提供。

如果 src != nil，ParseFile 将从 src 解析源代码，文件名仅在记录位置信息时使用。src 参数的参数类型必须是字符串、[]字节或 io.Reader。如果 src == nil，ParseFile 将解析 filename 指定的文件。

模式参数控制解析源文本的数量和其他可选的解析器功能。如果设置了 SkipObjectResolution 模式位，解析的对象解析阶段将被跳过，导致 File.Scope、File.Unresolved 和所有 Ident.Obj 字段为零。

位置信息记录在文件集 fset 中，该文件集不得为空。

如果无法读取源代码，返回的 AST 将为 nil，错误信息将指明具体的失败原因。如果源代码已读取，但发现了语法错误，则返回部分 AST（ast.Bad* 节点代表错误的源代码片段）。多个错误会通过一个按源代码位置排序的 scanner.ErrorList 返回。

### type Mode

```go
type Mode uint
```

模式值是一组标志（或 0）。它们控制着源代码解析量和其他可选的解析器功能。

```go
const (
	PackageClauseOnly Mode = 1 << iota // 包条款后停止解析
	ImportsOnly // 在导入声明后停止解析
	ParseComments // 解析注释并将其添加到 AST 中
	Trace // 打印已解析产品的跟踪
	DeclarationErrors // 报告声明错误
	SpuriousErrors // 与 AllErrors 相同，用于向后兼容
	SkipObjectResolution // 不将标识符解析为对象 - 参见 ParseFile
	AllErrors = SpuriousErrors // 报告所有错误（而不仅仅是不同行的前 10 个错误）
)
```