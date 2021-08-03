package filesystem

type Resource struct {
	Name    []byte `json:"name"`
	Type    string `json:"type"`
	Content []byte `json:"content"`
}

const DIRECTORY = "directory"
const FILE = "file"

type ResourceList struct {
	List []Resource
}

func CreateResource(filename []byte, resourceType string, content []byte) Resource {

	r := &Resource{
		Name:    filename,
		Type:    resourceType,
		Content: content}
	return *r
}

func (resources *ResourceList) AddResource(r Resource) []Resource {
	resources.List = append(resources.List, r)
	return resources.List
}

func (resources *ResourceList) RemoveResourceFromName(resourceName string) []Resource {
	for i, resource := range resources.List {
		if string(resource.Name) == resourceName {
			resources.List = append(resources.List[:i], resources.List[i+1:]...)
			break
		}
	}
	return resources.List
}

func (resources *ResourceList) ChangeResourceContentFromName(resourceName string, content string) []Resource {
	for i, resource := range resources.List {
		if string(resource.Name) == resourceName {
			resources.List[i].Content = []byte(content) // ou resource.content
			break
		}
	}
	return resources.List
}
