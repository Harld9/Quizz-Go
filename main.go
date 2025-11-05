package main

import (
	"Quizz-Go/affichage"
	"Quizz-Go/quizzcyber"
	"fmt"
	"os"
	"time"
)

func main() {

	// Initialisation du choix du menu
	var menuChoice int
	// Boucle principale du jeu
	for {

		affichage.AffichageMenu()
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case 1:
			fmt.Print("\033[H\033[2J")

			affichage.Choixquizz()
			fmt.Scan(&menuChoice)
			switch menuChoice {
			case 3:
				quizzcyber.Quizzcyber()
			case 4:

				// Effacer l'écran
				fmt.Print("\033[H\033[2J")

			}
			if menuChoice == 4 {
				menuChoice = 0
				return
			}

		case 2:
		case 3:
			// Quitter le jeu
			os.Exit(0)
		default:

			fmt.Print("\033[H\033[2J")

			time.Sleep(1 * time.Second)

			// Choix invalide
			fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
		}
	}
}
