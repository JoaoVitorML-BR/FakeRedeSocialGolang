package main

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

func findUserRegister() {
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

func Login() {

	for {
		fmt.Println("Logado com sucesso!")
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		fmt.Println("")
		fmt.Println("Deseja fazer alguma postagem ?")
		fmt.Println("Sim - 1")
		fmt.Println("Ver postagens - 2")
		fmt.Println("Não/sair - 3")
		var op int
		_, _ = fmt.Scan(&op)

		if op == 1 {
			Postage()
		} else if op == 2 {
			readPosts()
		} else if op == 3 {
			fmt.Println("Ok, obrigado.")
			return
		} else {
			fmt.Println("Opção invalida, tente novamente")
			continue
		}
	}

}

func Postage() {

	var title string
	var message string
	for i := 0; i < 1; i++ {
		fmt.Println("Faça sua postagem")
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		fmt.Println("Titulo: ")
		fmt.Scan(&title)
		fmt.Println("Message: ")
		fmt.Scan(&message)

		if title != "" && title != " " && message != "" && message != " " {
			fmt.Println("Postagem feita com sucesso!")
			post := models.Post{
				Title:   title,
				Message: message,
			}
			listposts = append(listposts, post)
		} else {
			fmt.Println("Você precisa preencher todos os campos.")
		}

	}

}

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

func readPosts() {
	for i := range listposts {
		fmt.Println("Titulo: " + listposts[i].GetTitle())
		fmt.Println("Message: " + listposts[i].GetMessage())
	}
}
