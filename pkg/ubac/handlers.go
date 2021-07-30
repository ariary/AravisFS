package ubac

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ariary/AravisFS/pkg/remote"
)

// ENDPOINT

//return all available endpoints
//test example: curl http://127.1:4444/endpoints
func Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := "endpoints\n"
	endpoints += "ls\n"
	fmt.Fprintf(w, endpoints)
}

// LS PART

// handler for ls function. Waiting request with JSON body with this structure {"name":"..."}
// where name is the name of the resource on which we apply the ls
// test it example:
// curl http://127.1:4444/ls -H "Content-Type: application/json" --request POST --data '{"name":"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le4+fERV81xq4dTDo0PnkM3M="}'
func RemoteLs(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//like this we could use path in RemoteLs Handler

		var body remote.BodyLs

		err := remote.DecodeJSONBody(w, r, &body)
		if err != nil {
			var mr *remote.MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		lsContent := Ls(body.ResourceName, path)

		fmt.Fprintf(w, lsContent)
	}
}
