package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/katoozi/go-devops-tools/nap"
)

var api = nap.NewAPI("https://httpbin.org")

func main() {
	list := flag.Bool("list", false, "Get List of all Api Resources.")
	flag.Parse()
	if *list {
		fmt.Println("Available Resources:")
		for _, name := range api.ResourceNames() {
			fmt.Println("\t", name)
		}
		return
	}
	resource := os.Args[1]
	if err := api.Call(resource, nil); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	router := nap.NewRouter()
	router.RegisterFunc(200, func(res *http.Response, _ interface{}) error {
		defer res.Body.Close()
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(content))
		return nil
	})
	rest := nap.NewResource("/get", "GET", router)
	api.AddResource("get", rest)
}
