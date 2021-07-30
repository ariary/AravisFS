package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ariary/AravisFS/pkg/ubac"
)

func main() {
	//ls
	lsCmd := flag.NewFlagSet("ls", flag.ExitOnError)
	lsResource := lsCmd.String("resource", "", "encrypted resource path to search in encrypted fs")
	lsPath := lsCmd.String("path", "", "path to the encrypted fs location (.arafs)")

	//cat
	catCmd := flag.NewFlagSet("cat", flag.ExitOnError)
	catResource := catCmd.String("resource", "", "encrypted resource path to search in encrypted fs")
	catPath := catCmd.String("path", "", "path to the encrypted fs location")

	//tree
	treeCmd := flag.NewFlagSet("tree", flag.ExitOnError)
	treePath := treeCmd.String("path", "", "path to the encrypted fs location")

	//listen
	listenCmd := flag.NewFlagSet("listen", flag.ExitOnError)
	listenPort := listenCmd.String("port", "", "listener port")
	listenPath := listenCmd.String("path", "", "path to the encrypted fs location (.arafs)")

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
	case "cat":
		catCmd.Parse(os.Args[2:])

		//path parsing
		if *catPath == "" {
			fmt.Println("No encrypted fs precised. expected path for cat subcommand. see 'ubac help' to get help")
			os.Exit(1)
		}
		//resource parsing
		if *catResource != "" {
			ubac.PrintCat(*catResource, *catPath)
			os.Exit(0)
		} else if len(catCmd.Args()) != 0 {
			ubac.PrintCat(catCmd.Arg(0), *catPath)
		} else {
			fmt.Println("expected resource to print content with cat subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
	case "tree":
		treeCmd.Parse(os.Args[2:])

		//path parsing
		if *treePath != "" {
			ubac.PrintTree(*treePath)
		} else if len(treeCmd.Args()) != 0 {
			ubac.PrintTree(treeCmd.Arg(0))
		} else {
			fmt.Println("No encrypted fs precised. expected path for tree subcommand. see 'ubac help' to get help")
			os.Exit(1)
		}
	case "listen":
		listenCmd.Parse(os.Args[2:])
		//path parsing
		if *listenPath == "" {
			fmt.Println("No encrypted fs precised. expected path for listen subcommand. see 'ubac help' to get help")
			os.Exit(1)
		}
		//resource parsing
		if *listenPort != "" {
			port_int, err := strconv.Atoi(*listenPort)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			ubac.Listen(port_int, *listenPath)
		} else if len(listenCmd.Args()) != 0 {
			port_int, err := strconv.Atoi(listenCmd.Arg(0))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			ubac.Listen(port_int, *listenPath)
		} else {
			//use default port
			ubac.Listen(4444, *listenPath)
		}
	case "help":
		ubac.PrintHelp()
	default:
		fmt.Printf("Unknown subcommand '%v', see 'ubac help' to get help\n", os.Args[1])
		os.Exit(1)
	}
}
