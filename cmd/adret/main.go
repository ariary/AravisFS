package main

import (
	"fmt"
	"log"

	"github.com/ariary/AravisFS/pkg/encrypt"
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

	// test darkenedPath and decrypt path
	// encrypted := encrypt.Encrypt([]byte("test/mytestfolder"), "toto")
	// fmt.Println("Encrypted path: ", encrypted)
	// darkpath_enc := base64.StdEncoding.EncodeToString(encrypted)
	// fmt.Println("base64 path: ", darkpath_enc)
	// j, _ := json.Marshal(encrypted)
	// fmt.Println("marshal ", string(j))

	// decrypted := encrypt.DecryptPath("AdyAMnISHHVnsmk/zydtskgbq8VWS8hR2rCGjywHbouh6wm8cR8vjk8B98Cy", "toto")
	// fmt.Println("Decrypted path: ", decrypted)

	// data, _ := base64.StdEncoding.DecodeString("RL9d71jr5lVRaW/ryhGTYLaoDBH8quBOM/fYfUlUIn8ngS7EEobdDHajWcfw")

	// fmt.Printf("Decoded in main %q\n", string(encrypt.DecryptByte(data, "toto")))

	// path := "test/mytestfolder"

	// fmt.Println("Path to retrieve: RL9d71jr5lVRaW/ryhGTYLaoDBH8quBOM/fYfUlUIn8ngS7EEobdDHajWcfw which is ", path)
	// var pathEncrypted = encrypt.EncryptString(path, "toto")
	// fmt.Println("Path encrypted ", pathEncrypted)
	// j, _ := json.Marshal(pathEncrypted)
	// fmt.Println("marshal ", string(j))
	// darkpath_enc := base64.StdEncoding.EncodeToString(pathEncrypted)
	// fmt.Println("base64 encrypted path: ", darkpath_enc)

	// fmt.Printf("DarkenPAth in main %q\n", encrypt.DarkenPath("test/mytestfolder", "toto"))
	// encrypted := encrypt.Encrypt([]byte("test/mytestfolder"), "toto")
	// j, _ := json.Marshal(encrypted)
	// fmt.Println("marshal ", string(j))
	// fmt.Println("Encrypted path: ", encrypted)
	// darkpath_enc := base64.StdEncoding.EncodeToString(encrypted)
	// fmt.Println("DArk a la mano : ", darkpath_enc)

	// j2, _ := json.Unmarshal([]byte("AdyAMnISHHVnsmk/zydtskgbq8VWS8hR2rCGjywHbouh6wm8cR8vjk8B98Cy"))
	// fmt.Println(encrypt.DecryptByte(j2))
	// encrypted := encrypt.DarkenPath("toto/tata/test.txt", "toto")
	// fmt.Println("Encrypted path: ", encrypted)
	// decrypted := encrypt.DecryptPath(encrypted, "toto")
	// fmt.Println("Decrypted path: ", decrypted)

	// //Browse fs to construct
	filesystem.CreateAravisFS("./test/mytestfolder", "toto")
	fmt.Println(encrypt.DarkenPath("test/mytestfolder", "toto"))

	// //test darkenPath
	// fmt.Println(encrypt.DarkenPath("./test/mytestfolder", "toto"))

	// //test decrypt
	// fmt.Println(string(encrypt.DecryptString("VTNLdlzLxDZd7S6vfADi2wXGpcVUE6jbK4B3t/qn//TcyOQwGe90OJ2ole1WrAtFenKRPMyF", "toto")))

	//test GetdirectoryContent
	// fmt.Println(filesystem.GetDirectoryContent("./test/"))
}
