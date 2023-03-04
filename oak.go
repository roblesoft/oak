package oak

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

func (o *Oak) GET(path string, handlerFn http.HandlerFunc) {
	o.server.Handle(path, handlerFn)
}

func (o *Oak) POST(path string, handlerFn http.HandlerFunc) {
	o.server.Handle(path, handlerFn)
}

func (o *Oak) PUT(path string, handlerFn http.HandlerFunc) {
	o.server.Handle(path, handlerFn)
}

func (o *Oak) DELETE(path string, handlerFn http.HandlerFunc) {
	o.server.Handle(path, handlerFn)
}

func (o *Oak) HEAD(path string, handlerFn http.HandlerFunc) {
	o.server.Handle(path, handlerFn)
}

func (o *Oak) PATCH(path string, handlerFn http.HandlerFunc) {
	o.server.Handle(path, handlerFn)
}

func (o *Oak) OPTIONS(path string, handlerFn http.HandlerFunc) {
	o.server.Handle(path, handlerFn)
}

func (o *Oak) Run() {
	s := http.Server{
		Addr:    ":3000",
		Handler: o.server,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			o.logger.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)
	sig := <-sigChannel
	o.logger.Println("Received terminate, graceful shutdown", sig)
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(timeoutContext)
}
