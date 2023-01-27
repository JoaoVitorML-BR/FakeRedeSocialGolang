package utils

import "fmt"

func Login() {
	for {
		fmt.Println("Logado.")
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		fmt.Println("")
		fmt.Println("Deseja fazer alguma postagem ?")
		fmt.Println("Sim - 1")
		fmt.Println("Ver postagens - 2")
		fmt.Println("Deletar um post - 3")
		fmt.Println("Atualizar um post - 4")
		fmt.Println("Não/sair - 5")
		var op int
		_, _ = fmt.Scan(&op)

		if op == 1 {
			Postage()
		} else if op == 2 {
			ReadPosts()
		} else if op == 3 {
			DeletePost()
		} else if op == 4 {
			updateIPostages()
		} else if op == 5 {
			fmt.Println("Ok, obrigado.")
			return
		} else {
			fmt.Println("Opção inválida, tente novamente")
			continue
		}
	}
}
