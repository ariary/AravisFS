package ubac

import "fmt"

func PrintHelpMessage() {
	fmt.Println("help: get help for ubac utility and example")
}

func PrintLsMessage() {
	fmt.Println(("ls: like ls in encrypted fs"))
	fmt.Println(("\tuse: ubac ls -path=<encryptedfs>.arfs <resource>"))
	fmt.Println(("\tparameters required: key (-path) and path (-resource). Resource is obtained by adret utility"))
	fmt.Println(("\texample: adret ls -path encrypted.arfs \"AAAAAAAAAAAAAAAAsFt3LbDTrKVIllwFtQLUTuE=\""))
}

func PrintHelp() {
	fmt.Println("ubac utility is used to interact with encrypted fs (provided by adret utility). As every data manipulate is encrypted you can use it in a non-trusted environement")
	fmt.Println()
	fmt.Println("Available commands:")
	PrintHelpMessage()
	fmt.Println()
	PrintLsMessage()
	fmt.Println()

}
