package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	const expr = "(6+2*3)/4"

	// 解析器.ParseExpr 解析参数并返回相应的 ast.Node.
	node, err := parser.ParseExpr(expr)
	if err != nil {
		log.Fatal(err)
	}

	// 为节点创建 FileSet。由于节点并非来自真正的源文件，因此 fset 将为空。
	fset := token.NewFileSet()

	var buf bytes.Buffer
	err = format.Node(&buf, fset, node)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String()) // (6 + 2*3) / 4

}
