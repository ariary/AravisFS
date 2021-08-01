package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ariary/AravisFS/pkg/adret"
	prompt "github.com/c-bata/go-prompt"
)

type FSContext struct {
	path    string
	urlUbac string
	key     string
}

var ctx *FSContext

// See https://github.com/eliangcs/http-prompt/blob/master/http_prompt/completion.py
var suggestions = []prompt.Suggest{
	// General
	{"exit", "Exit adret-prompt"},
	{"configkey", "set the key used to decrypt the fs"},
	{"help", "get help method"},

	// Command on ubac
	{"cd", "Change path"},
	{"connect", "Connect to the configured Ubac"}, //in fact launch get and see if there is result

	// Read Method

	{"ls", "list directory contents on remote encrypted fs"},
}

func livePrefix() (string, bool) {
	if ctx.path == "/" {
		return "", false
	}
	return ctx.path + "> ", true
}

func executor(in string) {
	in = strings.TrimSpace(in)

	// var method, body string
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	case "configkey":
		if len(blocks) < 2 {
			fmt.Println("please enter the key")
		} else {
			ctx.key = blocks[1]
		}
		return
	case "ls":
		fmt.Println(ctx.path)
		fmt.Println(ctx.key)
		adret.RemoteLs(ctx.path, ctx.key)
		// case "cd":
		// 	if len(blocks) < 2 {
		// 		ctx.url.Path = "/"
		// 	} else {
		// 		ctx.url.Path = path.Join(ctx.url.Path, blocks[1])
		// 	}
		// 	return
		// case "get", "delete":
		// 	method = strings.ToUpper(blocks[0])
		// case "post", "put", "patch":
		// 	if len(blocks) < 2 {
		// 		fmt.Println("please set request body.")
		// 		return
		// 	}
		// 	body = strings.Join(blocks[1:], " ")
		// 	method = strings.ToUpper(blocks[0])
	}
	// if method != "" {
	// 	req, err := http.NewRequest(method, ctx.url.String(), strings.NewReader(body))
	// 	if err != nil {
	// 		fmt.Println("err: " + err.Error())
	// 		return
	// 	}
	// 	req.Header = ctx.header
	// 	res, err := ctx.client.Do(req)
	// 	if err != nil {
	// 		fmt.Println("err: " + err.Error())
	// 		return
	// 	}
	// 	result, err := ioutil.ReadAll(res.Body)
	// 	if err != nil {
	// 		fmt.Println("err: " + err.Error())
	// 		return
	// 	}
	// 	fmt.Printf("%s\n", result)
	// 	ctx.header = http.Header{}
	// 	return
	// }

	// if h := strings.Split(in, ":"); len(h) == 2 {
	// 	// Handling HTTP Header
	// 	ctx.header.Add(strings.TrimSpace(h[0]), strings.Trim(h[1], ` '"`))
	// } else {
	// 	fmt.Println("Sorry, I don't understand.")
	// }
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Launch adret-interactive with hostname and port ('adret-interactive <ubac-hostname> <ubac_port>'")
		os.Exit(1)
	}
	// basePath := "" //retrieve root path of encrypted FS
	basePath := "test/mytestfolder"

	url := "http://" + os.Args[1] + ":" + os.Args[2] + "/"
	os.Setenv("REMOTE_UBAC_URL", url)
	ctx = &FSContext{
		path:    basePath,
		urlUbac: url,
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(basePath+"> "),
		prompt.OptionLivePrefix(livePrefix),
		prompt.OptionTitle("adret-prompt"),
	)
	p.Run()
}
