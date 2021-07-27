package ubac

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"github.com/ariary/AravisFS/pkg/filesystem"
)

//Provide cat result in string + resource type on which we operate the cat
// Return an error message if it is a directory or the file is not retrieved or the resource Type is not a file or directory
// Behind the scene: retrieve the list of resource, iterate over and compare each resource name with the
// one provided.
func Cat(resourcename string, filename string) (content string, err error) {

	resource := GetResourceInFS(resourcename, filename)
	// cat only it is a file
	if resource.Type == filesystem.FILE {
		content = base64.StdEncoding.EncodeToString(resource.Content)
		return content, nil

	} else {
		//error cases
		if resource.Type == "" { // resource == nil doesn't work
			err = errors.New(fmt.Sprintf("cat: %v: No such file or directory", resourcename))
			return "", err
		} else if resource.Type == filesystem.DIRECTORY {
			err = errors.New(fmt.Sprintf("cat: %v: Is a directory", resourcename))
			return "", err
		} else {
			err = errors.New(fmt.Sprintf("cat: invalid resource type %v for resource %v", resource.Type, resourcename))
			return "", err
		}
	}
}

func PrintCat(resourcename string, filename string) {
	content, err := Cat(resourcename, filename)

	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	} else {
		fmt.Println(content)
	}

}
