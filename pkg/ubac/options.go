package ubac

import "fmt"

func PrintCommandMessage(f func()) {
	f()
	fmt.Println()
}

func PrintHelpMessage() {
	fmt.Println("help: get help for ubac utility and example")
}

//READ CMD

func PrintLsMessage() {
	fmt.Println(("ls: like ls in encrypted fs"))
	fmt.Println(("\tuse: ubac ls -path=<encryptedfs>.arafs <resource>"))
	fmt.Println(("\tparameters required: path (-path) and resource (-resource). Resource is obtained by adret utility"))
	fmt.Println(("\texample: ubac ls -path encrypted.arafs \"AAAAAAAAAAAAAAAAsFt3LbDTrKVIllwFtQLUTuE=\""))
}

func PrintCatMessage() {
	fmt.Println(("cat: like cat in encrypted fs"))
	fmt.Println(("\tuse: ubac cat -path=<encryptedfs>.arafs <resource>"))
	fmt.Println(("\tparameters required: path (-path) and resource (-resource). Resource is obtained with adret utility"))
	fmt.Println(("\texample: ubac cat -path encrypted.arafs \"AAAAAAAAAAAAAAAAsFt3LbDTrKVIllwFtQLUTuE=\""))
}

func PrintTreeMessage() {
	fmt.Println(("tree: like tree command in encrypted fs. It shows the hierarchy of the filesytem"))
	fmt.Println(("\tuse: ubac tree -path=<encryptedfs>.arafs"))
	fmt.Println(("\tparameters required: path (-path), which is the .arafs location"))
	fmt.Println(("\texample: ubac tree encrypted.arafs"))
}

//WRITE CMD

func PrintRmMessage() {
	fmt.Println(("rm: like rm in encrypted fs (remove file and dir, no extra flag needed"))
	fmt.Println(("\tuse: ubac rm -path=encrypted.arafs -patch=<patch_from_adret>"))
	fmt.Println(("\tparameters required: path (-path) and the patch to apply. Patch is obtained with adret utility"))
	fmt.Println(("\texample: ubac rm -path=encrypted.arafs -patch=$(cat patch.json)"))
	fmt.Println(("\t💡 Example of how you can get a patch 'adret encryptrm -key=toto -tree=<ubac_tree_output> test_medium/ansible/slice > patch.json'"))
}

//REMOTECMD

func PrintListenMessage() {
	fmt.Println(("listen: launch http server waiting for adret request on specified port. It take the path of the encrypted fs on which we will apply commmand"))
	fmt.Println(("\tuse: ubac listen -path=<encryptedfs>.arafs <port>"))
	fmt.Println(("\tparameters required: path (-path), which is the .arafs location, and port (default 4444)"))
	fmt.Println(("\texample: ubac listen -path=encrypted.arafs"))
}

func PrintHelp() {
	fmt.Println("ubac utility is used to interact with encrypted fs (provided by adret utility). As every data manipulate is encrypted you can use it in a non-trusted environement")
	fmt.Println("Available commands:")
	PrintCommandMessage(PrintHelpMessage)

	//READ COMMAND
	fmt.Println()
	fmt.Println("READ ACCESS COMMAND")
	// Contain all command function help messsage
	mFunctionNameRead := map[string]func(){
		"PrintLsMessage":   PrintLsMessage,
		"PrintCatMessage":  PrintCatMessage,
		"PrintTreeMessage": PrintTreeMessage,
	}
	// oredered them for printing
	orderedFunctionNameRead := []string{"PrintLsMessage", "PrintCatMessage", "PrintTreeMessage"}

	//print help message for all
	for i := 0; i < len(orderedFunctionNameRead); i++ {
		functionName := orderedFunctionNameRead[i]
		PrintCommandMessage(mFunctionNameRead[functionName])
	}

	//WRITE COMMAND
	fmt.Println()
	fmt.Println("WRITE ACCESS COMMAND")
	// Contain all command function help messsage
	mFunctionNameWrite := map[string]func(){
		"PrintRmMessage": PrintRmMessage,
	}
	// oredered them for printing
	orderedFunctionNameWrite := []string{"PrintRmMessage"}

	//print help message for all
	for i := 0; i < len(orderedFunctionNameWrite); i++ {
		functionName := orderedFunctionNameWrite[i]
		PrintCommandMessage(mFunctionNameWrite[functionName])
	}

	//REMOTE COMMAND
	fmt.Println("REMOTE COMMAND")
	// Contain all command function help messsage
	mFunctionNameremote := map[string]func(){
		"PrintListenMessage": PrintListenMessage,
	}
	// oredered them for printing
	orderedFunctionNameRemote := []string{"PrintListenMessage"}

	//print help message for all
	for i := 0; i < len(orderedFunctionNameRemote); i++ {
		functionName := orderedFunctionNameRemote[i]
		PrintCommandMessage(mFunctionNameremote[functionName])
	}
}
