package viewer

import (
	"html/template"
	"net/http"
)

func UploadPageViewer(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("templates/upload.html"))
	page := Page{Title: "Test upload"}
	tpl.Execute(w, page)
}
