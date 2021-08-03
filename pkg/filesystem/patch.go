package filesystem

type Patch struct {
	AddList    []Resource        `json:"to_add"`
	RemoveList []string          `json:"to_remove"`
	ChangeMap  map[string]string `json:"to_change"` //
}

func CreatePatch(addList []Resource, removeList []string, changeMap map[string]string) Patch {

	p := &Patch{
		AddList:    addList,
		RemoveList: removeList,
		ChangeMap:  changeMap}
	return *p
}
