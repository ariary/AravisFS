package filesystem

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
)

// /!\ do not confuse with the Node & Tree struct of ubac package
type Node struct {
	Name   string
	Type   string
	Parent string
}

type Tree struct {
	Nodes []Node
}

// Create a node from its name, its type and its parent directory
func CreateNode(name string, nodeType string, dir string) Node {

	n := &Node{
		Name:   name,
		Type:   nodeType,
		Parent: dir}
	return *n
}

// Get a Node by providing its name, an error is thrown if the Node isn't found
func GetNodeByName(name string, nodes []Node) (node Node, err error) {

	for i := range nodes {
		if nodes[i].Name == name {
			node = nodes[i]
			return node, nil
		}
	}
	err = errors.New(fmt.Sprintf("getNodeByName: Node % v doesn't exist", name))
	return node, err
}

// Get tree structure from map. Map: key= resource name and value= resource type
func GetTreeStructFromResourcesMap(resources map[string]string) Tree {
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
		nodeTmp = CreateNode(name, resources[name], filepath.Dir(name))
		tree.Nodes = append(tree.Nodes, nodeTmp)
	}

	return tree
}

// Return all node with specific prefix/parent directory (ie prefix == node.Parent)
func GetNodesWithPrefix(prefix string, nodes []Node) []string {
	var nodesWithPrefix []string
	for i := range nodes {
		if nodes[i].Parent == prefix {
			nodesWithPrefix = append(nodesWithPrefix, nodes[i].Name)
		}
	}
	return nodesWithPrefix
}