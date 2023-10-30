package main

func main() {
	message := []byte("message to be signed")
	signature, _ := hex.DecodeString("ad2766728615cc7a746cc553916380ca7bfa4f8983b990913bc69eb0556539a350ff0f8fe65ddfd3ebe91fe1c299c2fac135bc8c61e26be44ee259f2f80c1530")

	// 只有较小的信息可以直接签名；因此签名的是信息的哈希值，而不是信息本身。这就要求散列函数具有抗碰撞性。在撰写本文时（2016 年），SHA-256 是用于此目的的强度最低的哈希函数。
	hashed := sha256.Sum256(message)

	hashed := sha256.Sum256(message)
	err := rsa.VerifyPKCS1v15(&rsaPrivateKey.PublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return
	}

	// 签名是公钥对信息的有效签名。
}