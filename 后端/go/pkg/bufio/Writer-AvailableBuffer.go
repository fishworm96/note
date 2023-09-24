package main

import (
	"bufio"
	"os"
	"strconv"
)

// AvailableBuffer返回一个具有b.Available（）容量的空缓冲区。此缓冲区旨在附加到并传递给紧接着的Write调用。缓冲区仅在b上的下一次写操作之前有效。
func main()  {
	w := bufio.NewWriter(os.Stdout)
	for _, i := range []int64{1, 2, 3, 4} {
		b := w.AvailableBuffer()
		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		w.Write(b)
	}
	w.Flush()
}