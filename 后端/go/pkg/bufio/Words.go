package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 使用扫描器扫描 输入为空格分隔的标记序列。
func main()  {
	// 人工输入源。
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 为扫描操作设置拆分函数。
	scanner.Split(bufio.ScanWords)
	// 统计单词数
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}