package main

import (
	"bytes"
	"os"
)

// Bytes 返回一个长度为 b.Len() 的片段，其中包含缓冲区的未读部分。该片仅在下一次修改缓冲区之前有效（也就是说，仅在下一次调用读取、写入、重置或截断等方法之前有效）。至少在下一次修改缓冲区之前，片段都是缓冲区内容的别名，因此立即更改片段会影响未来的读取结果。
func main() {
	buf := bytes.Buffer{}
	buf.Write([]byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'})
	os.Stdout.Write(buf.Bytes()) // hello world
}