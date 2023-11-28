## package gob

gob 包管理 gobs 流--在编码器（发送器）和解码器（接收器）之间交换的二进制值。典型的用途是传输远程过程调用（RPC）的参数和结果，如 net/rpc 提供的远程过程调用。

该实现为数据流中的每种数据类型编译一个自定义编解码器，当使用单个编码器传输数值流时效率最高，从而摊销了编译成本。

### Basics

数据流是自描述的。数据流中的每个数据项前面都有一个关于其类型的说明，该说明用一小套预定义类型来表示。指针不会被传送，但其指向的内容会被传送；也就是说，其值会被扁平化。不允许使用无指针，因为它们没有值。递归类型可以正常工作，但递归值（带循环的数据）就有问题了。这种情况可能会改变。

要使用 gobs，需要创建一个编码器，并将一系列数据项作为值或地址提交给编码器，这些数据项可以被反向引用为值。编码器确保在需要之前发送所有类型信息。在接收端，解码器会从编码流中获取值，并将其解压缩到本地变量中。

### Types and Values

源值和目标值/类型不必完全一致。对于结构体，源变量中存在但接收变量中没有的字段（用名称标识）将被忽略。接收变量中的字段，如果在传输类型或值中缺失，则在目的地中将被忽略。如果两个变量中都有同名字段，它们的类型必须兼容。接收器和发送器都将进行所有必要的间接和取消引用，以便在 gobs 和实际 Go 值之间进行转换。例如，一个 gob 类型示意为

```go
struct { A, B int }
```

可以从这些围棋类型中的任何一种发送或接收：

```go
struct { A, B int } // 相同
*struct { A, B int } // 结构的额外间接性
struct { *A, **B int } // 额外的间接字段
struct { A, B int64 } // 不同的具体值类型；见下文
```

也可以接收到其中任何一种：

```go
struct { A, B int } // 相同
struct { B, A int } // 排序并不重要；通过名称匹配
struct { A, B, C int } // 忽略额外字段（C）。
struct { B int } // 忽略缺失字段（A）；数据将被丢弃
struct { B, C int } // 忽略缺失字段（A）；忽略额外字段（C）。
```

尝试接收这些类型的数据将导致解码错误：

```go
struct { A int; B uint } // 更改 B 的符号性
struct { A int; B float } // 更改 B 的类型
struct { } // 没有共同的字段名
struct { C, D int } // 没有共同的字段名称
```

整数有两种传输方式：任意精度有符号整数或任意精度无符号整数。gob 格式中没有 int8、int16 等区分，只有带符号和无符号整数。如下所述，发送方以可变长度编码发送数值；接收方接受数值并将其存储到目标变量中。浮点数始终使用 IEEE-754 64 位精度发送（见下文）。

有符号整数可被接收到任何有符号整数变量中：int、int16 等；无符号整数可被接收到任何无符号整数变量中；浮点数可被接收到任何浮点数变量中。但是，目标变量必须能够表示数值，否则解码操作将失败。

此外，还支持结构体、数组和片段。结构体仅对导出字段进行编码和解码。字符串和字节数组支持特殊的高效表示法（见下文）。解码片段时，如果现有片段有容量，则会在原处扩展片段；如果没有，则会分配一个新数组。无论如何，结果片段的长度都会报告已解码元素的数量。

一般来说，如果需要分配，解码器会分配内存。如果不需要，解码器会使用从数据流中读取的值更新目标变量。解码器不会先对它们进行初始化，因此如果目标变量是一个复合值，如 map、struct 或 slice，解码后的值将按元素顺序合并到现有变量中。

函数和通道不会以 gob 形式发送。试图在顶层对此类值进行编码将失败。chan 或 func 类型的 struct 字段将被视为未导出字段并被忽略。

Gob 可以通过调用相应的方法（按优先级顺序排列），对实现了 GobEncoder 或 encoding.BinaryMarshaler 接口的任何类型的值进行编码。

Gob 可以通过调用相应的方法来解码实现 GobDecoder 或 encoding.BinaryUnmarshaler 接口的任何类型的值，同样按照优先级顺序进行。

### Encoding Details

本节记录编码，这些细节对大多数用户来说并不重要。详细内容自下而上呈现。

无符号整数有两种发送方式。如果小于 128，则以该值的字节形式发送。否则，它将作为一个最小长度的大二进制（高字节在前）字节流发送，该字节流包含数值，前面还有一个包含字节数的字节，但已被否定。因此，0 会以 (00) 的形式传输，7 会以 (07) 的形式传输，256 会以 (FE 01 00) 的形式传输。

布尔值用无符号整数编码：0 表示假，1 表示真。

有符号整数 i 在无符号整数 u 中编码。在 u 中，第 1 位向上包含数值；第 0 位表示接收时是否要补码。编码算法如下：

```go
var u uint
if i < 0 {
	u = (^uint(i) << 1) | 1 // 补充 i，第 0 位为 1
} else {
	u = (uint(i) << 1) // 不对 i 进行补码，第 0 位为 0
}
encodeUnsigned(u)
```

