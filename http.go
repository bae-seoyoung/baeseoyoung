package main

import (
	"log"
	"net/http"
	"text/template"
)

func webserver() {

	http.HandleFunc("/", handleInit)

	err := http.ListenAndServe(*flagHTTPPort, nil) //nil의 정확한 용도를 이슈로 만들어봅시당.
	if err != nil {
		log.Fatal(err)
	}
}

func handleInit(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles(
		"html/header.html",
		"html/init.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = i.ExecuteTemplate(w, "init", nil)
	if err != nil {
		log.Fatal(err)
	}
}
