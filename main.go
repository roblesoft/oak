package main

import (
	"fmt"

	"github.com/roblesoft/oak/oak"
)

func main() {
	oak := oak.New()

	oak.Get("/hello_world", func() {
		fmt.Println("Hello,")
	})
	oak.Run()
}
