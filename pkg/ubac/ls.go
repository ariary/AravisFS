package ubac

import (
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
	if resource == nil {
		// TO DO: == nil does'nt work
		err = errors.New(fmt.Sprintf("ls: cannot access %v: No such file or directory", resourcename))
		return "", "", err
	}
	if resource.Type == filesystem.DIRECTORY {
		// //TO DO: specify if it is a file or directory
		content = string(resource.Content)
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
		fmt.Print(resourceType + ":" + content)
	} else {
		log.SetFlags(0)
		log.Fatal(err)
	}

}

// Like lsResult but only take base
// ie toto/tatat/titi.txt ---> titi.txt
// It is more accurate with th edefault behavior of ls
// Otherwise with AravisFS we must specify resource  with the whole name
// when we xant to do actin action with so it may be helpful to have full path
// func lsResultBase(resourcename string, filename string) (string, error) {
// 	resource := GetResourceInFS(resourcename, filename)
// 	if resource == nil {
// 		return "", errors.New(fmt.Sprintf("ls: cannot access %v: No such file or directory", resourcename))
// 	}
// 	if resource.Type == filesystem.DIRECTORY {
// 		files := filesystem.ParseDirectoryContent(string(resource.Content))
// 		//take basename
// 		for i := range files {
// 			files[i] = filepath.Base(files[i])
// 		}

// 		//TO DO: specify if it is a file or directory
// 		output_files := strings.Join(files, " ")
// 		return output_files, nil

// 	} else if resource.Type == filesystem.FILE {
// 		return resource.Name, nil
// 	} else {
// 		return "", errors.New(fmt.Sprintf("ls: invalid resource type %v for resource %v", resource.Type, resourcename))
// 	}
// }

// Ls aim to act like standard ls (with less options)
// It is used to simulate an 'ls' call in the fs .arafs
// resourcename represent the file/directory which would be passed to ls command
// filename is the pathname of the fs .arafs
// use lsResult string and print it
// func Ls(resourcename string, filename string) {
// 	resourcename = filepath.Clean(resourcename)

// 	result, resourceType, err := lsResultBase(resourcename, filename)
// 	if err != nil {
// 		log.SetFlags(0)
// 		log.Fatal(err)
// 	}
// 	fmt.Println(resourceType + ":" + result)
// }
