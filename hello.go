package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 10
const intervalo = 5

func main() {

	Introducao()
	fmt.Println("")

	for {
		MostraMenu()
		op := Opcao()
		if op == 1 {
			IniciarMonitoraramento()
		} else if op == 2 {
			fmt.Println("Mostrando Logs...")
			MostarLogs()
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
	fmt.Println("A versão do seu sistema é " + versao + ".")
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

func IniciarMonitoraramento() {

	fmt.Println("--------------------------------------------------------")
	fmt.Println("Monitorando...")
	sites := LerSitesDoArquivo()

	for cont := 0; cont < monitoramentos; cont++ {
		for nsite, site := range sites {
			fmt.Println("Iniciando monitoramento no site:", nsite, "-> "+site)
			TestaSite(site)
		}
		time.Sleep(intervalo * time.Minute)
	}
}

func TestaSite(site string) {

	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		RegistraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. StatusCode:", resp.StatusCode)
		RegistraLog(site, false)
	}
	fmt.Println("--------------------------------------------------------")

}

func LerSitesDoArquivo() []string {

	sites := []string{}

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arq_lido := bufio.NewReader(arquivo)
	for {
		linha, err := arq_lido.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}

func RegistraLog(site string, status bool) {

	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro.")
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " Online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func MostarLogs() {

	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro")
	}

	fmt.Println(string(arquivo))
}
