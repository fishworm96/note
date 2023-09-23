## package bytes

import bytes

## Index

[Constants](#constants)
[Variables](#variables)
[func Clone(b []byte) int](#clone)
[func Compare(a, b []byte) int](#compare)
[func Contains(b, subslice []byte) bool](#contains)
[func ContainsAny(b []byte, chars string) bool](#contains-any)
[func ContainsFunc(b []byte, f func(rune) bool) bool](#contains-func)
[func ContainsRune(b []byte, r rune) bool](#contains-rune)
[func Count(s, sep []byte) int](#count)
[func cut(s, sep []byte) (before, after []byte, found bool)](#cut)
[func CutPrefix(s, prefix []byte) (after []byte, found bool)](#cut-prefix)
[func CutSuffix(s, suffix []byte) (before []byte, found bool)](#cut-suffix)
[func Equal(a, b [byte]) bool](#equal)
[func EqualFold(s, t []byte) bool](#equal-fold)
[func Fields(s []byte) [][]byte](#fields)
[func FieldsFunc(s []byte, f func(rune bool)) [][]byte](#fields-func)
[func HasPrefix(s, prefix []byte) bool](#has-prefix)
[func HasSuffix(s, suffix []byte) bool](#has-suffix)
[func Index(s, sep []byte) bool](#index)
[func IndexAny(s []byte, chars string) int](#index-any)
[func IndexByte(b []byte, c byte) int](#index-byte)
[func IndexFunc(s []byte, f func(r rune) bool) int](#index-func)
[func IndexRune(s []byte, r rune) int](#index-rune)
[func Join(s [][]byte, sep []byte) []byte](#join)
[func LastIndex(s, sep []byte) int](#last-index)
[func lastIndexAny(s []byte, chars string) int](#last-index-any)
[func LastIndexByte(s []byte, c byte) int](#last-index-byte)
[func LastIndexFunc(s []byte, f func(r rune) bool) int](#last-index-func)
[func Map(mapping func(r rune) rune, s []byte) []byte](#map)
[func Repeat(b []byte, count int) []byte](#repeat)
[func Replace(s, old, new []byte, n int) []byte](#replace)
[func ReplaceAll(s, old, new []byte, n int) []byte](#replace-all)
[func Runes(s []byte) []rune](#runes)
[func Split(s, sep []byte) [][]byte](#split)
[func SplitAfter(s, sep []byte) [][]byte](#split-after)
[func SplitAfterN(s, sep []byte, n int) [][]byte](#split-afnter-n)
[func SplitN(s, sep []byte, n int) [][]byte](#split-n)
[func Title(s []byte) []byte](#title) 废弃
[func ToLower(s []byte) []byte](#to-lower)
[func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte](#to-lower-special)
[func ToTitle(s []byte) []byte](#to-title)
[func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte](#to-title-special)
[func ToUpper(s []byte) []byte](#to-wupper)
[func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte](#to-upper-special)
[func ToValidUTF8(s, replacement []byte) []byte](#to-valid-utf8)
[func Trim(s []byte, cutset string) []byte](#trim)
[func TrimFunc(s []byte, f func(r rune) bool) []byte](#trim-func)
[func TrimLeft(s []byte, cutset string) []byte](#trim-left)
[func TrimLeftFunc(s []byte, f func(r rune) bool) []byte](#trim-left-func)
[func TrimPrefix(s, prefix []byte) []byte](#trim-prefix)
[func TrimRight(s []byte, cutset string) []byte](#trim-right)
[func TrimRightFunc(s []byte, f func(r rune) bool) []byte](#trim-right-func)
[func TrimSpace(s []byte) []byte](#trim-space)
[func TrimSuffix(s, suffix []byte) []byte](#trim-suffix)
[type Buffer](#type-Buffer)
    [func NewBuffer(buf []byte) *Buffer](#buffer-new-buffer)
    [func NewBufferString(s string) *Buffer](#buffer-new-buffer-string)
    [func (b *Buffer) Available() int](#buffer-available)
    [func (b *Buffer) AvailableBuffer() []byte](#buffer-available-buffer)
    [func (b *Buffer) Bytes() []byte](#buffer-bytes)
    [func (b *Buffer) Cap() int](#buffer-cap)
    [func (b *Buffer) Grow(n int)](#buffer-grow)
    [func (b *Buffer) Len() int](#buffer-Len)
    [func (b *Buffer) Next(n int) []byte](#buffer-next)
    [func (b *Buffer) Read(p []byte) (n int, err error)](#buffer-read)
    [func (b *Buffer) ReadByte() (byte, error)](#buffer-read-byte)
    [func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)](#buffer-read-bytes)
    [func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)](#buffer-read-from)
    [func (b *Buffer) ReadRune() (r rune, size int, err error)](#buffer-read-rune)
    [func (b *Buffer) ReadString(delim byte) (line string, err error)](#buffer-read-string)
    [func (b *Buffer) Reset()](#buffer-reset)
    [func (b *Buffer) String() string](#buffer-string)
    [func (b *Buffer) Truncate(n int)](#buffer-truncate)
    [func (b *Buffer) UnreadByte() error](#buffer-unread-byte)
    [func (b *Buffer) UnreadRune() error](#buffer-unread-rune)
    [func (b *Buffer) Write(p []byte) (n int, err error)](#buffer-write)
    [func (b *Buffer) WriteByte(c byte) error](#buffer-write-byte)
    [func (b *Buffer) WriteRune(r rune) (n int, err error)](#buffer-write-rune)
    [func (b *Buffer) WriteString(s string) (n int, err error)](#buffer-write-string)
    [func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)](#buffer-write-to)
[type Reader](#type-reader)
    [func NewReader(b []byte) *Reader](#reader-new-reader)
    [func (r *Reader) Len() int](#reader-len)
    [func (r *Reader) Read(b []byte) (n int, err error)](#reader-read)
    [func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)](#reader-read-at)
    [func (r *Reader) ReadByte() (byte, error)](#reader-read-byte)
    [func (r *Reader) ReadRune() (ch rune, size int, err error)](#reader-read-rune)
    [func (r *Reader) Reset(b []byte)](#reader-reset)
    [func (r *Reader) Seek(offset int64, whence int) (int64, error)](#reader-seek)
    [func (r *Reader) Size() int64](#reader-size)
    [func (r *Reader) UnreadByte() error](#reader-unread-byte)
    [func (r *Reader) UnreadRune() error](#reader-unread-rune)
    [func (r *Reader) WriteTo(w io.Writer) (n int64, err error)](#reader-write-to)

## **<a id="constants">Constants</a>**

```go
const MinRead = 512
```

MinRead 是 Buffer 传递给 Read 调用的最小切片大小。读取自。只要 Buffer 的 MinRead 字节数至少超过保存 r 内容所需的字节数，ReadFrom 就不会增长底层缓冲区。

## **<a id="variables">Variables</a>**

```go
var ErrTooLarge = errors.New("btyes.Buffer: too large")
```

如果不能分配内存来存储缓冲区中的数据，ErrTooLarge 将被传递给 panic。

## **<a id="functions">Functions</a>**

### **<a id="clone">func Clone  添加于1.20</a>**

```go
func Clone(b []byte) []byte
```

Clone 返回 b [:len(b)]的副本。结果可能有额外的未使用容量。Clone(nil)返回 null。

### **<a id="compare">func Compare</a>**

```go
func Compare(a, b []byte) int
```

Compare 返回一个整数，用于按字典顺序比较两个字节切片。如果 a == b，则结果为 0,；如果 a<b，则为-1；如果 a>b，则+1。nil 参数等价于一个空切片。

### **<a id="contains">func Contains</a>**

```go
func Contains(b, subslice []byte) bool
```

包含子切片是否再 b 内的报告。

### **<a id="contains-any">func ContainsAny  添加于1.7</a>**

```go
func ContainsAny(b []byte, chars string) bool
```

ContainsAny 报告字符串中是否有任何 UTF-8 编码的代码点在 b。

### **<a id="ContainsFunc">func ContainsFunc  添加于1.21</a>**

```go
func ContainsFunc(b []byte, f func(rune) bool) bool
```

ContainsFunc 报告 b 内的任何 UTF-8 编码的代码点 r 是否满足 f(r)。

### **<a id="containsRune">func ContainsRune  添加于1.7</a>**

```go
ContainsRune 报告符号是否包含在 UTF-8 编码的字节切片 b 中
```

**<a id="count">func Count</a>**

```go
func Count(s, sep []byte) int
```

Count 计算 s 中 sep 的非重叠实例的数量。如果 sep 是空片，Count 返回 1 + s 中 UTF-8编码的代码点数。

**<a id="cut">func Cut 添加于1.18</a>**

```go
func Cut(s, sep []byte) (before, after []byte, found bool)
```
在 sep 的第一个实例周围剪切切片，返回 sep. 之前和之后的文本。发现的结果报告了 sep 是否出现在 s 中。如果 sep 不出现在 s 中，cut 返回 s、 nil、 false。
Cut 返回原始切片的切片，而不是复制。

**<a id="cut-prefix">func CutPrefix 添加于1.20</a>**

```go
func CutPrefix(s, prefix []byte) (after []byte, found bool)
```

CutPrefix 返回没有提供前缀字节片的 s，并报告是否找到前缀。如果 s 不以前缀开头，则 CutPrefix 返回 s，false。如果前缀是空字节片，则 CutPrefix 返回 s，true。
CutPrefix 返回原始片的切片，而不是复制。

**<a id="cut-suffix">func CutSuffix 添加于1.20</a>**

```go
func CutSuffix(s, suffix []byte) (before []byte, found bool)
```

CutSuffix 返回不含提供的后缀结尾字节片段的 s，并报告是否找到了后缀。如果 s 不以 suffix 结尾，CutSuffix 会返回 s，false。如果后缀是空字节片段，CutSuffix 返回 s，true。
CutSuffix 返回原始切片 s 的切片，而不是副本。

### **<a id="equal">func Equal</a>**

```go
func Equal(a, b []byte) bool
```

Equal 报告 a 和 b 是否具有相同的长度和包含相同的字节。nil 参数等价于一个空切片。

### **<a id="equal-fold">func EqualFold</a>**

```go
func EqualFold(s, t []byte) bool
```

EqualFold 报告在简单的 Unicode 大小写折叠下，被解释为 UTF-8字符串的 s 和 t 是否相等，Unicode 大小写折叠是一种更普遍的不区分大小写的形式。

### **<a id="fields">func Fields</a>**

```go
func Fields(s []byte) [][]byte
```

Fields 将 s 解释为 UTF-8编码的代码点序列。如 unicode.IsSpace 所定义的，它将一个或多个连续的空白字符的每个实例周围的切片分割为 s，如果 s 只包含空白，则返回一个 s 的子切片切片或一个空切片。

### **<a id="fields-func">func FieldsFunc</a>**

```go
func FieldsFunc(s []byte, f func(rune) bool) [][]byte
```

FieldsFunc 将 s 解释为 UTF-8编码的代码点序列。它在满足 f (c)的每次运行的代码点 c 上分割切片 s，并返回 s 的子切片的一个切片。如果 s 中的所有代码都满足 f (c)或 len (s) = = 0，则返回一个空片。

FieldsFunc 不保证调用 f (c)的顺序，并假设 f 总是为给定的 c 返回相同的值。

### **<a id="has-prefix">func HasPrefix</a>**

```go
func HasPrefix(s, prefix []byte) bool
```

HasPrefix 测试字节片是否以前缀开头。

### **<a id="index">func Index</a>**

```go
func Index(s, sep []byte) int
```

Index 返回 s 中 sep 的第一个实例的索引，如果 s 中不存在 sep，则返回 -1。

### **<a id="index-any">func IndexAny</a>**

```go
func IndexAny(s []byte, chars string) int
```

IndexAny 将 s 解释为 UTF-8编码的 Unicode 代码点序列。它返回以字符为单位的任何 Unicode 代码点中 s 中第一个匹配项的字节索引。如果字符为空，或者没有共同的代码点，则返回 -1。

### **<a id="index-byte">func IndexByte</a>**

```go
func IndexByte(b []byte, c byte) int
```

IndexByte 返回 b 中 c 的第一个实例的索引，如果 b 中没有 c，返回 -1。

### **<a id="index-func">func IndexFunc</a>**

```go
func IndexFunc(s []byte, f func(r rune) bool) int
```

IndexFunc 将 s 解释为 UTF-8编码的代码点序列。它返回满足 f (c)的第一个 Unicode字符的 s 中的字节索引，如果没有，返回 -1。

### **<a id="index-rune">func IndexRune</a>**

```go
func IndexRune(s []byte, r rune) int
```

IndexRune 将 s 解释为 UTF-8编码的代码点序列。它返回给定符文中 s 中第一个匹配项的字节索引。如果 s 中没有符号，则返回 -1。如果 r 是 utf8.RuneError，则返回任何无效 UTF-8字节序列的第一个实例。

### **<a id="join">func Join</a>**

```go
func Join(s [][]byte, sep []byte) [] byte
```

Join 连接 s 的元素以创建一个新的字节片。分隔符 sep 放置在生成的切片中的元素之间。

### **<a id="last-index">func LastIndex</a>**

```go
func LastIndex(s, sep []byte) int
```

LastIndex 返回 s 中 sep 的最后一个实例的索引，如果 s 中没有 sep，返回 -1。

### **<a id="last-index-any">func lastIndexAny</a>**

```go
func LastIndexAny(s []byte, chars string) int
```

LastIndexAny将 s 解释为 UTF-8编码的 Unicode 代码点序列。它返回以字符为单位的任何 Unicode 代码点中最后一个出现在 s 中的字节索引。如果字符为空，或者没有共同的代码点，则返回 -1。

### **<a id="last-index-byte">func LastIndexByte 添加于1.5</a>**

```go
func LastIndexByte(s []byte, c byte) int
```

LastIndexByte 返回 s 中 c 的最后一个实例的索引，如果 s 中没有 c，返回 -1。

###  **<a id="last-index-func">func LastIndexFunc</a>**

```go
func LastIndexFunc(s []byte, f func(r rune) bool) int
```

LastIndexFunc 将 s 解释为 UTF-8编码的代码点序列。它返回满足 f (c)的最后一个 Unicode字符的 s 中的字节索引，如果没有，返回 -1。

### **<a id="map">func Map</a>**

```go
func Map(mapping func(r rune) rune, s []byte) []byte
```

Map 返回字节片的一个副本，其中包含根据映射函数修改的所有字符。如果映射返回一个负值，则字符将从字节片中删除，不进行替换。S 和输出中的字符被解释为 UTF-8编码的代码点。

### **<a id="Repeat">func Repeat</a>**

```go
func Repeat(b []byte, count int) []byte
```

重复返回一个新的字节片，该片由 b 的计数副本组成。

如果 count 为负，或者(len (b) * count)溢出的结果，它会感到恐慌。

### **<a id="replace">func Replace</a>**

```go
func Replace(s, old, new []byte, n int) []byte
```

Replace 返回切片 s 的一个副本，将前 n 个不重叠的旧实例替换为新实例。如果 old 为空，则在切片的开始和每个 UTF-8序列之后匹配，最多可以得到 k + 1替换 k-rune 切片。如果 n < 0，则对替换的数量没有限制。

### **<a id="replace-all">func ReplaceAll 添加于1.12</a>**

```go
func ReplaceAll(s, old, new []byte) []byte
```

ReplaceAll 返回切片 s 的一个副本，其中旧实例的所有非重叠实例都替换为新实例。如果 old 为空，则在切片的开始和每个 UTF-8序列之后匹配，最多可以得到 k + 1替换 k-rune 切片。

### **<a id="runes">func Runes</a>**

```go
func Runes(s []byte) []rune
```

Runes 将 s 解释为 UTF-8编码的代码点序列。它返回一个等效于 s 的符号片(Unicode 代码点)。

### **<a id="split">func Split</a>**

```go
func Split(s, sep []byte) [][]byte
```

将切片 s 分割成 sep 分隔的所有子切片，并返回这些分隔符之间的子切片。如果 sep 为空，Split 将在每个 UTF-8序列之后拆分。它等效于计数为 -1的 SplitN。

若要围绕分隔符的第一个实例拆分，请参见 Cut。

### **<a id="split-after">func SplitAfter</a>**

```go
func SplitAfter(s, sep []byte) [][]byte
```

SplitAfter 在每个 sep 实例之后将 s 切片到所有子切片中，并返回这些子切片的一个切片。如果 sep 为空，SplitAfter 将在每个 UTF-8序列之后拆分。它等效于计数为 -1的 SplitAfterN。

### **<a id="split-after-n">func SplitAfterN</a>**

```go
func SplitAfterN(s, sep []byte, n int) [][]byte
```

SplitAfterN 在 sep 的每个实例之后将 s 切片成子切片，并返回这些子切片的一个切片。如果 sep 为空，则 SplitAfterN 在每个 UTF-8序列之后分裂。计数器确定要返回的子切片数:

```go
n > 0: 最多有 n 个子切片; 最后一个子切片将是未拆分的余数。
n == 0: 结果是 nil (零子切片)
n < 0: 所有子切片
```

### **<a id="split-n">func SplitN</a>**

```go
func SplitN(s, sep []byte, n int) [][]byte
```

SplitN 将 s 分成由 sep 分隔的切片，并返回这些分隔符之间的切片的一个片。如果 sep 为空，则 SplitN 在每个 UTF-8序列之后分裂。计数器确定要返回的子切片数:

```go
N > 0: 最多 n 个切片; 最后一个切片将是未拆分的余数。
N = = 0: 结果是 nil (零子切片)
N < 0: 所有子切片
```

### **<a id="title">func Title 弃用</a>**

### **<a id="to-lower">func ToLower</a>**

```go
func ToLower(s []byte) []byte
```

ToLower 返回字节片 s 的一个副本，其中所有 Unicode 字母都映射到它们的小写。

### **<a id="to-lower-special">func ToLowerSpecial</a>**

```go
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
```

ToLowerSpecial 将 s 视为 UTF-8编码的字节，并返回一个包含所有映射到小写的 Unicode 字母的副本，优先考虑特殊的大小写规则。

### **<a id="to-title">func ToTitle</a>**

```go
func ToTitle(s []byte) []byte
```

ToTitle 将 s 视为 UTF-8编码的字节，并返回一个包含映射到它们的 title 大小写的所有 Unicode 字母的副本。

### **<a id="to-title-special">func ToTitleSpecial</a>**

```go
func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte
```

ToTitleSpecial 将 s 视为 UTF-8编码的字节，并返回一个包含映射到它们的 title case 的所有 Unicode 字母的副本，优先考虑特殊的大小写规则。

### **<a id="to-upper">func ToUpper</a>**

```go
func ToUpper(s []byte) []byte
```

ToUpper 返回字节片 s 的一个副本，其中所有 Unicode 字母都映射到它们的大写。

### **<a id="to-upper-special">func ToUpperSpecial</a>**

```go
func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte
```

ToUpperSpecial 将 s 视为 UTF-8编码的字节，并返回一个包含映射到大写的所有 Unicode 字母的副本，优先考虑特殊的大小写规则。

### **<a id="to-valid-utf8">func ToValidUTF8 添加于1.13</a>**

```go
func ToValidUTF8(s, replacement []byte) []byte
```

ToValidUTF8将 s 视为 UTF-8编码的字节，并返回一个副本，其中每次运行的字节代表无效的 UTF-8，替换的字节可能为空。

### **<a id="trim">func Trim</a>**

```go
func Trim(s []byte, cutset string) []byte
```

Trim 通过切除 cutset 中包含的所有前导和尾随 UTF-8编码的代码点来返回 s 的子切片。

### **<a id="trim-func">func TrimFunc</a>**

```go
func TrimFunc(s []byte, f func(r rune) bool) []byte
```

TrimFunc 通过切除满足 f (c)的所有前导和尾随 UTF-8编码的代码点 c 来返回 s 的切片。

### **<a id="trim-left">func TrimLeft</a>**

```go
func TrimLeft(s []byte, cutset string) []byte
```

TrimLeft 通过切除 cutset 中包含的所有前导 UTF-8编码的代码点来返回 s 的切片。

### **<a id="trim-left-func">func TrimLeftFunc</a>**

```go
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte
```

TrimLeftFunc 将 s 视为 UTF-8编码的字节，并通过切除满足 f (c)的所有前导 UTF-8编码的代码点 c 来返回 s 的切片。

### **<a id="trim-prefix">func TrimPrefix 添加于1.1</a>**

```go
func TrimPrefix(s, prefix []byte) []byte
```

TrimPrefix 返回没有提供前缀字符串的 s。如果 s 不以前缀开头，则不变地返回 s。

### **<a id="trim-right">func TrimRight</a>**

```go
func TrimRight(s []byte, cutset string) []byte
```

TrimRight 通过切除包含在 cutset 中的所有后面 UTF-8编码的代码点来返回 s 的切片。

### **<a id="trim-right-func">func TrimRightFunc</a>**

```go
func TrimRightFunc(s []byte, f func(r rune) bool) []byte
```

TrimRightFunc 通过切除满足 f (c)的所有后面 UTF-8编码的代码点 c 来返回 s 的切片。

### **<a id="trim-space">func TrimSpace</a>**

```go
func TrimSpace(s []byte) []byte
```

TrimSpace 通过切除所有由 Unicode 定义的前导和尾随空白来返回一个 s 的切片。

### **<a id="trim-suffix">func TrimSuffix 添加于1.1</a>**

```go
func TrimSpace(s []byte) []byte
```

TrimAffix 返回没有提供后缀字符串的 s。如果 s 没有以后缀结束，则不变地返回 s。

## Types

### **<a id="type-buffer">type Buffer</a>**

```go
type Buffer struct {
  // 包含经过筛选或未导出的字段
}
```

Buffer 是具有读写方法的可变大小的字节缓冲区。Buffer 的零值是一个准备使用的空缓冲区。

### **<a id="buffer-new-buffer">func NewBuffer</a>**

```go
func NewBuffer(buf []byte) *Buffer
```

NewBuffer 创建并初始化一个新的 Buffer，使用 buf 作为其初始内容。新的 Buffer 拥有 buf 的所有权，调用者在调用之后不应该使用 buf。NewBuffer 用于准备一个用于读取现有数据的 Buffer。它还可以用来设置内部缓冲区的初始大小，以便进行写操作。为此，buf 应该具有所需的容量，但长度为零。

在大多数情况下，new (Buffer)(或者只是声明一个 Buffer 变量)足以初始化 Buffer。

### **<a id="buffer-new-buffer-string">func NewBufferString</a>**

```go
func NewBufferString(s string) *Buffer
```

NewBufferString 创建并初始化一个新的 Buffer，使用字符串 s 作为其初始内容。它的目的是准备一个缓冲区来读取现有的字符串。

在大多数情况下，new (Buffer)(或者只是声明一个 Buffer 变量)足以初始化 Buffer。

### **<a id="buffer-available">func (*Buffer) Available 添加于1.21</a>**

```go
func (b *Buffer) Available() int
```

返回缓冲区中未使用的字节数。

### **<a id="buffer-available-buffer">func (*Buffer) AvailableBuffer 添加于1.21</a>**

```go
func (b *Buffer) AvailableBuffer() []byte
```

AvailableBuffer 返回具有 b.Available() 容量的空缓冲区。此缓冲区将被追加到并传递给紧接着的 Write 调用。缓冲区只有在 b 上的下一个写操作之前有效。

### **<a id="buffer-bytes">func (*Buffer) Bytes</a>**

```go
func (b *Buffer) Bytes() []byte
```

字节返回一个长度为 b.Len ()的片段，该片段保存缓冲区的未读部分。该切片仅在下一次缓冲区修改之前有效(也就是说，只在下一次调用 Read、 Write、 Reset 或 Truncate 等方法之前有效)。该切片至少在下一次缓冲区修改之前对缓冲区内容使用别名，因此对该切片的即时更改将影响未来读取的结果。

### **<a id="buffer-cap">func (*Buffer) Cap 添加于1.5</a>**

```go
func (b *Buffer) Cap() int
```

Cap 返回缓冲区底层字节片的容量，即分配给缓冲区数据的总空间。

### **<a id="buffer-grow">func (*Buffer) Grow 添加于1.1</a>**

```go
func (b *Buffer) Grow(n int)
```

必要时，Grow 会增加缓冲区的容量，以保证能再容纳 n 个字节。在 Grow(n) 之后，至少有 n 个字节可以写入缓冲区而无需再次分配。如果 n 为负数，Grow 就会崩溃。如果缓冲区无法增长，则会出现 ErrTooLarge 异常。

### **<a id="buffer-len">func (*Buffer) Len</a>**

```go
func (b *Buffer) Len() int
```

Len 返回缓冲区未读部分的字节数; b.Len () == Len(b. Bytes ())。

### **<a id="buffer-next">func (*Buffer) Next</a>**

```go
func (b *Buffer) Next(n int) []byte
```

Next 返回一个片段，其中包含缓冲区中下一个 n 字节，并将缓冲区向前推进，就像读取已返回的字节一样。如果缓冲区中的字节少于 n 个，Next 会返回整个缓冲区。该分片仅在下一次调用读或写方法之前有效。

### **<a id="buffer-read">func (*Buffer) Read</a>**

```go
func (b *Buffer) Read(p []byte) (n int, err error)
```

读取从缓冲区读取下一个 len(p) 字节或直到缓冲区耗尽。返回值 n 是读取的字节数。如果缓冲区没有要返回的数据，err 就是 io.EOF（除非 len(p) 为零）；否则就是 nil。

### **<a id="buffer-read-byte">func (*Buffer) ReadByte</a>**

```go
func (b *Buffer) ReadByte() (byte, error)
```

ReadByte 从缓冲区读取并返回下一个字节。如果没有可用的字节，它返回错误 io.EOF。

### **<a id="buffer-read-bytes">func (*Buffer) ReadBytes</a>**

```go
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
```

ReadBytes 一直读取到输入中第一次出现 delim 时为止，并返回一个包含分隔符（含分隔符）之前数据的片段。如果 ReadBytes 在找到分隔符前遇到错误，它会返回错误前读取的数据和错误本身（通常为 io.EOF）。只有当返回的数据不是以 delim 结束时，ReadBytes 才会返回 err != nil。

### **<a id="buffer-read-from">func (*Buffer) ReadFrom</a>**

```go
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
```

ReadFrom 从 r 读取数据，直到 EOF，然后将其添加到缓冲区，并根据需要扩大缓冲区。返回值 n 是读取的字节数。读取过程中遇到的除 io.EOF 以外的任何错误也会返回。如果缓冲区变得过大，ReadFrom 将以 ErrTooLarge 引起恐慌。

### **<a id="buffer-read-rune">func (*Buffer) ReadRune</a>**

```go
func (b *Buffer) ReadRune() (r rune, size int, err error)
```

ReadRune 从缓冲区读取并返回下一个 UTF-8 编码的 Unicode 代码点。如果没有可用字节，则返回错误信息 io.EOF。如果字节是错误的 UTF-8 编码，则会消耗一个字节并返回 U+FFFD, 1。

### **<a id="buffer-read-string">func (*Buffer) ReadString</a>**

```go
func (b *Buffer) ReadString(delim byte) (line string, err error)
```

### **<a id="buffer-reset">func (*Buffer) Reset</a>**

```go
func (b *Buffer) Reset()
```

重置会将缓冲区重置为空，但会保留底层存储空间，供将来写入时使用。重置与 Truncate(0) 相同。

### **<a id="buffer-string">func (*Buffer) String</a>**

```go
func (b *Buffer) String() string
```

String 以字符串形式返回缓冲区未读部分的内容。如果缓冲区是一个 nil 指针，则返回"<nil>"。

要更有效地构建字符串，请参阅 strings.Builder 类型。

### **<a id="buffer-truncate">func (*Buffer) Truncate</a>**

```go
func (b *Buffer) Truncate(n int)
```

Truncate 会丢弃缓冲区中除前 n 个未读字节之外的所有字节，但会继续使用已分配的存储空间。如果 n 为负数或大于缓冲区的长度，它就会崩溃。

### **<a id="buffer-unread-byte">func (*Buffer) UnreadByte</a>**

```go
func (b *Buffer) UnreadByte() error
```

UnreadByte 读取最近一次成功读取操作（至少读取一个字节）返回的最后一个字节。如果上次读取后发生了写操作，如果上次读取返回错误，或者如果读取的字节为零，UnreadByte 将返回错误信息。

### **<a id="buffer-unread-rune">func (*Buffer) UnreadRune</a>**

```go
func (b *Buffer) UnreadRune() error
```

UnreadRune 会读取 ReadRune 返回的最后一个符文。如果最近一次对缓冲区的读取或写入操作不是成功的 ReadRune，UnreadRune 将返回错误信息。(在这一点上，UnreadRune 比 UnreadByte 更严格，后者会从任何读取操作中解读最后一个字节）。

### **<a id="buffer-write">func (*Buffer) Write</a>**

```go
func (b *Buffer) Write(p []byte) (n int, err error)
```

Write 将 p 的内容追加到缓冲区，并根据需要扩大缓冲区。返回值 n 是 p 的长度；err 始终为零。如果缓冲区变得过大，Write 会以 ErrTooLarge 引起恐慌。

### **<a id="buffer-write-byte">func (*Buffer) WriteByte</a>**

```go
func (b *Buffer) WriteByte(c byte) error
```

WriteByte 将字节 c 附加到缓冲区，并根据需要扩大缓冲区。返回的错误总是 nil，但会包含在缓冲区中以匹配 bufio.Writer 的 WriteByte。如果缓冲区变得过大，WriteByte 会出现 ErrTooLarge 异常。

### **<a id="buffer-write-rune">func (*Buffer) WriteRune</a>**

```go
func (b *Buffer) WriteRune(r rune) (n int, err error)
```

WriteRune 会将 Unicode 代码点 r 的 UTF-8 编码追加到缓冲区，并返回其长度和错误信息，错误信息始终为零，但会包含在缓冲区中以匹配 bufio.Writer 的 WriteRune。缓冲区会根据需要不断扩大；如果缓冲区过大，WriteRune 会以 ErrTooLarge 引起恐慌。

### **<a id="buffer-write-string">func (*Buffer) WriteString</a>**

```go
func (b *Buffer) WriteString(s string) (n int, err error)
```

WriteString 将 s 的内容追加到缓冲区，并根据需要扩大缓冲区。返回值 n 是 s 的长度；err 始终为零。如果缓冲区过大，WriteString 会以 ErrTooLarge 引起恐慌。

### **<a id="buffer-write-to">func (*Buffer) WriteTo</a>**

```go
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)
```

WriteTo 会向 w 写入数据，直到缓冲区耗尽或发生错误。返回值 n 是写入的字节数；它总是适合一个 int，但为了与 io.WriterTo 接口匹配，它是 int64。写入过程中遇到的任何错误也会返回。

## **<a id="type-reader">type Reader</a>**

```go
type Reader struct {
  // 包含已筛选或未导出字段
}
```

阅读器通过从字节片中读取数据来实现 io.Reader、io.ReaderAt、io.WriterTo、io.Seeker、io.ByteScanner 和 io.RuneScanner 接口。与缓冲区不同，读取器是只读的，并且支持寻道。读取器的零值操作类似于读取空片段。

### **<a id="reader-new-reader">func NewReader</a>**

```go
func NewReader(b []byte) *Reader
```

NewReader 返回一个从 b 读取数据的新阅读器。

### **<a id="reader-len">func (*Reader) Len</a>**

```go
func (r *Reader) Len() int
```

Len 返回片段未读部分的字节数。

### **<a id="reader-read">func (*Reader) Read</a>**

```go
func (r *Reader) Read(b []byte) (n int, err error)
```

读取实现了 io.Reader 接口。

### **<a id="reader-read-at">func (*Reader) ReadAt</a>**

```go
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
```

ReadAt 实现了 io.ReaderAt 接口。

### **<a id="reader-read-byte">func (*Reader) ReadByte</a>**

```go
func (r *Reader) ReadByte() (byte, error)
```

ReadByte 实现了 io.ByteReader 接口。

### **<a id="reader-read-rune">func (*Reader) ReadRune</a>**

```go
func (r *Reader) ReadRune() (ch rune, size int, err error)
```

ReadRune 实现了 io.RuneReader 接口。

### **<a id="reader-reset">func (*Reader) Reset 添加于1.7</a>**

```go
func (r *Reader) Reset(b []byte)
```

重置读取器，从 b 读取数据。

### **<a id="reader-seek">func (*Reader) Seek</a>**

```go
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```

Seek 实现了 io.Seeker 接口。

### **<a id="reader-size">func (*Reader) Size 添加于1.5</a>**

```go
func (r *Reader) Size() int64
```

Size 返回底层字节片的原始长度。Size 是可通过 ReadAt 读取的字节数。除重置外，该结果不受任何方法调用的影响。

### **<a id="reader-unread-byte">func (*Reader) UnreadByte</a>**

```go
func (r *Reader) UnreadByte() error
```

在实现 io.ByteScanner 接口时，UnreadByte 与 ReadByte 互为补充。

### **<a id="reader-unread-rune">func (*Reader) UnreadRune</a>**

```go
func (r *Reader) UnreadRune() error
```

UnreadRune 与 ReadRune 在实现 io.RuneScanner 接口方面互为补充。

### **<a id="reader-write-to">func (*Reader) WriteTo 添加于1.1</a>**

```go
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```

WriteTo 实现了 io.WriterTo 接口。