package main

import (
	"fmt"
	"math"
	"time"
)

// 这些示例演示了使用格式字符串进行打印的基本方法。Printf、Sprintf 和 Fprintf 都使用一个格式字符串来指定后续参数的格式。例如，%d（我们称之为 "动词"）表示打印相应的参数，该参数必须是十进制整数（或包含整数的参数，如 int 的片段）。动词 %v（'v'表示'值'）总是以默认形式格式化参数，就像 Print 或 Println 所显示的那样。特殊动词 %T（'T'表示'类型'）打印参数的类型而不是值。示例并不详尽，所有细节请参阅软件包注释。
func main() {
	// 一组基本示例，说明 %v 是默认格式，在本例中是整数的十进制，可以用 %d 明确要求；输出只是 Println 产生的结果。
	integer := 23
	// 每一个都打印 "23"（不带引号）。
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	// 特殊动词 %T 显示项目的类型而不是其值。
	fmt.Printf("%T %T\n", integer, &integer)
	// 结果: int *int

	// Println(x) 与 Printf("%v\n", x) 相同，因此在下面的示例中我们将只使用 Printf。每个示例都演示了如何格式化特定类型的值，如整数或字符串。每个格式字符串都以 %v 开头，显示默认输出，然后是一个或多个自定义格式。

	// 布尔值用 %v 或 %t 打印为 "真 "或 "假"。
	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// 结果：true true

	// 整数用 %v 和 %d 打印为小数，或用 %x 打印为十六进制，用 %o 打印为八进制，或用 %b 打印为二进制。
	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// 结果： 42 42 2a 52 101010

	// 浮点数有多种格式：%v 和 %g 打印紧凑的表示法，而 %f 打印小数点，%e 使用指数表示法。这里使用的格式 %6.2f 演示了如何设置宽度和精度，以控制浮点数值的外观。在本例中，6 是数值打印文本的总宽度（注意输出中的额外空格），2 是要显示的小数点位数。
	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// 结果： 3.141592653589793 3.141592653589793 3.141592653589793 3.14 ( 3.14) 3.141593e+00

	// 复数格式为浮点数括号对，虚部后有一个 "i"。
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// 结果： (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)

	// 符文是整数，但用 %c 打印时会显示具有该 Unicode 值的字符。如果符文是可打印的，则 %q verb 显示为带引号的字符，%U 显示为十六进制 Unicode 代码点，而 %#U 则显示为代码点和带引号的可打印形式。
	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// 结果： 128512 128512 😀 '😀'U+1F600 U+1F600 '😀'。

	// 字符串的格式为 %v 和 %s 原样，%q 为带引号字符串，%#q 为反引号字符串。
	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// 结果： foo "bar" foo "bar" "foo （"bar/"）" `foo "bar"`

	// 使用 %v 格式的映射以默认格式显示键和值。 %#v 格式（这里的 # 称为 "标志"）以 Go 源格式显示映射。地图以一致的顺序打印，按键的值排序。
	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// 结果： map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}

	// 使用 %v 格式的结构体以默认格式显示字段值。 %+v 格式按名称显示字段，而 %#v 格式则以 Go 源格式显示结构体。
	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// 结果：{Kim 22}{Name:Kim Age:22} struct { Name string; Age int }{Name: "Kim", Age:22}

	// 指针的默认格式是在底层值前面加上一个 "括 "号。%p verb 以十六进制打印指针值。我们在这里使用 "nil "作为 %p 的参数，因为任何非 "nil "指针的值都会在运行过程中发生变化；请运行注释后的 Printf 调用查看。
	pointer := &person
	fmt.Printf("%v %p\n", pointer, (*int)(nil))
	// Result: &{Kim 22} 0x0
	// fmt.Printf("%v %p\n", pointer, pointer) 
	// Result: &{Kim 22} 0x010203 // See comment above.

	// 数组和片段的格式化是通过对每个元素应用格式来实现的。
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// 结果: [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)
	// Result: [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}

	// 字节片比较特殊。像 %d 这样的整数动词以这种格式打印元素。而 %s 和 %q 形式则像处理字符串一样处理切片。%x动词有一种特殊形式，带有空格标志，可以在字节之间加上空格。
	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
	// 结果: [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98

	// 实现 Stringer 的类型的打印方式与字符串相同。因为 Stringer 返回的是字符串，所以我们可以使用字符串专用动词打印它们，例如 %q。
	now := time.Unix(123456789, 0).UTC() // time.Time implements fmt.Stringer.
	fmt.Printf("%v %q\n", now, now)
	// 结果: 1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"

}
