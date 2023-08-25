package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 使用具有自定义拆分函数的扫描器来解析逗号分隔的 最后一个值为空的列表。
func main()  {
	// 逗号分隔的列表; 最后一项为空。
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 定义一个以逗号分隔的拆分函数。
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// 要传递的最后一个令牌可能是空字符串。
		// 返回 bufio. ErrFinalToken 这里告诉 Scan 在这之后没有更多的令牌了
		// 但不会触发从 Scan 本身返回的错误。
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	// 扫描
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}