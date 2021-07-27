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
// │   ├── bullit_conf
// │   │   ├── brain.txt
// │   │   ├── bullit_conf.yml.j2
// │   │   └── bullit.yml
// │   ├── cat.yaml
// │   ├── kube-hunter.yaml
// │   ├── report.j2
// │   ├── result.json
// │   ├── run.sh
// │   └── toto.log
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

//Print node name without prefix. If it is the last element of a directory it print it with a special character behind
func specialPrint(name string, last bool) {
	//calcul deepth (avec /)
	// print consequently filepath.Base((node.Name)
	depth := strings.Count(name, "/")
	if depth == 0 {
		//Base is not mandatory
		fmt.Println(filepath.Base(name))
	} else {
		if last {
			fmt.Println("|" + strings.Repeat("  ", depth) + "├── " + filepath.Base(name))
		} else {
			fmt.Println("|" + strings.Repeat("  ", depth) + "└── " + filepath.Base(name))
		}

	}

}

func getTreeStructFromResourcesMap(resources map[string]string) Tree {
	var tree Tree
	var nodeTmp Node
	for name, nodeType := range resources {
		nodeTmp = createNode(name, nodeType, filepath.Dir(name))
		tree.Nodes = append(tree.Nodes, nodeTmp)
	}
	return tree
}

func getNodeByName(name string, tree Tree) (node Node, err error) {

	for i := range tree.Nodes {
		if tree.Nodes[i].Name == name {
			node = tree.Nodes[i]
			return node, nil
		}
	}
	err = errors.New(fmt.Sprintf("getNodeByName: Node % v doesn't exist", name))
	return node, err
}

func PrintAll(tree Tree, root string) {
	rootNode, err := getNodeByName(root, tree)
	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}
	PrintNode(tree.Nodes, rootNode, false)
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

// Print all node with  a specific prefix ie node.Dir == prefix
// Retrieve a list of all node
// Theni terate over the list when we arrive at last PrintNode(node.Name,true)
func PrintNodeWithprefix(prefix string, nodes []Node) {

	//retrieve node with this prefix
	nodeWithPrefix := getNodeWithPrefix(prefix, nodes)
	//iterate to print the last one differently
	for i := range nodeWithPrefix {
		last := (len(nodeWithPrefix)-1 == i)
		specialPrint(nodeWithPrefix[i], last)
	}

}

func PrintNode(nodes []Node, node Node, last bool) {
	if node.Type == "file" {
		specialPrint(node.Name, last)
	} else if node.Type == "directory" {
		specialPrint(node.Name, last)
		PrintNodeWithprefix(node.Name, nodes)
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
	resources["test/ansible/cat.yaml"] = "file"
	resources["test/ansible/kube-hunter.yaml"] = "file"
	resources["test/ansible/bullit_conf"] = "directory"
	resources["test/ansible/bullit_conf/bullit.yml"] = "file"
	resources["test/ansible/bullit_conf/bullit_conf.yml.j2"] = "file"
	resources["test/ansible/bullit_conf/emptydir"] = "directory"
	resources["test/ansible/bullit_conf/brain.txt"] = "file"
	resources["test/ansible/result.json"] = "file"
	resources["test/ansible/report.j2"] = "file"
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
	specialPrint("test/ansible/bullit_conf/brain.txt", false)
	specialPrint("test/pentest/ftp-server.py", false)
	specialPrint("test/ansible/bullit_conf/brain.txt", true)
	specialPrint("test", true)
	specialPrint("test/ansible/toto.log", true)

	// general test
	fmt.Println()
	fmt.Println("Tree test")
	PrintAll(tree, "test")
}
