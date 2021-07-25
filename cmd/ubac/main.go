package main

import (
	"fmt"

	"github.com/ariary/AravisFS/pkg/ubac"
)

func main() {
	fmt.Println("Hello Aravis!")
	// resource := ubac.GetResourceInFS("test/mytestfolder/tata/atat", "./test/arafs/ceciestlav1_filename_unencrypted.arafs")
	// fmt.Println(resource.Name)

	//ls test
	ubac.PrintLs("AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA235ttcqjgwlrfQDRy+r2o07a", "test/arafs/encrypted.arafs")

}
