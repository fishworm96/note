package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	// src 是我们要标记化的输入内容。
	src := []byte("cos(x) + 1i*sin(x) // Eular")

	// 初始化扫描仪。
	var s scanner.Scanner
	fset := token.NewFileSet() // 位置相对于 fset
	file := fset.AddFile("", fset.Base(), len(src)) // 注册输入 "文件"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// 重复调用 "扫描 "会得到输入中的标记序列。
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
		// 1:1	IDENT	"cos"
		// 1:4	(	""
		// 1:5	IDENT	"x"
		// 1:6	)	""
		// 1:8	+	""
		// 1:10	IMAG	"1i"
		// 1:12	*	""
		// 1:13	IDENT	"sin"
		// 1:16	(	""
		// 1:17	IDENT	"x"
		// 1:18	)	""
		// 1:20	COMMENT	"// Euler"
		// 1:28	;	"\n"
	}
}