package ubac

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ariary/AravisFS/pkg/filesystem"
)

func GetResourceList(filename string) (filesystem.ResourceList, error) {
	fs, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer fs.Close()

	byteFS, _ := ioutil.ReadAll(fs)

	var rl filesystem.ResourceList
	json.Unmarshal(byteFS, &rl)

	if nil == rl.List {
		return rl, errors.New("empty resource list")
	}
	return rl, nil
}

// Take the name of the resource and search if the resource is in the list contains in
// the .arafs file pointed by the filename
func GetResourceInFS(resourcename string, filename string) *filesystem.Resource {
	resourcesList, err := GetResourceList(filename)
	if err != nil {
		log.Fatal(err)
	}
	resources := resourcesList.List
	for i := range resources { //for loop using index to avoid copy (perf)
		if string(resources[i].Name) == resourcename {
			return &resources[i]
		}
	}
	return nil
}
