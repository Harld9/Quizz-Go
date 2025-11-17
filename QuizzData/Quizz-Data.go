package data

import (
	"Quizz-Go/affichage"
	"Quizz-Go/logic"
)

var NomQuizz string = "Data"

var Questions []string = []string{
	"À quoi sert un modèle d’IA ?\n",
	"Que représente une donnée “bruitée” ?\n",
	"Pourquoi utilise‑t‑on des graphiques en Data ?\n",
	"Que signifie “entraîner un modèle” ?\n",
	"À quoi sert une base de données ?\n",
	"Qu’est‑ce qu’un dataset d’entraînement ?\n",
	"Que permet la prédiction d’un modèle ?\n",
	"Pourquoi nettoie‑t‑on les données avant analyse ?\n",
	"Que fait un modèle de classification ?\n",
	"À quoi sert un Data Engineer ?\n",
}
var Choix [][]string = [][]string{
	{"1 - À surveiller automatiquement l’activité réseau d’un ordinateur\n", "2 - À convertir des fichiers en différents formats\n", "3 - À effectuer une tâche en se basant sur des données apprises\n", "4 - À optimiser la qualité sonore d’un appareil\n"},                                  //Réponse 1
	{"1 - Une donnée compressée trop fortement\n", "2 - Une donnée envoyée par un micro ou un haut‑parleur\n", "3 - Une donnée utilisée pour tester les hautes fréquences\n", "4 - Une donnée qui contient des erreurs ou des valeurs incohérentes\n"},                                          //Réponse 2
	{"1 - Pour comprendre plus facilement des tendances dans les données\n", "2 - Pour optimiser l’affichage d’un jeu vidéo\n", "3 - Pour améliorer le rendu graphique d’une application\n", "4 - Pour mesurer la puissance d’une carte graphique\n"},                                           //Réponse 1
	{"1 - Ajuster le modèle grâce à des exemples pour qu’il donne de meilleurs résultats\n", "2 - L’envoyer sur un serveur externe pour le sécuriser\n", "3 - L’exécuter plusieurs fois pour tester sa stabilité\n", "4 - Modifier sa configuration pour consommer moins d’énergie\n"},          //Réponse 2
	{"1 - À gérer automatiquement les mises à jour d’un logiciel\n", "2 - À stocker et organiser les informations de manière structurée\n", "3 - À augmenter la vitesse d’un disque dur\n", "4 - À protéger un système contre les attaques\n"},                                                  //Réponse 1
	{"1 - Un fichier de configuration pour lancer des scripts IA\n", "2 - Un lot de données destiné uniquement aux tests réseau\n", "3 - Un ensemble de données utilisé pour apprendre à un modèle\n", "4 - Une copie temporaire des données système\n"},                                        //Réponse 3
	{"1 - De mesurer la température du processeur\n", "2 - De fournir un résultat basé sur ce qu’il a appris\n", "3 - De calculer la vitesse d’écriture d’un disque\n", "4 - D’installer automatiquement un pilote manquant\n"},                                                                 //Réponse 2
	{"1 - Pour améliorer la qualité visuelle des graphiques\n", "2 - Pour éviter que les résultats soient faussés\n", "3 - Pour permettre au modèle de s’entraîner plus rapidement, même sans données correctes\n", "4 - Pour réduire la taille d’un dataset afin qu’il occupe moins de RAM\n"}, //Réponse 1
	{"1 - Il gère la hiérarchie des permissions dans un serveur\n", "2 - Il divise automatiquement un dataset en plusieurs fichiers\n", "3 - Il place chaque élément dans une catégorie définie\n", "4 - Il vérifie la conformité des fichiers du système\n"},                                   //Réponse 4
	{"1 - À écrire les interfaces visuelles pour les logiciels d’entreprise\n", "2 - À créer des animations graphiques pour les applications IA\n", "3 - À optimiser les composants matériels utilisés par les modèles\n", "4 - À préparer, structurer et rendre les données exploitables\n"},   //Réponse 2
}
var RépCorrecte []int = []int{3, 4, 1, 1, 2, 3, 2, 2, 3, 4} // Numéros des réponses correctes

func QuizzData(u *logic.User) {
	if logic.PréQuizz(NomQuizz) {
		logic.QuestionnaireType(u, NomQuizz, Questions, Choix, RépCorrecte)
		return
	} else {
		affichage.ClearScreen()
		return
	}
}
