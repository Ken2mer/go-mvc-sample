package main

import (
	"fmt"
	"os"

	"github.com/Ken2mer/go-mvc/app"
	"github.com/Ken2mer/go-mvc/config"
)

func run() int {
	config.Routes()
	app.RunServer()

	return 0
}

func main() {
	os.Exit(run())
}
