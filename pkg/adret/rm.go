package adret

import (
	"path"
)

//provide the patch to remove a resource on ubac side
// the patch is a json string with 3 arrays: to_add,to_delete and to_change
//all the info within structure are encrypted after to be put in patch
func GetRmPatch(key string, tree Tree, resourceName string) Patch {
	resourceName = path.Clean(resourceName)
	var removeList []string
	changeMap := make(map[string]string)

	//add the resource to remove list
	removeList = append(removeList, resourceName)

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
		// newParentContentByte:= encrypt.EncryptString(newParentContent, key)
		// changeMap[parentName] = newParentContentByte
		changeMap[parentName] = newParentContent
	}

	//delete all resource under if it is a directory
	if IsDir(resourceName, tree.Nodes) {
		removeList = append(removeList, GetDescendantNodes(resourceName, tree.Nodes)...)
	}

	return createPatch(nil, removeList, changeMap)
}
