package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"main/internal/rod"
	"main/router"
)


const staticDir string = "/static/"


func main(){
	fmt.Println("Indeed scraper")

	godotenv.Load()

	port := os.Getenv("PORT")

	r := mux.NewRouter()
	router := router.NewRouter()

    staticHandler := http.StripPrefix(staticDir, http.FileServer(http.Dir("static/")))
    r.PathPrefix(staticDir).Handler(staticHandler)

	browser := rod.GetBrowser()
	defer browser.MustClose()

	r.HandleFunc("/", router.IndexPage)
	r.HandleFunc("/indeedtest1", router.IndeedScrapeTestPage)
	r.HandleFunc("/scrape/jobdesc", router.ScrapeJobDescription)

	fmt.Println("Now listening on port ", port)
	panic(http.ListenAndServe(":" + port, r))

}
