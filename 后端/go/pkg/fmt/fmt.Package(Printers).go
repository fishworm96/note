package main

import (
	"fmt"
	"math"
)

// Print、Println 和 Printf 的参数布局不同。在这个例子中，我们可以比较它们的行为。Println 总是在打印项之间添加空格，而 Print 只在非字符串参数之间添加空格，Printf 则完全按照指令行事。Sprint、Sprintln、Sprintf、Fprint、Fprintln 和 Fprintf 的行为与此处所示的相应 Print、Println 和 Printf 函数相同。
func main() {
	a, b := 3.0, 4.0
	h := math.Hypot(a, b)

	// 当参数都不是字符串时，Print 会在参数之间插入空格。 它不会在输出中添加换行符，因此我们要明确添加换行符。
	fmt.Print("The vector (", a, b, ") has length ", h, ".\n") // The vector (3 4) has length 5.

	// Println 总是在参数之间插入空格，因此在这种情况下不能使用它来产生与 Print 相同的输出；它的输出有额外的空格。 此外，Println 总是在输出中添加换行符。
	fmt.Println("The vector (", a, b, ") has length", h, ".") // The vector ( 3 4 ) has length 5 .

	// Printf 提供了完整的控制功能，但使用起来比较复杂。 它不会在输出中添加换行符，因此我们要在格式指定字符串的末尾明确添加换行符。
	fmt.Printf("The vector (%g %g) has length %g.\n", a, b, h) // The vector (3 4) has length 5.

}
