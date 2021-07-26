package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ariary/AravisFS/pkg/adret"
	"github.com/ariary/AravisFS/pkg/encrypt"
	"github.com/ariary/AravisFS/pkg/filesystem"
)

func main() {
	//darkenpath
	darkenpathCmd := flag.NewFlagSet("darkenpath", flag.ExitOnError)
	darkenpathPath := darkenpathCmd.String("path", "", "filepath to darkened (encrypt)")
	darkenpathKey := darkenpathCmd.String("key", "", "key used for encryption")

	//encryptfs
	encryptfsCmd := flag.NewFlagSet("encryptfs", flag.ExitOnError)
	encryptfsPath := encryptfsCmd.String("path", "", "directory to encrypt")
	encryptfsKey := encryptfsCmd.String("key", "", "key used for encryption")

	//decryptLs
	decryptlsCmd := flag.NewFlagSet("decryptls", flag.ExitOnError)
	decryptlsOutput := decryptlsCmd.String("output", "", "output of ubac ls")
	decryptlsKey := decryptlsCmd.String("key", "", "key used for decryption")

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
			result := encrypt.DarkenPath(darkenpathCmd.Arg(0), *darkenpathKey)
			fmt.Println(result)
		} else {
			fmt.Println("expected path for darkenpath subcommand. see 'adret help' to get help")
			os.Exit(1)
		}

	case "encryptfs":
		encryptfsCmd.Parse(os.Args[2:])

		//key parsing
		if *encryptfsKey == "" {
			fmt.Println("expected key for encryptfs subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		//path parsing
		if *encryptfsPath != "" {
			filesystem.CreateAravisFS(*encryptfsPath, *encryptfsKey)
			fmt.Println("done! Encrypted fs saved in encrypted.arafs")
		} else if len(encryptfsCmd.Args()) != 0 {
			filesystem.CreateAravisFS(encryptfsCmd.Arg(0), *encryptfsKey)
			fmt.Println("done! Encrypted fs saved in encrypted.arafs")
		} else {
			fmt.Println("expected path for encrypted subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
	case "decryptls":
		decryptlsCmd.Parse(os.Args[2:])

		//key parsing
		if *decryptlsKey == "" {
			fmt.Println("expected key for decryptls subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		//output parsing
		if *decryptlsOutput != "" {
			adret.PrintLS(*decryptlsOutput, *decryptlsKey)
		} else if len(decryptlsCmd.Args()) != 0 {
			adret.PrintLS(decryptlsCmd.Arg(0), *decryptlsKey)
		} else {
			fmt.Println("expected data to decrypt for decryptls subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
	case "help":
		adret.PrintHelp()
	default:
		fmt.Printf("Unknown subcommand '%v', see 'adret help' to get help\n", os.Args[1])
		os.Exit(1)
	}

	// test PrintLS
	// lsUbacResult := "directory:AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le21X9+Fky1Fb98v61k+DQJivbwJosBKJ8FSD4YitHoo9GZf40l3HLHGTDjc="
	// adret.PrintLS(lsUbacResult, key)
}
