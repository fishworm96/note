## package importer

软件包导入器提供对导出数据导入器的访问。

## Index

### func Default() types.importer

默认返回构建运行二进制文件的编译器的导入器。如果可用，结果将实现 types.ImporterFrom。

### func For(compiler string, lookup Lookup) types.Importer 废除

### func ForCompiler(fset *token.FileSet, compiler string, lookup Lookup) types.Importer 添加于1.12

ForCompiler 返回一个导入器，用于从已安装的编译器 "gc "和 "gccgo "包中导入；如果编译器参数为 "source"，则直接从源代码中导入。在后一种情况下，如果导出的 API 并非完全由纯 Go 源代码定义，则导入可能会失败（如果软件包 API 依赖于 cgo 定义的实体，则类型检查程序无法访问这些实体）。

每次导入程序需要解析导入路径时，都会调用查找函数。在此模式下，导入器只能调用规范导入路径（而不是相对或绝对导入路径）；假定导入器的客户端正在将导入路径转换为规范导入路径。

必须提供一个查找函数，以实现正确的模块感知操作。已废弃：如果 lookup 为 nil，为了向后兼容，导入器将尝试解析 $GOPATH 工作区中的导入。

### type Lookup

```go
type Lookup func(path string) (io.ReadCloser, error)
```

Lookup 函数返回一个阅读器，用于访问给定导入路径的软件包数据；如果没有找到匹配的软件包，则返回错误信息。