因此，低位类似于符号位，但将其作为补码位，可以保证最大负整数不是特例。例如，-129=^128=(^256>>1) 编码为 (FE 01 01)。

浮点数总是以 float64 值的形式发送。该值使用 math.Float64bits 转换为 uint64。然后将 uint64 进行字节反转，并作为普通无符号整数发送。字节反转意味着指数和尾数的高精度部分先行。由于低位通常为零，这可以节省编码字节数。例如，17.0 只用三个字节编码（FE 31 40）。

字符串和字节片段以无符号计数的形式发送，后面是数值的未解释字节数。

所有其他片段和数组都以无符号计数的形式发送，然后使用其类型的标准 gob 编码递归发送相应数量的元素。

映射会以无符号计数的形式发送，然后是相同数量的键和元素对。空但非零的映射会被发送，因此如果接收方尚未分配映射，除非发送的映射为零且不在顶层，否则接收方总会分配一个映射。

在片段和数组以及映射中，即使所有元素都为零，也会传输所有元素，甚至是零值元素。

结构体以（字段编号、字段值）对序列的形式发送。字段值使用其类型的标准 gob 编码递归发送。如果某个字段的类型值为零（数组除外，见上文），则在发送时省略该字段。字段编号由编码结构体的类型决定：编码类型的第一个字段是字段 0，第二个是字段 1，等等。在对数值进行编码时，为了提高效率，字段编号会进行 delta 编码，字段总是按照字段编号递增的顺序发送；因此 deltas 是无符号的。delta 编码的初始化将字段编号设置为-1，因此值为 7 的无符号整数字段 0 将以无符号 delta = 1、无符号值 = 7 或 (01 07) 的形式传输。最后，在发送完所有字段后，一个终止标记表示结构的结束。该标记为 delta=0 值，其表示形式为 (00)。

