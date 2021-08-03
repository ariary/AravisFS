package ubac

import (
	"encoding/json"
	"fmt"

	"github.com/ariary/AravisFS/pkg/filesystem"
)

//Apply the patch to the encrypted fs
//(it first get the patch and decompose it to remove,add and change resources)
func ApplyPatch(patch string, filename string) (err error) {

	var patchStruct filesystem.Patch
	//Unmarshall patch
	err = json.Unmarshal([]byte(patch), &patchStruct)
	//Apply Patch
	var resources filesystem.ResourceList
	resources, err = GetResourceList(filename)
	//apply resources change
	changeMap := patchStruct.ChangeMap
	for resourceName, content := range changeMap {
		resources.ChangeResourceContentFromName(resourceName, content)
	}
	//apply remove resources
	removeList := patchStruct.RemoveList
	for i := 0; i < len(removeList); i++ {
		fmt.Println(removeList[i])
		resources.RemoveResourceFromName(removeList[i])
	}
	//apply resource adding
	addList := patchStruct.AddList
	for i := 0; i < len(addList); i++ {
		resources.AddResource(addList[i])
	}
	filesystem.OverwriteFSFile(filename, resources)
	return err
}
