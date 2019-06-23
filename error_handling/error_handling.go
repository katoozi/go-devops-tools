package main

import (
	"log"
	"net/http"
	"os"
)

func handleError() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln("Unable to get response.")
	}
	router.Process(resp)
}
