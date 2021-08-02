package adret

import (
	"fmt"
	"log"
	"os"

	"github.com/ariary/AravisFS/pkg/encrypt"
	"github.com/ariary/AravisFS/pkg/remote"
)

// Parse the cat result received from ubac utility
// cat only work on file resource not directory, so we don't have to check the resourceType
// Error must be manage by ubac (if file doesn't exist or if it is a directory)
// result structure: base64(Encrypt(Content))
func ParseCatContent(result string, key string) string {

	// decrypt content and print it (it must be the file content)
	catOutput := string(encrypt.DecryptStringFromUbac(result, key))
	return catOutput
}

func PrintCat(result string, key string) {
	output := ParseCatContent(result, key)
	if output != "" {
		fmt.Println(output)
	} else {
		log.Fatal("Failed to parse cat result")
	}

}

// Perform cat on a remote listening ubac (proxing to encrypted fs)
// First craft the request, send it (the request instruct ubac to perform a cat)
// take the reponse and decrypt it
func RemoteCat(resourceName string, key string) string {
	url := os.Getenv("REMOTE_UBAC_URL")
	if url == "" {
		fmt.Println("Configure REMOTE_UBAC_URL envar with `adret configremote` before launching remotels. see `adret help`")
		os.Exit(1)
	}
	endpoint := url + "cat"

	darkenresourceName := encrypt.DarkenPath(resourceName, key)

	bodyRes := remote.SendReadrequest(darkenresourceName, endpoint)
	//decrypt the reponse to show cat result
	return ParseCatContent(bodyRes, key)
}

// Print the result of a cat on a remote listening ubac (proxing to encrypted fs)
func PrintRemoteCat(resourceName string, key string) {
	result := RemoteCat(resourceName, key)
	fmt.Println(result)
}
