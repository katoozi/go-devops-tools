package nap

import "fmt"

type API struct {
	BaseURL       string // https://httpbin.org/
	Resources     map[string]*RestResource
	DefaultRouter *CBRouter
	Client        *Client
}

func NewAPI(baseURL string) *API {
	return &API{
		BaseURL:       baseURL,
		Resources:     make(map[string]*RestResource),
		DefaultRouter: NewRouter(),
		Client:        NewClient(),
	}
}

func (api *API) SetAuth(auth Authentication) {
	api.Client.SetAuth(auth)
}

func (api *API) AddResource(name string, res *RestResource) {
	api.Resources[name] = res
}

func (api *API) Call(name string, params map[string]string) error {
	res, ok := api.Resources[name]
	if !ok {
		return fmt.Errorf("Resource Does not Exist: %s", name)
	}
	if err := api.Client.ProcessRequest(api.BaseURL, res, params); err != nil {
		return err
	}
	return nil
}

func (api *API) PrintRecources() string {
	resources := ""
	for k, v := range api.Resources {
		resources += fmt.Sprintf("\tname: %s, Endpoint: %s, Method: %s\n", k, v.Endpoint, v.Method)
	}
	return resources
}

func (api *API) ResourceNames() []string {
	names := []string{}
	for k := range api.Resources {
		names = append(names, k)
	}
	return names
}
