## package embed

包 embed 允许访问运行中的 Go 程序中嵌入的文件。

导入 "embed "的 Go 源文件可以使用 //go:embed 指令来初始化字符串、[]字节或 FS 类型的变量，使其包含编译时从包目录或子目录读取的文件内容。

例如，以下是嵌入名为 hello.txt 的文件并在运行时打印其内容的三种方法。

将一个文件嵌入一个字符串：

```go
import _ "embed"

//go:embed hello.txt
var s string
print(s)
```

将一个文件嵌入到一片字节中：

```go
import _ "embed"

//go:embed hello.txt
var b []byte
print(string(b))
```

将一个或多个文件嵌入文件系统：

```go
import "embed"

//go:embed hello.txt
var f embed.FS
data, _ := f.ReadFile("hello.txt")
print(string(data))
```

## Directives

变量声明上方的 //go:embed 指令使用一个或多个 path.Match 模式指定要嵌入的文件。

该指令必须紧接在包含单个变量声明的行之前。在指令和变量声明之间只允许有空行和"//"行注释。

变量类型必须是字符串类型，或字节类型的片段，或 FS（或 FS 的别名）。

例如:

```go
package server

import "embed"

// 内容包含我们的静态网络服务器内容。
//go:embed image/* template/*
//go:embed html/index.html
var content embed.FS
```

Go 构建系统会识别这些指令，并安排将文件系统中匹配的文件填充到声明的变量（在上面的例子中为 content）中。

为了简洁起见，//go:embed 指令接受多个空格分隔的模式，但也可以重复，以避免在模式较多时出现冗长的行。模式相对于包含源文件的软件包目录进行解释。路径分隔符是正斜线，即使在 Windows 系统上也是如此。模式不能包含"."、".. "或空路径元素，也不能以斜线开始或结束。要匹配当前目录中的所有内容，请使用 "*"而不是"."。为使文件命名时能包含空格，模式可以写成 Go 双引号或反引号字符串。

如果模式命名了一个目录，以该目录为根的子树中的所有文件都会被嵌入（递归），但名称以". "或"_"开头的文件除外。因此，上例中的变量几乎等同于:

```go
// 内容是我们的静态网络服务器内容。
//go:embed image template html/index.html
var content embed.FS
```

区别在于 "image/*"嵌入了 "image/.tempfile"，而 "image "没有。两者都不嵌入 "image/dir/.tempfile"。

如果一个模式以前缀 "all: "开头，那么行走目录的规则就会改为包含以". "或"_"开头的文件。例如，"all:image "同时包含 "image/.tempfile "和 "image/dir/.tempfile"。

//go:embed 指令既可用于导出变量，也可用于未导出变量，这取决于软件包是否希望将数据提供给其他软件包。它只能用于包作用域的变量，而不能用于本地变量。

模式不能匹配软件包模块之外的文件，如".git/*"或符号链接。模式不能匹配文件名包含特殊标点符号 " * < > ?` ' | / 和 :.空目录的匹配将被忽略。之后，//go:embed 行中的每个模式必须至少匹配一个文件或非空目录。

如果有任何模式无效或匹配无效，构建将失败。

**Strings and Bytes**

字符串或[]字节类型变量的 //go:embed 行只能有一个模式，且该模式只能匹配一个文件。字符串或[]字节将被初始化为该文件的内容。

//go:embed 指令要求导入 "embed"，即使使用的是字符串或[]字节。在不引用 embed.FS 的源文件中，请使用空白导入（import _ "embed"）。

**File Systems**

要嵌入单个文件，通常最好使用字符串或[]字节类型的变量。FS 类型可以嵌入文件树，例如上面例子中的静态网络服务器内容目录。

FS 实现了 io/fs 软件包的 FS 接口，因此可以与任何理解文件系统的软件包一起使用，包括 net/http、text/template 和 html/template。

例如，给定上例中的 content 变量，我们可以编写:

```go
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(content))))

template.ParseFS(content, "*.tmpl")
```

**Tools**

为了支持分析 Go 软件包的工具，"go list "输出中提供了 //go:embed 行中的模式。请参阅 "go help list "输出中的 EmbedPatterns、TestEmbedPatterns 和 XTestEmbedPatterns 字段。

## Index

### type FS

```go
type FS struct {
  // 包含已筛选或未导出字段
}
```

FS 是只读文件集合，通常使用 //go:embed 指令初始化。如果声明时没有 //go:embed 指令，FS 就是一个空文件系统。

FS 是只读值，因此可以在多个程序中同时使用，也可以将 FS 类型的值互相赋值。

FS 实现了 fs.FS，因此可以与任何理解文件系统接口的软件包一起使用，包括 net/http、text/template 和 html/template。

有关初始化 FS 的更多详情，请参阅软件包文档。

#### func (f FS) Open(name string) (fs.File, error)

Open 打开命名的文件以供读取，并以 fs.File 格式返回。

当文件不是目录时，返回的文件会实现 io.Seeker 和 io.ReaderAt。

#### func (f FS) ReadDir(name string) ([]fs.DirEntry, error)

ReadDir 读取并返回整个命名目录。

#### func (f FS) ReadFile(name string) ([]byte, error)

ReadFile 读取并返回指定文件的内容。