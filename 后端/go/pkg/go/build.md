## build

软件包构建会收集有关 Go 软件包的信息。

### Go Path

Go 路径是包含 Go 源代码的目录树列表。如果在标准 Go 目录树中找不到导入代码，就需要参考它来解决。默认路径是 GOPATH 环境变量的值，它被解释为适合操作系统的路径列表（在 Unix 上，该变量是一个用冒号分隔的字符串；在 Windows 上，是一个用分号分隔的字符串；在 Plan 9 上，是一个列表）。

Go 路径中列出的每个目录必须具有规定的结构：

src/ 目录存放源代码。src "下面的路径决定了导入路径或可执行文件的名称。

pkg/ 目录存放已安装的软件包对象。与 Go 树一样，每个目标操作系统和体系结构都有自己的 pkg 子目录（pkg/GOOS_GOARCH）。

如果 DIR 是 Go 路径中列出的一个目录，那么源代码位于 DIR/src/foo/bar 的软件包可以作为 "foo/bar "导入，其编译后的形式将被安装到 "DIR/pkg/GOOS_GOARCH/foo/bar.a"（或对于 gccgo，"DIR/pkg/gccgo/foo/libbar.a"）。

bin/ 目录存放已编译的命令。每条命令都以其源代码目录命名，但只使用最后一个元素，而不是整个路径。也就是说，源代码位于 DIR/src/foo/quux 的命令会被安装到 DIR/bin/quux，而不是 DIR/bin/foo/quux。去掉 foo/ 后，你就可以在 PATH 中添加 DIR/bin 来获取已安装的命令。

下面是一个目录布局示例：

```go
GOPATH=/home/user/gocode

/home/user/gocode/
    src/
        foo/
            bar/               (go code in package bar)
                x.go
            quux/              (go code in package main)
                y.go
    bin/
        quux                   (installed command)
    pkg/
        linux_amd64/
            foo/
                bar.a          (installed package object)
```

### Build Constraints

编译约束也称为编译标记，是将文件包含在软件包中的条件。编译约束由一行注释给出，注释以

```go
//go:build
```

编译限制也可能是文件名的一部分（例如，只有当目标操作系统是 windows 时，source_windows.go 才会被包含在内）。

详情请参阅 "go help buildconstraint"（https://golang.org/cmd/go/#hdr-Build_constraints）。

### Binary-Only Packages

在 Go 1.12 及更早版本中，可以分发二进制形式的软件包，而不包括编译软件包时使用的源代码。发布软件包时，源文件不会被编译约束排除在外，并且包含"//go:binary-only-package "注释。与构建约束一样，该注释出现在文件的顶部，前面只有空行和其他行注释，注释后面还有一个空行，以将其与软件包文档分开。与构建约束不同，该注释只在非测试 Go 源文件中被识别。

因此，纯二进制软件包的最小源代码为:

```go
//go:binary-only-package

package mypkg
```

源代码可能包括额外的 Go 代码。这些代码永远不会被编译，但会被 godoc 等工具处理，并可能成为有用的最终用户文档。

"go build "和其他命令不再支持纯二进制包。Import 和 ImportDir 仍会在包含这些注释的软件包中设置 BinaryOnly 标志，以便在工具和错误信息中使用。

## Index

### Variables

```go
var ToolDir = getToolDir()
```

ToolDir 是包含构建工具的目录。

### func ArchChar(goarch string) (string, error)

ArchChar 返回"? "和一个错误。在 Go 的早期版本中，返回的字符串用于推导编译器和链接器工具名称、默认对象文件后缀和默认链接器输出名称。从 Go 1.5 开始，这些字符串不再因架构而异；它们分别是 compile、link、.o 和 a.out。

### func IsLocalImport(path string) bool

IsLocalImport 报告导入路径是否为本地导入路径，如"."、".."、"./foo "或"../foo"。

### type Context

