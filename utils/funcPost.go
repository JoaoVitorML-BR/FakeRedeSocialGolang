package utils

import (
	"fmt"
	"modulo/models"
)

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

func ReadPosts() {

	if len(listposts) != 0 {	
		for i := range listposts {
			fmt.Println("Titulo: " + listposts[i].GetTitle())
			fmt.Println("Message: " + listposts[i].GetMessage())
		}	
	}else{
		fmt.Println("Não possui nenhum Post, você precisa postar algo.")
	}
}