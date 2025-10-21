package main

import (
	"Quizz-Go/affichage"
	"fmt"
)

func main() {

	// Initialisation du choix du menu
	var menuChoice int

	// Boucle principale du jeu
	for {
		// Effacer l'écran
		fmt.Print("\033[H\033[2J")

		// Affichage du menu de démarrage
		affichage.AffichageMenu()
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case 1:
			fmt.Println("On verra ...")

		case 2:
			fmt.Println("On verra ...")
		}

	}
}
