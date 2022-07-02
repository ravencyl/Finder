package viewer

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
}

func HomePageViewer(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("Templates/homepage.html"))
	page := Page{Title: "Create new Project"}
	tpl.Execute(w, page)
}
