package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Println("[ERROR] [renderTemplate] ParseFiles error:", err.Error())
		return
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("[ERROR] [renderTemplate] Execute error:", err.Error())
		return
	}

}
