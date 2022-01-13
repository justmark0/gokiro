package main

import (
	"database/sql"
	"fmt"
	"gokiro/src/app"
	"log"
	"net/http"
)

func main() {
	var err error

	app.DB, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/gokiro")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/upload/file", app.HandleUploadFile)
	http.HandleFunc("/upload/text", app.HandleUploadText)

	http.HandleFunc("/", app.Handler)

	fmt.Println("starting server at :8000")
	_ = http.ListenAndServe(":8000", nil)
}
