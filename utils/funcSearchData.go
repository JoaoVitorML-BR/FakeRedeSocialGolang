package utils

import "fmt"

func readName() {
	for i := range ListUsers {
		fmt.Println("Nome do usuário: " + ListUsers[i].Name())
	}
}

func readEmail() {
	for i := range ListUsers {
		fmt.Println("E-mail: " + ListUsers[i].Email())
	}
}
