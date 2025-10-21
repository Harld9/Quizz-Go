package affichage

import "fmt"

func Separator() {
	fmt.Println("==================================================")
}

func MenuAccueil() {
	Separator()
	fmt.Println("Bienvenue au Quiz Go")
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
	fmt.Println("2 - 🛡️  Quizz cyber sécurité")
	fmt.Println("3 - 🗄️  Quizz Data")
	fmt.Println("4 - 👋 Retour")
	Separator()

}

sdlijdkdgjqcyuwcxsqfi