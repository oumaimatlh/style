package handlers

import (
	"html/template"
	"net/http"
	"strings"
)

func AsciiController(w http.ResponseWriter, r *http.Request) {
	var errorMsg string

	if r.Method != http.MethodPost {
		http.Error(w, "405: Bad Method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	inputText := r.PostForm.Get("content")

	if inputText == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorMsg = "Vous devez taper quelque chose..."
	} else {
		for _, r := range inputText {
			if !(r >= 32 && r <= 126 || r == 13 || r == 10) {
				w.WriteHeader(http.StatusBadRequest)
				errorMsg = "Input non validée : les caractères doivent être en ASCII 32-126"
				break
			}
		}
	}
	if len(inputText) >= 3000 {
		w.WriteHeader(http.StatusBadRequest)
		errorMsg = "Votre Text a dépassée 3000 caractéres"
	}

	font := r.PostForm.Get("types")
	if errorMsg == "" && font == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorMsg = "Vous devez choisir un Art"
	}

	var result string
	if errorMsg == "" {
		var err error
		result, err = ApplyingFont(inputText, font)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// handle new line at the beginning
	if strings.HasPrefix(result, "\n") {
		result = "\n" + result
		inputText = "\n" + inputText
	}

	data := map[string]string{
		"font":   font,
		"text":   inputText,
		"result": result,
		"error":  errorMsg,
	}

	template, e := template.ParseFiles("./templates/home.html")
	if e != nil {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	err := template.Execute(w, data)
	if err != nil {
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}
}
