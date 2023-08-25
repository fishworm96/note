package main

import (
	"bufio"
	"fmt"
	"os"
)

// 编写器为io.writer对象实现缓冲。如果写入Writer时发生错误，则不再接受任何数据，并且所有后续写入和Flush都将返回错误。在写入所有数据之后，客户端应该调用Flush方法来保证所有数据都已转发到底层io.Writer。
func main()  {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush() // 不要忘记刷新
}