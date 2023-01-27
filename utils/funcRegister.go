package utils

import (
	"fmt"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"modulo/models"
)

var ListUsers []models.Register
var listposts []models.Post

func RegisterLister() {

	newUser := models.Register{}

	var name string
	var email string
	var password string

	fmt.Println("Bem-vindo a área de cadastro, digi te os seguintes valores para efetuar o seu cadastro.")

	fmt.Println("Nome: ")
	fmt.Scan(&name)
	newUser.SetName(name)

	fmt.Println("E-mail: ")
	fmt.Scan(&email)
	newUser.SetEmail(email)

	fmt.Println("Password: ")
	fmt.Scan(&password)
	newUser.SetPassword(password)

	if name != "" && email != "" && password != "" {
		register := models.Register{
			Name:     name,
			Email:    email,
			Password: password,
		}
		ListUsers = append(ListUsers, register)
		fmt.Println("Usuário registrado com sucesso!")
	} else {
		fmt.Println("Você precisa preencher todos os campos corretamente, por favor, digite 1 e tente novamente.")
	}

	ListUsers = append(ListUsers, newUser)

	conn, err := getConnection()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	insert := `
			insert into users (name, email, password)
			values ($1, $2, $3)`

	_, err = conn.Exec(insert, newUser.Name, newUser.Email, newUser.Password)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Usuário inserido")
	}
}

var listusers []models.Register

func FindUserRegister() {
	var findUser string
	var password string

	conn, err := getConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Faça seu Login:")
	fmt.Println("")
	fmt.Println("Digite seu e-mail: ")
	fmt.Scan(&findUser)
	fmt.Println("Digite sua senha: ")
	fmt.Scan(&password)

	rows, err := conn.Query("SELECT email, password FROM users WHERE email=$1", findUser)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var dbEmail string
	var dbPassword string

	for rows.Next() {
		err := rows.Scan(&dbEmail, &dbPassword)
		if err != nil {
			panic(err)
		}
		if dbEmail == findUser && dbPassword == password {
			Login()
			return
		}
	}
	fmt.Println("E-mail ou senha incorreto, tente novamente.")
}
