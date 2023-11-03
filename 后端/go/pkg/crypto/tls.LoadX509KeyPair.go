package main

import (
	"crypto/tls"
	"log"
)

// LoadX509KeyPair 可从一对文件中读取并解析一对公钥/私钥。文件必须包含 PEM 编码数据。证书文件可能包含叶子证书之后的中间证书，以形成证书链。成功返回时，Certificate.Leaf 将为零，因为证书的解析形式不会被保留。
func main() {
	cert, err := tls.LoadX509KeyPair("testdata/example-cert.pem", "testdata/example-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":2000", cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = listener
}