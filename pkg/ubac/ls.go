package ubac

import (
	"encoding/base64"
	"fmt"

	"github.com/ariary/AravisFS/pkg/filesystem"
)

// Provide ls result in string + resource type on which we operate ls
// ( it will help adret to parse the content)
// Return the filename (encrypted) if it is a file and the content (encrypted) it is a directory
// the directory content is a string of all the resources under the directory separated by "\"
// And an error if the if the file is not retrieved or the resource Type is not a file or directory
// Behind the scene: retrieve the list of resource, iterate over and compare each resource name with the
// one provided.
func Ls(resourcename string, filename string) (lscontent string) {

	resource := GetResourceInFS(resourcename, filename)

	if resource.Type == "" { // resource == nil doesn't work
		return fmt.Sprintf("ls: cannot access %v: No such file or directory", resourcename)
	}
	if resource.Type == filesystem.DIRECTORY || resource.Type == filesystem.FILE {
		// //TO DO: specify if it is a file or directory
		content := base64.StdEncoding.EncodeToString(resource.Content)
		resourceType := filesystem.DIRECTORY
		lscontent = resourceType + ":" + content
		return lscontent

	} else {
		return fmt.Sprintf("ls: invalid resource type %v for resource %v", resource.Type, resourcename)
	}
}

func PrintLs(resourcename string, filename string) {
	content := Ls(resourcename, filename)

	fmt.Println(content)

}
