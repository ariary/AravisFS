package adret

import "fmt"

func PrintHelpMessage() {
	fmt.Println("help: get help for adret utility and example")
}

func PrintDarkenpathMessage() {
	fmt.Println(("darkenpath: use it to encrypt a path with the key"))
	fmt.Println(("\tuse: adret darkenpath -key=<secret> <path>"))
	fmt.Println(("\tparameters required: key (-key) and path (-path"))
	fmt.Println(("\texample: adret darkenpath -key \"toto\" \"test/toto/titi\""))
}

func PrintEncryptfsMessage() {
	fmt.Println(("encryptfs: use it to encrypt a directory with the key. It will create a encrypted fs .arafs"))
	fmt.Println(("LIMITATION: encrypt a directory in the current one avoid ie  \"..\" etc in path parameters"))
	fmt.Println(("\tuse: adret encryptfs -key=<secret> <path>"))
	fmt.Println(("\tparameters required: key (-key) and path (-path"))
	fmt.Println(("\texample: adret encryptfs -key \"toto\" \"test/toto/titi\""))
	fmt.Println()
}

func PrintHelp() {
	fmt.Println("adret utility is used to perform FS encrytion and decrypt data from fs encrypted (from ubac utility)")
	fmt.Println("Available commands:")
	PrintHelpMessage()
	fmt.Println()
	PrintDarkenpathMessage()
	fmt.Println()
	PrintEncryptfsMessage()
	fmt.Println()

}
