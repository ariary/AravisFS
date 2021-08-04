package remote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type BodyRead struct {
	ResourceName string `json:"name"`
}

func CreateBodyRead(resourceName string) BodyRead {
	b := &BodyRead{
		ResourceName: resourceName}
	return *b
}

type BodyWrite struct {
	Patch string `json: "patch"`
}

func CreateBodyWrite(patch string) BodyWrite {
	b := &BodyWrite{
		Patch: patch}
	return *b
}

// Create a request for a read command (cat or ls or tree) and return the response body
// Errors are handled in the function
func SendReadRequest(darkenresourceName string, endpoint string) string {
	//Create body
	var body []byte
	if darkenresourceName == "" {
		//logically a tree request, we don't need a body request
		body = nil
	} else {
		//cat or ls
		var err error
		body, err = json.Marshal(CreateBodyRead(darkenresourceName))
		if err != nil {
			panic(err)
		}
	}
	//perform request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//decrypt the reponse to show read command result

	bodyRes, _ := ioutil.ReadAll(resp.Body)
	// bad status code (!=200, logically 404)
	if !strings.Contains(resp.Status, "200") {
		fmt.Println(string(bodyRes))
		os.Exit(1)
		//panic("Response code from ubac http server different from 200")
	}

	return string(bodyRes)
}

// Create a request for a write command (eg rm) and return the response body
// the body is a patch in json format which tells to ubac which ressource modified, removed or add
// Errors are handled in the function
func SendWriteRequest(patch string, endpoint string) string {
	//Create body
	var body []byte
	if patch == "" {
		//logically a tree request, we don't need a body request
		fmt.Println("SendWriteRequest: Patch is empty for request")
		os.Exit(1)
	} else {
		var err error
		body, err = json.Marshal(CreateBodyWrite(patch))
		if err != nil {
			panic(err)
		}
	}
	//perform request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//decrypt the reponse to show read command result

	bodyRes, _ := ioutil.ReadAll(resp.Body)
	// bad status code (!=200, logically 404)
	if !strings.Contains(resp.Status, "200") {
		fmt.Println(string(bodyRes))
		os.Exit(1)
		//panic("Response code from ubac http server different from 200")
	}

	return string(bodyRes)
}
