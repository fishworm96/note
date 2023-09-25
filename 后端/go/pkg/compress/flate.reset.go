package main

import (
	"bytes"
	"compress/flate"
	"io"
	"log"
	"os"
	"strings"
)

// 在对性能要求极高的应用中，重置可用于丢弃当前的压缩器或解压器状态，并利用先前分配的内存快速重新初始化。
func main() {
	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.\n",
		"Concurrency is not parallelism.\n",
		"The bigger the interface, the weaker the abstraction.\n",
		"Documentation is for users.\n",
	}

	var r strings.Reader
	var b bytes.Buffer
	buf := make([]byte, 32<<10)

	zw, err := flate.NewWriter(nil, flate.DefaultCompression)
	if err != nil {
		log.Fatal(err)
	}
	zr := flate.NewReader(nil)

	for _, s := range proverbs {
		r.Reset(s)
		b.Reset()

		// 重置压缩器，并对输入流进行编码。
		zw.Reset(&b)
		if _, err := io.CopyBuffer(zw, &r, buf); err != nil {
			log.Fatal(err)
		}
		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		// 重置解码器，并解码为某个输出流。
		if err := zr.(flate.Resetter).Reset(&b, nil); err != nil {
			log.Fatal(err)
		}
		if _, err := io.CopyBuffer(os.Stdout, zr, buf); err != nil {
			log.Fatal(err)
		}
		if err := zr.Close(); err != nil {
			log.Fatal(err)
		}
	}
// Don't communicate by sharing memory, share memory by communicating.
// Concurrency is not parallelism.
// The bigger the interface, the weaker the abstraction.
// Documentation is for users.
}