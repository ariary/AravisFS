package ubac

import (
	"fmt"
	"net/http"

	"github.com/ariary/AravisFS/pkg/remote"
)

// ENDPOINT

//return all available endpoints
//test example: curl http://127.1:4444/endpoints
func Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := "endpoints\n"
	endpoints += "ls\n"
	endpoints += "cat\n"
	fmt.Fprintf(w, endpoints)
}

// READ ACCESS HANDLER PART

// handler for ls function. Waiting request with JSON body with this structure {"name":"..."}
// where name is the name of the resource on which we apply the ls
// test it example:
// curl http://127.1:4444/ls -H "Content-Type: application/json" --request POST --data '{"name":"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le4+fERV81xq4dTDo0PnkM3M="}'
func RemoteLs(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//like this we could use path in RemoteLs Handler
		var body remote.BodyRead
		err := remote.DecodeBody(w, r, &body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		lsContent, err := Ls(body.ResourceName, path)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		fmt.Fprintf(w, lsContent)
	}
}

// CAT PART

// handler for cat function. Waiting request with JSON body with this structure {"name":"..."}
// where name is the name of the resource on which we apply the cat
func RemoteCat(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//read body request
		var body remote.BodyRead
		err := remote.DecodeBody(w, r, &body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		catContent, err := Cat(body.ResourceName, path)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		fmt.Fprintf(w, catContent)
	}
}

// handler for tree function. Waiting GET request with no body on /tree endpoint
// return a json which is the tree of the encrypted fs
// test example: curl -X GET http://localhost:4444/tree
func RemoteTree(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		treeContent := GetTreeFromFS(path)

		fmt.Fprintf(w, treeContent)
	}
}

// handler for apply patch function (ie modification of encrypted fs). Waiting POST request with the patch on the body
// on /patch endpoint
// return nothing apart if there is a problem
func RemoteApplyPatch(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//read body request
		var body remote.BodyWrite
		err := remote.DecodeBody(w, r, &body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		err = ApplyPatch(body.Patch, path)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, "encrypted filesytem changes done")
	}
}
