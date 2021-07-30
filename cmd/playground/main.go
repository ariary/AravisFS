package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ariary/AravisFS/pkg/adret"
)

///// Ubac side

func main() {
	//./ubac listen --path=./test/arafs/encrypted.arafs
	url := "http://localhost:4444/ls"
	key := "toto"

	var jsonStr = []byte(`{"name":"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le4+fERV81xq4dTDo0PnkM3M="}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	adret.PrintLs(string(body), key)
}
