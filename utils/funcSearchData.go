package utils

import "fmt"

func readName() {
	for i := range ListUsers {
		fmt.Println("Nome do usuário: " + ListUsers[i].GetName())
	}
}

func readEmail() {
	for i := range ListUsers {
		fmt.Println("E-mail: " + ListUsers[i].GetEmail())
	}
}
