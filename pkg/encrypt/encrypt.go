package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//from https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
// We just want symmetric encryption. The aim is not to have a solid one
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(data []byte, key string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(key)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	// /!\ AES encryption must be used we a unique nonce at each encryption
	// Here we don't want a different output for the same input that's why
	// (ie. AES ECB)
	// BUT IT IS UNSECURE !!(https://zachgrace.com/posts/attacking-ecb/)
	//uncomment this part to have a unique nonce each time
	// if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
	// 	panic(err.Error())
	// }
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func EncryptFile(filename string, key string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return Encrypt(data, key)
}

// Encrypt a String
func EncryptString(filename string, key string) []byte {
	filenameByte := []byte(filename)
	return Encrypt(filenameByte, key)
}

// Return the path encrypted as it appears in the encrypted fs
// As we use JSON.MArshall to embed resource in fs and the resource name,
// which is the path encrypted, is a []byte it is encoded with base64
// In that way , to obtain path as it appear in FS we encrypt it and then encode
// it using base64
func DarkenPath(path string, key string) string {
	path = filepath.Clean(path)
	darkpath := Encrypt([]byte(path), key)
	darkpath_enc := base64.StdEncoding.EncodeToString(darkpath)
	return darkpath_enc
}

// Decrypt a path of encrypted fs
// first it decode it using base64 and then decrypt it
// finally it convert it in a string
func DecryptPath(encryptedPath string, key string) string {
	decoded, err := base64.StdEncoding.DecodeString(encryptedPath)
	if err != nil {
		fmt.Printf("Error decoding encrypted path: %s ", err.Error())
		return ""
	}
	return string(DecryptByte(decoded, key))
}

func DecryptByte(data []byte, key string) []byte {
	passphrase := []byte(createHash(key))
	block, err := aes.NewCipher(passphrase)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

// Decrypt String from ubac with the key
// As it is from ubac 'data' is a base64 encoded byte array
func DecryptString(data string, key string) []byte {
	dataDecoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return DecryptByte([]byte(dataDecoded), key)
}
