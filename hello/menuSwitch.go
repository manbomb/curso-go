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

const OK_STATUS_CODE = 200

func main() {
	exibeIntro()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este programa")
			os.Exit(-1)
		}
	}
}

func exibeIntro() {
	fmt.Println("\nOlá senhor, escolha uma opção:")
}

func exibeMenu() {
	fmt.Println("\n1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair\n")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("\nO comando escolhido foi: ", comando, "\n")
	return comando
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")

	// SLICE
	sites := leSitesDoArquivo()

	//for i := 0; i < len(sites); i++ {
	for _, site := range sites {
		response := testaSite(site)
		if response {
			fmt.Println("Site:", site, "OK!")
			registraLog(site, true)
			continue
		}
		fmt.Println("Site:", site, "NOK...")
		registraLog(site, false)
	}
}

func testaSite(site string) bool {
	resp, err := http.Get(site)

	if err != nil {
		return false
	}

	if resp.StatusCode == OK_STATUS_CODE {
		return true
	} else {
		return false
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir log: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " :: " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
