package oak

import (
	"context"
	"log"
	"net/http"
	"os"
)

type Oak struct {
	router *Router
}

func New() *Oak {
	return &Oak{
		router: &Router{
			logger: log.New(os.Stdout, "Api: ", log.LstdFlags),
			trees:  nil,
		},
	}
}

type OakHandle func(ctx *Ctx)

func (o *Oak) GET(path string, handle OakHandle) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := &Ctx{
			Context:  context.Background(),
			response: w,
			request:  r,
		}
		handle(ctx)
	}
	o.router.GET(path, handler)
}

func (o *Oak) Run() {
	http.ListenAndServe(":3000", o.router)
}
