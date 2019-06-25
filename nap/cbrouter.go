package nap

import (
	"fmt"
	"net/http"
)

type RouterFunc func(client *http.Response, content interface{}) error

type CBRouter struct {
	Routers       map[int]RouterFunc // init refers to status code
	DefaultRouter RouterFunc
}

func NewRouter() *CBRouter {
	return &CBRouter{
		Routers: make(map[int]RouterFunc),
		DefaultRouter: func(resp *http.Response, content interface{}) error {
			return fmt.Errorf("From: %s received unknown status: %d",
				resp.Request.URL.String(), resp.StatusCode)
		},
	}
}

func (r *CBRouter) RegisterFunc(status int, fn RouterFunc) {
	r.Routers[status] = fn
}

func (r *CBRouter) CallFunc(resp *http.Response, content interface{}) error {
	fn, ok := r.Routers[resp.StatusCode]
	if !ok {
		fn = r.DefaultRouter
	}
	if err := fn(resp, content); err != nil {
		return err
	}
	return nil
}
