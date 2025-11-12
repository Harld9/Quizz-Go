package info

import (
	"Quizz-Go/affichage"
	"Quizz-Go/logic"
	"fmt"
)


NomQuizz := "Informatique"

Questions := []string{
			"Quel langage de programmation est principalement utilisé pour le développement d'applications iOS ?",
			"Quel protocole est utilisé pour sécuriser les communications sur Internet ?",
			"Quelle structure de données utilise le principe FIFO (First In, First Out) ?",
		}
Choix := []string{
			"1. Swift", "2. Java", "3. Python", "4. C#",
			"1. HTTP", "2. FTP", "3. SSL/TLS", "4. SMTP",
			"1. Pile", "2. File d'attente", "3. Arbre binaire", "4. Graphe",
		}
RépCorrecte := []int{1, 3, 2} // Numéros des réponses correctes

	
func QuizzInfo(u *logic.User) {
	if logic.PréQuizz(NomQuizz) {
		logic.QuestionnaireType(u, NomQuizz, Questions, Choix, RépCorrecte)
	}else {
		affichage.ClearScreen()
		return
	}
}