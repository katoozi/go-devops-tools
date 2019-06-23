package main

import (
	"log"
	"net/http"
)

type HandlerFuc func(*http.Response)

type ResponseRouter struct {
	Handlers       map[int]HandlerFuc
	DefaultHanderl HandlerFuc
}

func NewRouter() *ResponseRouter {
	return &ResponseRouter{
		Handlers: make(map[int]HandlerFuc),
		DefaultHanderl: func(r *http.Response) {
			log.Fatalln("Unhandeld Response: ", r.StatusCode)
		},
	}
}

func (r *ResponseRouter) Register(status int, handler HandlerFuc) {
	r.Handlers[status] = handler
}

func (r *ResponseRouter) Process(resp *http.Response) {
	f, ok := r.Handlers[resp.StatusCode]
	if !ok {
		r.DefaultHanderl(resp)
		return
	}
	f(resp)
}
