package oak

import (
	"log"
	"net/http"
)

type Router struct {
	logger   *log.Logger
	trees    map[string]*node
	NotFound http.Handler
}

type Handle func(http.ResponseWriter, *http.Request)

func (o *Router) Handler(handler http.HandlerFunc) Handle {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

func (o *Router) Handle(method string, path string, handler http.HandlerFunc) {
	if o.trees == nil {
		o.trees = make(map[string]*node)
	}

	if o.trees[method] != nil {
		root := o.trees[method]
		for len(root.children) != 0 {
			root = root.children[0]
		}
		root.children = append(root.children,
			&node{
				method:  method,
				path:    path,
				handler: o.Handler(handler),
			})
	} else {
		o.trees[method] = &node{
			method:  method,
			path:    path,
			handler: o.Handler(handler),
		}
	}
}

func (o *Router) GET(path string, handlerFn http.HandlerFunc) {
	o.Handle(http.MethodGet, path, handlerFn)
}

func (o *Router) POST(path string, handlerFn http.HandlerFunc) {
	o.Handle(http.MethodPost, path, handlerFn)
}

func (o *Router) PUT(path string, handlerFn http.HandlerFunc) {
	o.Handle(http.MethodPut, path, handlerFn)
}

func (o *Router) DELETE(path string, handlerFn http.HandlerFunc) {
	o.Handle(http.MethodDelete, path, handlerFn)
}

func (o *Router) HEAD(path string, handlerFn http.HandlerFunc) {
	o.Handle(http.MethodHead, path, handlerFn)
}

func (o *Router) PATCH(path string, handlerFn http.HandlerFunc) {
	o.Handle(http.MethodPatch, path, handlerFn)
}

func (o *Router) OPTIONS(path string, handlerFn http.HandlerFunc) {
	o.Handle(http.MethodOptions, path, handlerFn)
}

// ServeHTTP to implement http.Handler interface
func (o *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	o.logger.Println(req.Method, req.URL.Path)
	path := req.URL.Path
	root := o.trees[req.Method]

	// handle not found
	if root != nil {
		root = root.getNode(path)
	}

	if root != nil {
		root.getValue()(w, req)
		return
	}

	http.NotFound(w, req)
}
