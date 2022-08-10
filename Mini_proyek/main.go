package main

import (
	"Mini_proyek/handler"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/create", handler.CreateHandler)
	mux.HandleFunc("/browse", handler.BrowseHandler)
	mux.HandleFunc("/edit", handler.EditHandler)
	mux.HandleFunc("/about", handler.AboutHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
