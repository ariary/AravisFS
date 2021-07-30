package adret

import "fmt"

func PrintCommandMessage(f func()) {
	f()
	fmt.Println()
}

func PrintMapInOrder(m map[string]func(), order []string) {
	for i := 0; i < len(order); i++ {
		functionName := order[i]
		PrintCommandMessage(m[functionName])
	}
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

func PrintConfigremoteMessage() {
	fmt.Println(("configremote: enable to configurate the port and hostname of the remote ubac listener. It is equivalent to 'export REMOTE_UBAC_LISTER=<hostname>:<port>"))
	fmt.Println(("\tuse: eval `adret configremote -port=<ubac_port> -host=<ubac_hostanme>`"))
	fmt.Println(("\tparameters required: port (-port) and hostname (-host) (ubac is listening on hostname:port"))
	fmt.Println()
}

func PrintRemotelsMessage() {
	fmt.Println(("remotels: perform a ls on a remote encrypted fs. We use ubac on listening mode to proxify our request onto the fs"))
	fmt.Println(("\tuse: adret remotels -key=<key> <resource_name>"))
	fmt.Println(("\tparameters required: key (-key) use for encryption/decryption resource_name (-resource) which is the resource on which we want to perform the ls command"))
	fmt.Println(("\t💡 Lauch ubac listener and set REMOTE_UBAC_LISTER envar (w/ 'adret configremote') before"))
	fmt.Println()
}

//Print all help messages (all available command and their use)
func PrintHelp() {
	fmt.Println("adret utility is used to perform FS encrytion and decrypt data from fs encrypted (from ubac utility)")
	fmt.Println("Available commands:")
	PrintCommandMessage(PrintHelpMessage)

	//ENCRYPTION COMMAND
	fmt.Println("ENCRYPTION COMMAND")
	// Contain all command function help messsage
	mFunctionNameEncryption := map[string]func(){
		"PrintDarkenpathMessage": PrintDarkenpathMessage,
		"PrintEncryptfsMessage":  PrintEncryptfsMessage,
	}
	// oredered them for printing
	orderedFunctionNameEncryption := []string{"PrintDarkenpathMessage", "PrintEncryptfsMessage"}
	//print help message for all
	PrintMapInOrder(mFunctionNameEncryption, orderedFunctionNameEncryption)

	//READ COMMAND
	fmt.Println("READ ACCESS COMMAND")

	mFunctionNameRead := map[string]func(){
		"PrintDecryptlsMessage":   PrintDecryptlsMessage,
		"PrintDecryptcatMessage":  PrintDecryptcatMessage,
		"PrintDecrypttreeMessage": PrintDecrypttreeMessage,
	}

	orderedFunctionNameRead := []string{"PrintDecryptlsMessage", "PrintDecryptcatMessage", "PrintDecrypttreeMessage"}

	PrintMapInOrder(mFunctionNameRead, orderedFunctionNameRead)

	//REMOTE COMMAND
	fmt.Println("REMOTE COMMAND")
	// Contain all command function help messsage
	mFunctionNameRemote := map[string]func(){
		"PrintConfigremoteMessage": PrintConfigremoteMessage,
		"PrintRemotelsMessage":     PrintRemotelsMessage,
	}
	// oredered them for printing
	orderedFunctionNameRemote := []string{"PrintConfigremoteMessage", "PrintRemotelsMessage"}
	//print help message for all
	PrintMapInOrder(mFunctionNameRemote, orderedFunctionNameRemote)
}
