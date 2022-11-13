package main

import (
	"fmt"
	controller "mathe/golang/v2/controllers"
	"mathe/golang/v2/urls"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc(urls.UrlsV1.Index, controller.HandleIndex)

	http.HandleFunc("/search", controller.GetSearch)
	http.HandleFunc("/", controller.GetIndex)

	fmt.Println("Server up and runnning at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
