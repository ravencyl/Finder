package viewer

import (
	"html/template"
	"net/http"
)

func InstallPageViewer(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("templates/installPage.html"))
	page := Page{Title: "Create new Project"}
	tpl.Execute(w, page)
}
