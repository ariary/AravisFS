package adret

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ariary/AravisFS/pkg/encrypt"
	"github.com/ariary/AravisFS/pkg/filesystem"
	"github.com/ariary/AravisFS/pkg/remote"
)

// Parse the directory content retrieve from encrypted fs (w/ ubac for example)
// First decrypt the whole content
// Then parse the content as each resources is separated by "\"
func ParseLsDirectoryContent(content string, key string) []string {
	contentDecoded, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		log.Fatal("error:", err)
	}
	lsResult := string(encrypt.DecryptByte(contentDecoded, key))
	//TO DO: add info if file is directory or not
	return strings.Split(lsResult, "\\")
}

// Parse the ls result received from ubac utility
// result structure: resourceType:base64(Encrypt(Content))
func ParseLsContent(result string, key string) string {
	//file or directory output?
	resultParsed := strings.SplitN(result, ":", 2)
	if len(resultParsed) != 2 {
		log.SetFlags(0)
		log.Fatal("ParseLsContent: failed to parse the input (must be '<Type>:<ubac_ls_output>'")
	}
	resourceType := resultParsed[0]
	content := resultParsed[1]
	if resourceType == filesystem.DIRECTORY {
		// decrypt content, parse it to have all resource under and print it
		files := ParseLsDirectoryContent(content, key)
		//take basename
		// TO DO propose to print the full path (ie don't call Base ~ ls -d $PWD/*)
		for i := range files {
			files[i] = filepath.Base(files[i])
		}

		//TO DO: specify if resource is a file or directory
		lsOutput := strings.Join(files, " ")
		return lsOutput

	} else if resourceType == filesystem.FILE {
		// decrypt content and print it (it must be the filename)
		lsOutput := string(encrypt.DecryptStringFromUbac(content, key))
		return lsOutput
	} else {
		log.Fatal("Failed to retrieve content after decrypting result")
	}
	return ""
}

func PrintLs(result string, key string) {
	output := ParseLsContent(result, key)
	if output != "" {
		fmt.Println(output)
	} else {
		log.Fatal("Failed to parse ls result")
	}

}

// Perform ls on a remote listening ubac (proxing to encrypted fs)
// First craft the request, send it (the request instruct ubac to perform a ls)
// take the reponse and decrypt it
func RemoteLs(resourceName string, key string) string {
	url := os.Getenv("REMOTE_UBAC_URL")
	if url == "" {
		fmt.Println("Configure REMOTE_UBAC_URL envar with `adret configremote` before launching remotels. see `adret help`")
		os.Exit(1)
	}
	endpoint := url + "ls"

	darkenresourceName := encrypt.DarkenPath(resourceName, key)

	bodyRes := remote.SendReadrequest(darkenresourceName, endpoint)

	//decrypt the reponse to show cat result
	result := ParseLsContent(bodyRes, key)
	return result
}

func PrintRemoteLs(resourceName string, key string) {
	result := RemoteLs(resourceName, key)
	fmt.Println(result)
}
