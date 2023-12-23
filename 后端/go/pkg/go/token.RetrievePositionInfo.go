package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()

	const src = `package main

import "fmt"

import "go/token"

//line :1:5
type p = token.Pos

const bad = token.NoPos

//line fake.go:42:11
func ok(pos p) bool {
	return pos != bad
}

/*line :7:9*/func main() {
	fmt.Println(ok(bad) == bad.IsValid())
}
`

	f, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印 f 中每个声明的位置和种类。
	for _, decl := range f.Decls {
		// 通过文件集获取文件名、行和列。 我们同时获取相对位置和绝对位置。 相对位置是相对于最后一行指令的位置。 绝对位置是源文件中的确切位置。
		pos := decl.Pos()
		relPosition := fset.Position(pos)
		absPosition := fset.PositionFor(pos, false)

		// 要么是 FuncDecl，要么是 GenDecl，因为我们会在出错时退出。
		kind := "func"
		if gen, ok := decl.(*ast.GenDecl); ok {
			kind = gen.Tok.String()
		}

		// 如果相对位置和绝对位置不同，则同时显示这两个位置。
		fmtPosition := relPosition.String()
		if relPosition != absPosition {
			fmtPosition += "[" + absPosition.String() + "]"
		}

		fmt.Printf("%s: %s\n", fmtPosition, kind)
		// main.go:3:1: import
		// main.go:5:1: import
		// main.go:1:5[main.go:8:1]: type
		// main.go:3:1[main.go:10:1]: const
		// fake.go:42:11[main.go:13:1]: func
		// fake.go:7:9[main.go:17:14]: func
	}

}
