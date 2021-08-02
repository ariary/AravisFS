package adret

import (
	"path"
)

//provide the patch to remove a resource on ubac side
// the patch is a json string with 3 arrays: to_add,to_delete and to_change
//all the info within structure are encrypted after to be put in patch
func GetRmPatch(key string, tree Tree, resourceName string) Patch {

	var removeList []string
	changeMap := make(map[string]string)

	//add resource to remove list
	removeList = append(removeList, resourceName)

	//modify parent content (remove resource from it)
	parentName := path.Dir(resourceName)
	parentContent := GetChildrenNodes(parentName, tree.Nodes)
	//remove resource from content trick
	for i, resource := range parentContent {
		if resource == resourceName {
			parentContent = append(parentContent[:i], parentContent[i+1:]...)
			break
		}
	}
	//translate content in arafs form
	var newParentContent string
	for _, resource := range parentContent {
		//Create
		if newParentContent == "" {
			newParentContent = resource
		} else {
			newParentContent = newParentContent + "\\" + resource
		}
	}
	// newParentContentByte:= encrypt.EncryptString(newParentContent, key)
	// changeMap[parentName] = newParentContentByte
	changeMap[parentName] = newParentContent

	if IsDir(resourceName, tree.Nodes) {
		//add all resource under directory to remove list
		removeList = append(removeList, GetDescendantNodes(resourceName, tree.Nodes)...)
	}

	return createPatch(nil, removeList, changeMap)
}
