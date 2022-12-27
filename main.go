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

const monitoring = 5
const delay_minutes = 5

func main() {

	for {
		fmt.Println("----------------------------------------------")
		fmt.Println("                   MENU                       ")
		fmt.Println("----------------------------------------------")
		fmt.Println("1 - Iniciar Monitoramento")
		fmt.Println("2 - Mostrar Logs")
		fmt.Println("3 - Adicionar Sites")
		fmt.Println("4 - Fechar programa")
		fmt.Println("----------------------------------------------")

		fmt.Print("Opção: ")
		var option int
		fmt.Scan(&option)

		if option == 1 {
			testWebsites()
		} else if option == 2 {
			showLogs()
		} else if option == 3 {
			addWebsites()
		} else if option == 4 {
			os.Exit(1)
		} else {
			fmt.Println("----------------------------------------------")
			fmt.Println("Opção inválida.")
		}
	}
}

func testWebsites() {

	fmt.Println("----------------------------------------------")
	fmt.Println("Monitorando...")
	fmt.Println("----------------------------------------------")

	for {
		file, err := os.OpenFile("websites.txt", os.O_RDONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println("Ouve um erro, ao tentar abrir o arquivo.")
			break
		}

		sizeFile, err := file.Stat()

		if sizeFile.Size() == 0 {
			fmt.Println("Não há nenhum site adicionado a ser analizado.")
			break
		}

		if err != nil {
			fmt.Println("Ouve um erro, ao ler o arquivo.")
			break
		}

		readedFile := bufio.NewReader(file)

		for {

			site, err := readedFile.ReadString('\n')

			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println("Erro ao ler o arquivo.")
				break
			}

			site = strings.TrimSpace(site)
			response, err := http.Get("http://" + site)

			if err != nil {
				fmt.Println("Erro ao tentar monitorar o site:", err)
				continue
			}

			if response.StatusCode == 200 {
				fmt.Println(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "   Online: " + strconv.FormatBool(true))
				results := (time.Now().Format("02/01/2006 15:04:05") + " - " + site + "   Online: " + strconv.FormatBool(true) + "\n")
				registrateLogs(results)
			} else {
				fmt.Println(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "   Online: " + strconv.FormatBool(false) + " -> Status-Code: " + strconv.Itoa(response.StatusCode))
				results := (time.Now().Format("02/01/2006 15:04:05") + " - " + site + "   Online: " + strconv.FormatBool(false) + " -> Status-Code: " + strconv.Itoa(response.StatusCode) + "\n")
				registrateLogs(results)
			}

		}

		file.Close()
		break
	}
}

func showLogs() {
	fmt.Println("----------------------------------------------")
	fmt.Println("Mostrando Logs")
	fmt.Println("----------------------------------------------")

	file, err := os.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Erro ao ler arquivo de logs:", err)
	}

	fmt.Println(string(file))
}

func addWebsites() {
	fmt.Println("----------------------------------------------")
	fmt.Print("Digite o site a ser adicionado para análise: ")
	var site string
	fmt.Scan(&site)

	file, err := os.OpenFile("websites.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao adicionar site.")
	}

	_, err = file.WriteString(site + "\n")

	if err != nil {
		fmt.Println("Erro ao adicionar site.")
	}
}

func registrateLogs(results string) {

	file, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de logs.")
	}

	file.WriteString(results)
	file.Close()
}
