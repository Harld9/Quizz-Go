package main

import (
	cyber "Quizz-Go/QuizzCyber"
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

		case 3:
			fmt.Print("\033[H\033[2J")
			cyber.Quizzcyber()
		default:
			fmt.Print("\033[H\033[2J")
			// Choix invalide
			fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
		}

		// Reset de la variable menuChoice pour éviter les boucles infinies
		if menuChoice == 6 {
			menuChoice = 0
			// Retour au menu précédent (menu de démarrage)
			break
		}
	}
}
