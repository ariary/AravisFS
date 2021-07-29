package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

const LS = "ls"

type BodyLs struct {
	Path         string `json:"path"`
	ResourceName string `json:"name"`
}

func createBodyLs(path string, resourceName string) BodyLs {

	b := &BodyLs{
		Path:         path,
		ResourceName: resourceName}
	return *b
}

type RequestLs struct {
	Name string `json:"name"`
	Body BodyLs `json:"body"`
}

func CreateRequestLs(body BodyLs) RequestLs {

	r := &RequestLs{
		Name: LS,
		Body: body}
	return *r
}

//pour l'instant affiche plus tard enverra
func GetRequestLsJSON(path string, resourceName string) string {
	body := createBodyLs(path, resourceName)
	request := CreateRequestLs(body)

	requestLs, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	return string(requestLs)
}

type Request struct {
	Name string `json:"name"`
	Body string `json:"type"`
}

// Put a listener for adret call. It parse the request and call the appropriate function consequently
// ex: nohup adret listen -por=20080 &
func UbacListen(port int) {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:", port)
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency.
		go echo(conn)
	}
}

// echo is a handler function that simply echoes received data.
func echo(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store received data.
	b := make([]byte, 512)
	for {
		// Receive data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s", size, string(b))

		// Send data via conn.Write.
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	fmt.Println(GetRequestLsJSON("encyrpted.arafs", "toto/tata"))
	//equal {"name":"ls","body":{"path":"encyrpted.arafs","name":"toto/tata"}}

	//adret side
	request := GetRequestLsJSON("encyrpted.arafs", "toto/tata")

	//ubac side
	requestOnUbacSide := Request{}
	json.Unmarshal([]byte(request), &requestOnUbacSide)
	fmt.Println(requestOnUbacSide.Name)
	if requestOnUbacSide.Name == LS {
		//this is a ls request => build requestLS struct to parse it
		requestLs, err := json.Marshal(requestOnUbacSide.Body)
		if err != nil {
			panic(err)
		}
		bodyLs := BodyLs{}
		json.Unmarshal(requestLs, &bodyLs)

		//now call Printls
		fmt.Println("(on ubac side) path: ", bodyLs.Path)
		fmt.Println("(on ubac side) resource name: ", bodyLs.ResourceName)
	}

	//UbacListen(20080)

}
