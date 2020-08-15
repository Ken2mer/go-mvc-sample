package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Ken2mer/go-mvc/app/model"
)

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []model.User
	users = append(users, model.User{
		Name: "example",
	})
	json.NewEncoder(w).Encode(users)
}
