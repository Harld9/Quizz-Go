package logic

import (
	"Quizz-Go/affichage"
	"fmt"
)

type User struct {
	Name         string
	ScoreQ1      int
	ScoreQ2      int
	ScoreQ3      int
	Q1Percentage int
	Q2Percentage int
	Q3Percentage int
}

func InitUser() *User {
	affichage.NomUser()

	var userName string
	fmt.Scan(&userName)
	fmt.Println("Bienvenu :", userName)

	return &User{
		Name:         userName,
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
		user.Q1Percentage = (user.ScoreQ1 * 100) / TotalScore
	}

	if user.ScoreQ2 == 0 {
		user.Q2Percentage = 0
	} else {
		user.Q2Percentage = (user.ScoreQ2 * 100) / TotalScore
	}

	if user.ScoreQ3 == 0 {
		user.Q3Percentage = 0
	} else {
		user.Q3Percentage = (user.ScoreQ3 * 100) / TotalScore
	}
}
func UserStats(user *User) {
	var TotalScore int = user.ScoreQ1 + user.ScoreQ2 + user.ScoreQ3
	ScorePercentage(user, TotalScore)
	affichage.Separator()
	fmt.Println("Statistiques de l'utilisateur")
	affichage.Separator()
	fmt.Println("Nom :", user.Name)
	fmt.Println("Score Quizz Informatique :", user.ScoreQ1)
	fmt.Println("Score Quizz Cyber-Sécurité :", user.ScoreQ2)
	fmt.Println("Score Quizz Data :", user.ScoreQ3)
	fmt.Println("Score Total :", TotalScore)
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

func QuestionnaireType(u *User, NomQuizz string, Questions []string, Choix []string, RépCorrecte []int) {
	//mettre : les stats de l'User, la liste des questions, les choix de réponses, la réponse correcte (le numéros de la réponse correcte et non l'index)
	for i := range Questions {
		affichage.QuestionType(NomQuizz, i+1, Questions[i], Choix[0:4])
		MenuChoice := 0
		fmt.Scan(&MenuChoice)
		if MenuChoice == RépCorrecte[i] {
			affichage.ClearScreen()
			affichage.BonneRéponse(Questions[i], Choix[0:4], RépCorrecte[i])

		} else {
			// Réponse incorrecte
		}
	}
}

func AjoutScore(u *User, NomQuizz string, Score int) {
	switch NomQuizz {
	case "Informatique":
		u.ScoreQ1 += Score
	case "Cyber-Sécurité":
		u.ScoreQ2 += Score
	case "Data":
		u.ScoreQ3 += Score
	}
}
