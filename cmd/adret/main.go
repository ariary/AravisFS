package main

import (
	"fmt"
	"log"

	"github.com/ariary/AravisFS/pkg/filesystem"
)

func main() {

	log.Println("Hello adret!")
	// key := "toto"
	// // text := []byte("thisis/a/test")
	// // encrypted := encrypt.Encrypt(text, key)
	// // log.Println("  encrypted: " + string(encrypted))
	// // decrypted := encrypt.Decrypt(encrypted, key)
	// // log.Println("  decrypted: " + string(decrypted))
	// // file_content_encrypted := encrypt.EncryptFile("test/mytestfolder/tata/binary_hello", key)
	// file_content_encrypted := encrypt.EncryptFile("test/mytestfolder/toto.txt", key)
	// file_content_decrypted := encrypt.DecryptByte(file_content_encrypted, key)
	// //log.Println("  decrypted file content: " + string(file_content_decrypted))
	// f, _ := os.Create("hello")
	// defer f.Close()
	// f.Write(file_content_decrypted)

	//FS
	// r := filesystem.Test("test/mytestfolder/tata/binary_hello", "file", file_content_decrypted)
	// filesystem.Test2(r)

	// //Browse fs to construct
	filesystem.CreateAravisFS("./test/mytestfolder", "toto")

	// //test darkenPath
	// fmt.Println(encrypt.DarkenPath("./test/mytestfolder", "toto"))

	// //test decrypt
	// fmt.Println(string(encrypt.DecryptString("VTNLdlzLxDZd7S6vfADi2wXGpcVUE6jbK4B3t/qn//TcyOQwGe90OJ2ole1WrAtFenKRPMyF", "toto")))

	//test GetdirectoryContent
	fmt.Println(filesystem.GetDirectoryContent("./test/"))
}
