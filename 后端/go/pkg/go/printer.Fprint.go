package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
)

func parseFunc(filename, functionname string) (fun *ast.FuncDecl, fset *token.FileSet) {
	fset = token.NewFileSet()
	if file, err := parser.ParseFile(fset, filename, nil, 0); err == nil {
		for _, d := range file.Decls {
			if f, ok := d.(*ast.FuncDecl); ok && f.Name.Name == functionname {
				fun = f
				return
			}
		}
	}
	panic("function not found")
}

func printSelf() {
	// 为该函数解析源文件并提取不带注释的 AST，位置信息参考文件集 fset。
	funcAST, fset := parseFunc("example_test.go", "printSelf")

	// 将函数正文转入缓冲区 buf。提供给打印机的文件集可以让打印机了解原始源代码的格式，并在源代码中存在换行符的地方添加换行符。
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, funcAST.Body)

	// 删除包围函数体的大括号 {}，取消缩进、 并删除前导和尾部空白。
	s := buf.String()
	s = s[1 : len(s)-1]
	s = strings.TrimSpace(strings.ReplaceAll(s, "\n\t", "\n"))

	// 将清理后的正文文本打印到 stdout。
	fmt.Println(s)
}

func main() {
	printSelf()

}
