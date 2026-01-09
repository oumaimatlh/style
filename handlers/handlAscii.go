package handlers

import (
	"html/template"
	"net/http"
	"strings"
)

func AsciiController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		HandleError(w, http.StatusMethodNotAllowed, "405")
		return
	}

	if err := r.ParseForm(); err != nil {
		HandleError(w, http.StatusBadRequest, "400")
		return
	}

	inputText := r.PostForm.Get("content")
	font := r.PostForm.Get("types")

	var errorMsg string
	var result string

	if inputText == "" {
		errorMsg = "You need to write something..."
	} else if len(inputText) >= 3000 {
		errorMsg = "Your text has exceeded 3000 characters."
	} else {
		for _, ch := range inputText {
			if !(ch >= 32 && ch <= 126 || ch == '\n' || ch == '\r') {
				errorMsg = "Input not validated: characters must be in ASCII 32-126"
				break
			}
		}
	}

	// VALIDATION FONT
	if errorMsg == "" && font == "" {
		errorMsg = "You must choose an Art"
	}

	// APPLICATION FONT
	if errorMsg == "" {
		var err error
		result, err = ApplyingFont(inputText, font)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// GESTION DES NOUVELLES LIGNES
	if strings.HasPrefix(result, "\n") {
		result = "\n" + result
		inputText = "\n" + inputText
	}

	// SI ERREUR → on réaffiche HOME
	data := map[string]string{
		"text":   inputText,
		"font":   font,
		"result": result,
		"error":  errorMsg,
	}

	tmpl, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "500")
		return
	}

	if errorMsg != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	tmpl.Execute(w, data)
}
