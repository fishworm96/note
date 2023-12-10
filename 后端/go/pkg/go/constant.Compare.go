package main

import (
	"fmt"
	"go/constant"
	"go/token"
	"sort"
)

func main() {
	vs := []constant.Value{
		constant.MakeString("Z"),
		constant.MakeString("bacon"),
		constant.MakeString("go"),
		constant.MakeString("Frame"),
		constant.MakeString("defer"),
		constant.MakeFromLiteral(`"a"`, token.STRING, 0),
	}

	sort.Slice(vs, func(i, j int) bool {
		// Equivalent to vs[i] <= vs[j].
		return constant.Compare(vs[i], token.LEQ, vs[j])
	})

	for _, v := range vs {
		fmt.Println(constant.StringVal(v))
		// Frame
		// Z
		// a
		// bacon
		// defer
		// go
	}

}
