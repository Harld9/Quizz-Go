package data

import (
	"Quizz-Go/affichage"
	"Quizz-Go/logic"
)

var NomQuizz string = "Informatique"

var Questions []string = []string{
	"Qui a créé le langage Go (Golang) ?\n",
	"Le langage Go n'est pas conçu pour… ?\n",
	"Que fait l'instruction suivante ? :\n fmt.Println(Hello, world!)\n",
	"Quelle est la sortie du code suivant ? :\n nums := []int{1, 2, 3}\n fmt.Println(nums[1])\n",
	"Que veut dire HTML ?\n",
	"À quoi sert la balise <title> ?\n",
	"Le HTML ne sert pas à quoi ?\n",
	"Que veut dire CSS ?\n",
	"Que fait ce code CSS ? :\n body { background-color: blue; }\n",
	"Quelle est la différence entre class et id en CSS ?\n",
}
var Choix [][]string = [][]string{
	{"1 - Google\n", "2 - Microsoft\n", "3 - Apple\n", "4 - IBM\n"},                                                                                                                                                    //Réponse 1
	{"1 - Langage de programmation\n", "2 - Système d'exploitation\n", "3 - Base de données\n", "4 - Serveur web\n"},                                                                                                   //Réponse 2
	{"1 - Affiche du texte à l'écran\n", "2 - Lit une entrée utilisateur\n", "3 - Calcule une somme\n", "4 - Ouvre un fichier\n"},                                                                                      //Réponse 1
	{"1 - 1\n", "2 - 2\n", "3 - 3\n", "4 - Erreur d'exécution\n"},                                                                                                                                                      //Réponse 2
	{"1 - HyperText Markup Language\n", "2 - Hyperlinks and Text Markup Language\n", "3 - Home Tool Markup Language\n", "4 - Hyperlinking Text Marking Language\n"},                                                    //Réponse 1
	{"1 - Affiche une image\n", "2 - Crée un lien vers une autre page\n", "3 - Définit le titre de la page web\n", "4 - Ajoute un paragraphe\n"},                                                                       //Réponse 3
	{"1 - Structurer le contenu d'une page web\n", "2 - Styliser une page web\n", "3 - Ajouter des interactions à une page web\n", "4 - Créer des formulaires\n"},                                                      //Réponse 2
	{"1 - Cascading Style Sheets\n", "2 - Computer Style Sheets\n", "3 - Creative Style System\n", "4 - Colorful Style Sheets\n"},                                                                                      //Réponse 1
	{"1 - Aucune de ces réponses\n", "2 - Change la couleur du texte en bleu\n", "3 - Définit la police du texte en bleu\n", "4 - Change la couleur de fond du corps de la page en bleu\n"},                            //Réponse 4
	{"1 - Un id peut être utilisé plusieurs fois, une class est unique\n", "2 - Une class peut être utilisée plusieurs fois, un id est unique\n", "3 - Il n'y a pas de différence\n", "4 - Les deux sont obsolètes\n"}, //Réponse 2
}
var RépCorrecte []int = []int{1, 2, 1, 2, 1, 3, 2, 1, 4, 2} // Numéros des réponses correctes

func QuizzData(u *logic.User) {
	if logic.PréQuizz(NomQuizz) {
		logic.QuestionnaireType(u, NomQuizz, Questions, Choix, RépCorrecte)
		return
	} else {
		affichage.ClearScreen()
		return
	}
}
