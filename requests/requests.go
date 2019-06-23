package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func get() {
	resp, err := http.Get("https://httpbin.org/get")
	// resp, err := http.Get("https://httpbin.org/get") // you can send querystring
	if err != nil {
		log.Fatalln("Unable to get response")
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("Unable to read from response body.")
	}
	fmt.Println(string(content))
}

func post() {
	resp, err := http.Post("https://httpbin.org/post?search=something", "text/plain",
		strings.NewReader("this the request content"))
	if err != nil {
		log.Fatalln("Unable to get response")
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("Unable to read from response body.")
	}
	fmt.Println(string(content))
}

func customeClient() {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to create request.")
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Unable to get response")
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("Unable to read from response body.")
	}
	fmt.Println(string(content))
}
