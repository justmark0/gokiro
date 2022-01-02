package app

import (
	"fmt"
	"net/http"
)

// HandleUploadFile downloads file saves on disk, and save in db
func HandleUploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			_, _ = fmt.Fprintln(w, "Use form data")
			return
		}
		file, handler, err := r.FormFile("file") // Retrieve the file from form data

		if err != nil {
			_, _ = fmt.Fprintln(w, "Send in field \"file\"")
			return
		}
		defer file.Close() // Close the file when we finish

		path, err := SaveFileAndReturnPath(file, handler)
		if err != nil {
			_, _ = fmt.Fprintln(w, "Internal server error")
			return
		}
		id := SaveDataToDB(path, true)
		_, _ = fmt.Fprintln(w, "{\"id\":", id, "}")
	} else {
		_, _ = fmt.Fprintln(w, "Send POST method")
	}

}

// HandleUploadText saves text to db and returns id
func HandleUploadText(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data []byte
		_, _ = r.Body.Read(data)
		id := SaveDataToDB(string(data), false)
		_, _ = fmt.Fprintln(w, "{\"id\":", id, "}")
	} else {
		_, _ = fmt.Fprintln(w, "Send POST method")
	}
}

// Handler default root handler
func Handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Main page")
}
