package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	handleError()
}

var router = NewRouter()

func init() {
	router.Register(200, func(r *http.Response) {
		defer r.Body.Close()
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln("Unable to read body")
		}
		fmt.Println(string(content))
	})
	router.Register(404, func(r *http.Response) {
		log.Fatalln("Not Found(404): ", r.Request.URL.String())
	})
}
