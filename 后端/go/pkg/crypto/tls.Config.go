package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

// zeroSource 是一个 io.Reader，可返回无限数量的零字节。
type zeroSource struct{}

func (zeroSource) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 0
	}

	return len(b), nil
}

// Config 结构用于配置 TLS 客户端或服务器。在将 Config 传递给 TLS 函数后，不得对其进行修改。Config 可以重复使用；tls 软件包也不会修改它。
func main() {
	// 通过解密网络流量捕获调试 TLS 应用程序。
	// 警告：使用 KeyLogWriter 会影响安全性，只能用于调试。
	// 为示例使用不安全的随机 HTTP 服务器进行虚拟测试，以便输出结果可重现。
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.TLS = &tls.Config{
		Rand: zeroSource{}, // 仅作示例，请勿这样做。
	}
	server.StartTLS()
	defer server.Close()

	// 通常情况下，日志会存入一个打开的文件： w, err := os.OpenFile("tls-secrets.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	w := os.Stdout

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				KeyLogWriter: w,
				Rand: zeroSource{}, // 以获得可重复的输出；不要这样做。
				InsecureSkipVerify: true, // 测试服务器证书不可信。
			},
		},
	}
	resp, err := client.Get(server.URL)
	if err != nil {
		log.Fatal("Failed to get URL: %v", err)
	}
	resp.Body.Close()

	// 通过在 SSL 协议首选项中设置 (Pre)-Master-Secret 日志文件名，可以用 Wireshark 来解密 TLS 连接。
}