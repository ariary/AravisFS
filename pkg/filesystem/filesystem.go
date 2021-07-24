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

type resource struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content []byte `json:"content"`
}

type resourceList struct {
	List []resource
}

func Test(filename string, resourceType string, content []byte) resource {
	r := createResource(filename, resourceType, content)

	return r
}
func createResource(filename string, resourceType string, content []byte) resource {

	r := &resource{
		Name:    filename,
		Type:    resourceType,
		Content: content}
	return *r
	// rjson, _ := json.Marshal(r)
	// fmt.Println(string(rjson))
	// jsondat := &resourceList{List: []resource{*r, *r}}
	// encjson, _ := json.Marshal(jsondat)
	// fmt.Println(string(encjson))
}

func Test2(r resource) {
	// var resources *resourceList
	// resources = new(resourceList)

	rl := []resource{}
	resources := resourceList{rl}

	resources.Addresource(r)

	WriteFSFile(resources)
}
func (resources *resourceList) Addresource(r resource) []resource {
	resources.List = append(resources.List, r)
	return resources.List
}

func WriteFSFile(resources resourceList) {
	file, _ := json.MarshalIndent(resources, "", " ")
	// file, _ := json.Marshal(resources)

	_ = ioutil.WriteFile("ceciestlav1.arafs", file, 0644)
}

func PrintFSFile(resources resourceList) {
	file, _ := json.MarshalIndent(resources, "", " ")
	// file, _ := json.Marshal(resources)

	fmt.Println(string(file))
}

// CreateAravisFS take the path of a directory parameter and write the .arafs file representing the
// the encrypted fs of this directory with the key parameter
// Encrypted fs is a list of the resources. By resource we mean resource=[name,is it a dir?,content].
// Take into account that name and content are encrypted with the key
func CreateAravisFS(path string, key string) {

	rl := []resource{}
	resources := resourceList{rl}

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			var resourceType string
			if info.IsDir() {
				resourceType = "directory"
			} else {
				resourceType = "file"
			}
			r := createResource(path, resourceType, encrypt.EncryptFile(path, key))
			resources.Addresource(r)
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	WriteFSFile(resources)

}
