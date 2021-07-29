package ubac

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

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

// Provide all the resources within the fs in form of node list which represent the tree:
// each node  = (resourcename (encrypted), resource.Type)
// it also could have returned the content of each file, but it is like returning the whole fs (which is not a good idea)
// Node list is in json format
func GetTreeFromFS(filename string) string {
	resourcesList, err := GetResourceList(filename)
	if err != nil {
		log.Fatal(err)
	}
	resources := resourcesList.List

	nl := []Node{}
	tree := Tree{nl}

	// Tree construction
	for i := range resources {
		//Construct node and add it to the tree
		name := base64.StdEncoding.EncodeToString(resources[i].Name)
		node := CreateNode(name, resources[i].Type)
		tree.AddNode(node)
	}

	// Tree to JSON
	treeJSON, _ := json.Marshal(tree)

	return string(treeJSON)
}

func PrintTree(filename string) {
	treeJSON := GetTreeFromFS(filename)

	fmt.Println(treeJSON)
}
