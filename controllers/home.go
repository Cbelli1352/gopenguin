package controllers

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func (c *Controller) HomeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/home/index.html")
}