```go
type Context struct {
	GOARCH string // 目标架构
	GOOS string // 目标操作系统
	GOROOT string // Go 根目录
	GOPATH string // Go 路径

	// Dir 是调用者的工作目录，如果使用运行进程的当前目录，则为空字符串。在模块模式下，它用于定位主模块。
	//
	// 如果 Dir 为非空，则传给 Import 和 ImportDir 的目录必须是绝对目录。
	Dir string

	CgoEnabled  bool   // 是否包含 cgo 文件
	UseAllFiles bool   // 无论 go:build 行和文件名如何，都使用文件
	Compiler    string // 编译器在计算目标路径时

  // build、tool 和 release 标记指定了在处理 go:build 行时应视为已满足的构建约束。创建新上下文的客户端可以自定义 BuildTags（默认为空），但自定义 ToolTags 或 ReleaseTags 通常会出错。ToolTags 默认为适合当前 Go 工具链配置的构建标记。ReleaseTags 默认为当前版本兼容的 Go 版本列表。默认构建上下文不设置 BuildTags。除了 BuildTags、ToolTags 和 ReleaseTags 之外，联编约束还将 GOARCH 和 GOOS 的值视为已满足的标记。ReleaseTags 中的最后一个元素被假定为当前版本。
	BuildTags   []string
	ToolTags    []string
	ReleaseTags []string

  // install 后缀指定安装目录名称中使用的后缀。默认情况下它是空的，但需要将输出分开的定制编译可以设置 InstallSuffix。例如，在使用 race 检测器时，go 命令会使用 InstallSuffix = "race"，这样在 Linux/386 系统上，软件包就会被写入名为 "linux_386_race "的目录，而不是通常的 "linux_386"。
	InstallSuffix string

  // JoinPath 会将路径片段序列合并为一个路径。如果 JoinPath 为空，Import 将使用 filepath.Join。
	JoinPath func(elem ...string) string

  // SplitPathList 会将路径列表分割成一个个单独的路径片段。如果 SplitPathList 为空，Import 将使用 filepath.SplitList。
	SplitPathList func(list string) []string

	// IsAbsPath 报告路径是否为绝对路径。如果 IsAbsPath 为空，Import 将使用 filepath.IsAbs.Path。
	IsAbsPath func(path string) bool

	// IsDir 报告路径是否命名了一个目录。如果 IsDir 为空，Import 会调用 os.Stat 并使用结果的 IsDir 方法。
	IsDir func(path string) bool

	// HasSubdir 报告 dir 在词法上是否是 root 的子目录，也许是下面的多级子目录。它不会检查 dir 是否存在。如果存在，HasSubdir 会将 rel 设置为一个斜线分隔的路径，该路径可以与 root 连接，从而产生一个与 dir 相当的路径。如果 HasSubdir 为空，Import 将使用基于 filepath.EvalSymlinks 的实现。
	HasSubdir func(root, dir string) (rel string, ok bool)

	// ReadDir 返回一片 fs.FileInfo 文件，按名称排序，描述命名目录的内容。如果 ReadDir 为 nil，Import 将使用 os.ReadDir.FileInfo 返回文件名。
	ReadDir func(dir string) ([]fs.FileInfo, error)

	// OpenFile 打开一个文件（而不是目录）供读取。如果 OpenFile 为 nil，Import 会使用 os.Open.OpenFile 来打开文件。
	OpenFile func(path string) (io.ReadCloser, error)
}
```

上下文指定了构建的支持上下文。

```go
var Default Context = defaultContext()
```

Default 是用于构建的默认 Context。如果设置了 GOARCH、GOOS、GOROOT 和 GOPATH 环境变量，它将使用这些变量，否则将使用编译后代码的 GOARCH、GOOS 和 GOROOT。

#### func (ctxt *Context) Import(path string, srcDir string, mode ImportMode) (*Package, error)

Import 返回由导入路径命名的 Go 软件包的详细信息，并解释相对于 srcDir 目录的本地导入路径。如果路径是一个本地导入路径，命名了一个可以使用标准导入路径导入的软件包，返回的软件包将把 p.ImportPath 设置为该路径。

在包含软件包的目录中，.go、.c、.h 和 .s 文件被视为软件包的一部分，但以下文件除外：

