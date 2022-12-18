package utils

import (
	"fmt"
	"modulo/models"
)

var ListUsers []models.Register
var listposts []models.Post

func RegisterLister() {

	user := models.Register{}

	var name string
	var email string
	var password string

	fmt.Println("Bem-vindo a área de cadastro, digi te os seguintes valores para efetuar o seu cadastro.")

	fmt.Println("Nome: ")
	fmt.Scan(&name)
	user.SetName(name)

	fmt.Println("E-mail: ")
	fmt.Scan(&email)
	user.SetEmail(email)

	fmt.Println("Password: ")
	fmt.Scan(&password)

	if name != "" && email != "" && password != "" {
		fmt.Println("Usuário registrado com sucesso!")
	} else {
		fmt.Println("Você precisa preencher todos os campos corretamente, por favor, digite 1 e tente novamente.")
	}

	ListUsers = append(ListUsers, user)
}

func FindUserRegister() {
	var findUser string
	var password string

	if len(ListUsers) != 0 {
		for i := range ListUsers {
			fmt.Println("Faça seu Login:")
			fmt.Println("")
			fmt.Println("Digite seu e-mail: ")
			fmt.Scan(&findUser)
			fmt.Println("Digite sua senha: ")
			fmt.Scan(&password)
			if ListUsers[i].Email() == findUser {
				Login()
			} else {
				fmt.Println("E-mail ou senha incorreto, tente novamente.")
				continue
			}
		}
	} else {
		fmt.Println("Você precisa se cadastrar primeiro.")
	}
}