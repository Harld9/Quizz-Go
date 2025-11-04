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
