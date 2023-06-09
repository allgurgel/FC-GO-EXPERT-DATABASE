
<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://plataforma.fullcycle.com.br/static/media/logo.6d87ce09.svg" alt="Project logo"></a>
</p>

<h1 align="center">Full Cycle Go Expert</h1>

---

<p align="center"> Desafio Curso GO Expert 
    <br> 
</p>

## 🧐 About <a name = "about"></a>

Client-Server-API
Requisitos:

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

## ⛏️ Built Using <a name = "built_using"></a>

- [Go]() - Golang

   
## 🏁 Getting Started <a name = "getting_started"></a>

Instruções para rodar o projeto.

```
go mod tidy

go run server/server.go
go run client/client.go
```
A tabela do banco de dados será criada automaticamente

O arquivo cotação.txt pode ser encontrado na raiz do projeto


---
### Prerequisites

Software que você precisa instalar.

- [Go](https://go.dev/dl/) - 1.19




