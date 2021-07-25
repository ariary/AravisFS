package filesystem

type Resource struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content []byte `json:"content"`
}

const DIRECTORY = "directory"
const FILE = "file"

type ResourceList struct {
	List []Resource
}

func createResource(filename string, resourceType string, content []byte) Resource {

	r := &Resource{
		Name:    filename,
		Type:    resourceType,
		Content: content}
	return *r
}

func (resources *ResourceList) Addresource(r Resource) []Resource {
	resources.List = append(resources.List, r)
	return resources.List
}
