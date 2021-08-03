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

	//decryptcat
	decryptcatCmd := flag.NewFlagSet("decryptcat", flag.ExitOnError)
	decryptcatOutput := decryptcatCmd.String("output", "", "output of ubac cat")
	decryptcatKey := decryptcatCmd.String("key", "", "key used for decryption")

	//decrypttree
	decrypttreeCmd := flag.NewFlagSet("decrypttree", flag.ExitOnError)
	decrypttreeOutput := decrypttreeCmd.String("output", "", "output of ubac tree")
	decrypttreeKey := decrypttreeCmd.String("key", "", "key used for decryption")

	//encryptrm
	encryptrmCmd := flag.NewFlagSet("encryptrm", flag.ExitOnError)
	encryptrmOutput := encryptrmCmd.String("tree", "", "output of ubac tree")
	encryptrmKey := encryptrmCmd.String("key", "", "key used for decryption")
	encryptrmPath := encryptrmCmd.String("path", "", "resource to rm")

	//configremote
	configremoteCmd := flag.NewFlagSet("configremote", flag.ExitOnError)
	configremotePort := configremoteCmd.String("port", "", "listener ubac port")
	configremoteHost := configremoteCmd.String("host", "", "listener ubac hostname")

	//remotels
	remotelsCmd := flag.NewFlagSet("remotels", flag.ExitOnError)
	remotelsResource := remotelsCmd.String("resource", "", "name of the resource (in clear-text)")
	remotelsKey := remotelsCmd.String("key", "", "key used for encryption/decryption")

	//remotecat
	remotecatCmd := flag.NewFlagSet("remotecat", flag.ExitOnError)
	remotecatResource := remotecatCmd.String("resource", "", "name of the resource (in clear-text)")
	remotecatKey := remotecatCmd.String("key", "", "key used for encryption/decryption")

	//remotetree
	remotetreeCmd := flag.NewFlagSet("remotetree", flag.ExitOnError)
	remotetreeKey := remotetreeCmd.String("key", "", "key used for encryption/decryption")

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
			adret.PrintLs(*decryptlsOutput, *decryptlsKey)
		} else if len(decryptlsCmd.Args()) != 0 {
			adret.PrintLs(decryptlsCmd.Arg(0), *decryptlsKey)
		} else {
			fmt.Println("expected data to decrypt for decryptls subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
	case "decryptcat":
		decryptcatCmd.Parse(os.Args[2:])

		//key parsing
		if *decryptcatKey == "" {
			fmt.Println("expected key for decryptcat subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		//output parsing
		if *decryptcatOutput != "" {
			adret.PrintCat(*decryptcatOutput, *decryptcatKey)
		} else if len(decryptcatCmd.Args()) != 0 {
			adret.PrintCat(decryptcatCmd.Arg(0), *decryptcatKey)
		} else {
			fmt.Println("expected data to decrypt for decryptcat subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
	case "decrypttree":
		decrypttreeCmd.Parse(os.Args[2:])

		//key parsing
		if *decrypttreeKey == "" {
			fmt.Println("expected key for decrypttree subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		//output parsing
		if *decrypttreeOutput != "" {
			adret.PrintTree(*decrypttreeOutput, *decrypttreeKey)
		} else if len(decrypttreeCmd.Args()) != 0 {
			adret.PrintTree(decrypttreeCmd.Arg(0), *decrypttreeKey)
		} else {
			fmt.Println("expected data to decrypt for decrypttree subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
	case "encryptrm":
		encryptrmCmd.Parse(os.Args[2:])

		//key parsing
		if *encryptrmKey == "" {
			fmt.Println("expected key for encryptrm subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		//tree parsing
		if *encryptrmOutput == "" {
			fmt.Println("expected tree output from ubac for encryptrm subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		//path parsing
		if *encryptrmPath != "" {
			adret.PrintRmPatch(*encryptrmKey, *encryptrmOutput, *encryptrmPath)
		} else if len(encryptrmCmd.Args()) != 0 {
			adret.PrintRmPatch(*encryptrmKey, *encryptrmOutput, encryptrmCmd.Arg(0))
		} else {
			fmt.Println("expected resource to rmt for encryptrm subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
	case "configremote":
		configremoteCmd.Parse(os.Args[2:])

		//port parsing
		if *configremotePort == "" {
			fmt.Println("expected port for configremote subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		//host parsing
		if *configremoteHost != "" {
			adret.ConfigRemote(*configremoteHost, *configremotePort)
		} else if len(configremoteCmd.Args()) != 0 {
			adret.ConfigRemote(configremoteCmd.Arg(0), *configremotePort)
		} else {
			fmt.Println("expected hostname for configremote subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
	case "remotels":
		remotelsCmd.Parse(os.Args[2:])
		//key parsing
		if *remotelsKey == "" {
			fmt.Println("expected key for remotels subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		//output parsing
		if *remotelsResource != "" {
			adret.PrintRemoteLs(*remotelsResource, *remotelsKey)
		} else if len(remotelsCmd.Args()) != 0 {
			adret.PrintRemoteLs(remotelsCmd.Arg(0), *remotelsKey)
		} else {
			fmt.Println("expected data to decrypt for decryptls subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
	case "remotecat":
		remotecatCmd.Parse(os.Args[2:])
		//key parsing
		if *remotecatKey == "" {
			fmt.Println("expected key for remotecat subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		//output parsing
		if *remotecatResource != "" {
			adret.PrintRemoteCat(*remotecatResource, *remotecatKey)
		} else if len(remotecatCmd.Args()) != 0 {
			adret.PrintRemoteCat(remotecatCmd.Arg(0), *remotecatKey)
		} else {
			fmt.Println("expected data to decrypt for decryptcat subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
	case "remotetree":
		remotetreeCmd.Parse(os.Args[2:])
		//key parsing
		if *remotetreeKey == "" {
			fmt.Println("expected key for remotetree subcommand. see 'adret help' to get help")
			os.Exit(1)
		}
		adret.RemoteTree(*remotetreeKey)
	case "help":
		adret.PrintHelp()
	default:
		fmt.Printf("Unknown subcommand '%v', see 'adret help' to get help\n", os.Args[1])
		os.Exit(1)
	}

}
