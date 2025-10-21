package affichage

import "fmt"

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
	fmt.Println("2 - 🛡️  Quizz Cyber-Sécurité")
	fmt.Println("3 - 🗄️  Quizz Data")
	fmt.Println("4 - 👋 Retour")
	Separator()

}

func NomUser() {
	Separator()
	fmt.Println("Veuillez entrer votre nom :")
	Separator()
}
