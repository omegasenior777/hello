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

const monitoramentos = 3
const delay = 5

func main() {

	exibeInroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibeLogs()
		case 0:
			sair()
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeMenu() {
	fmt.Println("1-Iniciar monitoramento")
	fmt.Println("2-Exibir logs")
	fmt.Println("0-Sair")
}

func exibeInroducao() {
	nome := "NOME"
	versao := 1.0
	fmt.Println("Hello world!", nome)
	fmt.Println("Versão", versao)
}

func leComando() int {
	var comandolido int
	fmt.Scan(&comandolido)
	fmt.Println("O comando escolhido foi", comandolido)
	return comandolido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//sites := []string{}
	// sites = append(sites, "https://www.alura.com.br")
	// sites = append(sites, "https://www.caelum.com.br")
	// sites = append(sites, "https://www.google.com")
	sites := leSitesArquivo()

	for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }

        // adição AQUI!
        time.Sleep(delay * time.Second)
        fmt.Println("")
    }
    fmt.Println("")
}

func exibeLogs() {
	fmt.Println("Exibindo logs...")
	arquivo, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro:", err)
		 
	}
	fmt.Println(string(arquivo))
}

func sair() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}

func exibeNomes(){
	nomes := []string{""}
	fmt.Println(nomes)
}

func testaSite(site string) {

    resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro:", err)
		 
	}

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
    }
}
func leSitesArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := os.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Erro:", err)
			 
	}
	//fmt.Println(string(arquivo))
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		fmt.Println(linha)
		if err == io.EOF {
			fmt.Println("Erro:", err)
			break			 
		}
		
	}
	fmt.Println(sites)
	arquivo.Close()
	return sites
}
func registraLog(site string, status bool){
	arqivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		fmt.Println(err)
	}

	arqivo.WriteString(time.Now().Format("02/01/2006 15:04:05")+ " - " + site + "- Online: " + strconv.FormatBool(status) + "\n")
	arqivo.Close()
}