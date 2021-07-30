package remote

type BodyLs struct {
	ResourceName string `json:"name"`
}

func CreateBodyLs(resourceName string) BodyLs {

	b := &BodyLs{
		ResourceName: resourceName}
	return *b
}
