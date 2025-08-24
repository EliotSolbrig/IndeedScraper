package router

import (
	"fmt"
	"net/http"
	"html/template"
)

func (router *Router) IndexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, " request on /")

	templates := append([]string{"templates/pages/home.html",},basePaseFiles...)

	tmpl := template.Must(template.ParseFiles(templates...))

	err := tmpl.ExecuteTemplate(w, "base", map[string]any{
		"Data": nil,
	})
	if err != nil {
		panic(fmt.Errorf("Error executing home page template: %s", err))
	}
}
