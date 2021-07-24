package main

import (
	"log"
	"net/http"
	"text/template"
)

func webserver() {
	http.HandleFunc("/", handleInit)

	err := http.ListenAndServe(*flagHTTPPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleInit(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles(
		"assets/html/header.html",
		"assets/html/init.html",
		"assets/html/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = i.ExecuteTemplate(w, "init", nil)
	if err != nil {
		log.Fatal(err)
	}
}
