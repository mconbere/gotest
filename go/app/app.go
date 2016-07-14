package app

import (
	"fmt"
	"net/http"

	"github.com/mconbere/gotest/go/foo"
	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	rid := appengine.RequestID(ctx)
	fmt.Fprintf(w, "%s for request %s\n", foo.Foo(), rid)
}
