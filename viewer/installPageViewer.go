package viewer

import (
	"Finder/controller"
	"html/template"
	"net/http"
)

func InstallPageViewer(w http.ResponseWriter, r *http.Request) {
	controller.Install_DB()

	var tpl = template.Must(template.ParseFiles("templates/installPage.html"))
	page := Page{Title: "Create new Project"}
	tpl.Execute(w, page)
}
