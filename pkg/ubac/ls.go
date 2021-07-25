package ubac

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/ariary/AravisFS/pkg/filesystem"
)

// Provide ls result in string
// Return the filename if it is a file and the different file within it is a directory
// And an error if the if the file is not retrieve or the resource Type is not a file or directory
// Behind teh scene: retrieve the list of resource, iterate over and compare each resource name with the
// one provided. Then in function of the resource name it will adapt the output string
func lsResult(resourcename string, filename string) (string, error) {
	resource := GetResourceInFS(resourcename, filename)
	if resource == nil {
		return "", errors.New(fmt.Sprintf("ls: cannot access %v: No such file or directory", resourcename))
	}
	if resource.Type == filesystem.DIRECTORY {
		files := filesystem.ParseDirectoryContent(string(resource.Content))
		//TO DO: specify if it is a file or directory
		output_files := strings.Join(files, " ")
		return output_files, nil

	} else if resource.Type == filesystem.FILE {
		return resource.Name, nil
	} else {
		return "", errors.New(fmt.Sprintf("ls: invalid resource type %v for resource %v", resource.Type, resourcename))
	}
}

// Ls aim to act like standard ls (with less options)
// It is used to simulate an 'ls' call in the fs .arafs
// resourcename represent the file/directory which would be passed to ls command
// filename is the pathname of the fs .arafs
// use lsResult string and print it
func Ls(resourcename string, filename string) {
	resourcename = filepath.Clean(resourcename)

	result, err := lsResult(resourcename, filename)
	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}
	fmt.Println(result)
}
