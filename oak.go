package oak

import (
	"net/http"
)

type Oak struct {
	routes  []string
	AppName string
	server  *http.ServeMux
}

func New() *Oak {
	return &Oak{
		AppName: "Default",
		server:  &http.ServeMux{},
	}
}

func (o *Oak) Get(path string, handlerFn http.HandlerFunc) {
	o.server.Handle(path, handlerFn)
}

func (o *Oak) Run() {
	s := http.Server{
		Addr:    ":3000",
		Handler: o.server,
	}

	s.ListenAndServe()
}
