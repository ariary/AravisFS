package main

import (
	"log"

	"github.com/ariary/AravisFS/pkg/encrypt"
)

func main() {

	log.Println("Hello adret!")
	key := "toto"
	text := []byte("thisis/a/test")
	encrypted := encrypt.Encrypt(text, key)
	log.Println("  encrypted: " + string(encrypted))
	decrypted := encrypt.Decrypt(encrypted, key)
	log.Println("  decrypted: " + string(decrypted))
}
