package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ariary/AravisFS/pkg/ubac"
)

func main() {
	//ls
	lsCmd := flag.NewFlagSet("ls", flag.ExitOnError)
	lsResource := lsCmd.String("resource", "", "encrypted resource path to search in encrypted fs")
	lsPath := lsCmd.String("path", "", "path to encrypted fs (.arafs)")

	if len(os.Args) < 2 {
		fmt.Println("expected subcommands see 'ubac help' to get help")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "ls":
		lsCmd.Parse(os.Args[2:])

		//path parsing
		if *lsPath == "" {
			fmt.Println("No encrypted fs precised. expected path for ls subcommand. see 'ubac help' to get help")
			os.Exit(1)
		}
		//resource parsing
		if *lsResource != "" {
			ubac.PrintLs(*lsResource, *lsPath)
			os.Exit(0)
		} else if len(lsCmd.Args()) != 0 {
			ubac.PrintLs(lsCmd.Arg(0), *lsPath)
		} else {
			fmt.Println("expected resource to list for ls subcommand. see 'adret help' to get help")
			os.Exit(1)
		}

	case "help":
		ubac.PrintHelp()
	default:
		fmt.Printf("Unknown subcommand '%v', see 'ubac help' to get help\n", os.Args[1])
		os.Exit(1)
	}
	// resource := ubac.GetResourceInFS("test/mytestfolder/tata/atat", "./test/arafs/ceciestlav1_filename_unencrypted.arafs")
	// fmt.Println(resource.Name)

	//ls test
	// ubac.PrintLs("AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA235ttcqjgwlrfQDRy+r2o07a", "test/arafs/encrypted.arafs")

}
