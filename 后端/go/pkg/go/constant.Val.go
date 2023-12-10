package main

import (
	"fmt"
	"go/constant"
	"math"
)

func main() {
	maxint := constant.MakeInt64(math.MaxInt64)
	fmt.Printf("%v\n", constant.Val(maxint)) // 9223372036854775807

	e := constant.MakeFloat64(math.E)
	fmt.Printf("%v\n", constant.Val(e)) // 6121026514868073/2251799813685248

	b := constant.MakeBool(true)
	fmt.Printf("%v\n", constant.Val(b)) // true

	b = constant.Make(false)
	fmt.Printf("%v\n", constant.Val(b)) // false

}
