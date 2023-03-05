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
	router *Router
	Server *http.Server
}

func New() *Oak {
	return &Oak{
		Server: &http.Server{
			Addr: ":3000",
		},
		router: &Router{
			logger: log.New(os.Stdout, "Api: ", log.LstdFlags),
			trees:  nil,
		},
	}
}

type OakHandle func(ctx *Ctx)

func (o *Oak) oakHandleToHandlerFunc(handle OakHandle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &Ctx{
			Context:  context.Background(),
			response: w,
			request:  r,
		}
		handle(ctx)
	}
}

func (o *Oak) GET(path string, handle OakHandle) {
	handler := o.oakHandleToHandlerFunc(handle)
	o.router.GET(path, handler)
}

func (o *Oak) Run() {
	o.Server.Handler = o.router
	go func() {
		err := o.Server.ListenAndServe()
		if err != nil {
			o.router.logger.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)
	sig := <-sigChannel
	o.router.logger.Println("Received terminate, graceful shutdown", sig)
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	o.Server.Shutdown(timeoutContext)
}
