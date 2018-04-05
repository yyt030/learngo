package main

import (
	"log"
	"net/http"

	"learngo/crawler/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		"crawler/frontend/view/searchresult.html"))

	log.Panic(http.ListenAndServe(":8888", nil))
}
