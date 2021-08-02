package main

import (
	"fmt"

	"github.com/ariary/AravisFS/pkg/adret"
)

// test
// ├── ansible
// │   ├── bullit_conf					// "\t "
// │   │   ├── brain.txt
// │   │   ├── bullit_conf.yml.j2
// │   │   ├── fuldir
// │	 │   |	 ├── toto.c
// │   │   |	 └── bullit.yml
// │   │   └── bullit.yml
// │   ├── cat.yaml
// │   ├── kube-hunter.yaml
// │   ├── report.j2
// │   ├── result.json
// │   ├── run.sh
// │   ├── toto.log
// │   ├── slice
// │   |	 ├── slice2
// │   |	 |	 └── slice3
// │   |	 └── slice2bis
// │	 │    	 ├── toto.c
// │   |       └── slice2bis.txt
// ├── go
// │   ├── hello-world
// │   ├── hello-world.go
// │   └── slice.go
// └── pentest
//     └── ftp-server.py

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
	tree := adret.GetTreeStructFromResourcesMap(resources)
	//print tree struct test
	// treeJSON, _ := json.Marshal(tree)
	// fmt.Println(string(treeJSON))

	//isDir
	dir := adret.IsDir("test/pentest", tree.Nodes)
	notdir := adret.IsDir("test/ansible/slice/slice2bis/slice2bis.txt", tree.Nodes)
	fmt.Println("dir:", dir, "not dir:", notdir)

	//GetRmPatch
	patch := adret.GetRmPatch("toto", tree, "test/ansible/bullit_conf")
	fmt.Println("Remove list:", patch.RemoveList)
	fmt.Println("Change map:", patch.ChangeMap)
	fmt.Println("Add list:", patch.AddList)

}
