package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"main/internal/rod"
	"main/router"
)


func main(){
	fmt.Println("Indeed scraper")

	router := router.NewRouter()

	browser := rod.GetBrowser()
	defer browser.MustClose()

	page := browser.MustPage("https://google.com")
	fmt.Println("page: ", page)

}
