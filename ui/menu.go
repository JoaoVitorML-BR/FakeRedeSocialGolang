package ui

import (
	"fmt"
	. "modulo/utils"
)

func openMainMenu() {
	for {
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		fmt.Println("Bem Vindo(a)!")
		fmt.Println()

		fmt.Println("Digite a opção que deseja fazer: ")

		fmt.Printf("Fazer cadastro - 1 \n")
		fmt.Printf("Fazer Login - 2 \n")
		fmt.Printf("Sair - 0 \n")
		fmt.Println("Digite a opção desejada: ")

		var op int
		_, _ = fmt.Scan(&op)

		if op == 1 {
			RegisterLister()
			continue
		} else if op == 2 {
			FindUserRegister()
			continue
		} else if op == 0 {
			fmt.Println("Ok, obrigado.")
			return
		}

	}
}
