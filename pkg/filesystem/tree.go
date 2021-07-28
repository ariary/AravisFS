package filesystem

type Node struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Tree struct {
	List []Node
}

func CreateNode(name string, resourceType string) Node {

	n := &Node{
		Name: name,
		Type: resourceType}
	return *n
}

func (nodes *Tree) AddNode(n Node) []Node {
	nodes.List = append(nodes.List, n)
	return nodes.List
}
