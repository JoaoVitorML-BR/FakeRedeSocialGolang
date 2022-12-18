package utils

import (
	"fmt"
)

func Login() {

	for {
		fmt.Println("Logado.")
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
			ReadPosts()
		} else if op == 3 {
			fmt.Println("Ok, obrigado.")
			return
		} else {
			fmt.Println("Opção invalida, tente novamente")
			continue
		}
	}

}