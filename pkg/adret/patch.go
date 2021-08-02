package adret

type Patch struct {
	AddList    []string          `json:"to_add"`
	RemoveList []string          `json:"to_remove"`
	ChangeMap  map[string]string `json:"to_change"` //
}

func createPatch(addList []string, removeList []string, changeMap map[string]string) Patch {

	p := &Patch{
		AddList:    addList,
		RemoveList: removeList,
		ChangeMap:  changeMap}
	return *p
}
