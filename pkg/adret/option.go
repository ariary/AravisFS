package adret

import "fmt"

func PrintCommandMessage(f func()) {
	f()
	fmt.Println()
}

func PrintHelpMessage() {
	fmt.Println("help: get help for adret utility and example")
}

func PrintDarkenpathMessage() {
	fmt.Println(("darkenpath: use it to encrypt a path with the key"))
	fmt.Println(("\tuse: adret darkenpath -key=<secret> <path>"))
	fmt.Println(("\tparameters required: key (-key) and path (-path)"))
	fmt.Println(("\texample: adret darkenpath -key \"toto\" \"test/toto/titi\""))
}

func PrintEncryptfsMessage() {
	fmt.Println(("encryptfs: use it to encrypt a directory with the key. It will create a encrypted fs .arafs"))
	fmt.Println(("\tLIMITATION: encrypt a directory in the current one avoid ie  \"..\" etc in path parameters"))
	fmt.Println(("\tuse: adret encryptfs -key=<secret> <path>"))
	fmt.Println(("\tparameters required: key (-key) and path (-path)"))
	fmt.Println(("\texample: adret encryptfs -key \"toto\" \"test/toto/titi\""))
	fmt.Println()
}

func PrintDecryptlsMessage() {
	fmt.Println(("decryptls: use it to decrypt output of 'ubac ls' command. It enable us to perform a ls in encryted fs"))
	fmt.Println(("\tuse: adret decryptls -key=<secret> <ubac_ls_output>"))
	fmt.Println(("\tparameters required: key (-key) and ubac_ls_output (-output)"))
	fmt.Println(("\t💡 get output with 'ubac ls -path=<encryptedfs>.arfs <resource>'"))
	fmt.Println(("\texample: adret decryptls -key \"toto\" \"directory:AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le21X9+Fky1Fb98v61k+DQJivbwJosBKJ8FSD4YitHoo9GZf40l3HLHGTDjc=\""))
	fmt.Println()
}

func PrintDecryptcatMessage() {
	fmt.Println(("decryptcat: use it to decrypt output of 'ubac cat' command. It enable us to perform a cat in encryted fs"))
	fmt.Println(("\tuse: adret decryptcat -key=<secret> <ubac_cat_output>"))
	fmt.Println(("\tparameters required: key (-key) and ubac_cat_output (-output)"))
	fmt.Println(("\t💡 get output with 'ubac cat -path=<encryptedfs>.arfs <resource>'"))
	fmt.Println(("\texample: adret decryptcat -key \"toto\" \"directory:AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le21X9+Fky1Fb98v61k+DQJivbwJosBKJ8FSD4YitHoo9GZf40l3HLHGTDjc=\""))
	fmt.Println()
}

func PrintDecrypttreeMessage() {
	fmt.Println(("decrypttree: use it to decrypt output of 'ubac tree' command. It show the hierarchy of the encryted fs"))
	fmt.Println(("\tuse: adret decrypttree -key=<secret> <ubac_tree_output>"))
	fmt.Println(("\tparameters required: key (-key) and ubac_tree_output (-output)"))
	fmt.Println(("\t💡 get output with 'ubac tree -path=<encryptedfs>.arfs'"))
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
	PrintDecryptlsMessage()
	fmt.Println()
	PrintDecryptcatMessage()
	fmt.Println()

	// Contain all command function help messsage
	mFunctionName := map[string]func(){
		"PrintHelpMessage":        PrintHelpMessage,
		"PrintDarkenpathMessage":  PrintDarkenpathMessage,
		"PrintEncryptfsMessage":   PrintEncryptfsMessage,
		"PrintDecryptlsMessage":   PrintDecryptlsMessage,
		"PrintDecryptcatMessage":  PrintDecryptcatMessage,
		"PrintDecrypttreeMessage": PrintDecrypttreeMessage,
	}
	// oredered them for printing
	orderedFunctionName := []string{"PrintHelpMessage", "PrintDarkenpathMessage", "PrintEncryptfsMessage", "PrintDecryptlsMessage", "PrintDecryptcatMessage", "PrintDecrypttreeMessage"}

	//print help message for all
	for i := 0; i < len(orderedFunctionName); i++ {
		functionName := orderedFunctionName[i]
		PrintCommandMessage(mFunctionName[functionName])
	}

}
