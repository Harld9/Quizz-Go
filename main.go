package main

import (
	"Quizz-Go/affichage"
	"fmt"
	"os"
)

func main() {

	// Initialisation du choix du menu
	var menuChoice int
	affichage.ClearScreen()
	// Boucle principale du jeu
	for {
		menuChoice = 0
		// Affichage du menu de démarrage
		affichage.MenuAccueil()
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case 1:
			affichage.ClearScreen()
			for {
				menuChoice = 0
				affichage.MenuPrincipale()
				fmt.Scan(&menuChoice)
				switch menuChoice {
				case 1:
					affichage.ClearScreen()
					for {
						menuChoice = 0
						affichage.MenuQuizz()
						fmt.Scan(&menuChoice)
						switch menuChoice {
						case 1:
							affichage.ClearScreen()
							fmt.Println("Quizz Informatique")
							// Appeler la fonction du quizz info
						case 2:
							affichage.ClearScreen()
							fmt.Println("Quizz Cyber-Sécurité")
							// Appeler la fonction du quizz cyber sécurité
						case 3:
							affichage.ClearScreen()
							fmt.Println("Quizz Data")
							// Appeler la fonction du quizz data
						case 4:
							affichage.ClearScreen()
							menuChoice = 4
						default:
							affichage.ClearScreen()
							fmt.Println("Choix invalide, veuillez réessayer.")
						}
						if menuChoice == 4 {
							break
						}
					}
				case 2:
					affichage.ClearScreen()
					fmt.Println("Statistiques")
					// Faire une fonction qui montre le nom + score des quizz déjà fait + le score total
				case 3:
					affichage.ClearScreen()
					menuChoice = 3
				default:
					affichage.ClearScreen()
					fmt.Println("Choix invalide, veuillez réessayer.")
				}
				if menuChoice == 3 {
					break
				}
			}
		case 2:
			affichage.ClearScreen()
			fmt.Println("On verra ...")
			fmt.Println("Florian est un gros nul en Go 😂 et Harold adore sa copine avec une baguette de 7cm attaché sur elle *Miam*")
		case 3:
			affichage.ClearScreen()
			fmt.Println("Adios cheerios !")
			os.Exit(0)
		default:
			affichage.ClearScreen()
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
