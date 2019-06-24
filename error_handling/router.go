package main

import (
	"log"
	"net/http"
)

// HandlerFuc is a simple function that handle the response
type HandlerFuc func(*http.Response)

// ResponseRouter is a struct that store the response handlers
type ResponseRouter struct {
	Handlers       map[int]HandlerFuc
	DefaultHanderl HandlerFuc
}

// NewRouter is a factory function that create new ResponseRouter object
func NewRouter() *ResponseRouter {
	return &ResponseRouter{
		Handlers: make(map[int]HandlerFuc),
		DefaultHanderl: func(r *http.Response) {
			log.Fatalln("Unhandeld Response: ", r.StatusCode)
		},
	}
}

// Register will save the response handler
func (r *ResponseRouter) Register(status int, handler HandlerFuc) {
	r.Handlers[status] = handler
}

// Process will detect the response status code and choose the handler function
func (r *ResponseRouter) Process(resp *http.Response) {
	f, ok := r.Handlers[resp.StatusCode]
	if !ok {
		r.DefaultHanderl(resp)
		return
	}
	f(resp)
}
