package ui

import (
	"fmt"
	."modulo/utils"
)

func OpenMainMenu() {
	for {
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		fmt.Println("Bem Vindo(a)!")
		fmt.Println()

		fmt.Println("Digite a opção que deseja fazer: ")

		fmt.Printf("Fazer cadastro - 1 \n")
		fmt.Printf("Fazer Login - 2 \n")
		fmt.Printf("Sair - 3 \n")
		fmt.Println("Digite a opção desejada: ")

		var op int
		_, _ = fmt.Scan(&op)

		if op == 1 {
			RegisterLister()
			continue
		} else if op == 2 {
			FindUserRegister()
			continue
		} else if op == 3 {
			fmt.Println("Ok, obrigado.")
			return
		}else{
			fmt.Println("Opção invalida, tente novamente.")
			continue
		}

	}
}
