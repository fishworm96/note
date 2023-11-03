package main

import (
	"crypto/tls"
	"crypto/x509"
)

func main() {
	// VerifyConnection 可用于替换和自定义连接验证。本示例展示的 VerifyConnection 实现与 crypto/tls 通常验证对等方证书的方式大致相同。

	// 客户端配置
	_ = &tls.Config{
		// 设置 InsecureSkipVerify 以跳过我们要替换的默认验证。这不会禁用 VerifyConnection。
		InsecureSkipVerify: true,
		VerifyConnection: func(cs tls.ConnectionState) error {
			opts := x509.VerifyOptions{
				DNSName: cs.ServerName,
				Intermediates: x509.NewCertPool(),
			}
			for _, cert := range cs.PeerCertificates[1:] {
				opts.Intermediates.AddCert(cert)
			}
			_, err := cs.PeerCertificates[0].Verify(opts)
			return err
		},
	}

	// 服务器端配置
	_ = &tls.Config{
		// 需要客户端证书（否则无论如何都会运行 VerifyConnection，并在访问 cs.PeerCertificates[0] 时发生恐慌），但不使用默认验证器验证。这不会禁用 VerifyConnection。
		ClientAuth: tls.RequireAnyClientCert,
		VerifyConnection: func(cs tls.ConnectionState) error {
			opts := x509.VerifyOptions{
				DNSName: cs.ServerName,
				Intermediates: x509.NewCertPool(),
				KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			}
			for _, cert := range cs.PeerCertificates[1:] {
				opts.Intermediates.AddCert(cert)
			}
			_, err := cs.PeerCertificates[0].Verify(opts)
			return err
		},
	}

	// 请注意，当证书不是由默认验证器处理时，ConnectionState.VerifiedChains 将为零。
}