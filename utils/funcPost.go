package utils

import (
	"bufio"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"modulo/models"
	"os"
	"strconv"
)

func Postage() {
	newPost := models.Post{}

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
			newPost.Title = title
			newPost.Message = message
		} else {
			fmt.Println("Você precisa preencher todos os campos.")
			continue
		}

		conn, err := getConnection()
		if err != nil {
			panic(err)
		}

		defer conn.Close()

		insert := `
		insert into postages (title, message)
		values ($1, $2)`

		_, err = conn.Exec(insert, newPost.Title, newPost.Message)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Postagem inserida!")
		}

	}
}

func ReadPosts() {
	conn, err := getConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT id, title, message FROM postages")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var id, title, message string

	for rows.Next() {
		err := rows.Scan(&id, &title, &message)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID: " + id)
		fmt.Println("Titulo: " + title)
		fmt.Println("Message: " + message)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}

func updateIPostages() {
	Scan := bufio.NewScanner(os.Stdin)
	conn, err := getConnection()
	if err != nil {
		returnErr(err)
	}
	defer conn.Close()

	updateLine := `update postages SET title = $2, message = $3 where id = $1`

	fmt.Print("Digite o id do titulo atual: ")
	Scan.Scan()
	IdTitle, _ := strconv.Atoi(Scan.Text())

	fmt.Print("Digite o titulo atualizado: ")
	Scan.Scan()
	currentTitle := Scan.Text()

	fmt.Print("Digite o a mensagem atualizada: ")
	Scan.Scan()
	currentMessage := Scan.Text()

	_, err = conn.Exec(updateLine, IdTitle, currentTitle, currentMessage)
	if err != nil {
		returnErr(err)
	}
}

func DeletePost() {
	Scan := bufio.NewScanner(os.Stdin)

	conn, err := getConnection()
	if err != nil {
		returnErr(err)
	}

	defer conn.Close()

	deleteLine := `delete from postages where id = $1`

	fmt.Print("Digite o id do post para ser apagado: ")
	Scan.Scan()
	finder := Scan.Text()

	_, err = conn.Exec(deleteLine, finder)
	if err != nil {
		returnErr(err)
	}

	println("Postagem deletada com sucesso.")
}

func getConnection() (*sqlx.DB, error) {
	conn, err := sqlx.Open("postgres", "postgres://postgres:12345678@localhost:5432/exemplo_db?sslmode=disable")
	return conn, err
}

func returnErr(err error) {
	log.Fatal(err)
}
