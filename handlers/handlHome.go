package handlers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		HandleError(w, http.StatusNotFound, "404")
		return
	}

	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "405")
		return
	}

	tmpl, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "500")
		return
	}

	data := map[string]string{
		"text": "",
		"font": "standard",
	}

	tmpl.Execute(w, data)
}
