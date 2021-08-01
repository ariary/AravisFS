package adret

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ariary/AravisFS/pkg/encrypt"
	"github.com/ariary/AravisFS/pkg/remote"
	"github.com/ariary/AravisFS/pkg/ubac"
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
	err = errors.New(fmt.Sprintf("getNodeByName: Node %s doesn't exist", name))
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

// Take the Tree from ubac(JSON format) as input and return it in a struct that help to work with it
func GetTreeStructFromTreeJson(treeJSON string, key string) (tree Tree) {
	var ubacTree ubac.Tree
	json.Unmarshal([]byte(treeJSON), &ubacTree)
	ubacNodes := ubacTree.List

	nodesMap := make(map[string]string)
	// fill nodesMap: key = name, value = type
	for i := 0; i < len(ubacNodes); i++ {
		//don't forget to decrypt it
		name := string(encrypt.DecryptStringFromUbac(ubacNodes[i].Name, key))
		resourceType := ubacNodes[i].Type
		nodesMap[name] = resourceType
	}

	tree = GetTreeStructFromResourcesMap(nodesMap)
	return tree
}

// Return all nodes with specific prefix/parent directory (ie prefix == node.Parent)
// It enables us to retrieve all nodes irectly under a specified one (with depth=depth_nodes+1)
func GetNodesWithPrefix(prefix string, nodes []Node) []string {
	var nodesWithPrefix []string
	for i := range nodes {
		if nodes[i].Parent == prefix {
			nodesWithPrefix = append(nodesWithPrefix, nodes[i].Name)
		}
	}
	return nodesWithPrefix
}

// Return all nodes under the prefix/parent directory (ie node.Parent begin w/ prefix)
// It enables us to retrieve all nodes under a specified one
func GetNodesUnder(prefix string, nodes []Node) []string {
	var nodesUnder []string
	for i := range nodes {
		if strings.HasPrefix(nodes[i].Parent, prefix) {
			nodesUnder = append(nodesUnder, nodes[i].Name)
		}
	}
	return nodesUnder
}

// Function wich aim to imitate the tree command output
// It prints the node name without prefix with the right indentation compare to its position in the Tree.
// If it is the last element of a directory it print it with a special character behind
// If it is in a directory which is "the last element" it print one "|" less
func specialPrint(name string, last bool, inlast bool) {
	// compute depth (using / counter)
	depth := strings.Count(name, "/")

	output := ""

	//determine the appropriate characters
	var lastCharacter string
	var inlastCharacter string

	tab := "   " //tabulate character
	if last {
		lastCharacter = "└── "
	} else {
		lastCharacter = "├── "
	}

	if inlast {
		inlastCharacter = ""
	} else {
		inlastCharacter = "|"
	}

	if depth != 0 {
		//if not root path
		if depth == 1 {
			output += lastCharacter
		}
		if depth == 2 {
			output += inlastCharacter + tab + lastCharacter
		}
		if depth > 2 {
			output += strings.Repeat("|"+tab, depth-2) + inlastCharacter + tab + lastCharacter
		}

	}
	output += filepath.Base(name) //whatever happens

	fmt.Println(output)
}

// (recursive) Print the tree under the Node (except the node itself)
// Retrieve all node under if it is a directory and print it, nothing if it is a file
func PrintNode(nodes []Node, node Node, last bool, inlast bool) {
	if node.Type == "file" {
		// the Node has already been printed
	} else if node.Type == "directory" {
		// Recursivity: print node under this node
		// Print all node with  a specific prefix ie node.Dir == prefix
		// Retrieve a list of all node
		// Then iterate over the list when we arrive at last PrintNode(node.Name,true)
		nodeWithPrefix := GetNodesWithPrefix(node.Name, nodes)

		inlast = last //if we are in last we must now call PrintNode with inlast at true, and conversely

		//iterate to print the last one differently
		for i := range nodeWithPrefix {
			last := (len(nodeWithPrefix)-1 == i)
			specialPrint(nodeWithPrefix[i], last, inlast)
			//!recursivity
			node, err := GetNodeByName(nodeWithPrefix[i], nodes)
			if err != nil {
				log.SetFlags(0)
				log.Fatal(err)
			}
			PrintNode(nodes, node, last, inlast)
		}
	} else {
		log.Fatal("Node/Resource with undefined type")
	}
}

// Print the Tree struct (input) in a fashion way (as tree command would do.. I hope)
func PrintTree(treeJSON string, key string) {
	tree := GetTreeStructFromTreeJson(treeJSON, key)
	if len(tree.Nodes) == 0 {
		log.SetFlags(0)
		log.Fatal("PrintTree: Failed to convert JSON to Tree structure")
	}
	rootSlice := GetNodesWithPrefix(".", tree.Nodes) // Normally only one
	if len(rootSlice) == 0 {
		log.SetFlags(0)
		log.Fatal("PrintTree: Could not find root node in Tree structure")
	} else if len(rootSlice) > 1 {
		fmt.Println("WARNING: found multiple root nodes")
	}
	root := rootSlice[0]
	rootNode, err := GetNodeByName(root, tree.Nodes)
	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}

	specialPrint(root, true, false)               //bool have no real impact for this case
	PrintNode(tree.Nodes, rootNode, false, false) //bool have no real impact for this case
}

// Perform tree on a remote listening ubac (proxing to encrypted fs)
// First craft the request, send it (the request instruct ubac to perform a tree)
// take the reponse and decrypt it
func RemoteTree(key string) {
	url := os.Getenv("REMOTE_UBAC_URL")
	if url == "" {
		fmt.Println("Configure REMOTE_UBAC_URL envar with `adret configremote` before launching remotels. see `adret help`")
		os.Exit(1)
	}
	endpoint := url + "tree"

	bodyRes := remote.SendReadrequest("", endpoint)
	//decrypt the reponse to show cat result
	PrintTree(bodyRes, key)
}
