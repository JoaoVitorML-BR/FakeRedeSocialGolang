package main

import (
	"fmt"
	"modulo/models"
)

var ListUsers []models.Register

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

	if(name != "" && email != "" && password != ""){
		fmt.Println("Usuário registrado com sucesso!")
	}else{
		fmt.Println("Você precisa preencher todos os campos corretamente, por favor, digite 1 e tente novamente.")
	}

	ListUsers = append(ListUsers, user)
}

func Login(){
	fmt.Println("Logado com sucesso!")
}

func readList() {
	fmt.Println(ListUsers)
}

func readName() {
	for i := range ListUsers {
		fmt.Println("Nome do usuário: " + ListUsers[i].Name())
	}
}

func readEmail(){
	for i := range ListUsers {
		fmt.Println("E-mail: " + ListUsers[i].Email())
	}
}

func findUserRegister(){
	var findUser string
	var password string

	fmt.Println("Faça seu Login:")

	if !(ListUsers == []){
		for i := range ListUsers{
			fmt.Println("Faça seu Login:")
			fmt.Println("")
			fmt.Println("Digite seu e-mail: ")
			fmt.Scan(&findUser)
			fmt.Println("Digite sua senha: ")
			fmt.Scan(&password)
			if (ListUsers[i].Email() == findUser){
				Login()
			}else{
				fmt.Println("E-mail não existe.")
			}
		}
	}else{
		fmt.Println("Você precisa se cadastrar primeiro.")
	}

}