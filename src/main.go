package main

import (
	"fmt"
	"gokiro/src/app"
	"net/http"
)

func main() {
	http.HandleFunc("/upload/file", app.HandleUploadFile)
	http.HandleFunc("/upload/text", app.HandleUploadText)

	http.HandleFunc("/", app.Handler)

	fmt.Println("starting server at :8000")
	_ = http.ListenAndServe(":8000", nil)
}
