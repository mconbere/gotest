package main

import (
	"fmt"

	"github.com/mconbere/gotest/go/foo"
	"google.golang.org/appengine"
)

func main() {
	fmt.Printf("%s on %s\n", foo.Foo(), appengine.ServerSoftware())
}
