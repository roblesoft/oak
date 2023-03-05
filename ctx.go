package oak

import (
	"context"
	"encoding/json"
	"net/http"
)

type Ctx struct {
	Context  context.Context
	response http.ResponseWriter
	request  *http.Request
}

func (c *Ctx) JSON(data interface{}) error {
	c.Response().Header().Set("Content-Type", "application/json")
	return json.NewEncoder(c.Response()).Encode(data)
}

func (c *Ctx) Response() http.ResponseWriter {
	return c.response
}

func (c *Ctx) Request() *http.Request {
	return c.request
}
