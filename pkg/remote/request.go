package remote

type BodyRead struct {
	ResourceName string `json:"name"`
}

func CreateBodyRead(resourceName string) BodyRead {

	b := &BodyRead{
		ResourceName: resourceName}
	return *b
}
