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
	mux.HandleFunc("/leaderboard", controller.Pinfo)
	mux.HandleFunc("/contact", controller.Pcyber)
	mux.HandleFunc("/jeu", controller.Pdata)

	// Ajout des fichiers statiques
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
