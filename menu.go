package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const delay_minutos = 5

func main() {
	for {
		Menu()
	}
}

func Menu() {

	fmt.Println("----------------------------------------------")
	fmt.Println("                   MENU                       ")
	fmt.Println("----------------------------------------------")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Mostrar Logs")
	fmt.Println("3 - Sites")
	fmt.Println("4 - Configurações")
	fmt.Println("5 - Sair do programa")
	fmt.Println("----------------------------------------------")

	fmt.Print("Opção: ")
	var op int
	fmt.Scan(&op)

	if op == 1 {
		IniciarMonitoramento()
	} else if op == 2 {
		MostarLogs()
	} else if op == 3 {
		Sites()
		var op string
		fmt.Scan(&op)
	} else if op == 4 {
		Configurar()
	} else if op == 5 {
		os.Exit(1)
	}
}

func IniciarMonitoramento() {

	fmt.Println("----------------------------------------------")
	fmt.Println("Monitorando...")
	fmt.Println("----------------------------------------------")
	TestaSites()
}

func MostarLogs() {

	fmt.Println("----------------------------------------------")
	fmt.Println("Mostrando Logs")
	fmt.Println("----------------------------------------------")
}

func Sites() {

	fmt.Println("----------------------------------------------")
	fmt.Print("Digite o site, a ser adicionado para análise: ")
}

func TestaSites() {

	for {
		arquivo, err := os.OpenFile("sites.txt", os.O_RDONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println("Ouve um erro, ao tentar abrir o arquivo.")
			break
		}

		tamanho_arq, err := arquivo.Stat()

		if tamanho_arq.Size() == 0 {
			fmt.Println("Nã há nenhum site, a ser analizado.")
			fmt.Println("Adicione, na opção sites do menu.")
			break
		}

		if err != nil {
			fmt.Println("Ouve um erro, ao ler o arquivo.")
			break
		}

		arq_lido := bufio.NewReader(arquivo)

		for {

			site, err := arq_lido.ReadString('\n')

			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println("Erro ao ler o arquivo.")
				break
			}

			site = strings.TrimSpace(site)
			resp, err := http.Get(site)

			if err != nil {
				fmt.Println("Erro ao tentar monitorar o site.")
			}

			if resp.StatusCode == 200 {
				fmt.Println(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "   Online: " + strconv.FormatBool(true))
				resultado := (time.Now().Format("02/01/2006 15:04:05") + " - " + site + "   Online: " + strconv.FormatBool(true) + "\n")
				RegistraLogs(string(resultado))
			} else {
				fmt.Println(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "   Online: " + strconv.FormatBool(false) + " -> Status-Code: " + string(resp.StatusCode))
				resultado := (time.Now().Format("02/01/2006 15:04:05") + " - " + site + "   Online: " + strconv.FormatBool(false) + " -> Status-Code: " + strconv.Itoa(resp.StatusCode) + "\n")
				RegistraLogs(string(resultado))
			}

		}

		arquivo.Close()
		break
	}
}

func RegistraLogs(resultado string) {

	arquivo, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de logs.")
	}

	arquivo.WriteString(resultado)
	arquivo.Close()
}

func Configurar() {

}
