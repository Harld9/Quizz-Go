package logic

import (
	"Quizz-Go/affichage"
	"fmt"
)

type User struct {
	Name       string
	ScoreQ1    int
	ScoreQ2    int
	ScoreQ3    int
	TotalScore int
}

func InitUser() *User {
	affichage.NomUser()

	var userName string
	fmt.Scan(&userName)
	fmt.Println("Bienvenu :", userName)

	return &User{
		Name:       userName,
		ScoreQ1:    0,
		ScoreQ2:    0,
		ScoreQ3:    0,
		TotalScore: 0,
	}
}
