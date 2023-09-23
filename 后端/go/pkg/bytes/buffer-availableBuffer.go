package main

import (
	"bytes"
	"os"
	"strconv"
)

// AvailableBuffer 返回一个容量为 b.Available() 的空缓冲区。该缓冲区将被追加并传递给紧随其后的 "写 "调用。该缓冲区只在下次对 b 进行写操作之前有效。
func main() {
	var buf bytes.Buffer
	for i := 0; i < 4; i++ {
		b := buf.AvailableBuffer()
		b = strconv.AppendInt(b, int64(i), 10)
		b = append(b, ' ')
		buf.Write(b)
	}
	os.Stdout.Write(buf.Bytes()) // 0 1 2 3
}