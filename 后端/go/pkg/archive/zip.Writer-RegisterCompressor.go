package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"io"
)

// RegisterCompressor为特定的方法ID注册或重写自定义压缩器。如果没有找到给定方法的压缩器，Writer将默认在包级别查找压缩器。
func main()  {
	// 使用更高的压缩级别重写默认的 Deflate 压缩器。

	// 创建一个缓冲区，将我们的归档文件写入其中。
	buf := new(bytes.Buffer)

	// 创建一个新的压缩文件
	w := zip.NewWriter(buf)

	// 注册一个自定义 Deflate 压缩器。
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriterCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// 继续将文件添加到 w。
}