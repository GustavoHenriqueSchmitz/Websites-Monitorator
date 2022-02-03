package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	Introducao()
	fmt.Println("")

	for {
		MostraMenu()
		op := Opcao()
		if op == 1 {
			Monitorar()
		} else if op == 2 {
			fmt.Println("Mostrando Logs...")
		} else if op == 3 {
			os.Exit(0)
		} else {
			fmt.Println("Comando Invalido")
			os.Exit(-1)
		}
	}
}

func Introducao() {

	nome := "Gustavo"
	versao := "1.2"

	fmt.Println("Olá,", nome)
	fmt.Println("A versão do seu sistema é", versao)
}

func MostraMenu() {

	fmt.Println("1 - Ativar Monitoramento")
	fmt.Println("2 - Mostrar Logs")
	fmt.Println("3 - Sair do Programa")
}

func Opcao() int {

	fmt.Print("Opção: ")
	var op int

	fmt.Scan(&op)
	fmt.Println("A opção escolhida foi:", op)

	return op
}

func Monitorar() {

	fmt.Println("--------------------------------------------------------")
	fmt.Println("Monitorando...")
	site := "https://alura.com.br/hello"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. StatusCode:", resp.StatusCode)
	}
	fmt.Println("--------------------------------------------------------")

}
