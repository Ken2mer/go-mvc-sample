package app

import (
	"fmt"
	"net/http"
)

func RunServer() int {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
