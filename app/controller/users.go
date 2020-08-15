package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Ken2mer/go-mvc/app"
	"github.com/Ken2mer/go-mvc/app/model"
)

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []model.User
	app.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
}
