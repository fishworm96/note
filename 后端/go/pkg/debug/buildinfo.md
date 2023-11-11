## package buildinfo

软件包 buildinfo 提供了访问嵌入在 Go 二进制文件中的信息的途径，包括 Go 工具链版本和使用的模块集（适用于以模块模式构建的二进制文件）。其中包括 Go 工具链版本和使用的模块集（适用于以模块模式构建的二进制文件）。

当前运行的二进制文件可在 runtime/debug.ReadBuildInfo 中获取编译信息。

## Index

### type BuildInfo

```go
type BuildInfo = debug.BuildInfo
```

构建信息的类型别名。我们不能将类型移到这里，因为运行时/调试时需要导入此软件包，这将使其依赖性大大增加。

#### func Read(r op/ReaderAt) (*BuildInfo, error)

读取返回通过给定的 ReaderAt 访问的 Go 二进制文件中嵌入的构建信息。大多数信息只适用于支持模块的二进制文件。

#### func ReadFile(name string) (info *BuildInfo, err error)

ReadFile 返回嵌入在给定路径下 Go 二进制文件中的构建信息。大多数信息只适用于支持模块的二进制文件。