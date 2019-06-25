package nap

import (
	"net/http"
	"testing"
)

func TestApiCall(t *testing.T) {
	api := NewAPI("https://httpbin.org")
	router := NewRouter()
	router.RegisterFunc(200, func(res *http.Response, _ interface{}) error {
		return nil
	})
	rest := NewResource("/get", "GET", router)
	api.AddResource("get", rest)
	if err := api.Call("get", nil); err != nil {
		t.Fail()
	}
	// fmt.Println(api.PrintRecources())
	resources := api.ResourceNames()
	if len(resources) != 1 || resources[0] != "get" {
		t.Fail()
	}
}

func TestAPIAuth(t *testing.T) {
	api := NewAPI("https://httpbin.org")
	router := NewRouter()
	router.RegisterFunc(200, func(res *http.Response, _ interface{}) error {
		return nil
	})
	rest := NewResource("/basic-auth/{{.user}}/{{.pass}}", "GET", router)
	api.AddResource("basicauth", rest)
	api.SetAuth(&AuthBasic{
		Username: "user",
		Password: "passwOrd",
	})
	if err := api.Call("basicauth", map[string]string{
		"user": "user",
		"pass": "passwOrd",
	}); err != nil {
		t.Fail()
	}
}
