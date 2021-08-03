package adret

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/ariary/AravisFS/pkg/encrypt"
	"github.com/ariary/AravisFS/pkg/filesystem"
)

//provide the patch to remove a resource on ubac side
// the patch is a json string with 3 arrays: to_add,to_delete and to_change
//all the info within structure are encrypted after to be put in patch
func GetRmPatch(key string, tree Tree, resourceName string) filesystem.Patch {
	resourceName = path.Clean(resourceName)
	var removeList []string
	changeMap := make(map[string]string)

	//check if resource exist
	if !Exist(resourceName, tree.Nodes) {
		fmt.Println(fmt.Sprintf("rm: cannot remove '%v': No such file or directory", resourceName))
		os.Exit(1)
	}

	//add the resource to remove list
	resourceNameEnc := encrypt.DarkenPath(resourceName, key)
	removeList = append(removeList, resourceNameEnc)

	//modify parent content (remove resource from it)
	if tree.rootNode != resourceName { //check if resource isn't the root node (no parent)
		parentName := path.Dir(resourceName)

		parentContent := GetChildrenNodes(parentName, tree.Nodes)
		//remove resource from parent content trick
		for i, resource := range parentContent {
			if resource == resourceName {
				parentContent = append(parentContent[:i], parentContent[i+1:]...)
				break
			}
		}
		//translate the new parent content in arafs format
		var newParentContent string
		for _, resource := range parentContent {
			//Create
			if newParentContent == "" {
				newParentContent = resource
			} else {
				newParentContent = newParentContent + "\\" + resource
			}
		}
		newParentContentEnc := encrypt.DarkenPath(newParentContent, key)
		parentNameEnc := encrypt.DarkenPath(parentName, key)
		// changeMap[parentName] = newParentContentByte
		changeMap[parentNameEnc] = newParentContentEnc
	}

	//delete all resource under if it is a directory
	if IsDir(resourceName, tree.Nodes) {
		descendantNodes := GetDescendantNodes(resourceName, tree.Nodes)
		//add encruypted name of descendant to removeList
		for _, nodeName := range descendantNodes {
			removeList = append(removeList, encrypt.DarkenPath(nodeName, key))
		}

	}
	//enc

	return filesystem.CreatePatch(nil, removeList, changeMap)
}

//provide the patch (string) to remove a resource on ubac side
func GetRmPatchString(key string, tree Tree, resourceName string) string {
	patch := GetRmPatch(key, tree, resourceName)
	patchJSON, err := json.Marshal(patch)

	if err != nil {
		fmt.Println("Error:", err)
	}

	return string(patchJSON)
}
