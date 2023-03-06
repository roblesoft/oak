# Oak

Oak is a http web framework

# Example

```go
package main

import (
	"github.com/roblesoft/oak"
)

type Json struct {
	Body string
}

func main() {
	app := oak.New()

	app.GET("/hello_world", func(ctx *oak.Ctx) {
		ctx.JSON(&Json{Body: "Hello world"})
	})

	app.Run()
}
```