软件包文档中的 .go 文件
以 _ 或 .开头的文件（可能是编辑器临时文件）
上下文中不符合构建约束的文件
如果发生错误，Import 将返回一个非零错误和一个包含部分信息的非零 *Package 文件。

#### func (ctxt *Context) ImportDir(dir string, mode ImportMode) (*Package, error)

ImportDir 与 Import 类似，但处理的是在指定目录中找到的 Go 软件包。

#### func (ctxt *Context) MatchFile(dir, name string) (match bool, err error) 添加于1.2

MatchFile 报告给定目录中带有给定名称的文件是否与上下文相匹配，是否会包含在由 ImportDir 在该目录下创建的软件包中。

MatchFile 会考虑文件名，并可能使用 ctxt.OpenFile 读取部分或全部文件内容。

#### func (ctxt *Context) SrcDirs() []string

SrcDirs 返回软件包源代码根目录列表。它取自当前 Go 根目录和 Go 路径，但会忽略不存在的目录。

### type Directive 添加于1.21.0

```go
type Directive struct {
	Text string // 整行注释，包括前导斜线
	Pos token.Position // 注释的位置
}
```

指令是源文件中的 Go 指令注释 (//go:zzz...)。

### type ImportMode

```go
type ImportMode uint
```

导入模式控制导入方法的行为。

```go
const (
	// 如果设置了 FindOnly，Import 会在找到包含软件包源代码的目录后停止。它不会读取目录中的任何文件。
	FindOnly ImportMode = 1 << iota

	// 如果设置了 AllowBinary，则可以通过编译后的 包对象满足导入要求。
	//
	// 过时：创建仅编译包的支持方法是编写源代码，并在文件顶部包含 go:binary-only-package 注释。这样的软件包将被识别（因为它有源代码），并且在返回的软件包中将 BinaryOnly 设置为 true。
	AllowBinary

	// 如果设置了 ImportComment，则会解析包语句中的导入注释。如果 Import 发现无法理解的注释或在多个源文件中发现相互冲突的注释，则会返回错误。更多信息请参阅 golang.org/s/go14customimport。
	ImportComment

	// 默认情况下，"导入 "会搜索供应商目录 然后再搜索 GOROOT 和 GOPATH 根目录。如果 Import 找到并返回一个使用供应商目录的软件包，那么生成的 ImportPath 就是该软件包的完整路径，包括 "vendor "之前的路径元素。例如，如果 Import("y", "x/subdir", 0) 找到 "x/vendor/y"，返回软件包的 ImportPath 就是 "x/vendor/y"，而不是简单的 "y"。更多信息请参见 golang.org/s/go15vendor。
	//
	// 设置 IgnoreVendor 会忽略供应商目录。
	//
  // 与软件包的 ImportPath 相反，返回的软件包的 Imports、TestImports 和 XTestImports 始终是源文件的精确导入路径： Import 不会尝试解析或检查这些路径。
	IgnoreVendor
)
```

### type MultiplePackageError 添加于1.4

```go
type MultiplePackageError struct {
	Dir string // 包含文件的目录
	Packages []string // 找到的软件包名称
	Files []string // 对应的文件：Files[i] 声明包 Packages[i]
}
```

MultiplePackageError 描述了包含多个软件包的多个可编译 Go 源文件的目录。

#### func (e *MultiplePackageError) Error() string 添加于1.4

### type NoGoError

```go
type NoGoError struct {
	Dir string
}
```

NoGoError 是 Import 用来描述不包含可编译 Go 源文件的目录的错误。(它可能仍然包含测试文件、被编译标记隐藏的文件等）。

#### func (e *NoGoError) Error() string

### type Package

```go
type Package struct {
	Dir string // 包含软件包源代码的目录
	Name string // 软件包名称
	ImportComment string // 包声明中导入注释的路径
	Doc string // 文档概要
	ImportPath string // 软件包的导入路径（如果未知，则为""）
	Root string // 该软件包所在 Go 树的根目录
	SrcRoot string // 软件包源根目录（如果未知，则为""）
	PkgRoot string // 软件包安装根目录（如果未知，则为""）
	PkgTargetRoot string // 依赖于架构的安装根目录（如果未知，则为""）
	BinDir string // 命令安装目录（如果未知，则为""）
	Goroot bool // 在 Go 根目录中找到软件包
	PkgObj string // 安装的 .a 文件
	AllTags []string // 可影响此目录中文件选择的标记
	ConflictDir string // 此目录与 $GOPATH 中的 Dir 有阴影
	BinaryOnly bool // 不能从源代码重建（有 //go:binary-only-package 注释）

	// 源文件
	GoFiles []string // .go 源文件（不包括 CgoFiles、TestGoFiles、XTestGoFiles）
	CgoFiles []string // 导入 "C "的 .go 源文件
	IgnoredGoFiles []string // 本次构建忽略的 .go 源文件（包括忽略的 _test.go 文件）
	InvalidGoFiles []string // 检测到问题（解析错误、软件包名称错误等）的 .go 源文件
	IgnoredOtherFiles []string // 本次构建忽略的非 .go 源文件
	CFiles []string // .c 源文件
	CXXFiles []string // .cc、.pp 和 .cxx 源文件
	MFiles []string // .m (Objective-C) 源文件
	HFiles []string // .h、.hh、.hpp 和 .hxx 源文件
	FFiles []string // .f、.F、.for 和 .f90 Fortran 源文件
	SFiles []string // .s 源文件
	SwigFiles []string // .swig 文件
	SwigCXXFiles []string // .swigcxx 文件
	SysoFiles []string // 要添加到归档中的 .syso 系统对象文件

	// Cgo 指令
	CgoCFLAGS []string // Cgo CFLAGS 指令
	CgoCPPFLAGS []string // Cgo CPPFLAGS 指令
	CgoCXXFLAGS []string // Cgo CXXFLAGS 指令
	CgoFFLAGS []string // Cgo FFLAGS 指令
	CgoLDFLAGS []string // Cgo LDFLAGS 指令
	CgoPkgConfig []string // Cgo pkg-config 指令

	// 测试信息
	TestGoFiles []string // 软件包中的 _test.go 文件
	XTestGoFiles []string // 软件包外的 _test.go 文件

	// 在源文件中发现的 Go 指令注释 (//go:zzz...)。
	Directives []Directive
	TestDirectives []Directive
	XTestDirectives []Directive

	// 依赖关系信息
	导入 []string // 从 GoFiles、CgoFiles 导入路径
	ImportPos map[string][]token.Position // 进口的行信息
	TestImports []string // 从 TestGoFiles 导入路径
	TestImportPos map[string][]token.Position // TestImports 的行信息
	XTestImports []string // 从 XTestGoFiles 导入路径
	XTestImportPos map[string][]token.Position // XTestImports 的行信息

	// 在 Go 源文件中发现的 //go:embed 模式
	// 例如，如果一个源文件中写道
	//go:embed a* b.c
	// 那么列表中将包含这两个单独的字符串。
	// 关于 //go:embed 的更多详情，请参阅包 embed。
	EmbedPatterns []string // 来自 GoFiles、CgoFiles 的模式
	EmbedPatternPos map[string][]token.Position // EmbedPatterns 的行信息
	TestEmbedPatterns []string // TestGoFiles 中的模式
	TestEmbedPatternPos map[string][]token.Position // TestEmbedPatterns 的行信息
	XTestEmbedPatterns []string // XTestGoFiles 中的模式
	XTestEmbedPatternPos map[string][]token.Position // XTestEmbedPatternPos 的行信息
}
```

软件包描述在某个目录中找到的 Go 软件包。

#### func Import(path, srcDir string, mode ImportMode) (*Package, error)

Import 是 Default.Import 的简称。

#### func ImportDir(dir string, mode ImportMode) (*Package, error)

ImportDir 是 Default.ImportDir 的简称。

#### func (p *Package) IsCommand() bool

IsCommand 报告软件包是否被视为要安装的命令（而不仅仅是库）。名为 "main "的软件包被视为命令。