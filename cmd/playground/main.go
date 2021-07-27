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
	"fmt"
	"log"
	"path/filepath"
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
	for name, nodeType := range resources {
		nodeTmp = createNode(name, nodeType, filepath.Dir(name))
		tree.Nodes = append(tree.Nodes, nodeTmp)
	}
	return tree
}

func PrintAll() {
	PrintNode("")
}
func PrintNode(node Node) {
	if node.Type == "file" {
		print(filepath.Base((node.Name)))
	} else if node.Type == "directory" {
		print(filepath.Base((node.Name)))
		PrintNode(node.Name)
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

}
