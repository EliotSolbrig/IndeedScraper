package router

import (
	"fmt"
	"net/http"
	"html/template"
)

var basePaseFiles []string = []string{
	"templates/base.html",
	"templates/header.html",
	"templates/footer.html",
}

type Router struct {
}

func NewRouter() *Router {
	return &Router{
	}
}

func (router *Router) IndeedScrapeTestPage(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Method, " request on /indeedtest1")

	templates := append([]string{"templates/pages/testpage.html",},basePaseFiles...)
	tmpl := template.Must(template.ParseFiles(templates...))

	err := tmpl.ExecuteTemplate(w, "base", map[string]any{
		"Data": nil,
	})
	
	if err != nil {
		panic(fmt.Errorf("Error executing test page template: %s", err))
	}


}

