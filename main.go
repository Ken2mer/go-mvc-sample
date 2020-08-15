package main

import (
	"fmt"
	"os"

	"github.com/Ken2mer/go-mvc/app"
	"github.com/Ken2mer/go-mvc/config"
)

var (
	// set postgres url
	dbURL = ""
)

func run() int {
	if err := app.OpenDB(dbURL); err != nil {
		fmt.Println(err)
		return 1
	}
	defer app.CloseDB()

	config.Routes()
	app.RunServer()

	return 0
}

func main() {
	os.Exit(run())
}
