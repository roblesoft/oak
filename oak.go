package oak

import (
	"fmt"
	"net/http"
	"time"
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
	o.server.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Time.String(time.Now()), path)
		handlerFn(w, r)
	})
}

func (o *Oak) Run() {
	s := http.Server{
		Addr:    ":3000",
		Handler: o.server,
	}

	s.ListenAndServe()
}
