package options

import "fmt"

func PrintHelpMessage() {
	fmt.Println("help: get help for adret utility and example")
}

func PrintDarkenpathMessage() {
	fmt.Println(("darkenpath: use it to encrypt a path with the key"))
	fmt.Println(("\tuse: adret darken path -key=<secret> <path>"))
	fmt.Println(("\tparameters required: key (-key) and path (-path"))
	fmt.Println(("\texample: adret darkenpath -key \"toto\" \"test/toto/titi\""))
}

func PrintHelp() {
	fmt.Println("adret utility is used to perform FS encrytion and decrypt data from fs encrypted (from ubac utility)")
	fmt.Println("Available commands:")
	PrintHelpMessage()
	PrintDarkenpathMessage()

}
