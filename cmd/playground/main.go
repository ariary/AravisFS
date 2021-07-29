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

// Use to get the ls parameter of a request (as string format) ie  get body of a RequestLs or BodyLs struct
// Behind the scene it transforms the request in requestLs struct. Then retrieve the body
// (by Unmarshall it into a BodyLs struct)
func GetBodyLsFromRequest(request string) BodyLs {
	//We now that the request is (or should be) a RequestLs so transform it
	requestLs := RequestLs{}
	json.Unmarshal([]byte(request), &requestLs)

	bodyJson, err := json.Marshal(requestLs.Body)
	if err != nil {
		panic(err)
	}

	bodyLs := BodyLs{}
	json.Unmarshal(bodyJson, &bodyLs)

	return bodyLs
}

type Request struct {
	Name string `json:"name"`
	Body []byte `json:"body"`
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
	fmt.Println(GetRequestLsJSON("encrypted.arafs", "toto/tata"))
	//equal {"name":"ls","body":{"path":"encrypted.arafs","name":"toto/tata"}}

	//adret side
	request := GetRequestLsJSON("encyrpted.arafs", "toto/tata")

	//ubac side
	requestOnUbacSide := Request{}
	json.Unmarshal([]byte(request), &requestOnUbacSide)
	fmt.Println(requestOnUbacSide.Name)
	if requestOnUbacSide.Name == LS {
		//this is a ls request => build requestLs struct to parse it and get parameter
		bodyLs := GetBodyLsFromRequest(request)

		path := bodyLs.Path
		resourceName := bodyLs.ResourceName
		//now call Printls
		fmt.Println("(on ubac side) path: ", path)
		fmt.Println("(on ubac side) resource name: ", resourceName)
	}

	//UbacListen(20080)

}
