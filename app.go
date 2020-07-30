// Test http package
package main

import (
	"fmt"
	"net/http"
)

// RouteHandler export
func RouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Welcome To GO</h1>")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":3000", nil)
}
