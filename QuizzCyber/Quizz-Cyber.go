package quizzcyber

import (
	"Quizz-Go/affichage"
	"fmt"
)

type Question struct {
	Texte   string
	Choix   [4]string
	Correct int
}

func Cyberquizz() {
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

// Niveau facile
var QuestionsFacile = []Question{
	{
		Texte: "Qu'est-ce qu'un antivirus ?",
		Choix: [4]string{
			"1 - Un logiciel pour créer des virus",
			"2 - Un logiciel pour protéger contre les virus informatiques",
			"3 - Un logiciel pour accélérer l'ordinateur",
			"4 - Un logiciel pour stocker des photos",
		},
		Correct: 2,
	},
	{
		Texte: "Que signifie le mot 'phishing' ?",
		Choix: [4]string{
			"1 - Pêcher des poissons",
			"2 - Envoyer des mails frauduleux pour voler des informations",
			"3 - Hacker un site web",
			"4 - Crypter des fichiers",
		},
		Correct: 2,
	},
	{
		Texte: "Lequel de ces mots de passe est le plus sûr ?",
		Choix: [4]string{
			"1 - 123456",
			"2 - motdepasse",
			"3 - H7@lP9z!",
			"4 - azerty",
		},
		Correct: 3,
	},
	{
		Texte: "Qu'est-ce qu'un firewall (pare-feu) ?",
		Choix: [4]string{
			"1 - Un programme pour faire des graphiques",
			"2 - Un logiciel qui bloque certains accès Internet",
			"3 - Un jeu vidéo sur le feu",
			"4 - Une application pour transférer des fichiers",
		},
		Correct: 2,
	},
	{
		Texte: "Que fait un ransomware ?",
		Choix: [4]string{
			"1 - Il supprime tous vos fichiers",
			"2 - Il vous vole vos mots de passe",
			"3 - Il bloque vos fichiers et demande une rançon",
			"4 - Il accélère votre ordinateur",
		},
		Correct: 3,
	},
	{
		Texte: "Que signifie HTTPS dans une adresse web ?",
		Choix: [4]string{
			"1 - HyperText Transfer Protocol Secure",
			"2 - Hyper Transfer Text Protocol Standard",
			"3 - High Tech Transfer Protocol Secure",
			"4 - HyperText Transfer Private System",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un malware ?",
		Choix: [4]string{
			"1 - Un virus ou logiciel malveillant",
			"2 - Un navigateur web",
			"3 - Une application bancaire",
			"4 - Une clé USB",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un cookie sur un site web ?",
		Choix: [4]string{
			"1 - Un fichier contenant des informations sur votre navigation",
			"2 - Un virus informatique",
			"3 - Une image sur le site",
			"4 - Un mot de passe",
		},
		Correct: 1,
	},
	{
		Texte: "Que fait le chiffrement (cryptage) des données ?",
		Choix: [4]string{
			"1 - Il supprime les données",
			"2 - Il transforme les données pour qu'elles soient illisibles sans clé",
			"3 - Il les copie automatiquement",
			"4 - Il les imprime",
		},
		Correct: 2,
	},
	{
		Texte: "Qu'est-ce qu'un VPN ?",
		Choix: [4]string{
			"1 - Un logiciel pour pirater des sites",
			"2 - Un réseau privé virtuel pour sécuriser sa connexion",
			"3 - Une application de messagerie",
			"4 - Un type de virus",
		},
		Correct: 2,
	},
}

// Niveau Moyen
var questionsMoyen = []Question{
	{
		Texte: "Qu'est-ce qu'un VPN ?",
		Choix: [4]string{
			"1 - Un logiciel pour pirater des sites",
			"2 - Un réseau privé virtuel pour sécuriser sa connexion",
			"3 - Une application de messagerie",
			"4 - Un type de virus",
		},
		Correct: 2,
	},
	{
		Texte: "Que signifie '2FA' ?",
		Choix: [4]string{
			"1 - Fast File Access",
			"2 - Two-Factor Authentication",
			"3 - True File Algorithm",
			"4 - Top Firewall Access",
		},
		Correct: 2,
	},
	{
		Texte: "Qu'est-ce qu'un botnet ?",
		Choix: [4]string{
			"1 - Un groupe de robots physiques",
			"2 - Un réseau d'ordinateurs piratés contrôlé à distance",
			"3 - Une application antivirus",
			"4 - Un protocole internet",
		},
		Correct: 2,
	},
	{
		Texte: "Qu'est-ce qu'un keylogger ?",
		Choix: [4]string{
			"1 - Un logiciel qui enregistre les frappes au clavier",
			"2 - Un logiciel de navigation web",
			"3 - Un antivirus avancé",
			"4 - Une application de messagerie",
		},
		Correct: 1,
	},
	{
		Texte: "Quel protocole est sécurisé pour envoyer un email ?",
		Choix: [4]string{
			"1 - SMTP",
			"2 - HTTP",
			"3 - HTTPS",
			"4 - SMTPS",
		},
		Correct: 4,
	},
	{
		Texte: "Qu'est-ce qu'une attaque par 'brute force' ?",
		Choix: [4]string{
			"1 - Deviner un mot de passe par toutes les combinaisons possibles",
			"2 - Envoyer un virus sur un ordinateur",
			"3 - Copier des fichiers",
			"4 - Effacer le disque dur",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce que le 'social engineering' ?",
		Choix: [4]string{
			"1 - Utiliser des techniques pour tromper des personnes et obtenir des infos",
			"2 - Programmer des logiciels",
			"3 - Créer des réseaux sociaux",
			"4 - Développer un site web",
		},
		Correct: 1,
	},
	{
		Texte: "Que signifie DDoS ?",
		Choix: [4]string{
			"1 - Distributed Denial of Service",
			"2 - Direct Data on System",
			"3 - Digital Domain of Security",
			"4 - Dynamic Database of Servers",
		},
		Correct: 1,
	},
	{
		Texte: "Quel est le rôle d'un certificat SSL ?",
		Choix: [4]string{
			"1 - Chiffrer la connexion entre le navigateur et le serveur",
			"2 - Supprimer les virus",
			"3 - Accélérer la navigation",
			"4 - Sauvegarder les fichiers",
		},
		Correct: 1,
	},
	{
		Texte: "Quel est l'objectif d'un honeypot ?",
		Choix: [4]string{
			"1 - Piéger les pirates pour les analyser",
			"2 - Augmenter la vitesse d'un réseau",
			"3 - Sauvegarder les données",
			"4 - Supprimer les virus",
		},
		Correct: 1,
	},
}

// Niveau Difficile
var questionsDifficile = []Question{
	{
		Texte: "Qu'est-ce qu'une attaque par 'man-in-the-middle' ?",
		Choix: [4]string{
			"1 - Intercepter les communications entre deux parties",
			"2 - Créer un virus informatique",
			"3 - Supprimer des fichiers sur un serveur",
			"4 - Pirater un mot de passe par force brute",
		},
		Correct: 1,
	},
	{
		Texte: "Quel est le rôle d'un IDS (Intrusion Detection System) ?",
		Choix: [4]string{
			"1 - Détecter des intrusions dans un réseau",
			"2 - Bloquer tous les emails",
			"3 - Sauvegarder les données",
			"4 - Chiffrer les fichiers",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un exploit zero-day ?",
		Choix: [4]string{
			"1 - Une faille inconnue exploitée par des hackers",
			"2 - Un antivirus avancé",
			"3 - Un type de ransomware",
			"4 - Une application de sauvegarde",
		},
		Correct: 1,
	},
	{
		Texte: "Que signifie 'SQL Injection' ?",
		Choix: [4]string{
			"1 - Injection de code malveillant dans une base de données",
			"2 - Un logiciel antivirus",
			"3 - Un protocole réseau",
			"4 - Une méthode de chiffrement",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'pharming' ?",
		Choix: [4]string{
			"1 - Rediriger les utilisateurs vers un faux site pour voler des infos",
			"2 - Programmer des applications web",
			"3 - Installer un antivirus",
			"4 - Crypter les fichiers",
		},
		Correct: 1,
	},
	{
		Texte: "Quel protocole sécurise les communications Wi-Fi ?",
		Choix: [4]string{
			"1 - WPA2 / WPA3",
			"2 - HTTP",
			"3 - FTP",
			"4 - SMTP",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un rootkit ?",
		Choix: [4]string{
			"1 - Un logiciel permettant à un hacker de garder un accès secret à un ordinateur",
			"2 - Une application de sécurité",
			"3 - Un antivirus cloud",
			"4 - Une extension de navigateur",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un hash ?",
		Choix: [4]string{
			"1 - Une fonction qui transforme des données en une valeur fixe et unique",
			"2 - Un virus informatique",
			"3 - Une clé USB sécurisée",
			"4 - Un protocole de messagerie",
		},
		Correct: 1,
	},
	{
		Texte: "Que signifie 'MITRE ATT&CK' ?",
		Choix: [4]string{
			"1 - Une base de connaissances sur les techniques de cyberattaques",
			"2 - Un antivirus",
			"3 - Un type de ransomware",
			"4 - Une norme de chiffrement",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'une attaque par 'DNS spoofing' ?",
		Choix: [4]string{
			"1 - Modifier les adresses DNS pour rediriger les utilisateurs vers de faux sites",
			"2 - Voler des mots de passe en masse",
			"3 - Installer un malware automatiquement",
			"4 - Déchiffrer les emails",
		},
		Correct: 1,
	},
}

// Niveau Très Difficile
var questionsTresDifficile = []Question{
	{
		Texte: "Qu'est-ce qu'un 'buffer overflow' ?",
		Choix: [4]string{
			"1 - Une vulnérabilité qui permet d'exécuter du code arbitraire",
			"2 - Un type de virus",
			"3 - Un protocole de chiffrement",
			"4 - Une fonction de sauvegarde automatique",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'sandbox' en cybersécurité ?",
		Choix: [4]string{
			"1 - Un environnement isolé pour tester du code suspect",
			"2 - Un type de firewall",
			"3 - Une méthode de phishing",
			"4 - Un protocole de messagerie sécurisée",
		},
		Correct: 1,
	},
	{
		Texte: "Que signifie 'APT' ?",
		Choix: [4]string{
			"1 - Advanced Persistent Threat",
			"2 - Automatic Password Theft",
			"3 - Antivirus Protection Tool",
			"4 - Active Proxy Tunnel",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce que le 'Cross-Site Scripting (XSS)' ?",
		Choix: [4]string{
			"1 - Injection de scripts malveillants dans une page web",
			"2 - Vol de mots de passe",
			"3 - Chiffrement des données",
			"4 - Analyse des logs",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'certificate pinning' ?",
		Choix: [4]string{
			"1 - Vérification que le certificat SSL du serveur est celui attendu",
			"2 - Une méthode de chiffrement de fichier",
			"3 - Une attaque réseau",
			"4 - Une configuration de firewall",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce que le 'side-channel attack' ?",
		Choix: [4]string{
			"1 - Exploiter des informations indirectes comme le timing ou la consommation d'énergie",
			"2 - Pirater un mot de passe en force brute",
			"3 - Installer un malware",
			"4 - Modifier les DNS",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'une 'clickjacking' ?",
		Choix: [4]string{
			"1 - Tromper un utilisateur pour qu'il clique sur un élément invisible",
			"2 - Une attaque par injection SQL",
			"3 - Un ransomware",
			"4 - Une méthode de cryptage",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'token CSRF' ?",
		Choix: [4]string{
			"1 - Un mécanisme pour prévenir les attaques Cross-Site Request Forgery",
			"2 - Un antivirus",
			"3 - Une clé USB sécurisée",
			"4 - Un protocole réseau",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'une attaque 'pass-the-hash' ?",
		Choix: [4]string{
			"1 - Utiliser le hash d'un mot de passe pour s'authentifier sans connaître le mot de passe",
			"2 - Chiffrer les mots de passe",
			"3 - Voler les cookies",
			"4 - Pirater un VPN",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'white hat hacker' ?",
		Choix: [4]string{
			"1 - Un hacker éthique qui teste la sécurité pour la corriger",
			"2 - Un hacker qui vole des données",
			"3 - Un virus informatique",
			"4 - Une attaque réseau",
		},
		Correct: 1,
	},
}

// Niveau Hardcor
var questionsHardcore = []Question{
	{
		Texte: "Qu'est-ce qu'une attaque par 'Return-Oriented Programming (ROP)' ?",
		Choix: [4]string{
			"1 - Exploiter des morceaux de code existants pour exécuter du code malveillant",
			"2 - Créer un virus à partir de zéro",
			"3 - Pirater un mot de passe par force brute",
			"4 - Installer un VPN automatiquement",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'memory leak' exploitable en sécurité ?",
		Choix: [4]string{
			"1 - Une fuite mémoire pouvant permettre l'exécution de code malveillant",
			"2 - Une méthode pour effacer un virus",
			"3 - Une application de stockage",
			"4 - Un type de firewall",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'timing attack' ?",
		Choix: [4]string{
			"1 - Exploiter le temps de réponse d'un système pour récupérer des informations",
			"2 - Un protocole réseau sécurisé",
			"3 - Une attaque par malware",
			"4 - Une technique de chiffrement",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'heap spray' ?",
		Choix: [4]string{
			"1 - Remplir la mémoire avec du code malveillant pour faciliter l'exploitation",
			"2 - Installer un antivirus",
			"3 - Créer une clé USB sécurisée",
			"4 - Chiffrer les fichiers",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'cold boot attack' ?",
		Choix: [4]string{
			"1 - Récupérer des données sensibles à partir de la RAM après un redémarrage",
			"2 - Une attaque par phishing",
			"3 - Un protocole SSL",
			"4 - Un type de malware",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'Rowhammer attack' ?",
		Choix: [4]string{
			"1 - Exploiter des faiblesses physiques de la RAM pour modifier des bits",
			"2 - Voler des mots de passe en ligne",
			"3 - Chiffrer des fichiers",
			"4 - Installer un firewall",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'spectre/meltdown attack' ?",
		Choix: [4]string{
			"1 - Exploiter des failles matérielles dans les processeurs pour accéder à la mémoire",
			"2 - Un ransomware avancé",
			"3 - Une attaque réseau classique",
			"4 - Une méthode de chiffrement",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'logic bomb' ?",
		Choix: [4]string{
			"1 - Un programme qui s'exécute quand une condition spécifique est remplie",
			"2 - Un virus qui détruit tous les fichiers",
			"3 - Un antivirus",
			"4 - Une attaque réseau",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'advanced evasion technique' ?",
		Choix: [4]string{
			"1 - Méthode utilisée pour contourner les systèmes de détection",
			"2 - Un type de firewall",
			"3 - Une attaque simple de phishing",
			"4 - Un protocole de sécurité",
		},
		Correct: 1,
	},
	{
		Texte: "Qu'est-ce qu'un 'supply chain attack' ?",
		Choix: [4]string{
			"1 - Compromettre un fournisseur pour atteindre ses clients",
			"2 - Voler des mots de passe",
			"3 - Pirater un serveur FTP",
			"4 - Installer un antivirus",
		},
		Correct: 1,
	},
}
