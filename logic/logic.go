package logic

import (
	"Quizz-Go/affichage"
	"fmt"
)

type User struct {
	Name          string
	ScoreQ1       int
	NbQuestionsQ1 int
	ScoreQ2       int
	NbQuestionsQ2 int
	ScoreQ3       int
	NbQuestionsQ3 int
	Q1Percentage  int
	Q2Percentage  int
	Q3Percentage  int
}

func InitUser() *User {
	var username string
	var valid bool
	for {
		var discard string
		fmt.Scanln(&discard)
		affichage.NomUser()
		fmt.Scan(&username)
		valid = true
		result := []rune(username)

		// Vérification que le pseudo n'est pas vide et ne contient que des lettres
		if len(result) == 0 {
			fmt.Println("❌ Le pseudo ne peut pas être vide.")
			valid = false
			continue
		}

		// Vérification que chaque caractère est une lettre
		for _, r := range result {
			if r < 65 || (r > 90 && r < 97) || r > 122 {
				fmt.Println("❌ Le pseudo ne doit contenir que des lettres.")
				valid = false
				break
			}
		}

		// Si le pseudo est valide, on le formate (première lettre majuscule, le reste en minuscule)
		if valid {
			if result[0] >= 97 && result[0] <= 122 {
				result[0] = result[0] - ('a' - 'A')
			}
			for i := 1; i < len(result); i++ {
				if result[i] >= 65 && result[i] <= 90 {
					result[i] = result[i] + ('a' - 'A')
				}
			}

			// Assignation du pseudo au personnage
			username = string(result)
			// Sortie de la boucle
			break
		}
	}
	affichage.ClearScreen()
	affichage.Separator()
	fmt.Println("Bienvenu :", username)

	return &User{
		Name:         username,
		ScoreQ1:      0,
		ScoreQ2:      0,
		ScoreQ3:      0,
		Q1Percentage: 0,
		Q2Percentage: 0,
		Q3Percentage: 0,
	}
}
func ScorePercentage(user *User, TotalScore int) {
	if user.ScoreQ1 == 0 {
		user.Q1Percentage = 0
	} else {
		user.Q1Percentage = (user.ScoreQ1 * 100) / user.NbQuestionsQ1
	}

	if user.ScoreQ2 == 0 {
		user.Q2Percentage = 0
	} else {
		user.Q2Percentage = (user.ScoreQ2 * 100) / user.NbQuestionsQ2
	}

	if user.ScoreQ3 == 0 {
		user.Q3Percentage = 0
	} else {
		user.Q3Percentage = (user.ScoreQ3 * 100) / user.NbQuestionsQ3
	}
}
func UserStats(user *User) {
	var TotalScore int = user.ScoreQ1 + user.ScoreQ2 + user.ScoreQ3
	ScorePercentage(user, TotalScore)
	affichage.Separator()
	fmt.Println("Statistiques de l'utilisateur")
	affichage.Separator()
	fmt.Println("Nom :", user.Name)
	fmt.Println("Score Quizz Informatique :", user.ScoreQ1, "/", user.NbQuestionsQ1)
	fmt.Println("Score Quizz Cyber-Sécurité :", user.ScoreQ2, "/", user.NbQuestionsQ2)
	fmt.Println("Score Quizz Data :", user.ScoreQ3, "/", user.NbQuestionsQ3)
	fmt.Println("Score Total :", TotalScore)
	fmt.Println("Nombre de questions répondues :", user.NbQuestionsQ1+user.NbQuestionsQ2+user.NbQuestionsQ3)
	affichage.Separator()
	fmt.Println("Tu as : ", user.Q1Percentage, "% de score en Informatique")
	if user.Q1Percentage == 100 { //100 %
		fmt.Println("Psartek mon frère ! Tu est un crackhead de l'Informatique !")
	} else if user.Q1Percentage < 100 && user.Q1Percentage >= 80 { //80 à 99 %
		fmt.Println("Bravo ! Tu maîtrises bien l'Informatique !")
	} else if user.Q1Percentage < 80 && user.Q1Percentage >= 50 { //50 à 79 %
		fmt.Println("Bravo ! Tu comprends bien l'Informatique !")
	} else if user.Q1Percentage < 50 && user.Q1Percentage > 0 { //0 à 49 %
		fmt.Println("Tu es en difficulté en Informatique mais ne perd pas courage, tu vas y arriver !")
	} else { //0 %
		fmt.Println("Nul Nul Nul ! Tu est éclaté en Info ! Va falloir retourner chez ta mère et redémarer de 0 !")
	}
	affichage.Separator()
	fmt.Println("Tu as : ", user.Q2Percentage, "% de score en Cyber-Sécurité")
	if user.Q2Percentage == 100 { //100 %
		fmt.Println("Psartek mon frère ! Tu est un cyber-crack de l'espace !") //A compléter
	} else if user.Q2Percentage < 100 && user.Q2Percentage >= 80 { //80 à 99 %
		fmt.Println("Bravo ! Tu maîtrises bien la Cyber-Sécurité !")
	} else if user.Q2Percentage < 80 && user.Q2Percentage >= 50 { //50 à 79 %
		fmt.Println("Bravo ! Tu comprends bien la Cyber-Sécurité !")
	} else if user.Q2Percentage < 50 && user.Q2Percentage > 0 { //0 à 49 %
		fmt.Println("Tu es en difficulté en Cyber-Sécurité mais ne perd pas courage, tu vas y arriver !")
	} else { //0 %
		fmt.Println("Merde alors ! Tu est a chier en Cyber-Sécurité ! T'arriveras même pas rentrer chez toi avec tes compétences !")
	}
	affichage.Separator()
	fmt.Println("Tu as : ", user.Q3Percentage, "% de score en Data")
	if user.Q3Percentage == 100 { //100 %
		fmt.Println("Psartek mon frère ! Tu est la super IA de la Data !")
	} else if user.Q3Percentage < 100 && user.Q3Percentage >= 80 { //80 à 99 %
		fmt.Println("Bravo ! Tu maîtrises bien la Data !")
	} else if user.Q3Percentage < 80 && user.Q3Percentage >= 50 { //50 à 79 %
		fmt.Println("Bravo ! Tu comprends bien la Data !")
	} else if user.Q3Percentage < 50 && user.Q3Percentage > 0 { //0 à 49 %
		fmt.Println("Tu es en difficulté en Data mais ne perd pas courage, tu vas y arriver !")
	} else { //0 %
		fmt.Println("MDR ! Tu pues du fion en Data ! Retourne dans les jupettes de ChatGPT tu comprendra peut être comment ça marche !")
	}
	affichage.Separator()
}
func PréQuizz(nomQuizz string) bool { //Choix de commencer le quizz ou pas
	MenuChoice := 0
	for {
		affichage.PréQuizz(nomQuizz)
		fmt.Scan(&MenuChoice)
		switch MenuChoice {
		case 1:
			affichage.ClearScreen()
			fmt.Printf("Lancement du quizz %s...\n", nomQuizz)
			affichage.ClearScreen()
			return true

		case 2:
			affichage.ClearScreen()
			return false
		default:
			affichage.ClearScreen()
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
func QuestionnaireType(u *User, NomQuizz string, Questions []string, Choix [][]string, RépCorrecte []int) {
	//mettre : les stats de l'User, la liste des questions, les choix de réponses, la réponse correcte (le numéros de la réponse correcte et non l'index)
	scoreSession := 0
	for i := range Questions {
		for {
			affichage.QuestionType(NomQuizz, i+1, Questions[i], Choix[i])
			MenuChoice := 0
			fmt.Scan(&MenuChoice)
			if MenuChoice == RépCorrecte[i] { // Réponse correcte
				affichage.ClearScreen()
				affichage.BonneRéponse(Questions[i], Choix[i], RépCorrecte[i])
				AjoutScore(u, NomQuizz)
				AjoutNbQuestions(u, NomQuizz)
				scoreSession++
				break
			} else if MenuChoice != RépCorrecte[i] && MenuChoice <= 4 && MenuChoice >= 1 { // Réponse incorrecte
				affichage.ClearScreen()
				affichage.MauvaiseRéponse(Questions[i], Choix[i], MenuChoice, RépCorrecte[i])
				AjoutNbQuestions(u, NomQuizz)
				break
			}
		}
	}
	for {
		affichage.ClearScreen()
		affichage.FinQuizz(scoreSession, len(Questions))
		fmt.Println("Tapper sur '0' pour continuer.")
		MenuChoice := -1
		fmt.Scan(&MenuChoice)
		if MenuChoice == 0 {
			affichage.ClearScreen()
			break
		} else {
			affichage.Separator()
			fmt.Println("Bravo ! Tu n'a pas mis 0 ! Sais-tu ce que c'est le nombre 0 ? Demande à ton voisin tu verras c'est simple !")
			affichage.Separator()
		}
	}
}
func AjoutScore(u *User, NomQuizz string) {
	switch NomQuizz {
	case "Informatique":
		u.ScoreQ1++
	case "Cyber-Sécurité":
		u.ScoreQ2++
	case "Data":
		u.ScoreQ3++
	}
}

func AjoutNbQuestions(u *User, NomQuizz string) {
	switch NomQuizz {
	case "Informatique":
		u.NbQuestionsQ1++
	case "Cyber-Sécurité":
		u.NbQuestionsQ2++
	case "Data":
		u.NbQuestionsQ3++
	}
}
