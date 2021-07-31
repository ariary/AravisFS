package ubac

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"github.com/ariary/AravisFS/pkg/filesystem"
)

// Provide ls result in string + resource type on which we operate ls
// ( it will help adret to parse the content)
// Return the filename (encrypted) if it is a file and the content (encrypted) it is a directory
// the directory content is a string of all the resources under the directory separated by "\"
// And an error if the if the file is not retrieved or the resource Type is not a file or directory
// Behind the scene: retrieve the list of resource, iterate over and compare each resource name with the
// one provided.
func Ls(resourcename string, filename string) (lscontent string, err error) {

	resource := GetResourceInFS(resourcename, filename)

	if resource.Type == "" { // resource == nil doesn't work
		err = errors.New(fmt.Sprintf("ls: cannot access %v: No such file or directory", resourcename))
		return "", err
	}
	if resource.Type == filesystem.DIRECTORY {
		content := base64.StdEncoding.EncodeToString(resource.Content)
		resourceType := filesystem.DIRECTORY
		lscontent = resourceType + ":" + content
		return lscontent, nil

	} else if resource.Type == filesystem.FILE {
		content := base64.StdEncoding.EncodeToString(resource.Name)
		resourceType := filesystem.FILE
		lscontent = resourceType + ":" + content
		return lscontent, nil
	} else {
		err = errors.New(fmt.Sprintf("ls: invalid resource type %v for resource %v", resource.Type, resourcename))
		return "", err
	}
}

func PrintLs(resourcename string, filename string) {
	content, err := Ls(resourcename, filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(content)

}
