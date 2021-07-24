package main

import (
	"fmt"
	"log"
	"os"

	"github.com/a8m/tree/ostree"
	"github.com/ariary/AravisFS/pkg/encrypt"
	"github.com/ariary/AravisFS/pkg/filesystem"
)

func main() {

	log.Println("Hello adret!")
	key := "toto"
	// text := []byte("thisis/a/test")
	// encrypted := encrypt.Encrypt(text, key)
	// log.Println("  encrypted: " + string(encrypted))
	// decrypted := encrypt.Decrypt(encrypted, key)
	// log.Println("  decrypted: " + string(decrypted))
	// file_content_encrypted := encrypt.EncryptFile("test/mytestfolder/tata/binary_hello", key)
	file_content_encrypted := encrypt.EncryptFile("test/mytestfolder/toto.txt", key)
	file_content_decrypted := encrypt.Decrypt(file_content_encrypted, key)
	//log.Println("  decrypted file content: " + string(file_content_decrypted))
	f, _ := os.Create("hello")
	defer f.Close()
	f.Write(file_content_decrypted)

	//FS
	// r := filesystem.Test("test/mytestfolder/tata/binary_hello", "file", file_content_decrypted)
	// filesystem.Test2(r)

	//Browse fs to construct
	// dir := new(ostree.FS)
	fmt.Println(ostree.Print("."))
	// fmt.Println(dir.ReadDir("."))
	// fmt.Println(dir.Stat("./README.md"))
	filesystem.MyVisitTree()
	filesystem.MyVisitWalk()

}
