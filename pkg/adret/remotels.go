package adret

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/ariary/AravisFS/pkg/encrypt"
	"github.com/ariary/AravisFS/pkg/remote"
)

// Perform ls on a remote listening ubac (proxing to encrypted fs)
// First craft the request, send it (the request instruct ubac to perform a ls)
// take the reponse and decrypt it
func RemoteLs(resourceName string, key string) {
	url := os.Getenv("REMOTE_UBAC_URL")
	if url == "" {
		fmt.Println("Configure REMOTE_UBAC_URL envar with `adret configremote` before launching remotels. see `adret help`")
		os.Exit(1)
	}
	endpoints := url + "ls"

	darkenresourceName := encrypt.DarkenPath(resourceName, key)

	//create body
	body, err := json.Marshal(remote.CreateBodyLs(darkenresourceName))
	if err != nil {
		panic(err)
	}

	//perform request
	req, err := http.NewRequest("POST", endpoints, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//decrypt the reponse to show ls result

	bodyRes, _ := ioutil.ReadAll(resp.Body)

	if !strings.Contains(resp.Status, "200") {
		fmt.Println(string(bodyRes))
		//panic("Response code from ubac http server different from 200")
	}

	PrintLs(string(bodyRes), key)
}
