## package hmac

import "crypto/hmac"

软件包 hmac 实现了美国联邦信息处理标准出版物 198 中定义的密钥散列信息验证码（HMAC）。HMAC 是一种使用密钥对信息进行签名的加密散列。接收方通过使用相同的密钥重新计算哈希值来验证。

## Index

### func Equal(mac1, mac2 []byte) bool 添加于1.1

Equal 在不泄露定时信息的情况下，比较两个 MAC 是否相等。

### func New(h func() hash.Hash, key []byte) hash.Hash

New 使用给定的 hash.Hash 类型和密钥返回新的 HMAC 哈希值。每次调用 h 时，它必须返回一个新的哈希值。请注意，与标准库中的其他散列实现不同，返回的 Hash 并不实现 encoding.BinaryMarshaler 或 encoding.BinaryUnmarshaler。