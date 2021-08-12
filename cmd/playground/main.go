package main

import (
	"fmt"

	"github.com/ariary/AravisFS/pkg/encrypt"
)

func main() {
	key := "toto955748deded!"
	path := "referfref541fe8r4er84ferfer7%58s\""
	pathEnc := encrypt.Encrypt([]byte(path), key)
	fmt.Println("pathEnc:", string(pathEnc))
	pathDec := encrypt.DecryptByte(pathEnc, key)
	fmt.Println("pathDec:", string(pathDec))

	pathEnc2 := encrypt.Encrypt([]byte(path), key)
	fmt.Println("pathEnc2:", string(pathEnc2))

	pathEncN, nonce := encrypt.EncryptAndGetNonce([]byte(path), key)
	fmt.Println("pathEncN:", string(pathEncN))
	fmt.Println("nonce:", string(nonce))
	pathEncWithNonce := encrypt.EncryptWithNonce([]byte(path), nonce, key)
	fmt.Println("pathEncN==pathEncWithNonce:", string(pathEncN) == string(pathEncWithNonce))
}
