package adret_test

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"testing"

	"github.com/ariary/AravisFS/pkg/adret"
	"github.com/ariary/AravisFS/pkg/encrypt"
	"github.com/ariary/AravisFS/pkg/filesystem"
)

func TestEncrypt(t *testing.T) {
	key := "toto"

	// test decrypt with byte and string
	textString := "thisis/a/test"
	text := []byte(textString)
	encrypted := encrypt.Encrypt(text, key)
	decrypted := string(encrypt.DecryptByte(encrypted, key))
	decryptedString := string(encrypt.DecryptStringFromUbac(base64.StdEncoding.EncodeToString(encrypted), key))
	if decrypted != textString {
		t.Errorf("Decryption was incorrect, got: %v, want: %v.", decrypted, textString)
	}

	if decryptedString != textString {
		t.Errorf("Decryption was incorrect, got: %v, want: %v.", decryptedString, textString)
	}

	//test encryptFile
	path := "../../../test/mytestfolder/toto.txt"
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	file_content_encrypted := encrypt.EncryptFile(path, key)
	file_content_decrypted := encrypt.DecryptByte(file_content_encrypted, key)

	if string(content) != string(file_content_decrypted) {
		t.Errorf("Decryption was incorrect, got: %v, want: %v.", string(content), string(file_content_decrypted))
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func TestGetNodeUnder(t *testing.T) {
	prefix := "je/suis/"

	nodes := []adret.Node{
		adret.CreateNode("je/suis", "", "je"),
		adret.CreateNode("je/suisaussi/tmp.txt", "file", "je/suisaussi"),
		adret.CreateNode("je/suis/ton/fils/oo", "", "je/suis/ton/fils"),
		adret.CreateNode("fake/je/suis/ton/fils/oo", "", "fake/je/suis/ton/fils"),
		adret.CreateNode("je/suis/aussi/tonfils", "", "je/suis/aussi"),
		adret.CreateNode("je/suis/tonpapa/oupas", "", "je/suis/tonpapa"),
		adret.CreateNode("je/suis/doncjepense", "", "je/suis/"),
		adret.CreateNode("je/nesuis/pas/ton/fils/pointtxt", "", "je/nesuis/pas/ton/fils"),
	}

	children := adret.GetNodesUnder(prefix, nodes)

	expected := []string{
		"je/suis/ton/fils/oo",
		"je/suis/aussi/tonfils",
		"je/suis/tonpapa/oupas",
		"je/suis/doncjepense",
	}

	if !Equal(expected, children) {
		t.Errorf("GetNodeUnder failed, got: %v, want: %v.", children, expected)
	}
}

func TestTree(t *testing.T) {
	path := "../../../test/mytestfolder/titi"
	prefix := path + "/"
	key := "riruhferiuh9898"
	contentEnc := filesystem.GetDirectoryContent(path, key)
	content := encrypt.DecryptByte(contentEnc, key)

	expected := prefix + "tutu.txt" + "\\" + prefix + "utut.txt"

	if expected != string(content) {
		t.Errorf("GetDirectoryContent failed, got: %v, want: %v.", string(content), expected)
	}
}

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
// filesystem.CreateAravisFS("./test/mytestfolder", "toto")
// fmt.Println(encrypt.DarkenPath("test/mytestfolder", "toto"))

// //test darkenPath
// fmt.Println(encrypt.DarkenPath("./test/mytestfolder", "toto"))

// //test decrypt
// fmt.Println(string(encrypt.DecryptString("VTNLdlzLxDZd7S6vfADi2wXGpcVUE6jbK4B3t/qn//TcyOQwGe90OJ2ole1WrAtFenKRPMyF", "toto")))

//test GetdirectoryContent
// fmt.Println(filesystem.GetDirectoryContent("./test/"))
