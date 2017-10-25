package wego

var (
	UrlRouter map[string]WegoRouter = make(map[string]WegoRouter)
)

type WegoRouter struct {
	Controller    interface{}
	Method        string
	RequestMethod string
}

func AddRouter(url string, r WegoRouter) {
	if _, ok := UrlRouter[url]; !ok {
		UrlRouter[url] = r
	}
}

func Router(requestUrl string, controller interface{}, requestMethod string, method string) {
	var r WegoRouter = WegoRouter{Controller: controller, Method: method, RequestMethod: requestMethod}
	AddRouter(requestUrl, r)
}
