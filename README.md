# Website Monitoring Script

Este é um script em Go que monitora uma lista de sites e registra os resultados dos testes em um arquivo de log.

## Funcionalidades

1. Monitoramento de sites: verifica a disponibilidade de sites listados em um arquivo.
2. Registro de logs: registra os resultados dos testes em um arquivo de log, incluindo data, hora, URL do site, status online e código HTTP.
3. Exibição de logs: exibe o conteúdo do arquivo de log.

## Como usar

### Pré-requisitos

- Go instalado no seu sistema.

### Passos

1. Clone o repositório ou copie o arquivo `script.go` para o seu ambiente local.
2. Crie um arquivo chamado `sites.txt` no mesmo diretório que contém o script `script.go`. Liste os sites que deseja monitorar, um por linha. Por exemplo:

- <https://www.google.com>
- <https://www.exemplo.com>

3. Compile e execute o script:

```bash
go run script.go
```

### Siga as instruções no console

- Digite **1** para iniciar o monitoramento dos sites.
- Digite **2** para exibir os logs.
- Digite **0** para encerrar o programa.

## Estrutura do Código

### Constantes

- ``monitoringCount``: número de vezes que cada site será monitorado.
- ``delay``: intervalo de tempo (em segundos) entre cada rodada de monitoramento.

### Funções

- ``main()``: função principal que apresenta o menu e lida com a entrada do usuário.
- ``readInput()``: lê a entrada do usuário.
- ``monitor(sitesFile string)``: realiza o monitoramento dos sites listados no arquivo sites.txt.
- ``testSite(site string)``: testa a disponibilidade de um site.
- ``readSites(sitesFile string) []string``: lê a lista de sites a partir de um arquivo.
- ``registerLogs(logsFile string, site string, status bool, statusCode int)``: registra os resultados dos testes em um arquivo de log.
- ``showLogs(logsFile string)``: exibe o conteúdo do arquivo de log.

### Tratamento de Fuso Horário

Os logs são registrados usando o fuso horário de São Paulo **(America/Sao_Paulo)**.

## Observações

- Certifique-se de que o arquivo ``logs.txt`` seja criado no diretório onde o script está sendo executado.
- Se o arquivo ``logs.txt`` não existir quando você tentar exibir os logs, o programa solicitará que você inicie o monitoramento primeiro, pois é a partir do monitoramento que esse arquivo é criado.
