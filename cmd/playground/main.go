package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ariary/AravisFS/pkg/adret"
	"github.com/ariary/AravisFS/pkg/encrypt"
	"github.com/ariary/AravisFS/pkg/remote"
)

///// Ubac side
// Use to config the information aboute remote ubac listener to avoid to mentionned them in
// CLI after
func ConfigRemote(hostname string, port string) {
	urlUbac := "http://" + hostname + ":" + port + "/"
	os.Setenv("REMOTE_UBAC_URL", urlUbac)
	fmt.Println("REMOTE_UBAC_URL:", os.Getenv("REMOTE_UBAC_URL"))
}

//Perform ls on a remote listening ubac (proxing to encrypted fs)
func RemoteLs(resourceName string, key string) {
	endpoints := os.Getenv("REMOTE_UBAC_URL") + "ls"
	fmt.Println(endpoints)

	darkenresourceName := encrypt.DarkenPath(resourceName, key)

	//create body
	body, err := json.Marshal(remote.CreateBodyLs(darkenresourceName))
	if err != nil {
		panic(err)
	}

	//perform requets
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

	//do smtg with response
	fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	bodyRes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(bodyRes))
	adret.PrintLs(string(bodyRes), key)
}

func main() {
	//./ubac listen --path=./test/arafs/encrypted.arafs
	// url := "http://localhost:4444/ls"
	// key := "toto"

	// var jsonStr = []byte(`{"name":"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le4+fERV81xq4dTDo0PnkM3M="}`)
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// // fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	// adret.PrintLs(string(body), key)

	ConfigRemote("localhost", "4444")
	RemoteLs("test/mytestfolder", "toto")
}
