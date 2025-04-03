# Full Cycle - Desafio Multithreading 🧵⚡

Este projeto é um desafio proposto no curso **Full Cycle**, com o objetivo de implementar chamadas concorrentes a serviços externos utilizando a linguagem **Go (Golang)**.

## :computer: Desafio

Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas. As duas requisições serão feitas simultaneamente para as seguintes APIs:

- [BrasilAPI](https://brasilapi.com.br)
`https://brasilapi.com.br/api/cep/v1/ + cep`

- [ViaCEP](https://viacep.com.br)
`http://viacep.com.br/ws/" + cep + "/json/`

###### Os requisitos para este desafio são

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

---

## 🚀 Funcionalidades

- Rota `/cep?cep={valor}` que recebe um CEP e retorna os dados de endereço.
- As APIs `ViaCEP` e `BrasilAPI` são chamadas em paralelo.
- A primeira resposta válida é utilizada; a outra é descartada.
- Timeout configurável via contexto (`context.WithTimeout`).
- Logs estruturados com `slog`.

---

## 📦 Estrutura do Projeto

```plaintext
├── cmd/api/          # Entry point da aplicação 
│ └── main.go         # Definição das rotas e inicialização do servidor 
├── internal/ 
│ └── handler/        # Funções de integração com as APIs externas 
│ └── address.go      # Structs e métodos de consulta (GetViaCep, GetBrasilApi) 
├── go.mod / go.sum   # Gerenciador de dependências do Go 
└── README.md         # Documentação do projeto
```

---

## 🔧 Tecnologias Utilizadas

- [Golang](https://golang.org)
- [Go Chi](https://github.com/go-chi/chi) — router leve e idiomático
- [slog](https://pkg.go.dev/log/slog) — logging estruturado
- APIs públicas: [ViaCEP](https://viacep.com.br) e [BrasilAPI](https://brasilapi.com.br)

---

## ▶️ Como executar localmente

```bash
# Clone o repositório
git clone https://github.com/JeanGrijp/Full-Cycle-Desafio-Multithreading.git
cd Full-Cycle-Desafio-Multithreading

# Execute o servidor
go run cmd/api/main.go
```
