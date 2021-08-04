package ubac

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

//Start an http server waiting for action request over encrypted fs (pointed by path)
func Listen(port int, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(fmt.Sprintf("Encrypted fs not found (cannot access '%s': No such file or directory)", path))
	}
	mux := http.NewServeMux()
	//Add handlers
	mux.HandleFunc("/endpoints", Endpoints)
	mux.HandleFunc("/ls", RemoteLs(path))            //ls
	mux.HandleFunc("/cat", RemoteCat(path))          //cat
	mux.HandleFunc("/tree", RemoteTree(path))        //tree
	mux.HandleFunc("/patch", RemoteApplyPatch(path)) //apply patch (for write access command)

	log.Println("Waiting for remote command over encrypted fs (", path, ") on port", port, ":...")
	err := http.ListenAndServe(":"+strconv.Itoa(port), mux)
	log.Fatal(err)
}
