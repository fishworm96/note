## package zip
import "archive/zip"

ZIP包支持读写 ZIP 归档文件。

本包不支持跨硬盘的压缩。

关于ZIP64：
为了向后兼容，FileHeader具有32位和64位大小字段。64位字段将始终包含正确的值，并且对于正常存档，两个字段将相同。对于需要ZIP64格式的文件，32位字段将为0xffffffff，必须使用64位字段。

## Index

[Constants](#constants)
[Variables](#variables)
    [func RegisterCompressor(method uint16, comp Compressor)](#register-compressor)
    [func RegisterDecompressor(method uint16, dcomp Decompressor)](#register-decompressor)
[type Compressor](#compressor)
[type Decompressor](#decompressor)
[type File](#file)
    [func (f *File) DataOffset() (offset int64, err error)](#data-offset)
    [func (f *File) Open() (io.ReadCloser, error)](#open)
    [func (f *File) OpenRaw() (io.Reader, error)](#open-raw)
[type FileHeader](#file-header)
    [func FileInfoHeader(fi fs.FileInfo) (*FileHeader, error)](#file-info-header)
    [func (h *FileHeader) FileInfo() fs.FileInfo](#file-info)
    [func (h *FileHeader) ModTime() fs.FileInfo](#mod-time) 弃用
    [func (h *FileHeader) Mode() (mode fs.FileMode)](#mode)
    [func (h *FileHeader) SetModTime(t time.Time)](#set-mod-time) 弃用
    [func (h *FileHeader) SetMode(mode fs.FileMode)](#set-mode)
[type ReadCloser](#read-closer)
    [func OpenReader(name string) (*ReadCloser, error)](#open-reader)
    [func (rc *ReadCloser) Close() error](#read-close)
[type Reader](#reader)
    [func NewReader(r io.ReaderAt, size int64) (*Reader, error)](#new-reader)
    [func (r *Reader) Open(name string) (fs.FIle, error)](#open)
    [func (r *Reader) RegisterDecompressor(method uint16, dcomp Decompressor)](#register-decompress)
[type Writer](#writer)
    [func NewWriter(w io.Writer) *Writer](#new-writer)
    [func (w *Writer) Close() error](#writer-close)
    [func (w *Writer) Copy(f *File) error](#copy)
    [func (w *Writer) Create(name string) (io.Writer, error)](#create)
    [func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)](#create-header)
    [func (w *Writer) CreateRaw(fh *FileHeader) (io.Writer, error)](#create-raw)
    [func (w *Writer) Flush() error](#flush)
    [func (w *Writer) RegisterCompressor(method uint16, comp Compressor)](#register-compressor)
    [func (w *Writer) SetComment(comment string) error](#set-comment)
    [func (w *Writer) SetOffset(n int64)](#set-offset)

## <a id="constants">Constants</a>

```go
const (
  Store uint16 = 0 // 无压缩
  Deflate uint16 = 8 // 压缩
)
```

压缩方法

## **<a id="variables">Variables</a>**

```go
var (
	ErrFormat       = errors.New("zip: not a valid zip file")
	ErrAlgorithm    = errors.New("zip: unsupported compression algorithm")
	ErrChecksum     = errors.New("zip: checksum error")
	ErrInsecurePath = errors.New("zip: insecure file path")
)
```

## **Functions**

### **<a id="register-compressor">func RegisterCompressor</a>**

```go
func RegisterCompressor(method uint16, comp Compressor)
```

RegisterCompressor为指定的方法ID注册自定义压缩器。内置了常用的Store和Deflate方法。

### **<a id="register-decompressor">func RegisterDecompressor</a>**

```go
func RegisterDecompressor(method uint16, dcomp Decompressor)
```

RegisterDecompressor允许为指定的方法ID定制解压缩器。内置了常用的Store和Deflate方法。

## Types

### **<a id="compressor">type Compressor</a>**

```go
type Compressor func(w io.Writer) (io.WriteCloser, error)
```

Compressor返回一个新的compressing writer，写入w。必须使用WriteCloser的Close方法将挂起的数据刷新到w。Compressor本身必须安全地同时从多个goroutine调用，但是每个返回的writer一次只能由一个goroutine使用。

### **<a id="decompressor">type Decompressor</a>**

```go
type Decompressor func(r io.Reader) io.ReadCloser
```

解压缩器返回一个新的解压缩读取器，从r阅读。必须使用ReadCloser的Close方法来释放关联的资源。解压缩器本身必须安全地同时从多个goroutine调用，但每个返回的读取器一次只能由一个goroutine使用。

### **<a id="file">type File</a>**

```go
type File struct {
  FileHeader
  // 包含经过筛选或未导出的字段
}
```

文件是ZIP存档中的单个文件。文件信息位于嵌入的FileHeader中。可以通过调用Open来访问文件内容。

#### **<a id="data-offset">func (*File) DataOffset</a>**

```go
func (f *File) DataOffset() (offset int64, err error)
```

DataOffset返回文件可能压缩的数据相对于 zip 文件开头的偏移量。

大多数调用者应该使用Open，它透明地解压缩数据并验证校验和。

#### **<a id="open">func (*File) Open</a>**

```go
func (f *File) Open() (io.ReadCloser, error)
```

Open返回一个ReadCloser，它提供对File内容的访问。可以同时读取多个文件。

#### **<a id="open-raw">func (*File) OpenRaw</a>**

```go
func (f *File) OpenRaw() (io.Reader, error)
```

OpenRaw返回一个读取器，该读取器提供对文件内容的访问而无需解压缩。

### **<a id="file-header">type FileHeader</a>**

```go
type FileHeader struct {
  // Name 是文件的名称。
  // 
  // 它必须是一个相对路径，而不是以驱动器字母开头(如“ C:”) ,
  // 必须使用正斜杠而不是反斜杠
  // 表示该文件是一个目录，应该没有数据。
	Name string

	// “注释”为任意短于 64KiB 的用户定义字符串。
	Comment string

	// NonUTF8表示 Name 和 comments 不是用 UTF-8编码的。
	//
	// 根据规范，唯一允许的其他编码应该是 CP-437,
  // 但从历史上看，许多 ZIP 读者将 Name 和 comments 解释为任何东西系统的本地字符编码正好是。
  // 
  // 只有当用户打算编码一个不可移植的为特定的本地化区域提供
  // ZIP 文件
  // 自动为有效的 UTF-8字符串设置 ZIP 格式的 UTF-8标志。
	NonUTF8 bool

	CreatorVersion uint16
	ReaderVersion  uint16
	Flags          uint16

	// Method 为压缩方法。如果为零，则使用存储。
	Method uint16

  // Modified 是文件的修改时间。
  // 
  // 读取时，扩展时间戳优于传统 MS-DOS
  // date 字段，时间之间的偏移量用作时区。
  // 如果只有 MS-DOS 日期，则假定时区为 UTC。
  // 
  // 在编写时，扩展时间戳(与时区无关)是
  // 总是发出。遗留的 MS-DOS 日期字段根据
  // 更改时间的位置。
	Modified time.Time

	// 修改时间是 MS-DOS 编码的时间。
	//
	// 弃用: 改为 Modified 代替。
	ModifiedTime uint16

	// 修改日期是 MS-DOS 编码的日期。
	//
	// 弃用: 改为 Modified 代替。
	ModifiedDate uint16

	// CRC32 是 CRC32 的文件内容校验总和。
	CRC32 uint32

  // CompressedSize 是文件的压缩大小(以字节为单位)。
  // 如果文件的未压缩或压缩大小
  // 不适合32位，Compressed Size 被设置为 ^uint32(0)。
  // 
  // 启弃用: 改为 CompressedSize64。
	CompressedSize uint32

  // 未压缩的大小是文件的压缩大小(以字节为单位)。
  // 如果文件的未压缩或压缩大小
  // 不适合32位，CompressedSize 被设置为 ^uint32(0)。
  //
	// 弃用: 改为 UncompressedSize64.
	UncompressedSize uint32

	// CompressedSize64 是文件的压缩大小(以字节为单位)
	CompressedSize64 uint64

	// UncompressedSize64 是文件的未压缩大小(以字节为单位)。
	UncompressedSize64 uint64

	Extra         []byte
	ExternalAttrs uint32 // 意义取决于创造者版本
}
```

FileHeader描述ZIP文件中的文件。有关详细信息，请参阅ZIP规范。

#### **<a id="file-info-header">func FileInfoHeader</a>**

```go
func FileInfoHeader(fi fs.FileInfo) (*FileHeader, error)
```

FileInfoHeader从fs. FileInfo创建一个部分填充的FileHeader。因为fs.FileInfo的Name方法只返回它描述的文件的基名称，所以可能需要修改返回的头的Name字段，以提供文件的完整路径名。如果需要压缩，调用方应设置FileHeader.Method字段;默认情况下未设置。

#### **<a id="file-info">func (*FileHeader) FileInfo</a>**

```go
func (h *FileHeader) FileInfo() fs.FileInfo
```

FileInfo 返回 FileHeader 的 fs.FileInfo。

func (*FileHeader) ModTime **弃用**

#### **<a id="mode">func (*FileHeader) Mode</a>**

```go
func (h *FileHeader) Mode() (mode fs.FileMode)
```

Mode 返回 FileHeader 的权限和模式位。

func (*FileHeader) SetModTime **弃用**

#### **<a id="set-mode">func (*FileHeader) SetMode</a>**

```go
func (h *FileHeader) SetMode(mode fs.FileMode)
```

SetMode 更改 FileHeader 的权限和模式位。

### **<a id="read-closer">type ReadCloser</a>**

```go
type ReadCloser struct {
  Reader
  // 包含经过筛选或未导出的字段
}
```

ReadCloser 是一个 Reader，当不再需要时是必须关闭。

#### **<a id="open-reader">func OpenReader</a>**

```go
func OpenReader(name string) (*ReadCloser, error)
```

OpenReader 将打开指定的 ZIP 文件并返回 ReadCloser。

如果归档文件中的任何文件使用非本地名称（如 filepath.IsLocal 所定义）或包含反斜杠的名称并且 GODEBUG 环境变量包含 `zipinsecurepath=0`，OpenReader 将返回读取器并显示 ErrInsecurePath 错误。Go 语言的未来版本可能会默认引入这种行为。想要接受非本地名称的程序可以忽略 ErrInsecurePath 错误并使用返回的读取器。

#### **<a id="read-close">func (*ReadCloser) Close</a>**

```go
func (rc *ReadCloser) Close() error
```

Close 关闭 ZIP 文件，使其无法使用 I/O。

### **<a id="reader">type Reader</a>**

```go
type Reader struct {
  File []*File
  Comment string
  // 包含经过筛选或未导出的字段
}
```

阅读器提供来自 ZIP 存档的内容。

#### **<a id="new-reader">func NewReader</a>**

```go
func NewReader(r io.ReaderAt, size int64) (*Reader, error)
```

NewReader 返回从 r 阅读的新 Reader，假定 r 具有给定的字节大小。

如果归档文件中的任何文件使用非本地名称（如 filepath.IsLocal 定义的）或包含反斜杠的名称，并且 GODEBUG 环境变量包含 `zipinsecurepath=0`，NewReader 将返回读取器并显示 ErrInsecurePath 错误。Go 语言的未来版本可能会默认引入这种行为。想要接受非本地名称的程序可以忽略 ErrInsecurePath 错误并使用返回的读取器。

#### **<a id="open">func (*Reader) Open</a>**

```go
func (r *Reader) Open(name string) (fs.File, error)
```

Open 使用 fs.FS 的语义打开 ZIP 归档文件中的命名文件。Open：路径始终以斜线分割，没有前导 / 或 ../ 元素。

#### **<a id="register-decompressor">func (*Reader) RegisterDecompressor</a>**

```go
func (r *Reader) RegisterDecompressor(method uint16, dcomp Decompressor)
```

RegisterDecompressor 注册或重写特定方法 ID 的自定义解释压缩器。如果未找到给定方法的解压缩器，Reader 将默认在包级别查找解压缩器。

### **<a id="writer">type Writer</a>**

```go
type Writer struct {
  // 包含经过筛选或未导出的字段
}
```

Writer 实现了 zip 文件编写器

#### **<a id="new-writer">func NewWriter</a>**

```go
func NewWriter(w io.Writer) *Writer
```

NewWriter 返回一个新的 Writer，将 zip 文件写入 w。

#### **<a id="writer-close">func (*Writer) Close</a>**

```go
func (w *Writer) Close() error
```

关闭通过写入中央目录完成 zip 文件的写入。它不会关闭底层编写器。

#### **<a id="copy">func (*Writer) Copy</a>**

```go
func (w *Writer) Copy(f *File) error
```

Copy 将文件 f（从 Reader 获得）复制到 w 中。它直接复制原始表单，绕过解压缩、压缩和验证。

#### **<a id="create">func (*Writer) Create</a>**

```go
func (w *Writer) Create(name string) (io.Writer, error)
```

Create 使用提供的名称向 zip 文件添加文件。它返回一个 Writer，文件内容应该写入其中。将使用 Deflate 方法压缩文件内容。名称必须是相对路径：他不能以驱动器号开头（例如，C:）或前导斜杠，并且只允许使用正斜杠。要创建目录而不是文件，请在名称后面添加一个斜杠。在下次调用 Create、CreateHeader 或 Close 之前，必须将文件内容写入 io.Writer。

#### **<a id="create-header">func (*Writer) CreateHeader</a>**

```go
func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)
```

CreateHeader 使用为文件元数据提供的 FileHeader 将文件添加到 zip 存档。Writer 获得 fh 的所有权，并可以改变其字段。调用方在调用 CreateHeader 后不得修改 fh。

这将返回一个 Writer，文件内容应写入该 Writer。在下一次调用 Create、CreateHeader、CreateRaw 或 Close 之前，必须将文件的内容写入 io.Writer。

#### **<a id="create-raw">func (*Writer) CreateRaw</a>**

```go
func (w *Writer) CreateRaw(fh *FileHeader) (io.Writer, error)
```

CreateRaw 使用提供的 FileHeader 将文件添加到 zip 存档中，并返回一个写入文件内容的 Writer。在下一次调用 Create、CreateHeader、CreateRaw 或 Close 之前，必须将文件的内容写入 io.Writer。

与 CreateHeader 相反，传递给 Writer 的字节不被压缩

#### **<a id="flush">func (*Writer) Flush</a>**

```go
func (w *Writer) Flush() error
```

刷新将所有缓存冲数据刷新到基础写入器。调用 Flush 通常不是必须的；关闭就足够了。

#### **<a id="register-compressor">func (*Writer) RegisterCompressor</a>**

```go
func (w *Writer) RegisterCompressor(method uint16, comp Compressor)
```

#### **<a id="set-comment">func (*Writer) SetComment</a>**

```go
func (w *Writer) SetComment(comment string) error
```

SetComment 设置中心目录结束注释字段。只能在关闭前调用。

#### **<a id="set-offset">func (*Writer) SetOffset</a>**

```go
func (W *Writer) SetOffset(n int64)
```

SetOffset 设置底层 write 中 zip 数据开头的偏移量。当 zip 数据附加到现有文件（如二进制可执行文件）时，应使用此选项。必须在写入任何数据之前调用它。