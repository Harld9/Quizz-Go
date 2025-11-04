package cyber

import (
	"Quizz-Go/affichage"
	"fmt"
)

func Quizzcyber() {
	affichage.Separator()
	fmt.Println("👾 Bienvenue au Cyber Quizz :")
	affichage.Separator()
	fmt.Println("1 - L'Innocent du Web (Facile)")
	fmt.Println("2 - Le Curieux Connecté (Moyen)")
	fmt.Println("3 - Le Surfeur Méfiant (Dur)")
	fmt.Println("4 - Le Gardien du Wifi (Très dur)")
	fmt.Println("5 - Le Maître du Cyber-Kung-Fu (Hardcore)")
	affichage.Separator()
}
