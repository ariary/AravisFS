package adret

//provide the patch to remove a resource on ubac side
// the patch is a json string with 3 arrays: to_add,to_delete and to_change
func GetRmPatch(key string, tree Tree, resourceParentName string, resourceName string) Patch {

	var removeList []string
	changeMap := make(map[string]string)

	//add resource to remove list
	removeList = append(removeList, resourceName)

	//modify parent content (remove resource from it)

	if IsDir(resourceName, tree.Nodes) {
		//add all resource under directory to remove list
		removeList = append(removeList, GetNodesUnder(resourceName, tree.Nodes)...)
	}

	return createPatch(nil, removeList, changeMap)
}
