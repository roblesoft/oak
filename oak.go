package oak

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Oak struct {
	routes  []string
	AppName string
	server  *http.ServeMux
	logger  *log.Logger
}

func New() *Oak {
	return &Oak{
		AppName: "Default",
		server:  &http.ServeMux{},
		logger:  log.New(os.Stdout, "Api: ", log.LstdFlags),
	}
}

func (o *Oak) saveRoute(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

func (o *Oak) GET(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(handlerFn)
}

func (o *Oak) POST(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(handlerFn)
}

func (o *Oak) PUT(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(handlerFn)
}

func (o *Oak) DELETE(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(handlerFn)
}

func (o *Oak) HEAD(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(handlerFn)
}

func (o *Oak) PATCH(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(handlerFn)
}

func (o *Oak) OPTIONS(path string, handlerFn http.HandlerFunc) {
	o.saveRoute(handlerFn)
}

func (o *Oak) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, world")
	fmt.Fprint(w, req.Method)
}
