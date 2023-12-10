package main

import (
	"fmt"
	"go/constant"
	"go/token"
)

func main() {
	zero := constant.MakeInt64(0)
	one := constant.MakeInt64(1)
	negOne := constant.MakeInt64(-1)

	mkComplex := func(a, b constant.Value) constant.Value {
		b = constant.MakeImag(b)
		return constant.BinaryOp(a, token.ADD, b)
	}

	vs := []constant.Value{
		negOne,
		mkComplex(zero, negOne),
		mkComplex(one, negOne),
		mkComplex(negOne, one),
		mkComplex(negOne, negOne),
		zero,
		mkComplex(zero, zero),
		one,
		mkComplex(zero, one),
		mkComplex(one, one),
	}

	for _, v := range vs {
		fmt.Printf("% d %s\n", constant.Sign(v), v)
		// -1 -1
		// -1 (0 + -1i)
		// -1 (1 + -1i)
		// -1 (-1 + 1i)
		// -1 (-1 + -1i)
		//  0 0
		//  0 (0 + 0i)
		//  1 1
		//  1 (0 + 1i)
		//  1 (1 + 1i)
	}

}
