## package format

包格式实现了 Go 代码的标准格式化。

请注意，Go 源代码的格式会随着时间的推移而改变，因此依赖于一致格式的工具应执行特定版本的 gofmt 二进制程序，而不是使用此软件包。这样，格式化将保持稳定，而且每次 gofmt 更改时，工具都无需重新编译。

例如，直接使用此软件包的预提交检查会因每个开发者使用的 Go 版本不同而表现不同，从而导致检查固有的脆弱性。

## Index

### func Node(dst io.Writer, fset *token.FileSet, node any) error

节点以规范的 gofmt 风格格式化节点，并将结果写入 dst。

节点类型必须是 *ast.File、*printer.CommentedNode、[]ast.Decl、[]ast.Stmt，或与赋值兼容的 ast.Expr、ast.Decl、ast.Spec 或 ast.Stmt。Node 不修改节点。对于表示部分源文件的节点（例如，如果节点不是 *ast.File 或 *printer.CommentedNode 未封装 *ast.File），则不会对导入进行排序。

函数可能会提前返回（在写入整个结果之前）并返回格式错误，例如由于 AST 不正确。

### func Source(src []byte) ([]byte, error)

源代码会以规范的 gofmt 风格格式化 src，并返回结果或（I/O 或语法）错误。

如果 src 是部分源文件，则会将 src 的前导和尾部空格应用于结果（使其具有与 src 相同的前导和尾部空格），并且结果的缩进量与包含代码的 src 第一行相同。对于部分源文件，不会对导入进行排序。