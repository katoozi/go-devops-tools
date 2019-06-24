package nap

type RestResource struct {
	Endpoint string // <BaseUrl>/get/
	Method   string // GET, POST, PUT, DELETE, OPTION, etc...
	Router   *CBRouter
}
