package main

import (
	"log"
	"net/http"
	"text/template"
)

func webserver() {

	TEMPLATES = template.Must(template.ParseGlob("assets/html/*.html"))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", handleInit)

	err := http.ListenAndServe(*flagHTTPPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleInit(w http.ResponseWriter, r *http.Request) {
	err := TEMPLATES.ExecuteTemplate(w, "init", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
