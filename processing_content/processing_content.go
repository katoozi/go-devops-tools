package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetResponse will used by json standard library
type GetResponse struct {
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

// ToString is string represent of GetResponse struct
func (r *GetResponse) ToString() string {
	s := fmt.Sprintf("GET Response \nOrigin IP: %s\nRequest URL: %s", r.Origin, r.URL)
	for k, v := range r.Headers {
		s += fmt.Sprintf("Header: %s = %s\n", k, v)
	}
	return s
}

func processContent() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to get response.")
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read body.")
	}
	respContent := GetResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Println(respContent.ToString())
}
