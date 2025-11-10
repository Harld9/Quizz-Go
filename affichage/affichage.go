package affichage

import "fmt"

func Separator() {
	fmt.Println("==================================================")
}

func AffichageMenu() {
	Separator()
	fmt.Println("Bienvenue au Quizz Go")
	Separator()
	fmt.Println("1 - ▶️  Commencer un Quizz")
	fmt.Println("2 - 🪺  Crédits")
	fmt.Println("3 - 👋 Quitter")
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
