package remote

type BodyLs struct {
	ResourceName string `json:"name"`
}

func createBodyLs(path string, resourceName string) BodyLs {

	b := &BodyLs{
		ResourceName: resourceName}
	return *b
}
