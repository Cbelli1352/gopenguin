package controllers

import (
	"encoding/json"
	"net/http"

	"gopenguin/auth"
	"gopenguin/models"
)

func (c Controller) accountLoginGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/account/login.html")
}

func (c Controller) accountLoginPost(w http.ResponseWriter, r *http.Request) {
	loginRequest := models.LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}

	user := auth.User{ID: "hello", Username: loginRequest.Username, Role: "default"}

	response := models.LoginResponse{}

	response.Token, err = c.Auth.Encode(&user)

	responseString, err := json.Marshal(response)
	w.Write(responseString)
}

func (c Controller) AccountLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.accountLoginGet(w, r)
	case http.MethodPost:
		c.accountLoginPost(w, r)
	default:
		http.Error(w, "Invalid method for operation", 405)
	}
}
