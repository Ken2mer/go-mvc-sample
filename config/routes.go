package config

import (
	"fmt"
	"html"
	"net/http"

	"github.com/Ken2mer/go-mvc/controller"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Routes() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/users", controller.Users)
}
