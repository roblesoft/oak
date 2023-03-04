package oak

import (
	"log"
	"net/http"
	"os"
)

type Oak struct {
	routes   []string
	AppName  string
	server   *http.ServeMux
	logger   *log.Logger
	trees    map[string]*node
	NotFound http.Handler
}

type Handle func(http.ResponseWriter, *http.Request)

type node struct {
	method  string
	path    string
	handler Handle
}

func New() *Oak {
	return &Oak{
		AppName: "Default",
		server:  &http.ServeMux{},
		logger:  log.New(os.Stdout, "Api: ", log.LstdFlags),
		trees:   nil,
	}
}

func (o *Oak) Handler(handler http.HandlerFunc) Handle {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

func (o *Oak) saveRoute(method string, path string, handler http.HandlerFunc) {
	if o.trees == nil {
		o.trees = make(map[string]*node)
	}

	o.trees[method] = &node{
		method:  method,
		path:    path,
		handler: o.Handler(handler),
	}
}

func (o *Oak) GET(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(http.MethodGet, path, handlerFn)
}

func (o *Oak) POST(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(http.MethodPost, path, handlerFn)
}

func (o *Oak) PUT(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(http.MethodPut, path, handlerFn)
}

func (o *Oak) DELETE(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(http.MethodDelete, path, handlerFn)
}

func (o *Oak) HEAD(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(http.MethodHead, path, handlerFn)
}

func (o *Oak) PATCH(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(http.MethodPatch, path, handlerFn)
}

func (o *Oak) OPTIONS(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(http.MethodOptions, path, handlerFn)
}

// ServeHTTP to implement http.Handler interface
func (o *Oak) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	o.logger.Println(req.Method, req.URL.Path)
	handler := o.trees[req.Method]

	// handle not found
	if handler == nil {
		http.NotFound(w, req)
		return
	}

	handler.handler(w, req)

}
