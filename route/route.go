package route

import (
	"log"
	"net/http"

	"gopenguin/config"
	"gopenguin/controllers"
)

func Load() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static", http.StripPrefix("/static", fs))

	config, err := config.GetConfig("./config.json")
	if err != nil {
		log.Fatal("Invalid config: ", err)
	}

	c, _ := controllers.NewController(config)

	http.Handle("/", http.HandlerFunc(c.HomeIndex))
	http.Handle("/account/login", http.HandlerFunc(c.AccountLogin))
}
