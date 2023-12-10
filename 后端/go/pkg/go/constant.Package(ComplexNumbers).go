package main

import (
	"fmt"
	"go/constant"
	"go/token"
)

func main() {
	// 创建复数 2.3 + 5i。
	ar := constant.MakeFloat64(2.3)
	ai := constant.MakeImag(constant.MakeInt64(5))
	a := constant.BinaryOp(ar, token.ADD, ai)

	// 计算 (2.3 + 5i) * 11。
	b := constant.MakeUint64(11)
	c := constant.BinaryOp(a, token.MUL, b)

	// 将 c 转换为复数128。
	Ar, exact := constant.Float64Val(constant.Real(c))
	if !exact {
		fmt.Printf("Could not represent real part %s exactly as float64\n", constant.Real(c)) // Could not represent real part 25.3 exactly as float64
	}
	Ai, exact := constant.Float64Val(constant.Imag(c))
	if !exact {
		fmt.Printf("Could not represent imaginary part %s as exactly as float64\n", constant.Imag(c))
	}
	C := complex(Ar, Ai)

	fmt.Println("literal", 25.3+55i) // literal (25.3+55i)
	fmt.Println("go/constant", c) // go/constant (25.3 + 55i)
	fmt.Println("complex128", C) // complex128 (25.299999999999997+55i)

}
