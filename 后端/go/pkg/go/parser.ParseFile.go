package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet() // 位置是相对于 FSET 的

	src := `package foo

import (
	"fmt"
	"time"
)

func bar() {
	fmt.Println(time.Now())
}`

	// 解析 src，但在处理导入后停止。
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印从文件的 AST 导入的内容。
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
		// fmt
		// time
	}

}
