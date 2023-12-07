package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// 本例演示了如何检查 Go 程序的 AST。
func main() {
	// src 是我们要检查 AST 的输入。
	src := `
package p
const c = 1.0
var X = f(3.14)*2 + c
`

	// 通过解析 src 创建 AST。
	fset := token.NewFileSet() // 位置相对于 fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	// 检查 AST 并打印所有标识符和字面量。
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
		case *ast.Ident:
			s = x.Name
		}
		if s != "" {
			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
			// src.go:2:9:	p
			// src.go:3:7:	c
			// src.go:3:11:	1.0
			// src.go:4:5:	X
			// src.go:4:9:	f
			// src.go:4:11:	3.14
			// src.go:4:17:	2
			// src.go:4:21:	c
		}
		return true
	})

}
