## package tar
import "archive/tar"

磁带归档(tar)是一种文件格式，用于存储一系列可以以流方式读写的文件。这个包旨在涵盖格式的大多数变体，包括那些由 GNU 和 BSD tar 工具产生的变体。

## Index
[Constants](#constants)
[Variables](#variables)
[type Format](#format)
    [func (f Format) String() string](#string)
[type Header](#header)
    [func FileInfoHeader(fi fs.FileInfo, link string) (*Header, error)](#file-info-header)
    [func (h *Header) FileInfo() fs.FileInfo](#file-info)
[type Reader](#reader)
    [func NewReader(r io.Reader) *Reader](#new-reader)
    [func (tr *Reader) Next() (*Header, error)](#next)
    [func (tr *Reader) Read(b []byte) (int, error)](#read)
[type Writer](#writer)
    [func NewWriter(w io.Writer) *Writer](#new-writer)
    [func (tw *Writer) Close() error](#close)
    [func (tw *Writer) Flush() error](#flush)
    [func (tw *Writer) Write(b []byte) (int, error)](#write)
    [func (tw *Writer) WriteHeader(hdr *Header) error](#write-header)

## <a id="constants">Constants</a>

```go
const (
	TypeReg = '0' // 0表示普通文件
	TypeRegA = '\x00' // 普通文件 不推荐：改为使用 TypeReg

	// 类型“1”到“6”是只有标头的标志，可能没有数据体。
	TypeLink    = '1' // 硬链接
	TypeSymlink = '2' // 符号连接
	TypeChar    = '3' // 字符设备节点
	TypeBlock   = '4' // 块设备节点
	TypeDir     = '5' // 目录
	TypeFifo    = '6' // 先进先出队列节点

	TypeCont = '7' // 保留位

	//类型‘x’由 PAX 格式用于存储键值记录
  //只与下一个文件相关。
  //此包透明地处理这些类型。
	TypeXHeader = 'x'

  //类型‘g’由 PAX 格式用于存储键值记录
  //与所有后续文件相关。
  //这个包只支持解析和组合这样的头文件,
  //但当前不支持跨文件持久化全局状态。
	TypeXGlobalHeader = 'g'

	//类型‘S’表示 GNU 格式的稀疏文件。
	TypeGNUSparse = 'S'

  //类型‘L’和‘K’由 GNU 格式用于元文件
  //用于存储下一个文件的路径或链接名。
  //此包透明地处理这些类型。
	TypeGNULongName = 'L'
	TypeGNULongLink = 'K'
)
```

## **<a id="variables">Variables</a>**

```go
var (
	ErrHeader          = errors.New("archive/tar: invalid tar header")
	ErrWriteTooLong    = errors.New("archive/tar: write too long")
	ErrFieldTooLong    = errors.New("archive/tar: header field too long")
	ErrWriteAfterClose = errors.New("archive/tar: write after close")
	ErrInsecurePath    = errors.New("archive/tar: insecure file path")
)
```

## Types

### **<a id="type-format">type Format</a>**

```go
type Format int
```

标识各种 tar 格式的常数。

#### **<a id="string">func (Format) String</a>**

```go
func (f Format) String() string
```

### type Header

```go
type Header struct {
  // TypeFlag 是头条目的类型。
  // 零值自动提升为 TypeReg 或 TypeDir
  // 取决于 Name 中是否存在尾部斜杠。
	Typeflag byte

	Name     string // 记录头域的文件名
	Linkname string // 链接的目标名称(对 TypeLink 或 TypeSymlink 有效)

	Size  int64  // 逻辑文件大小(以字节为单位)
	Mode  int64  // 权限和模式位
	Uid   int    // 所有者的用户 ID
	Gid   int    // 所有者的组 ID
	Uname string // 所有者的用户名
	Gname string // 所有者的组名

  // 如果未指定 Format，则 Writer.WriteHeader 将调整 ModTime
  // 设置为最接近的秒，并忽略 AccessTime 和 ChangeTime 字段。
  // 要使用 AccessTime 或 ChangeTime，请将格式指定为 PAX 或 GNU。
  // 要使用亚秒级分辨率，请将格式指定为 PAX。
	ModTime    time.Time // 修改时间
	AccessTime time.Time // 存取时间 (需要 PAX 或 GNU 支持)
	ChangeTime time.Time // 更改时间(需要 PAX 或 GNU 支持)

	Devmajor int64 // 主设备号(对 TypeChar 或 TypeBlock 有效)
	Devminor int64 // 次要设备号(对 TypeChar 或 TypeBlock 有效)

  // Xattrs 将扩展属性作为 PAX 记录存储在
  // “ SCHILY.xattr.”名称空间。
  // 
  // 以下内容在语义上是等同的:
  // h. Xattrs [ key ] = value
  // h.PAXRecords [“ SCHILY.xattr.”+ key ] = value
  // 
  // 当 Writer.WriteHeader 被调用时，Xattrs 的内容将采用
  // 优先于 PAXRecords。
  // 
  // 已弃用: 改用 PAXRecords。
	Xattrs map[string]string

	// PAXRecords 是 PAX 扩展头记录的映射。
	//
	// 用户定义的记录应有以下表格的键:
	//	VENDOR.keyword
	// 其中 VENDOR 是某个名称空间的全部大写字母，而关键字可以
	// 不包含’=’字符(例如，“ GOLANG.pkg.version”)。
	// 键和值应该是非空的 UTF-8字符串。
	//
	// 调用 Writer.WriteHeader 时，派生自
	// Header 中的其他字段优先于 PAXRecords。
	PAXRecords map[string]string

	// Format 指定 tar 头的格式。
	//
	// 这是由 Reader 设置的。
	// 因为 Reader 会阅读一些不兼容的文件,
	// 这可能是未知格式。
	//
	// /如果在调用 Writer.WriteHeader 时未指定格式,
	// 然后使用第一种格式(按照 USTAR、 PAX、 GNU 的顺序)
	// 能够编码此标题(请参阅格式)。
	Format Format
}
```

Header 表示 tar 归档文件中的一个头。有些字段可能没有填充。

为了向前兼容，从 Reader 检索 Header 的用户。接下来，以某种方式对其进行变异，然后将其传递回 Writer。WriteHeader 应该通过创建一个新的 Header 并复制它们感兴趣保留的字段来完成。

#### **<a id="file-info-header">func FileInfoHeader</a>**

```go
func FileInfoHeader(fi fs.FileInfo, link string) (*Header, error)
```

FileInfoHeader从fi创建一个部分填充的Header。如果fi描述符号链接，则FileInfoHeader将link记录为链接目标。如果fi描述一个目录，则在名称后面附加一个斜杠。

由于fs.FileInfo的Name方法只返回它所描述的文件的基名称，因此可能需要修改Header.Name以提供文件的完整路径名。

#### **<a id="file-info">func (*Header) FileInfo</a>**

```go
func (h *Header) FileInfo() fs.FileInfo
```

FileInfo为Header返回一个fs.FileInfo。

### type Reader

```go
type Reader struct {
	// 包含经过筛选或未导出的字段
}
```

Reader提供对tar归档文件内容的顺序访问。Reader.Next前进到归档文件中的下一个文件（包括第一个文件），然后Reader可以被视为io.Reader来访问文件的数据。

#### **<a id="new-reader">func NewReader</a>**

```go
func NewReader(r io.Reader) *Reader
```

NewReader创建一个新的Reader，从r读取数据。

#### **<a id="next">func (*Reader) Next</a>**

```go
func (tr *Reader) Next() (*Header, error)
```

#### **<a id="read">Read</a>**

```go
func (tr *Reader) Read(b []byte) (int, error)
```

读取从tar归档中的当前文件中读取。当它到达该文件的末尾时，它将返回（0，io.EOF），直到调用Next前进到下一个文件。

如果当前文件是稀疏的，则标记为孔的区域将作为NULL字节读回。

对TypeLink、TypeSymlink、TypeChar、TypeBlock、TypeDir和TypeFifo等特殊类型调用Read将返回（0，io.EOF），而不管Header.Size声明什么。

### type Writer

```go
type Writer struct {
  // 包含经过筛选或未导出的字段
}
```

Writer提供对tar存档的顺序写入。WriteHeader以提供的Header开始一个新文件，然后Writer可以被当作一个io.Writer来提供该文件的数据。

#### **<a id="new-writer">func NewWriter</a>**

```go
func NewWriter(w io.Writer) *Writer
```

NewWriter创建一个新的Writer写入w。

#### **<a id="close">func (*Writer) Close</a>**

```go
func (tw *Writer) Close() error
```

Close通过刷新填充和写入页脚来关闭tar存档。如果当前文件（来自对WriteHeader的先前调用）未完全写入，则返回错误。

#### **<a id="flush">func (*Writer) Flush</a>**

```go
func (tw *Writer) Flush() error
```

刷新完成当前文件的块填充的写入。在调用Flush之前，必须完全写入当前文件。

这是不必要的，因为下一次调用WriteHeader或Close将隐式地清除文件的填充。

#### **<a id="write">func (*Writer)</a>**

```go
func (tw *Writer) Write(b []byte) (int, error)
```

写入tar归档中的当前文件。如果超过Header，则Write返回错误ErrWriteTooLong。在WriteHeader之后写入大小字节。

对TypeLink、TypeSymlink、TypeChar、TypeBlock、TypeDir和TypeFifo等特殊类型调用Write将返回（0，ErrWriteTooLong），而不管Header.Size声明什么。

#### **<a id="write-header">func (*Writer) WriteHeader</a>**

```go
func (tw *Writer) WriteHeader(hdr *Header) error
```

WriteHeader写入hdr并准备接受文件的内容。Header.Size决定下一个文件可以写入多少字节。如果当前文件未完全写入，则返回错误。这会在写入标头之前隐式地刷新任何必要的填充。