package controller

import (
	"Quizz-Go/logic"
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
	User    *logic.User
}

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Accueil",
		Message: "Bienvenue Au Jeu Puissance 4 🎉",
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func Pinfo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Accueil",
		Message: "Bienvenue Au Jeu Puissance 4 🎉",
	}
	tmpl := template.Must(template.ParseFiles("template/QuizzInfo.html"))
	tmpl.Execute(w, data)
}

func Pcyber(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Accueil",
		Message: "Bienvenue Au Jeu Puissance 4 🎉",
	}
	tmpl := template.Must(template.ParseFiles("template/QuizzCyber.html"))
	tmpl.Execute(w, data)
}

func Pdata(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Accueil",
		Message: "Bienvenue Au Jeu Puissance 4 🎉",
	}
	tmpl := template.Must(template.ParseFiles("template/QuizzData.html"))
	tmpl.Execute(w, data)
}
