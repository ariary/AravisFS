package adret

import (
	"fmt"
	"log"

	"github.com/ariary/AravisFS/pkg/encrypt"
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
