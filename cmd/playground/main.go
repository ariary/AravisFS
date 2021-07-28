// test
// test/ansible
// test/ansible/toto.log
// test/ansible/run.sh
// test/ansible/cat.yaml
// test/ansible/kube-hunter.yaml
// test/ansible/bullit_conf
// test/ansible/bullit_conf/bullit.yml
// test/ansible/bullit_conf/bullit_conf.yml.j2
// test/ansible/bullit_conf/brain.txt
// test/ansible/result.json
// test/ansible/report.j2
// test/go
// test/go/slice.go
// test/go/hello-world.go
// test/go/hello-world
// test/pentest
// test/pentest/ftp-server.py

// output wanted:
// test
// ├── ansible
// │   ├── bullit_conf					// "\t "
// │   │   ├── brain.txt
// │   │   ├── bullit_conf.yml.j2
// │   │   ├── fuldir
// │	 │   |	 ├── toto.c
// │   │   |	 └── bullit.yml
// │   │   └── bullit.yml
// │   ├── cat.yaml
// │   ├── kube-hunter.yaml
// │   ├── report.j2
// │   ├── result.json
// │   ├── run.sh
// │   ├── toto.log
// │   ├── slice
// │   |	 ├── slice2
// │   |	 |	 └── slice3
// │   |	 └── slice2bis
// │	 │    	 ├── toto.c
// │   |       └── slice2bis.txt
// ├── go
// │   ├── hello-world
// │   ├── hello-world.go
// │   └── slice.go
// └── pentest
//     └── ftp-server.py

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

type Node struct {
	Name string
	Type string
	Dir  string
}

type Tree struct {
	Nodes []Node
}

func createNode(name string, nodeType string, dir string) Node {

	n := &Node{
		Name: name,
		Type: nodeType,
		Dir:  dir}
	return *n
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

// Print node name without prefix with the right indentation its position in the Tree.
// If it is the last element of a directory it print it with a special character behind
// If it is in a directory which is "the last element" it print one "|" less
func specialPrint(name string, last bool, inlast bool) {
	// compute depth (using / counter)
	// print consequently filepath.Base((node.Name)
	depth := strings.Count(name, "/")
	//fmt.Println(name + ":" + strconv.Itoa(depth) + "," + strconv.FormatBool(inlast))
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

func PrintAll(tree Tree, root string) {
	rootNode, err := getNodeByName(root, tree.Nodes)
	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}
	specialPrint(root, true, false)
	PrintNode(tree.Nodes, rootNode, false, false)
}

func getNodeWithPrefix(prefix string, nodes []Node) []string {
	var nodeWithPrefix []string
	for i := range nodes {
		if nodes[i].Dir == prefix {
			nodeWithPrefix = append(nodeWithPrefix, nodes[i].Name)
		}
	}
	return nodeWithPrefix
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
		nodeWithPrefix := getNodeWithPrefix(node.Name, nodes)

		inlast = last //if we are in last we must now call PrintNode with inlast at true, and conversely

		//iterate to print the last one differently
		for i := range nodeWithPrefix {
			last := (len(nodeWithPrefix)-1 == i)
			specialPrint(nodeWithPrefix[i], last, inlast)
			//!recursivity
			node, err := getNodeByName(nodeWithPrefix[i], nodes)
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

func main() {
	resources := make(map[string]string)
	resources["test"] = "directory"
	resources["test/ansible"] = "directory"
	resources["test/ansible/toto.log"] = "file"
	resources["test/ansible/run.sh"] = "file"
	resources["test/ansible/bullit_conf/notemptydir"] = "directory"
	resources["test/ansible/bullit_conf/notemptydir/brain.txt"] = "file"
	resources["test/ansible/bullit_conf/notemptydir/emptydir"] = "directory"
	resources["test/ansible/cat.yaml"] = "file"
	resources["test/ansible/kube-hunter.yaml"] = "file"
	resources["test/ansible/bullit_conf"] = "directory"
	resources["test/ansible/bullit_conf/bullit.yml"] = "file"
	resources["test/ansible/bullit_conf/bullit_conf.yml.j2"] = "file"
	resources["test/ansible/bullit_conf/emptydir"] = "directory"
	resources["test/ansible/bullit_conf/brain.txt"] = "file"
	resources["test/ansible/result.json"] = "file"
	resources["test/ansible/report.j2"] = "file"
	resources["test/ansible/slice"] = "directory"
	resources["test/ansible/slice/slice2"] = "directory"
	resources["test/ansible/slice/slice2/slice3"] = "directory"
	resources["test/ansible/slice/slice2bis"] = "directory"
	resources["test/ansible/slice/slice2bis/toto.c"] = "file"
	resources["test/ansible/slice/slice2bis/slice2bis.txt"] = "file"
	resources["test/ansible/bullit_conf/hello_world"] = "file"
	resources["test/go"] = "directory"
	resources["test/go/slice.go"] = "file"
	resources["test/go/hello-world.go"] = "file"
	resources["test/pentest"] = "directory"
	resources["test/pentest/ftp-server.py"] = "file"

	fmt.Println(resources["test/ansible"])
	tree := getTreeStructFromResourcesMap(resources)
	//print tree struct test
	treeJSON, _ := json.Marshal(tree)
	fmt.Println(string(treeJSON))

	//test specialPrint
	// specialPrint("test/ansible/bullit_conf/brain.txt", false)
	// specialPrint("test/pentest/ftp-server.py", false)
	// specialPrint("test/ansible/bullit_conf/brain.txt", true)
	// specialPrint("test", true)
	// specialPrint("test/ansible/toto.log", true)

	// general test
	fmt.Println()
	fmt.Println("Tree test")
	PrintAll(tree, "test")
}
