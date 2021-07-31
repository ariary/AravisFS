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
		err := remote.DecodeBodyRead(w, r, &body)

		if err != nil {
			return
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
// where name is the name of the resource on which we apply the ls
// test it example:
func RemoteCat(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//like this we could use path in RemoteLs Handler
		var body remote.BodyRead
		err := remote.DecodeBodyRead(w, r, &body)

		if err != nil {
			return
		}

		catContent, err := Cat(body.ResourceName, path)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		fmt.Fprintf(w, catContent)
	}
}
