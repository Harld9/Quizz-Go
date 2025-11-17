package affichage

import (
	"fmt"
)

func Separator() {
	fmt.Println("==================================================")
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
	/*
		- \033 :
		C'est la séquence d'échappement ASCII pour ESC (escape),
		utilisée pour envoyer des commandes de contrôle au terminal.
		- [H :
		Après ESC, [H est une commande ANSI qui déplace le curseur à la position "home",
		c'est-à-dire en haut à gauche du terminal (ligne 1, colonne 1).
		- \033[2J :
		ESC suivi de [2J est une commande ANSI pour effacer tout le contenu de l'écran.
	*/
}

func MenuAccueil() {
	Separator()
	fmt.Println("Bienvenue au Quizz Go")
	Separator()
	fmt.Println("1 - ▶️  Commencer une nouvelle partie")
	fmt.Println("2 - 🪺  On verra plus tard")
	fmt.Println("3 - 👋 Quitter le jeu")
	Separator()
}

func MenuPrincipale() {
	Separator()
	fmt.Println("Choix principal")
	Separator()
	fmt.Println("1 - 🆕 Nouveaux quizz")
	fmt.Println("2 - 📚 Statistiques")
	fmt.Println("3 - 👋 Retour à l'accueil")
	Separator()
}
func MenuQuizz() {
	Separator()
	fmt.Println("Choix du quizz")
	Separator()
	fmt.Println("1 - 🖥️  Quizz Informatique")
	fmt.Println("2 - 🛡️  Quizz Cyber-Sécurité")
	fmt.Println("3 - 🗄️  Quizz Data")
	fmt.Println("4 - 👋 Retour")
	Separator()

}

func NomUser() {
	Separator()
	fmt.Println("Entres ton nom :")
	Separator()
}

func Statistiques() {
	Separator()
	fmt.Println("Statistiques de l'utilisateur")
	Separator()
}

func PréQuizz(Nom string) {
	Separator()
	fmt.Printf("Début du quizz %s\n", Nom)
	fmt.Println("Es-tu sûr ?")
	Separator()
	fmt.Println("1 - 👍  Oui ! Let's get this party rocking !")
	fmt.Println("2 - 👎  Non, je veux choisir un autre quizz.")
	Separator()
}

func QuestionType(nomQuizz string, numQuestion int, question string, listeChoix []string) {
	Separator()
	fmt.Printf("Quizz %s - Question n°%d\n", nomQuizz, numQuestion)
	Separator()
	fmt.Println(question)
	Separator()
	for _, choix := range listeChoix {
		fmt.Print(choix)
	}
	Separator()
}

func BonneRéponse(question string, choix []string, répCorrecte int) {
	Separator()
	fmt.Println("✅ Bonne réponse !")
	Separator()
	fmt.Println("La question :")
	fmt.Println(question)
	fmt.Println("Ta réponse : ")
	fmt.Println(choix[répCorrecte-1])
	Separator()
}

func MauvaiseRéponse(question string, choix []string, choixJoueur int, répCorrecte int) {
	Separator()
	fmt.Println("❌ Mauvaise réponse !")
	Separator()
	fmt.Println("La question :")
	fmt.Println(question)
	fmt.Println("Ta réponse : ")
	fmt.Println(choix[choixJoueur-1])
	fmt.Println("La bonne réponse est : ")
	fmt.Println(choix[répCorrecte-1])
	Separator()
}

func FinQuizz(scoreSession int, totalQuestions int) {
	Separator()
	fmt.Println("🎉 Fin du quizz ! 🎉")
	Separator()
	fmt.Printf("Votre score : %d/%d\n", scoreSession, totalQuestions)
	Separator()
}

func Choixquizz() {
	Separator()
	fmt.Println("Bienvenue au Quizz Go")
	Separator()
	fmt.Println("1 - Quizz Data")
	fmt.Println("2 - Quizz Info")
	fmt.Println("3 - Quizz Cyber")
	fmt.Println("4 - 👋 Retour")
	Separator()
}

func AffichageMenudata() {
	Separator()
	fmt.Println("Bienvenue au Quiz Data et IA")
	Separator()
	fmt.Println("1 - ▶️  Commencer une nouvelle partie")
	fmt.Println("2 - 🪺  On verra plus tard")
	fmt.Println("3 - 👋 Quitter")
	Separator()
}
