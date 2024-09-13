package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct{
	Nome string
	Email string
	contato string
}

func home (wt http.ResponseWriter, req *http.Request) {

	u := usuario{
		"Thiago Eleuterio", "thiago-eleuterio@testando.go", "2345678", 
	}

	templates.ExecuteTemplate(wt, "index.html", u)
}

func main() {
	
	templates = template.Must(template.ParseGlob("static/*.html"))

	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}