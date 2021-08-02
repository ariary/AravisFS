package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/ariary/AravisFS/pkg/adret"
	prompt "github.com/c-bata/go-prompt"
)

type UrlUbac struct {
	hostname string
	port     string
	url      string
}
type FSContext struct {
	path    string
	urlUbac UrlUbac
	key     string
}

var ctx *FSContext

// suggestions list
var suggestions = []prompt.Suggest{
	// General
	{"exit", "Exit adretctl"},
	{"help", "get help method"},

	//key
	{"keyconfig", "set the key used to decrypt the fs"},
	{"keyprint", "print the current key"},

	// Command on ubac
	{"connect", "Connect to the configured Ubac"}, //in fact launch get and see if there is result
	{"cd", "Change the  working directory in encrypted fs. (Do not support full path)"},

	// Read Method
	{"ls", "list directory contents on remote encrypted fs"},
	{"cat", "print file content on remote encrypted fs resource"},
	{"tree", "print tree of remote encrypte fs"},
}

//Prefix for the prompt
func livePrefix() (string, bool) {
	if ctx.path == "/" {
		return "", false
	}
	return ctx.path + "> ", true
}

//Check if context key is defined
func hasKey(ctx FSContext) bool {
	if ctx.key == "" {
		fmt.Println("Please set the key used for encryption/decryption with keyconfig")
		return false
	}
	return true
}

//perform at each loop
func executor(in string) {
	in = strings.TrimSpace(in)

	// var method, body string
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "keyconfig":
		if len(blocks) < 2 {
			fmt.Println("please enter the key")
		} else {
			ctx.key = blocks[1]
		}
		return
	case "keyprint":
		fmt.Println(ctx.key)
		return
	case "connect":
		//check if host is alive
		fmt.Printf("Checking if host   (%s) is alive...", ctx.urlUbac.url)
		fmt.Println()
		_, err := http.Get(ctx.urlUbac.url) //could also use net.DialTimeout
		if err != nil {
			fmt.Println("Get: connect: connection refused. Please make sure you're ubac listener is up and adrectl well configured")
			return
		}
		//retrieve root dir
		fmt.Println("Retrieve root dir of encrypted fs...")
		if !hasKey(*ctx) {
			return
		}
		root, err := adret.RemoteRootDir(ctx.key)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.path = root
		return
	case "ls":
		if !hasKey(*ctx) {
			return
		}
		var path string
		if len(blocks) == 2 {
			path = ctx.path + "/" + blocks[1]
		} else {
			path = ctx.path
		}
		adret.PrintRemoteLs(path, ctx.key)
		return
	case "cd":
		//TODO: special case: "cd"-> root et "cd -"--->previous
		if len(blocks) < 2 {
			return
		} else {
			//test if dir exist (TODO test if it is effectively a directory)
			//ie have a function Exist + IsDir renvoie vrai si la resource est de type dir
			// ubac & adret side
			if !hasKey(*ctx) {
				return
			}
			newPath := path.Clean(ctx.path + "/" + blocks[1])
			if !adret.RemoteExist(newPath, ctx.key) {
				fmt.Printf("cd: %v: No such file or directory", blocks[1])
				fmt.Println()
			} else if !adret.RemoteIsDir(newPath, ctx.key) {
				fmt.Printf("cd: %v: Not a directory", blocks[1])
				fmt.Println()
			} else {
				ctx.path = newPath
			}
		}
	case "cat":
		if !hasKey(*ctx) {
			return
		}
		var path string
		if len(blocks) == 2 {
			path = ctx.path + "/" + blocks[1]
		} else {
			fmt.Println("Please provide a file to cat command to print its result")
			return
		}
		adret.PrintRemoteCat(path, ctx.key)
		return
	case "tree":
		if !hasKey(*ctx) {
			return
		}
		adret.RemoteTree(ctx.key)
		return
	case "help":
		fmt.Println("available commands: keyconfig, keyprint, connect, ls, cat, tree")
		return
	case "exit":
		fmt.Println("Bye!🕶")
		handleExit()
		os.Exit(0)
	default:
		fmt.Printf("Unknown command: %s", blocks[0])
		fmt.Println()
		return
	}

}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

//Function launch when adrectl exit. Mainly use to prevent https://github.com/c-bata/go-prompt/issues/228
func handleExit() {
	rawModeOff := exec.Command("/bin/stty", "-raw", "echo")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()
}

func main() {
	defer handleExit()

	if len(os.Args) != 3 && os.Getenv("REMOTE_UBAC_URL") == "" {
		fmt.Println("Launch adretctl with hostname and port ('adretclt <ubac-hostname> <ubac_port>' or defined envar REMOTE_UBAC_URL ")
		os.Exit(1)
	}

	//launch config loading
	basePath := ""
	var url, port, hostname string
	if os.Getenv("REMOTE_UBAC_URL") != "" {
		//envar exist
		url = os.Getenv("REMOTE_UBAC_URL")
		fqdn := strings.Split(url, "/")[2]
		hostname = strings.Split(fqdn, ":")[0]
		port = strings.Split(fqdn, ":")[1]
	} else {
		//retrieve root path of encrypted FS
		hostname = os.Args[1]
		port = os.Args[2]
		url = "http://" + hostname + ":" + port + "/"
		os.Setenv("REMOTE_UBAC_URL", url)
	}

	u := &UrlUbac{
		hostname: hostname,
		port:     port,
		url:      url,
	}

	ctx = &FSContext{
		path:    basePath,
		urlUbac: *u,
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(basePath+"> "),
		prompt.OptionLivePrefix(livePrefix),
		prompt.OptionTitle("adretctl"),
	)
	p.Run()
}
