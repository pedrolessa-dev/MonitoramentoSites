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

const monitoringCount = 3
const delay = 2

func main() {
	for {
		fmt.Println("1 - Monitorar\n2 - Exibir logs\n0 - Encerrar")
		input := readInput()
		switch input {
		case 1:
			fmt.Println("iniciando monitoração...")
			monitor("sites.txt")
		case 2:
			fmt.Println("carregando logs...")
			showLogs("logs.txt")
		case 0:
			fmt.Println("saindo...")
			os.Exit(0)
		default:
			fmt.Print(input, " não é um valor válido\n\n")
		}
	}

}

func readInput() int {
	var input int = 1
	fmt.Scan(&input)
	return input
}

func monitor(pathFile string) {
	sites := readSites(pathFile)

	for x := 0; x < monitoringCount; x++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
}

func testSite(site string) {
	res, err := http.Get(site)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if res.StatusCode == 200 {
		fmt.Println("Site: '" + site + "' retornou com sucesso")
		registerLogs("logs.txt", site, true, res.StatusCode)
	} else {
		fmt.Println("Site: '"+site+"' retornou com erro", res.StatusCode)
		registerLogs("logs.txt", site, false, res.StatusCode)
	}
}

func readSites(pathFile string) []string {
	var sites []string
	file, err := os.Open(pathFile)
	if err != nil {
		fmt.Println("Error:", err)
		return sites
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func registerLogs(pathFile string, site string, status bool, statusCode int) {
	file, err := os.OpenFile(pathFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	file.WriteString(time.Now().In(location).Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + " - código HTTP: " + strconv.Itoa(statusCode) + "\n")
	file.Close()
}

func showLogs(logsFile string) {
	file, err := os.ReadFile(logsFile)
	if _, err := os.Stat(logsFile); os.IsNotExist(err) {
		fmt.Print("arquivo de logs não encontrado. Por favor, inicie a monitoração primeiro\n\n")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(file))
}
