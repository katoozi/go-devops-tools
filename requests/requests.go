package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln("There is a Error during the Get Request.")
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("Unable to read from response body.")
	}
	fmt.Println(string(content))
}
