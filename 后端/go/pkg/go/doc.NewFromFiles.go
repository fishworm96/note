package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
)

// 本例通过示例说明如何使用 NewFromFiles 计算软件包文档。
func main() {
	// src 和 test 是组成软件包的两个源文件，将对其文档进行计算。
	const src = `
// This is the package comment.
package p

import "fmt"

// This comment is associated with the Greet function.
func Greet(who string) {
	fmt.Printf("Hello, %s!\n", who)
}
`
	const test = `
package p_test

// This comment is associated with the ExampleGreet_world example.
func ExampleGreet_world() {
	Greet("world")
}
`

	// 通过解析 src 和 test 创建 AST。
	fset := token.NewFileSet()
	files := []*ast.File{
		mustParse(fset, "src.go", src),
		mustParse(fset, "src_test.go", test),
	}

	// 利用示例计算软件包文档。
	p, err := doc.NewFromFiles(fset, files, "example.com/p")
	if err != nil {
		panic(err)
	}

	fmt.Printf("package %s - %s", p.Name, p.Doc) // package p - This is the package comment.
	fmt.Printf("func %s - %s", p.Funcs[0].Name, p.Funcs[0].Doc) // func Greet - This comment is associated with the Greet function.
	fmt.Printf(" ⤷ example with suffix %q - %s", p.Funcs[0].Examples[0].Suffix, p.Funcs[0].Examples[0].Doc) //  ⤷ example with suffix "world" - This comment is associated with the ExampleGreet_world example.

}

func mustParse(fset *token.FileSet, filename, src string) *ast.File {
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return f
}
