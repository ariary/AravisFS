package filesystem

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
)

// /!\ do not confuse with the Node & Tree struct of resources package
type Node struct {
	Name string
	Type string
	Dir  string
}

type Tree struct {
	Nodes []Node
}

// Create a node from its name, its type and its parent directory
func createNode(name string, nodeType string, dir string) Node {

	n := &Node{
		Name: name,
		Type: nodeType,
		Dir:  dir}
	return *n
}

// Get a Node by providing its name, an error is thrown if the Node isn't found
func getNodeByName(name string, nodes []Node) (node Node, err error) {

	for i := range nodes {
		if nodes[i].Name == name {
			node = nodes[i]
			return node, nil
		}
	}
	err = errors.New(fmt.Sprintf("getNodeByName: Node % v doesn't exist", name))
	return node, err
}

// Take the Tree (JSON format, from ubac) as input and return it in a struct that help to work with it
func GetTreeStructFromTreeJson(treeJSON string, key string) (tree Tree, err error) {

}

func getTreeStructFromResourcesMap(resources map[string]string) Tree {
	var tree Tree
	var nodeTmp Node

	// Browse map alphabetically
	// first contruct key list in alphabtical order
	keys := make([]string, 0)
	for k, _ := range resources {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, name := range keys {
		nodeTmp = createNode(name, resources[name], filepath.Dir(name))
		tree.Nodes = append(tree.Nodes, nodeTmp)
	}

	return tree
}
