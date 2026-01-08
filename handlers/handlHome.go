package handlers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {

		http.Error(w, "405", http.StatusMethodNotAllowed)
		return
	}

	template, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"text": "",
		"font": "standard",
	}

	er := template.Execute(w, data)
	if er != nil {
		http.Error(w, "Erreur lors de l'affichage du template", http.StatusInternalServerError)
		return
	}
}
