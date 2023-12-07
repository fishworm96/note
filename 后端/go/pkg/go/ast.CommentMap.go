package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

// 本例说明了如何使用 ast.CommentMap 删除 Go 程序中的变量声明，同时保持正确的注释关联。
func main() {
	// src 是我们创建 AST 的输入，我们将对其进行操作。
	src := `
// This is the package comment.
package main

// This comment is associated with the hello constant.
const hello = "Hello, World!" // line comment 1

// This comment is associated with the foo variable.
var foo = hello // line comment 2

// This comment is associated with the main function.
func main() {
	fmt.Println(hello) // line comment 3
}
`

	// 通过解析 src 创建 AST。
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// 这有助于保持注释和 AST 节点之间的关联。
	cmap := ast.NewCommentMap(fset, f, f.Comments)

	//从变量声明列表中删除第一个变量声明。
	for i, decl := range f.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.VAR {
			copy(f.Decls[i:], f.Decls[i+1:])
			f.Decls = f.Decls[:len(f.Decls)-1]
			break
		}
	}

	// 使用注释映射过滤不再属于变量的注释（与变量声明相关的注释），并创建新的注释列表。
	f.Comments = cmap.Filter(f).Comments()

	// 打印修改后的 AST。
	var buf strings.Builder
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}
	fmt.Printf("%s", buf.String())
	// This is the package comment.
	// package main

	// This comment is associated with the hello constant.
	// const hello = "Hello, World!" // line comment 1

	// This comment is associated with the main function.
	// func main() {
	// 	fmt.Println(hello) // line comment 3
	// }

}
