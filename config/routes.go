package config

import (
	"fmt"
	"html"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Routes() {
	http.HandleFunc("/", hello)
}
