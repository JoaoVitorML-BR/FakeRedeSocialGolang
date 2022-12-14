package main

import (
	"fmt"
	"modulo/models"
)

var ListUsers []models.Register

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
	

	ListUsers = append(ListUsers, user)

}
