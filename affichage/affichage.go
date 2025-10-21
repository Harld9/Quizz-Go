package affichage

import "fmt"

func Separator() {
	fmt.Println("==================================================")
}

func AffichageMenu() {
	Separator()
	fmt.Println("Bienvenue au Quiz Go")
	Separator()
	fmt.Println("1 - ▶️  Commencer une nouvelle partie")
	fmt.Println("2 - 🪺  On verra plus tard")
	fmt.Println("3 - 👋 Quitter")
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
