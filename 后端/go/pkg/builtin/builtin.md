## package builtin

内置包提供了Go语言预先声明的标识符的文档。这里记录的项目实际上并不是内置的包，但它们的描述允许godoc为语言的特殊标识符提供文档。

## Index

[Constants](#constants)
[Variables](#variables)
[func append(slice []Type, elems ...Type) [Type]](#append)
[func cap(v Type) int](#cap)
[func clear(t T)](#clear)
[func close(c chan<- Type)](#close)
[func complex(r, i FloatType) ComplexType](#complex)
[func copy(dst, src []Type) int](#copy)
[func delete(m map[Type]Type1, key Type)](#delete)
[func imag(c ComplexType) FloatType](#imag)
[func len(v Type) int](#len)
[func make(t Type, size ...IntegerType) Type](#make)
[func max(x T, Y ...T) T](#max)
[func min(x T, y ...T) T](#min)
[func new(Type) *Type](#new)
[func panic(v any)](#panic)
[func print(args ...Type)](#print)
[func println(args ...Type)](#println)
[func real(c ComplexType) FloatType](#real)
[func recover() any](#recover)
[type ComplexType](#ComplexType)
[type FloatType](#float-type)
[type IntegerType](#integer-type)
[type Type](#type)
[type Type1](#type1)
[type any](#any)
[type bool](#bool)
[type byte](#byte)
[type comparable](#comparable)
[type complex128](#complex128)
[type complex64](#complex64)
[type error](#error)
[type float32](#float32)
[type float64](#float64)
[type int](#int)
[type int16](#int16)
[type int32](#int32)
[type int64](#int64)
[type int8](#int8)
[type rune](#rune)
[type string](#string)
[type uint](#uint)
[type uint16](#uint16)
[type uint32](#uint32)
[type uint8](#uint8)
[type uintptr](#uinptr)

## **<a id="constants">Constants</a>**

```go
const (
  true = 0 == 0 // 无类型的布尔值
  false = 0 != 0 // 无类型的布尔值
)
```

true和false是两个非类型化的布尔值。

```go
const iota = 0 // 无类型整数值
```

iota是一个预先声明的标识符，表示（通常带括号的）const声明中当前const规范的无类型整数序数。它是零索引。

## **<a id="variables">Variables</a>**

```go
var nil Type // 类型必须是指针、通道、函数、接口、映射或片类型
```

nil是预先声明的标识符，表示指针、通道、func、接口、map或slice类型的零值。

## Functions

### **<a id="append">func append</a>**

```go
func append(slice []Type, elems ...Type) []Type
```

append内置函数将元素追加到切片的末尾。如果它有足够的容量，则对目的地重新切片以容纳新元素。如果没有，则将分配新的底层阵列。Append返回更新后的切片。因此，有必要存储append的结果，通常在保存切片本身的变量中：

```go
slice = append(slice, elem1, elem2)
slice = append(slice, anotherSlice...)
```

作为特殊情况，将字符串附加到字节切片是法律的的，如下所示：

```go
slice = append([]byte("hello"), "world"...)
```

### **<a id="cap">func cap</a>**

```go
func cap(v Type) int
```

cap内置函数根据类型返回v的容量：

```go
数组：v 中元素的数量，与 len(v) 相同
数组指针：*v 中元素的数量，与len(v) 相同
切片：切片的容量（底层数组的长度）；若 v为nil，cap(v) 即为零
如果 v 为 nil，则 cap(v) 为 0
信道：按照元素的单元，相应信道缓存的容量；若v为nil，cap(v)即为零
如果 v 为 nil，则 cap(v) 为 0
```

对于某些参数（如简单数组表达式），结果可以是常量。有关详细信息，请参阅Go语言规范的“长度和容量”部分。

### **<a id="func clear">clear</a>**

```go
func clear[T ~[]Type | ~map[Type]Type1](t T)
```

clear 内置函数用于清除字典和切片。对于字典，清除将删除所有条目，从而生成空字典。对于切片，clear 将切片长度以内的所有元素设置为相应元素类型的零值。如果参数类型是类型参数，则类型参数的类型集必须仅包含 map 或 slice 类型，clear 执行类型参数所隐含的操作。

### **<a id="close">func close</a>**

```go
func close(c chan<- Type)
```

close内置函数关闭通道，该通道必须是双向的或只发送的。它应该只由发送方执行，而不是由接收方执行，并且具有在接收到最后一个发送值后关闭通道的效果。在从关闭的通道c接收到最后一个值之后，来自c的任何接收都将成功而不阻塞，从而为通道元素返回零值。形式

```go
x, ok := <-c
```

对于关闭和空的通道，也将 OK 作为假值

### **<a id="complex">func complex</a>**

```go
func complex(r, i FloatType) ComplexType
```

复杂内置函数从两个浮点值构造一个复杂值。真实的部和虚部必须大小相同，float 32 或 float 64（或可分配给它们），返回值将是相应的复数类型（float 32为complex 64，float 64为complex 128）。

### **<a id="copy">func copy</a>**

```go
func copy(dst, src []Type) int
```

copy 内置函数将元素从源切片复制到目标切片。(As 在一个特殊的情况下，它也会将字节从字符串复制到字节片。）源和目的地可以重叠。Copy 返回复制的元素数，它将是 len（src）和 len（dst）中的最小值。

### **<a id="delete">func delete</a>**

```go
func delete(m map[Type]Type1, key Type)
```

delete 内置函数从 map 中删除具有指定键（m[key]）的元素。如果 m 是 nil 或者没有这样的元素，delete 是一个 no-op。

### **<a id="imag">func imag</a>**

```go
func imag(c ComplexType) FloatType
```

imag 内置函数返回复数 c 的虚部返回值将是对应于 c 的类型的浮点类型。

### **<a id="len">func len</a>**

```go
func len(v Type) int
```

len 内置函数根据类型返回v的长度：

```go
数组：v 中元素的数量
数组指针：*v 中元素的数量（v为nil时panic）
切片、映射：v 中元素的数量；若 v 为 nil，len(v)即为零
字符串：v中字节的数量
通道：通道缓存中队列（未读取）元素的数量；若v为 nil，len(v)即为零
```

### **<a id="make">func make</a>**

```go
func make(t Type, size ...IntegerType) Type
```

make 内置函数分配并初始化 slice、map 或 chan（仅限）类型的对象。与 new 一样，第一个参数是类型，而不是值。与 new 不同的是，make 的返回类型与其参数的类型相同，而不是指向它的指针。结果的规格取决于类型：

```go
切片：size 指定了其长度。该切片的容量等于其长度。切片支持第二个整数实参可用来指定不同的容量；
     它必须不小于其长度，因此 make([]int, 0, 10) 会分配一个长度为 0，容量为 10 的切片。
映射：初始分配的创建取决于 size，但产生的映射长度为 0。size 可以省略，这种情况下就会分配一个
     小的起始大小。
通道：通道的缓存根据指定的缓存容量初始化。若 size 为零或被省略，该信道即为无缓存的。
```

### **<a id="max">func max</a>**

```go
func max[T cmp.Ordered](x T, y ...T) T
```

max 内置函数返回 cmp.Ordered 类型的固定数量参数的最大值。必须至少有一个参数。如果T是浮点类型，并且任何参数都是 NaN，则 max 将返回 NaN。

### **<a id="min">func min</a>**

```go
func min[T cmp.Ordered](x T, y ...T) T
```

min 内置函数返回 cmp.Ordered 类型的固定数量参数中的最小值。必须至少有一个参数。如果T是浮点类型，并且任何参数都是 NaN，则 min 将返回 NaN。

### **<a id="new">func new</a>**

```go
func new(Type) *Type
```

新的内置函数分配内存。第一个参数是一个类型，而不是一个值，返回的值是一个指向该类型的新分配的零值的指针。

### **<a id="panic">func panic</a>**

```go
func panic(v any)
```

panic 内置函数停止当前 goroutine 的正常执行。当函数 F 调用 panic 时，F 的正常执行立即停止。任何被F延迟执行的函数都以通常的方式运行，然后 F 返回给它的调用者。对于调用者 G 来说，调用 F 的行为就像调用 panic，终止 G 的执行并运行任何延迟的函数。这将继续下去，直到执行 goroutine 中的所有函数都停止，顺序相反。在这一点上，程序终止与非零退出代码。这种终止序列称为 panicking，可以通过内置函数 recover 来控制。

从 Go 1.21 开始，使用 nil 接口值或无类型 nil 调用 panic 会导致运行时错误（另一种panic）。GODEBUG 设置 panicnil=1 禁用运行时错误。

### **<a id="print">func print</a>**

```go
func print(args ...Type)
```

print 内置函数以特定于实现的方式格式化其参数，并将结果写入标准错误。Print 对于引导和调试很有用;它不能保证保持在语言中。

### **<a id="println">func println</a>**

```go
func println(args ...Type)
```

println内置函数以特定于实现的方式格式化其参数，并将结果写入标准错误。参数之间总是添加空格，并追加一个换行符。Println对于引导和调试很有用;它不能保证保持语言。

### **<a id="real">func real</a>**

```go
func real(c ComplexType) FloatType
```

真实的内置函数返回复数 c 的实部。返回值将是对应于 c 的类型的浮点类型。

### **<a id="recover">func recover</a>**

```go
func recover() any
```

recover 内置函数允许程序管理 panicking goroutine 的行为。在延迟函数（但不是它调用的任何函数）内执行恢复调用，通过恢复正常执行来停止 panicking 序列，并检索传递给 panic 调用的错误值。如果在 deferred 函数之外调用 recover，它将不会停止恐慌序列。在这种情况下，或者当 goroutine 没有 panicking 时，或者如果提供给 panic 的参数是 nil，recover 返回nil。因此，recover 的返回值报告了 goroutine 是否出现了 panicking。

## Types

### **<a id="complex-type">type ComplexType</a>**

```go
type ComplexType complex64
```

ComplexType 仅用于文档目的。它是两种复杂类型的替代：complex64 或 complex128。

### **<a id="float-type">type FloatType</a>**

```go
type FloatType float32
```

FloatType 仅用于文档目的。它是两种 float 类型的替代：float32 或 float64。

### **<a id="integer-type">type IntegerType</a>**

```go
type IntegerType int
```

IntegerType 仅用于文档母的。它是任何整数类型的替代；int、uint、int8 等。

### **<a id="type">type Type</a>**

```go
type Type int
```

类型仅用于文档编制。它是任何 Go 类型的替身，但表示任何给定函数调用的想同类型。

### **<a id="type1">type Type1</a>**

```go
type Type1 int
```

Type1 仅用于文档目的。它是任何 Go 类型的替身，但表示任何给定函数调用的相同类型。

### **<a id="any">type any</a>**

```go
type any = interface{}
```

any 是 interface{} 的别名，在任何方面都等效于 interface{}。

### **<a id="byte">type byte</a>**

```go
type byte = uint8
```

byte 是 uint8 的别名，在所有方面都等同于 uint8。按照惯例，它用于区分字节值和 8 位无符号整数值。

### **<a id="comparable">type comparable</a>**

```go
type comparable interface{ comparable }
```

comparable 是一个由所有可比较类型（布尔值、数字、字符串、指针、通道、可比较类型的数组、字段都是可比较类型的结构）实现的接口。可比较接口只能用作类型参数约束，而不能用作变量的类型。

### **<a id="complex128">type complex128</a>**

```go
type complex128 complex128
```

complex128 是具有 float64 实数部分和虚数部分的所有复数的集合。

### **<a id="complex64">type complex64</a>**

```go
type complex64 complex64
```

Complex64 是具有 float32 实数部分和虚数部分的所有复数的集合。

### **<a id="error">type error</a>**

```go
type error interface {
  Error() string
}
```

错误内置接口类型是表示错误条件的常规接口，nil 值表示没有错误。

### **<a id="float32">type float32</a>**

```go
type float32 float32
```

float32 是所有 IEEE-754 32 位浮点数的集合。

### **<a id="float64">type float64</a>**

```go
type float64 float64
```

float 64 是所有 IEEE-754 64 位浮点数的集合。

### **<a id="int">type int</a>**

```go
type int int
```

int是一个有符号整数类型，大小至少为 32 位。然而，它是一个独特的类型，而不是例如 int32 的别名。

### **<a id="int16">type int16</a>**

```go
type int16 int16
```

int 16 是所有有符号 16 位整数的集合。范围：-32768 至 32767。

### **<a id="int32">type int32</a>**

```go
type int32 int32
```

int 32 是所有有符号 32 位整数的集合。范围：-2147483648 至 2147483647。

### **<a id="int64">type int64</a>**

```go
type int64 int64
```

int 64 是所有有符号的 64 位整数的集合。范围：-9223372036854775808 至 9223372036854775807。

### **<a id="int8">type int8</a>**

```go
type int8 int8
```

int 8 是所有有符号 8 位整数的集合。范围：-128 到 127。

### **<a id="rune">type rune</a>**

```go
type rune = int32
```

rune 是 int32 的别名，在所有方面都等同于 int32。按照惯例，它用于区分字符值和整数值。

### **<a id="string">type string</a>**

```go
type string string
```
字符串是所有 8 位字节的字符串的集合，通常但不一定表示 UTF-8 编码的文本。字符串可以为空，但不能为 nil。字符串类型的值是不可变的。

### **<a id="uint">type uint</a>**

```go
type uint uint
```

uint 是大小至少为 32 位的无符号整数类型。然而，它是一个独特的类型，而不是 uint32 的别名。

### **<a id="uint16">type uint16</a>**

```go
type uint16 uint16
```

uint16 是所有无符号 16 位整数的集合。范围：0 到 65535。

### **<a id="uint32">type uint32</a>**

```go
type uint32 uint32
```

uint32 是所有无符号 32 位整数的集合。范围：0 到 4294967295。

### **<a id="uint64">type uint64</a>**

```go
type uint64 uint64
```

uint64 是所有无符号 64 位整数的集合。范围：0 到 18446744073709551615。

### **<a id="uint8">type uint8</a>**

```go
type uint8 uint8
```

uint8 是所有无符号 8 位整数的集合。范围：0 到 255。

### **<a id="uinptr">type uintptr</a>**

```go
type uinptr uintptr
```

uintptr 是一个整数类型，它足够大，可以保存任何指针的位模式。