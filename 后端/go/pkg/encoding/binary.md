## package binary

二进制包实现了数字和字节序列之间的简单转换，以及变量的编码和解码。

数字是通过读写固定大小的值进行转换的。固定大小值可以是固定大小的算术类型（bool、int8、uint8、int16、float32、complex64......），也可以是只包含固定大小值的数组或结构体。

varint 函数使用可变长度编码对单整数值进行编码和解码；较小的数值需要较少的字节。有关规范，请参见 https://developers.google.com/protocol-buffers/docs/encoding。

本软件包更倾向于简单而非高效。需要高性能序列化的客户端，尤其是大型数据结构，应考虑更先进的解决方案，如编码/gob 包或协议缓冲区。

## Index

### Constants

```go
const (
	MaxVarintLen16 = 3
	MaxVarintLen32 = 5
	MaxVarintLen64 = 10
)
```

MaxVarintLenN 是变量编码 N 位整数的最大长度。

### Variables

```go
var BigEndian bigEndian
```

BigEndian 是 ByteOrder 和 AppendByteOrder 的 big-endian 实现。

```go
var LittleEndian littleEndian
```

LittleEndian 是 ByteOrder 和 AppendByteOrder 的小二进制实现。

```go
var NativeEndian nativeEndian
```

NativeEndian 是 ByteOrder 和 AppendByteOrder 的本地对数实现。

### func AppendUvarint(buf []byte, x uint64) []byte 添加于1.19

AppendUvarint 将 PutUvarint 生成的 x 的 varint 编码形式追加到 buf 中，并返回扩展缓冲区。

### func AppendVarint(buf []byte, x int64) []byte 添加于1.19

AppendVarint 将 PutVarint 生成的 x 的 varint 编码形式追加到 buf 中，并返回扩展缓冲区。

### func PutUvarint(buf []byte, x int64) int

PutUvarint 将 uint64 编码到 buf 中，并返回写入的字节数。如果缓冲区太小，PutUvarint 就会崩溃。

### func PutVarint(buf []byte, x int 64) int

PutVarint 将 int64 编码到 buf 中，并返回写入的字节数。如果缓冲区太小，PutVarint 就会崩溃。

### func Read(r io.Reader, order ByteOrder, data any) error

读取将结构化二进制数据从 r 读入数据。数据必须是指向固定大小值或固定大小值片段的指针。从 r 中读取的字节将按照指定的字节顺序解码，并写入数据的连续字段。在解码布尔值时，0 字节会被解码为 false，而任何其他非 0 字节都会被解码为 true。读入结构体时，会跳过字段名为空白 (_) 的字段数据；也就是说，空白字段名可用于填充。在读入结构体时，必须导出所有非空白字段，否则读取时可能会出错。

只有在没有读取字节的情况下，才会出现 EOF 错误。如果在读取部分字节而非全部字节后发生 EOF，Read 会返回 ErrUnexpectedEOF。

### func ReadUvarint(r io.ByteReader) (uint64, error)

ReadUvarint 从 r 中读取一个已编码的无符号整数，并以 uint64 的形式返回。只有在没有读取任何字节的情况下，才会出现 EOF 错误。如果读取了部分字节而不是全部字节后出现 EOF，ReadUvarint 将返回 io.ErrUnexpectedEOF。

### func ReadVarint(r io.ByteReader) (int64, error)

ReadVarint 从 r 中读取编码后的带符号整数，并以 int64 的形式返回。只有在没有读取任何字节的情况下，才会出现 EOF 错误。如果在读取部分字节而非全部字节后出现 EOF，ReadVarint 会返回 io.ErrUnexpectedEOF。

### func Size(v any) int

值 v 必须是固定大小的值或固定大小值的片段，或者是指向此类数据的指针。如果 v 都不是，Size 将返回-1。

### func Uvarint(buf []byte) (uint64, int)

Uvarint 从 buf 解码 uint64，并返回该值和读取的字节数（> 0）。如果发生错误，则返回值为 0，字节数 n <= 0：

```go
n == 0: buf too small
n  < 0: value larger than 64 bits (overflow)
        and -n is the number of bytes read
```

### func Varint(buf []byte) (int64, int)

Varint 从 buf 解码 int64，并返回该值和读取的字节数（> 0）。如果发生错误，则返回值为 0，字节数 n <= 0，含义如下：

```go
n == 0: buf too small
n  < 0: value larger than 64 bits (overflow)
        and -n is the number of bytes read
```

### func Write(w io.Writer, order ByteOrder, data any) error

数据必须是固定大小的值或固定大小值的片段，或者是指向此类数据的指针。布尔值编码为一个字节：1 表示 "真"，0 表示 "假"。写入 w 的字节按照指定的字节顺序编码，并从数据的连续字段中读取。在写入结构体时，对于字段名为空白 (_) 的字段，将写入 0 值。

### type AppendByteOrder 添加于1.19

```go
type AppendByteOrder interface {
	AppendUint16([]byte, uint16) []byte
	AppendUint32([]byte, uint32) []byte
	AppendUint64([]byte, uint64) []byte
	String() string
}
```

AppendByteOrder 指定如何将 16、32 或 64 位无符号整数追加到字节片中。

### ByteOrder

```go
type ByteOrder interface {
	Uint16([]byte) uint16
	Uint32([]byte) uint32
	Uint64([]byte) uint64
	PutUint16([]byte, uint16)
	PutUint32([]byte, uint32)
	PutUint64([]byte, uint64)
	String() string
}
```

ByteOrder 指定如何将字节片转换为 16、32 或 64 位无符号整数。