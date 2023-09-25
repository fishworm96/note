package main

import (
	"compress/flate"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	// 使用 io.Pipe 模拟网络连接。
	// 真正的网络应用程序应该注意正确关闭
	// 底层连接。
	rp, wp := io.Pipe()

		// 启动一个 goroutine 作为发送器。
		wg.Add(1)
		go func() {
			defer wg.Done()

			zw, err := flate.NewWriter(wp, flate.BestSpeed)
			if err != nil {
				log.Fatal(err)
			}

			b := make([]byte, 256)
			for _, m := range strings.Fields("A long time ago in a galaxy far, far away...") {
				// 我们使用一种简单的成帧格式，第一个字节是 
				// 报文长度，然后是报文本身。
				b[0] = uint8(copy(b[1:], m))

				if _, err := zw.Write(b[:1+len(m)]); err != nil {
					log.Fatal(err)
				}

				// 刷新确保接收器能读取到目前为止发送的所有数据。
				if err := zw.Flush(); err != nil {
					log.Fatal(err)
				}
			}

			if err := zw.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		// 启动一个 goroutine 作为接收器。
		wg.Add(1)
		go func() {
			defer wg.Done()

			zr := flate.NewReader(rp)

			b := make([]byte, 256)
			for {
				// 读取报文长度。 
				// 保证发送端每次相应的 
				// 清除和关闭操作都会返回此值。
				if _, err := io.ReadFull(zr, b[:1]); err != nil {
					if err == io.EOF {
						break // 发送机关闭了数据流
					}
					log.Fatal(err)
				}

				// 读取信息内容。
				n := int(b[0])
				if _, err := io.ReadFull(zr, b[:n]); err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Received %d bytes: %s\n", n, b[:n])
			}
			fmt.Println()

			if err := zr.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		// Received 1 bytes: A
		// Received 4 bytes: long
		// Received 4 bytes: time
		// Received 3 bytes: ago
		// Received 2 bytes: in
		// Received 1 bytes: a
		// Received 6 bytes: galaxy
		// Received 4 bytes: far,
		// Received 3 bytes: far
		// Received 7 bytes: away...
}