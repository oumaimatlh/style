package main

import (
	"fmt"
	"net/http"
	"ASCII-ART-WEB/handlers"
)


func main() {
	fmt.Println("Serveur lancé sur : http://localhost:8080")

	http.Handle("/static/",
		http.StripPrefix("/static/", 
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/ascii-art", handlers.AsciiController)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur:", err)
	}
}
