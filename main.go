package main

import (
	"Quizz-Go/affichage"
	"fmt"
	"os"
)

func main() {

	// Initialisation du choix du menu
	var menuChoice int

	// Boucle principale du jeu
	for {
		menuChoice = 0
		// Effacer l'écran
		fmt.Print("\033[H\033[2J")

		// Affichage du menu de démarrage
		affichage.MenuAccueil()
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case 1:
			for {
				menuChoice = 0
				fmt.Print("\033[H\033[2J")

				affichage.MenuPrincipale()
				fmt.Scan(&menuChoice)
				switch menuChoice {
				case 1:
					fmt.Println("Nouveaux quizz")
					for {
						menuChoice = 0
						fmt.Print("\033[H\033[2J")

						affichage.MenuQuizz()
						fmt.Scan(&menuChoice)
						switch menuChoice {
						case 1:
							fmt.Println("Quizz Informatique")
							// Appeler la fonction du quizz info
						case 2:
							fmt.Println("Quizz cyber sécurité")
							// Appeler la fonction du quizz cyber sécurité
						case 3:
							fmt.Println("Quizz Data")
							// Appeler la fonction du quizz data
						case 4:
							menuChoice = 4
						default:
							fmt.Println("Choix invalide, veuillez réessayer.")
						}

						if menuChoice == 4 {
							break
						}
					}
				case 2:
					fmt.Println("Statistiques")
					// Faire une fonction qui montre le nom + score des quizz déjà fait + le score total
				case 3:
					menuChoice = 3
				default:
					fmt.Println("Choix invalide, veuillez réessayer.")
				}
				if menuChoice == 3 {
					break
				}
			}
		case 2:
			fmt.Println("On verra ...")
		case 3:
			fmt.Println("Adios cheerios !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
