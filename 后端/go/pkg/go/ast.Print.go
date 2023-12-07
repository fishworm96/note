package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// 该示例展示了 AST 在调试时的打印效果。
func main() {
	// src 是我们要打印 AST 的输入。
	src := `
package main
func main() {
	println("Hello, World!")
}
`

	// 通过解析 src 创建 AST。
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// 打印 AST。
	ast.Print(fset, f)
	//  0  *ast.File {
	// 	1  .  Package: 2:1
	// 	2  .  Name: *ast.Ident {
	// 	3  .  .  NamePos: 2:9
	// 	4  .  .  Name: "main"
	// 	5  .  }
	// 	6  .  Decls: []ast.Decl (len = 1) {
	// 	7  .  .  0: *ast.FuncDecl {
	// 	8  .  .  .  Name: *ast.Ident {
	// 	9  .  .  .  .  NamePos: 3:6
	//  10  .  .  .  .  Name: "main"
	//  11  .  .  .  .  Obj: *ast.Object {
	//  12  .  .  .  .  .  Kind: func
	//  13  .  .  .  .  .  Name: "main"
	//  14  .  .  .  .  .  Decl: *(obj @ 7)
	//  15  .  .  .  .  }
	//  16  .  .  .  }
	//  17  .  .  .  Type: *ast.FuncType {
	//  18  .  .  .  .  Func: 3:1
	//  19  .  .  .  .  Params: *ast.FieldList {
	//  20  .  .  .  .  .  Opening: 3:10
	//  21  .  .  .  .  .  Closing: 3:11
	//  22  .  .  .  .  }
	//  23  .  .  .  }
	//  24  .  .  .  Body: *ast.BlockStmt {
	//  25  .  .  .  .  Lbrace: 3:13
	//  26  .  .  .  .  List: []ast.Stmt (len = 1) {
	//  27  .  .  .  .  .  0: *ast.ExprStmt {
	//  28  .  .  .  .  .  .  X: *ast.CallExpr {
	//  29  .  .  .  .  .  .  .  Fun: *ast.Ident {
	//  30  .  .  .  .  .  .  .  .  NamePos: 4:2
	//  31  .  .  .  .  .  .  .  .  Name: "println"
	//  32  .  .  .  .  .  .  .  }
	//  33  .  .  .  .  .  .  .  Lparen: 4:9
	//  34  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
	//  35  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
	//  36  .  .  .  .  .  .  .  .  .  ValuePos: 4:10
	//  37  .  .  .  .  .  .  .  .  .  Kind: STRING
	//  38  .  .  .  .  .  .  .  .  .  Value: "\"Hello, World!\""
	//  39  .  .  .  .  .  .  .  .  }
	//  40  .  .  .  .  .  .  .  }
	//  41  .  .  .  .  .  .  .  Ellipsis: -
	//  42  .  .  .  .  .  .  .  Rparen: 4:25
	//  43  .  .  .  .  .  .  }
	//  44  .  .  .  .  .  }
	//  45  .  .  .  .  }
	//  46  .  .  .  .  Rbrace: 5:1
	//  47  .  .  .  }
	//  48  .  .  }
	//  49  .  }
	//  50  .  FileStart: 1:1
	//  51  .  FileEnd: 5:3
	//  52  .  Scope: *ast.Scope {
	//  53  .  .  Objects: map[string]*ast.Object (len = 1) {
	//  54  .  .  .  "main": *(obj @ 11)
	//  55  .  .  }
	//  56  .  }
	//  57  .  Unresolved: []*ast.Ident (len = 1) {
	//  58  .  .  0: *(obj @ 29)
	//  59  .  }
	//  60  .  GoVersion: ""
	//  61  }
}
