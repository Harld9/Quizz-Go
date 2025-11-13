package router

import (
	"Quizz-Go/controller"
	"net/http"
)

// New crée et retourne un nouvel objet ServeMux configuré avec les routes de l'application
func New() *http.ServeMux {
	mux := http.NewServeMux()

	// Routes de ton app
	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/QuizzInfo", controller.Pinfo)
	mux.HandleFunc("/QuizzCyber", controller.Pcyber)
	mux.HandleFunc("/QuizzData", controller.Pdata)

	// Ajout des fichiers statiques
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
