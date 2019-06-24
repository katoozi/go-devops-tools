package nap

type API struct {
	BaseURL       string // https://httpbin.org/
	Resources     map[string]RestResource
	DefaultRouter *CBRouter
	Client        *Client
}
