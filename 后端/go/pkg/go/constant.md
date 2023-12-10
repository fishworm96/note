## package constant

包 constant 实现了代表非类型 Go 常量及其相应操作的值。

当一个值由于错误而未知时，可以使用一个特殊的未知值。除非另有说明，否则对未知值的操作会产生未知值。

## Index

### func BitLen(x Value) int

BitLen 返回用二进制表示绝对值 x 所需的位数；x 必须是 Int 或未知数。如果 x 是未知数，结果为 0。

### func BoolVal(x Value) bool

BoolVal 返回 x 的布尔值，x 必须是布尔值或未知值。如果 x 是未知数，结果为 false。

### func Bytes(x Value) []Byte

Bytes 返回 x 的绝对值的字节数（小二进制）；x 必须是 Int。

### func Compare(x_ Value, op token.Token, y_ Value) bool

Compare 返回 x op y 的比较结果。如果操作数之一为未知数，则结果为 false。

### func Float32Val(x Value) (float32, bool)

Float32Val 类似于 Float64Val，但对象是 float32 而不是 float64。

### func Float64Val(x Value) (float64, bool)

Float64Val 返回与 x 值最接近的 Go float64 值，以及结果是否精确；x 必须是数值或未知数，但不能是复数。如果数值太小（太接近 0），无法用 float64 表示，Float64Val 会默默地向下溢出为 0。

### func Int64Val(x Value) (int64, bool)

Int64Val 返回 x 的 Go int64 值以及结果是否精确；x 必须是 Int 或未知数。如果结果不精确，则其值未定义。如果 x 是未知数，则结果为（0，false）。

### func Sign(x Value) int

Sign 根据 x < 0、x == 0 或 x > 0 返回-1、0 或 1；x 必须是数值或未知数。对于复数值 x，如果 x == 0，则符号为 0，否则为 != 0。

### func StringVal(x Value) string

StringVal 返回 x 的 Go 字符串值，x 必须是字符串或未知字符串。如果 x 是未知数，结果就是""。

### func Uint64(x Value) (Uint64, bool)

Uint64Val 返回 x 的 Go uint64 值以及结果是否精确；x 必须是 Int 或 Unknown。如果结果不精确，则其值未定义。如果 x 是未知数，则结果为（0，false）。

### func Val(x Value) any 添加于1.13

Val 返回给定常量的基本值。由于 Val 返回的是一个接口，因此调用者有责任将结果类型断言为预期类型。可能的动态返回类型有：

```go
x Kind             type of result
-----------------------------------------
Bool               bool
String             string
Int                int64 or *big.Int
Float              *big.Float or *big.Rat
everything else    nil
```

### type Kind

```go
type Kind int
```

Kind 指定值所代表的值的类型。

```go
const (
	// 未知值
	Unknown Kind = iota

	// 非数值
	Bool
	String

	// 数值
	Int
	Float
	Complex
)
```

#### func (i Kind) String() string 添加于1.18

### type Value

```go
type Value interface {
	// Kind returns the value kind.
	Kind() Kind

	// String 返回值的简短加引号（人类可读）形式。对于数值，结果可能是一个近似值；对于字符串值，结果可能是一个缩短的字符串。对于精确表示数值的字符串，请使用 ExactString。
	String() string

	// ExactString 返回值的精确引号（人类可读）形式。如果值的类型是字符串，则使用 StringVal 获取未加引号的字符串。
	ExactString() string
	// 包含已过滤或未导出的方法
}
```

Value 表示 Go 常量的值。

#### func BinaryOp(x_ Value, op token.Token, y_ Value) Value

BinaryOp 返回二进制表达式 x op y 的结果。如果操作数之一是未知数，结果就是未知数。BinaryOp 不处理比较或移位；请使用比较或移位。

要强制对 Int 操作数进行整除，请使用 op == token.QUO_ASSIGN 代替 token.QUO；在这种情况下，结果保证是 Int。除以 0 会导致运行时恐慌。

#### func Denom(x Value) Value

Denom 返回 x 的分母；x 必须是 Int、Float 或 Unknown。如果 x 是未知数，或者它太大或太小而无法表示为分数，则结果为未知数。否则，结果为 Int >=1。

#### func Imag(x Value) Value

Imag 返回 x 的虚部，必须是数值或未知数。如果 x 为未知数，则结果为未知数。

#### func Make(x any) Value 添加于1.13

返回 x 的值。

```go
type of x        result Kind
----------------------------
bool             Bool
string           String
int64            Int
*big.Int         Int
*big.Float       Float
*big.Rat         Float
anything else    Unknown
```

#### func MakeBool(b bool) Value

MakeBool 返回 b 的 Bool 值。

#### func MakeFloat64(x float64) Value

MakeFloat64 返回 x 的浮点数值。如果 x 为 -0.0，结果为 0.0。如果 x 不是有限值，则结果为未知数。

#### func MakeFromBytes(bytes []bytes) Value

MakeFromBytes 返回 Int 值的小二进制字节数。空字节片参数表示 0。

#### func MakeFromLiteral(lit string, tok token.Token, zero uint) Value

MakeFromLiteral 返回 Go 字面字符串的相应整数、浮点、虚数、字符或字符串值。tok 值必须是 token.INT、token.FLOAT、token.IMAG、token.CHAR 或 token.STRING 中的一个。最后一个参数必须为零。如果字面字符串语法无效，结果将是未知。

#### func MakeImag(x Value) Value

MakeImag 返回复数值 x*i；x 必须是 Int、Float 或 Unknown。如果 x 为未知数，则结果为未知数。

#### func MakeInt64(x int64) Value

MakeInt64 返回 x 的 Int 值。

#### func MakeString(s string) Value

MakeString 返回 s 的字符串值。

#### func MakeUint64(x Uint64) Value

MakeUint64 返回 x 的 Int 值。

#### func MakeUnknown() Value

MakeUnknown 返回未知值。

#### func Num(x Value) Value

Num 返回 x 的分子；x 必须是 Int、Float 或 Unknown。如果 x 是未知数，或者太大或太小而无法表示为分数，则结果为未知数。否则，结果将是一个与 x 符号相同的 Int。

#### func Real(x Value) Value

Real 返回 x 的实部，必须是数值或未知数。如果 x 是未知数，结果就是未知数。

#### func Shift(x Value, op token.Token, s uint) Value

Shift 返回 op == token.SHL 或 token.SHR (<< 或 >>)的移位表达式 x op s 的结果。如果 x 是未知数，结果就是 x。

#### func ToComplex(x Value) Value 添加于1.6

如果 x 可以用复数表示，则 ToComplex 会将 x 转换为复数值。否则返回未知数。

#### func ToFloat(x Value) Value 添加于1.6

如果 x 可以用浮点表示，ToFloat 会将 x 转换为浮点值。否则返回未知数。

#### func ToInt(x Value) Value 添加于1.6

如果 x 可以用 Int 表示，则 ToInt 会将 x 转换为 Int 值。否则返回未知数。

#### func UnaryOp(op token.Token, y Value, prec uint) Value

UnaryOp 返回一元表达式 op y 的结果。如果 prec > 0，则指定以位为单位的 ^ (xor) 结果大小。如果 y 为未知数，则结果为未知数。