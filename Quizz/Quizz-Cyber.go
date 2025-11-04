package cyber

import (
	"Quizz-Go/affichage"
	"fmt"
)

func Quizzcyber() {
	affichage.Separator()
	fmt.Println("👾 Bienvenue au Cyber Quizz :")
	affichage.Separator()
	fmt.Print("1 - L'Innocent du Web (Facile)")
	fmt.Print("2 - Le Curieux Connecté (Moyen)")
	fmt.Print("3 - Le Surfeur Méfiant (Dur)")
	fmt.Print("4 - Le Gardien du Wifi (Très dur)")
	fmt.Print("5 - Le Maître du Cyber-Kung-Fu (Hardcore)")
	affichage.Separator()
}
