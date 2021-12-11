package main

import (
	"gokiro/app"
	"log"
	"net/http"
	"os"
)

func main() {
	app := app.New()
	http.HandleFunc("/", app.Router.ServeHTTP)

	port := "80"
	log.Println("App running on port: " + port)
	err := http.ListenAndServe(":" + port, nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}

