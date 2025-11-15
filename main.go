package main

import (
	quizzcyber "Quizz-Go/QuizzCyber"
	info "Quizz-Go/QuizzInfo"
	"Quizz-Go/affichage"
	"Quizz-Go/logic"
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
			User := logic.InitUser()
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
							info.QuizzInfo(User)
						case 2:
							affichage.ClearScreen()
							fmt.Println("Quizz Cyber-Sécurité")
							quizzcyber.QuestionCyberGlobal(User)
						case 3:
							affichage.ClearScreen()
							fmt.Println("Quizz Data")
							// Appeler la fonction du quizz data
						case 4:
							affichage.ClearScreen()
							menuChoice = 4
						default:
							affichage.ClearScreen()
							fmt.Println("Choix invalide, réessayes.")
						}
						if menuChoice == 4 {
							break
						}
					}
				case 2:
					affichage.ClearScreen()
					logic.UserStats(User)
				case 3:
					affichage.ClearScreen()
					menuChoice = 3
				default:
					affichage.ClearScreen()
					fmt.Println("Choix invalide, réessayes.")
				}
				if menuChoice == 3 {
					break
				}
			}
		case 2:
			affichage.ClearScreen()
			fmt.Println("Bienvenu ! Voici un jeu de quizz dans les thèmes de l'informatique/dévelopement, de la cyber sécurité, et de la data !\n Il y a 10 questions par quizz et tu pourras voir tes statistiques dans un menu !")
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
