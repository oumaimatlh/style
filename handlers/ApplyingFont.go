package handlers

import (
	"fmt"
	"os"
	"strings"
)

func ApplyingFont(inputText string, font string) (string, error) {
	fileName := font + ".txt"

	g, err := os.ReadFile("./fonts/" + fileName)
	if err != nil {
		return "", fmt.Errorf("erreur lecture fichier: %v", err)
	}

	graphicAscii := strings.ReplaceAll(string(g), "\r", "")
	lines := strings.Split(graphicAscii, "\n")

	if len(lines) != 856 {
		return "", fmt.Errorf("données incorrectes dans le fichier de font %s", font)
	}

	// Création du map des caractères
	f := make(map[rune][]string)
	for i := 32; i < 127; i++ {
		start := (i - 32) * 9
		f[rune(i)] = lines[start+1 : start+9]
	}

	// Gestion des sauts de ligne
	text := strings.Split(strings.ReplaceAll(inputText, "\r\n", "\n"), "\n")
	result := ""

	for _, c := range text {
		if c == "" {
			result += "\n"
			continue
		}

		for i := 0; i < 8; i++ {
			for _, b := range c {
				artC := f[b]
				result += artC[i]
			}
			result += "\n"
		}
	}
	return result, nil
}
