package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Ken2mer/go-mvc/config"
)

func run() int {
	config.Routes()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
