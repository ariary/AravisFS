package filesystem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ariary/AravisFS/pkg/encrypt"
)

func WriteFSFile(resources ResourceList) {
	file, _ := json.MarshalIndent(resources, "", " ")
	// file, _ := json.Marshal(resources)

	_ = ioutil.WriteFile("encrypted.arafs", file, 0644)
}

func PrintFSFile(resources ResourceList) {
	file, _ := json.MarshalIndent(resources, "", " ")
	// file, _ := json.Marshal(resources)

	fmt.Println(string(file))
}

// Return the content of a directory as a byte array
// the content is the list of files & directories within the directory
// content structure: "file1/file2/../fileN"
// it is return as a list of these resources in a string, each resource is seperated by the character "\"
// which is then converted in byte array
func GetDirectoryContent(dirname string, key string) []byte {
	dirname = filepath.Clean(dirname)
	f, err := ioutil.ReadDir(dirname) // (we use ReadDir instead of Walk to avoid recursively browsing the directory)
	if err != nil {
		log.Fatal(err)
	}
	var files string
	for _, file := range f {
		// Add path encrypted as it appears in the fs
		filename := dirname + "/" + file.Name()
		//Create
		if files == "" {
			files = filename
		} else {
			files = files + "\\" + filename
		}

	}

	fmt.Println(files)
	return encrypt.EncryptString(files, key)
}

// CreateAravisFS take the path of a directory parameter and write the .arafs file representing the
// the encrypted fs of this directory with the key parameter
// Encrypted fs is a list of the resources. By resource we mean resource=[name,is it a dir?,content].
// Take into account that name and content are encrypted with the key
func CreateAravisFS(path string, key string) {
	rl := []Resource{}
	resources := ResourceList{rl}

	path = filepath.Clean(path) // avoid "./ type path"

	// get list of all resources within this path (recursively)
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Encrypt path and add it

			// fmt.Println("Raw ", path)
			var pathEncrypted = encrypt.EncryptString(path, key)
			// fmt.Println("Byte ", pathEncrypted)
			// fmt.Println("encr raw ", string(pathEncrypted))
			// fmt.Println("Base64 ", base64.StdEncoding.EncodeToString(pathEncrypted))

			// Determine resource type and add content accordingly
			var resourceType string
			var resourceContent []byte
			if info.IsDir() {
				resourceType = DIRECTORY
				resourceContent = GetDirectoryContent(path, key)
				//encrypt it

			} else {
				resourceType = FILE
				resourceContent = encrypt.EncryptFile(path, key)
			}
			r := createResource(pathEncrypted, resourceType, resourceContent)
			resources.Addresource(r)
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	WriteFSFile(resources)

}
