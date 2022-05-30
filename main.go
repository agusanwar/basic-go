package main

import (
	"go-web/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", handler.HomeHandler) //root
	mux.HandleFunc("/about", handler.AboutHandler)
	mux.HandleFunc("/profile", handler.ProfileHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	// GET POST METODE
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.ProsesHandler)
	
	// load assets
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}


