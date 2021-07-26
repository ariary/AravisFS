package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ariary/AravisFS/cmd/options"
	"github.com/ariary/AravisFS/pkg/encrypt"
)

func main() {
	//darkenpath
	darkenpathCmd := flag.NewFlagSet("darkenpath", flag.ExitOnError)
	darkenpathPath := darkenpathCmd.String("path", "", "filepath to darkened (encrypt)")
	darkenpathKey := darkenpathCmd.String("key", "", "key used for encryption")

	if len(os.Args) < 2 {
		fmt.Println("expected subcommands see 'adret help' to get help")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "darkenpath":
		darkenpathCmd.Parse(os.Args[2:])

		//key parsing
		if *darkenpathKey == "" {
			fmt.Println("expected key for darkenpath subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		//path parsing
		if *darkenpathPath != "" {
			result := encrypt.DarkenPath(*darkenpathPath, *darkenpathKey)
			fmt.Println(result)
		} else if len(darkenpathCmd.Args()) != 0 {
			darkenpathCmd := flag.NewFlagSet("darkenpath", flag.ExitOnError)
			result := encrypt.DarkenPath(darkenpathCmd.Arg(0), *darkenpathKey)
			fmt.Println(result)
		} else {
			fmt.Println("expected path for darkenpath subcommand. see 'adret help' to get help")
			os.Exit(1)
		}

	case "help":
		options.PrintHelp()
	default:
		fmt.Printf("Unknown subcommand '%v', see 'adret help' to get help\n", os.Args[1])
		os.Exit(1)
	}

	// test PrintLS
	// lsUbacResult := "directory:AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le21X9+Fky1Fb98v61k+DQJivbwJosBKJ8FSD4YitHoo9GZf40l3HLHGTDjc="
	// adret.PrintLS(lsUbacResult, key)
}
