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
			case 1: //Data
			case 2: //Info
			case 3:
				fmt.Print("\033[H\033[2J")
				quizzcyber.Cyberquizz()
				fmt.Scan(&menuChoice)
				switch menuChoice {
				case 1:
					quizzcyber.Quizz1()
				case 2:
					quizzcyber.Quizz2()
				case 3:
					quizzcyber.Quizz3()
				case 4:
					quizzcyber.Quizz4()
				case 5:
					quizzcyber.Quizz5()
				case 6:
					fmt.Print("\033[H\033[2J")
					if menuChoice == 6 {
						menuChoice = 0
						return
					}

				}
			case 4:
				fmt.Print("\033[H\033[2J")
				if menuChoice == 4 {
					menuChoice = 0
					return
				}

		case 2: //Crédits
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
}
