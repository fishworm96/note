package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

// Reader 提供来自 ZIP 归档的内容。
func main() {
	// 打开压缩文件进行读取。
	r, err := zip.OpenReader("testdata/readme.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// 迭代归档中的文件,
	// 打印其中一些内容。
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}