package main

import (
	"crypto/ed25519"
	"log"
)

func main() {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatal(err)
	}

	msg := []byte("The quick brown fox jumps over the lazy dog")

	sig, err := priv.Sign(nil, msg, &ed25519.Options{
		Context: "Example_ed25519",
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := ed25519.VerifyWithOptions(pub, msg, sig, &ed25519.Options{
		Context: "Example_ed25510ctx",
	}); err != nil {
		log.Fatal("invalid signature")
	}
}