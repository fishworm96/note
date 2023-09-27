package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func main() {
	// 这是一个编写压缩阅读器的例子。
	// 如图所示，这对 HTTP 客户端正文很有用。

	const testdata = "the data to be compressed"

	// 此 HTTP 处理程序仅用于测试目的。
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		zr, err := gzip.NewReader(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		// 只需输出示例的数据即可。
		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// 其余为示例代码。

	// 我们要压缩的数据，以 io.Reader 格式显示
	dataReader := strings.NewReader(testdata)

	// bodyReader 是 HTTP 请求的正文，作为 io.Reader.BodyReader 的一部分。
	// httpWriter 是 HTTP 请求的正文，作为一个 io.Writer 文件。
	bodyReader, httpWriter := io.Pipe()

	// 确保 bodyReader 始终处于关闭状态，以便下面的 goroutine 将始终退出。
	defer bodyReader.Close()

	// gzipWriter 会压缩 httpWriter 中的数据。
	gzipWriter := gzip.NewWriter(httpWriter)

	// errch 从写入例程中收集任何错误。
	errch := make(chan error, 1)

	go func() {
		defer close(errch)
		sentErr := false
		sendErr := func(err error) {
			if !sentErr {
				errch <- err
				sentErr = true
			}
		}

		// 将数据复制到 gzipWriter，它将数据压缩为 gzipWriter 会将其输入 bodyReader。
		if _, err := io.Copy(gzipWriter, dataReader); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
		if err := gzipWriter.Close(); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
		if err := httpWriter.Close(); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
	}()

	// 向测试服务器发送 HTTP 请求。
	req, err := http.NewRequest("PUT", ts.URL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	// 请注意，将 req 传递给 http.Client.Do 承诺将关闭 body，在本例中是 bodyReader。
	resp, err := ts.Client().Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// 检查压缩数据时是否出错。
	if err := <-errch; err != nil {
		log.Fatal(err)
	}

	// 在本例中，我们不关心响应。
	resp.Body.Close()

	// the data to be compressed
}