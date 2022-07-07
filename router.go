package main

import (
	"Finder/viewer"
	"net/http"
)

func Router(mux *http.ServeMux) *http.ServeMux {

	// mux.HandleFunc("/test", handler)
	// mux.HandleFunc("/project/create", controller.ProjectCreate)
	// mux.HandleFunc("/api/project/", controller.APIProject)
	// mux.HandleFunc("/", controller.Homehandler)
	// // Assets resouce direction
	// fs := http.FileServer(http.Dir("assets"))
	// mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", viewer.HomePageViewer)

	return mux
}
