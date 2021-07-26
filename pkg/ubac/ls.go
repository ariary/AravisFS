package ubac

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"github.com/ariary/AravisFS/pkg/filesystem"
)

// Provide ls result in string + string specifying resource type on which we operate ls
// ( ti will help adret to parse the content)
// Return the filename (encrypted) if it is a file and the content (encrypted) it is a directory
// the directory content is a string of all the resources under the directory separated by "\"
// And an error if the if the file is not retrieve or the resource Type is not a file or directory
// Behind teh scene: retrieve the list of resource, iterate over and compare each resource name with the
// one provided.
func Ls(resourcename string, filename string) (content string, resourceType string, err error) {

	resource := GetResourceInFS(resourcename, filename)

	if resource.Type == "" { // resource == nil doesn't work
		err = errors.New(fmt.Sprintf("ls: cannot access %v: No such file or directory", resourcename))
		return "", "", err
	}
	if resource.Type == filesystem.DIRECTORY {
		// //TO DO: specify if it is a file or directory
		content = base64.StdEncoding.EncodeToString(resource.Content)
		resourceType = filesystem.DIRECTORY
		return content, resourceType, nil

	} else if resource.Type == filesystem.FILE {
		content = string(string(resource.Name))
		resourceType = filesystem.FILE
		return content, resourceType, nil
	} else {
		err = errors.New(fmt.Sprintf("ls: invalid resource type %v for resource %v", resource.Type, resourcename))
		return "", "", err
	}
}

func PrintLs(resourcename string, filename string) {
	content, resourceType, err := Ls(resourcename, filename)

	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	} else {
		fmt.Println(resourceType + ":" + content)
	}

}
