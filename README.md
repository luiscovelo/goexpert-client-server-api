# Desafio Go Expert - Client-Server-API

## Proposta

Olá dev, tudo bem?
 
Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.
 
Você precisará nos entregar dois sistemas em Go:
- client.go
- server.go
 
Os requisitos para cumprir este desafio são:
 
O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
 
O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
 
Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
 
O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.
 
O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

## Solução para possível problema com o sqlite3

Caso ao tentar executar o servidor e ocorrer o erro abaixo:

```
# github.com/mattn/go-sqlite3
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in $PATH
```

Caso esteja no linux, execute este comando para instalar as ferramentas úteis de desenvolvimento:

```
apt-get install build-essential -y
```

Caso esteja em outro sistema operacional, procure como instalar o compilador GCC.

## Como rodar

O projeto contem dois diretórios, um para a aplicação do servidor e outro para a aplicação do cliente.

Esteja na raiz do diretório `client-server-api`, abra dois terminais ou abra somente um e faça o split para dividir em duas partes.

### 1. Servidor

Execute o servidor com o seguinte comando abaixo:

```
cd server/ && go run cmd/api/server.go
```

Caso resulte em succeso, nosso servidor rodando em http://localhost:8080.

### 2. Cliente

Execute o cliente com o seguinte comando abaixo:

```
cd client/ && go run client.go
```

Caso resulte em sucesso, será printado no console a seguinte informação:

```
quotation save successful
```

Para visualizar o resultado gravado em arquivo, execute o comando abaixo:

```
cat client/cotacao.txt
```