接口类型不进行兼容性检查；在传输时，所有接口类型都被视为单一 "接口 "类型的成员，类似于 int 或 []byte - 实际上，它们都被视为 interface{}。接口值以字符串形式传输，字符串标识发送的具体类型（名称必须由调用 Register 预先定义），后面是后面数据长度的字节数（以便在无法存储时跳过该值），后面是存储在接口值中的具体（动态）值的常规编码。(接口值为 nil 时，用空字符串标识，不传输任何值）。接收时，解码器会验证解压缩的具体项目是否满足接收变量的接口。

如果传递给编码器的值的类型不是结构体（或结构体指针等），为简化处理，会将其表示为一个字段的结构体。这样做的唯一明显效果是在值后面编码一个零字节，就像在编码结构体的最后一个字段后面一样，这样解码算法就知道顶层值何时完成。

类型的表示方法如下。在编码器和解码器之间的给定连接上定义类型时，会分配一个带符号的整数类型 id。调用 Encoder.Encode(v) 时，它将确保为 v 类型及其所有元素分配了一个 id，然后发送一对（typeid、encoded-v），其中 typeid 是 v 的编码类型 id，encoded-v 是值 v 的 gob 编码。

要定义一个类型，编码器会选择一个未使用的正类型 id，然后发送一对（-type id, encoded-type），其中 encoded-type 是由这些类型构建的 wireType 描述的 gob 编码：

```go
type wireType struct {
	ArrayT           *ArrayType
	SliceT           *SliceType
	StructT          *StructType
	MapT             *MapType
	GobEncoderT      *gobEncoderType
	BinaryMarshalerT *gobEncoderType
	TextMarshalerT   *gobEncoderType

}
type arrayType struct {
	CommonType
	Elem typeId
	Len  int
}
type CommonType struct {
	Name string // 结构类型的名称
	Id  int    // 类型的 id，重复使用，使其位于类型内部
}
type sliceType struct {
	CommonType
	Elem typeId
}
type structType struct {
	CommonType
	Field []*fieldType // 结构的字段。
}
type fieldType struct {
	Name string // 字段的名称。
	Id   int    // 字段的类型 id，必须已经定义
}
type mapType struct {
	CommonType
	Key  typeId
	Elem typeId
}
type gobEncoderType struct {
	CommonType
}
```

如果存在嵌套类型 ID，则必须先定义所有内部类型 ID 的类型，然后再使用顶层类型 ID 来描述编码后的 v。

为简化设置，连接被定义为先验理解这些类型，以及基本的 gob 类型 int、uint 等。它们的 id 是

```go
bool        1
int         2
uint        3
float       4
[]byte      5
string      6
complex     7
interface   8
// 保留 ID 的间隙。
WireType    16
ArrayType   17
CommonType  18
SliceType   19
StructType  20
FieldType   21
// 22 是 fieldType 的片段。
MapType     23
```

最后，调用 Encode 创建的每条信息前面都有一个编码后的无符号整数，表示信息中剩余的字节数。在初始类型名称之后，接口值以同样的方式进行包装；实际上，接口值就像 Encode 的递归调用。

总之，一个 gob 流看起来像

```go
(byteCount (-type id, encoding of a wireType)* (type id, encoding of a value))*
```

其中 * 表示零次或多次重复，值的类型 id 必须是预定义的或在数据流中的值之前定义。

兼容性：今后对软件包的任何修改都将努力保持与使用以前版本编码的流的兼容性。也就是说，本软件包的任何已发布版本都应能解码用以前发布的任何版本写入的数据，但须考虑安全修复等问题。有关背景信息，请参阅 Go 兼容性文档：https://golang.org/doc/go1compat。

有关 gob 线格式的设计讨论，请参阅 "大量数据"：https://blog.golang.org/gobs-of-data

### Security

该软件包的设计初衷并不是为了对抗恶意输入，也不属于 https://go.dev/security/policy 的范围。特别是，解码器只对解码后的输入大小进行基本的正确性检查，而且其限制是不可配置的。在解码来自不可信来源的 gob 数据时应小心谨慎，因为这可能会消耗大量资源。

## Index

### func Register(value any)

注册器在其内部类型名称下记录一种类型，并由该类型的值标识。该名称将标识作为接口变量发送或接收的值的具体类型。只有作为接口值的实现而传输的类型才需要注册。如果类型和名称之间的映射关系不是双射关系，它就会慌乱。

### func RegisterName(name string, value any)

RegisterName 与 Register 类似，但使用的是提供的名称而不是类型的默认名称。

### type CommonType

```go
type CommonType struct {
  Name string
  Id typeId
}
```

CommonType 包含所有类型的元素。它是一个历史产物，为二进制兼容性而保留，仅为软件包的类型描述符编码而导出。客户端无法直接使用它。

### type Decoder

```go
type Decoder struct {
  // 包含已筛选或未导出字段
}
```

解码器管理从连接远程端读取的类型和数据信息的接收。解码器可安全地供多个程序并发使用。

解码器只对解码后的输入大小进行基本的正确性检查，其限制是不可配置的。在解码来自不可信来源的数据时，请务必谨慎。

#### func NewDecoder(r io.Reader) *Decoder

NewDecoder 返回一个从 io.Reader 读取数据的新解码器。如果 r 没有同时实现 io.ByteReader，它将被封装在 bufio.Reader 中。

#### func (dec *Decoder) Decoder(e any) error

解码从输入流中读取下一个值，并将其存储到空接口值所代表的数据中。如果 e 为空，该值将被丢弃。否则，e 的底层值必须是一个指针，指向接收到的下一个数据项的正确类型。如果输入已到 EOF，解码会返回 io.EOF，而不会修改 e。

#### func (dec *Decoder) DecodeValue(v reflect.Value) error

DecodeValue 从输入流中读取下一个值。如果 v 是零 reflect.Value (v.Kind() == Invalid)，DecodeValue 会丢弃该值。在这种情况下，v 必须是一个指向数据的非零指针，或者是一个可赋值的 reflect.Value (v.CanSet()) 如果输入已到 EOF，DecodeValue 返回 io.EOF，不会修改 v。

### type Encoder

```go
type Encoder struct {
  // 包含已筛选或未导出字段
}
```

编码器负责将类型和数据信息传输到连接的另一端。多个程序可以同时使用编码器。

#### func NewEncoder(w io.Writer) *Encoder

NewEncoder 返回一个新的编码器，该编码器将在 io.Writer.NX 上进行传输。

#### func (enc *Encoder) Encode(e any) error

编码器会传输空接口值所代表的数据项，确保所有必要的类型信息都已先行传输。向 Encoder 传递 nil 指针会引起恐慌，因为 gob 无法传输它们。

#### func (enc *Encoder) EncodeValue(value reflect.Value) error

EncodeValue 会传输反射值所代表的数据项，确保所有必要的类型信息都已先行传输。如果向 EncodeValue 传递一个 nil 指针，由于无法通过 gob 传输，因此会出现问题。

### type GobDecoder

```go
type GobDecoder interface {
  // GobDecode 会用 GobEncode 写入的字节片段所代表的值覆盖接收器（接收器必须是指针），通常是针对相同的具体类型。
  GobDecode([]byte) error
}
```

GobDecoder 是描述数据的接口，它为解码 GobEncoder 发送的传输值提供了自己的例程。

### type GobEncoder

```go
type GobEncoder interface {
  // GobDecode 会用 GobEncode 写入的字节片段所代表的值覆盖接收器（接收器必须是指针），通常是针对相同的具体类型。
  GobEncode() ([]byte, error)
}
```

GobEncoder 是描述数据的接口，它提供了自己的表示方法，用于将值编码后传输到 GobDecoder。实现了 GobEncoder 和 GobDecoder 的类型可以完全控制其数据的表示，因此可以包含私有字段、通道和函数等内容，这些内容通常不能在 gob 流中传输。

注意：由于 gobs 可以永久存储，因此在设计时要确保 GobEncoder 使用的编码能随着软件的发展而保持稳定。例如，GobEncode 可以在编码中包含版本号。