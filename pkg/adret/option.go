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

//Encrypt
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
	fmt.Println("\t⚠ To be able to print the tree later, launch the command on the same location of the directory you want to encrypt (ie. path must be in you current directory when you launch the command")
}

//Read access
func PrintDecryptlsMessage() {
	fmt.Println(("decryptls: use it to decrypt output of 'ubac ls' command. It enable us to perform a ls in encryted fs"))
	fmt.Println(("\tuse: adret decryptls -key=<secret> <ubac_ls_output>"))
	fmt.Println(("\tparameters required: key (-key) and ubac_ls_output (-output)"))
	fmt.Println(("\t💡 get output with 'ubac ls -path=<encryptedfs>.arfs <resource>'"))
	fmt.Println(("\texample: adret decryptls -key \"toto\" \"directory:AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le21X9+Fky1Fb98v61k+DQJivbwJosBKJ8FSD4YitHoo9GZf40l3HLHGTDjc=\""))
}

func PrintDecryptcatMessage() {
	fmt.Println(("decryptcat: use it to decrypt output of 'ubac cat' command. It enable us to perform a cat in encryted fs"))
	fmt.Println(("\tuse: adret decryptcat -key=<secret> <ubac_cat_output>"))
	fmt.Println(("\tparameters required: key (-key) and ubac_cat_output (-output)"))
	fmt.Println(("\t💡 get output with 'ubac cat -path=<encryptedfs>.arafs <resource>'"))
	fmt.Println(("\texample: adret decryptcat -key \"toto\" \"directory:AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le21X9+Fky1Fb98v61k+DQJivbwJosBKJ8FSD4YitHoo9GZf40l3HLHGTDjc=\""))
}

func PrintDecrypttreeMessage() {
	fmt.Println(("decrypttree: use it to decrypt output of 'ubac tree' command. It show the hierarchy of the encryted fs"))
	fmt.Println(("\tuse: adret decrypttree -key=<secret> <ubac_tree_output>"))
	fmt.Println(("\tparameters required: key (-key) and ubac_tree_output (-output)"))
	fmt.Println(("\t💡 get output with 'ubac tree -path=<encryptedfs>.arafs'"))
	fmt.Println("\t⚠ To be able to print the tree, the fs has to be built in the same directory where the `adret encryptfs` command was run(see adret encryptfs")
}

//Write access

func PrintRmPatchMessage() {
	fmt.Println(("encryptrm: return a patch to provide to ubac to remove a resource on encrypted fs."))
	fmt.Println(("\tuse: adret encryptrm -key=<secret> -tree=<ubac_tree_output> <path>"))
	fmt.Println(("\tparameters required: key (-key), ubac_tree_output (-output) and path (-path) of the resource you want to remove"))
	fmt.Println(("\t💡 You first need to retrieve the tree of the encrypted fs using ubac 'ubac tree -path=<encryptedfs>.arafs'"))
}

//Remote
func PrintConfigremoteMessage() {
	fmt.Println(("configremote: enable to configurate the port and hostname of the remote ubac listener. It is equivalent to 'export REMOTE_UBAC_URL=<hostname>:<port>"))
	fmt.Println(("\tuse: eval `adret configremote -port=<ubac_port> -host=<ubac_hostname>`"))
	fmt.Println(("\tparameters required: port (-port) and hostname (-host) (ubac is listening on hostname:port"))
}

func PrintRemotelsMessage() {
	fmt.Println(("remotels: perform a ls on a remote encrypted fs. We use ubac on listening mode to proxify our request onto the fs"))
	fmt.Println(("\tuse: adret remotels -key=<key> <resource_name>"))
	fmt.Println(("\tparameters required: key (-key) use for encryption/decryption resource_name (-resource) which is the resource on which we want to perform the ls command"))
	fmt.Println(("\t💡 Lauch ubac listener and set REMOTE_UBAC_URL envar (w/ 'adret configremote') before"))
}

func PrintRemotecatMessage() {
	fmt.Println("remotecat: perform a cat on a remote encrypted fs. We use ubac on listening mode to proxify our request onto the fs")
	fmt.Println("\tuse: adret remotecat -key=<key> <resource_name>")
	fmt.Println("\tparameters required: key (-key) use for encryption/decryption and resource_name (-resource) which is the resource on which we want to perform the cat command")
	fmt.Println("\t💡 Lauch ubac listener and set REMOTE_UBAC_URL envar (w/ 'adret configremote') before")
}

func PrintRemotetreeMessage() {
	fmt.Println("remotetree: perform a tree on a remote encrypted fs. We use ubac on listening mode to proxify our request onto the fs")
	fmt.Println("\tuse: adret remotree -key=<key>")
	fmt.Println("\tparameters required: key (-key) use for encryption/decryption ")
	fmt.Println(("\t💡 Lauch ubac listener and set REMOTE_UBAC_URL envar (w/ 'adret configremote') before"))
	fmt.Println("\t⚠ To be able to print the tree, the fs has to be built in the same directory where the `adret encryptfs` command was run(see adret encryptfs")
}

func PrintRemotermMessage() {
	fmt.Println("remoterm: perform a rm on a remote encrypted fs. We use ubac on listening mode to proxify our request onto the fs")
	fmt.Println("\tuse: adret remoterm -key=<key> <resource_name>")
	fmt.Println("\tparameters required: key (-key) use for encryption/decryption and resource_name (-resource) which is the resource on which we want to perform the cat command")
	fmt.Println(("\t💡 Lauch ubac listener and set REMOTE_UBAC_URL envar (w/ 'adret configremote') before"))
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
	fmt.Println()
	fmt.Println("READ ACCESS COMMAND")

	mFunctionNameRead := map[string]func(){
		"PrintDecryptlsMessage":   PrintDecryptlsMessage,
		"PrintDecryptcatMessage":  PrintDecryptcatMessage,
		"PrintDecrypttreeMessage": PrintDecrypttreeMessage,
	}

	orderedFunctionNameRead := []string{"PrintDecryptlsMessage", "PrintDecryptcatMessage", "PrintDecrypttreeMessage"}

	PrintMapInOrder(mFunctionNameRead, orderedFunctionNameRead)

	//WRITE COMMAND
	fmt.Println()
	fmt.Println("WRITE ACCESS COMMAND")

	mFunctionNameWrite := map[string]func(){
		"PrintRmPatchMessage": PrintRmPatchMessage,
	}

	orderedFunctionNameWrite := []string{"PrintRmPatchMessage"}

	PrintMapInOrder(mFunctionNameWrite, orderedFunctionNameWrite)

	//REMOTE COMMAND
	fmt.Println()
	fmt.Println("REMOTE COMMAND")
	// Contain all command function help messsage
	mFunctionNameRemote := map[string]func(){
		"PrintConfigremoteMessage": PrintConfigremoteMessage,
		"PrintRemotelsMessage":     PrintRemotelsMessage,
		"PrintRemotecatMessage":    PrintRemotecatMessage,
		"PrintRemotetreeMessage":   PrintRemotetreeMessage,
		"PrintRemotermMessage":     PrintRemotermMessage,
	}
	// oredered them for printing
	orderedFunctionNameRemote := []string{"PrintConfigremoteMessage", "PrintRemotelsMessage", "PrintRemotecatMessage", "PrintRemotetreeMessage", "PrintRemotermMessage"}
	//print help message for all
	PrintMapInOrder(mFunctionNameRemote, orderedFunctionNameRemote)
}
