package main

import (
	"archive/zip"
	"bytes"
	"log"
)

// Writer实现了zip文件编写器。
func main()  {
	// 创建一个缓冲区来写入我们的归档文件。
	buf := new(bytes.Buffer)

	// C创建一个新的压缩文件。
	w := zip.NewWriter(buf)

	// 向存档添加一些文件。
	var files = []struct {
		Name, Body string
	} {
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 确保在 Close 上检查错误。
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}