package ubac

import "fmt"

func PrintCommandMessage(f func()) {
	f()
	fmt.Println()
}

func PrintHelpMessage() {
	fmt.Println("help: get help for ubac utility and example")
}

func PrintLsMessage() {
	fmt.Println(("ls: like ls in encrypted fs"))
	fmt.Println(("\tuse: ubac ls -path=<encryptedfs>.arfs <resource>"))
	fmt.Println(("\tparameters required: path (-path) and resource (-resource). Resource is obtained by adret utility"))
	fmt.Println(("\texample: ubac ls -path encrypted.arfs \"AAAAAAAAAAAAAAAAsFt3LbDTrKVIllwFtQLUTuE=\""))
}

func PrintCatMessage() {
	fmt.Println(("cat: like cat in encrypted fs"))
	fmt.Println(("\tuse: ubac cat -path=<encryptedfs>.arfs <resource>"))
	fmt.Println(("\tparameters required: path (-path) and resource (-resource). Resource is obtained by adret utility"))
	fmt.Println(("\texample: ubac cat -path encrypted.arfs \"AAAAAAAAAAAAAAAAsFt3LbDTrKVIllwFtQLUTuE=\""))
}

func PrintTreeMessage() {
	fmt.Println(("tree: like tree command in encrypted fs. It shows the hierarchy of the filesytem"))
	fmt.Println(("\tuse: ubac tree -path=<encryptedfs>.arfs"))
	fmt.Println(("\tparameters required: path (-path), which is the .arafs location"))
	fmt.Println(("\texample: ubac tree encrypted.arfs"))
}

func PrintHelp() {
	fmt.Println("ubac utility is used to interact with encrypted fs (provided by adret utility). As every data manipulate is encrypted you can use it in a non-trusted environement")
	fmt.Println("Available commands:")

	// Contain all command function help messsage
	mFunctionName := map[string]func(){
		"PrintHelpMessage": PrintHelpMessage,
		"PrintLpMessage":   PrintLsMessage,
		"PrintCatMessage":  PrintCatMessage,
		"PrintTreeMessage": PrintTreeMessage,
	}
	// oredered them for printing
	orderedFunctionName := []string{"PrintHelpMessage", "PrintLspMessage", "PrintCatMessage", "PrintTreeMessage"}

	//print help message for all
	for i := 0; i < len(orderedFunctionName); i++ {
		functionName := orderedFunctionName[i]
		PrintCommandMessage(mFunctionName[functionName])
	}

}
