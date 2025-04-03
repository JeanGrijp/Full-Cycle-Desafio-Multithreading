package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/JeanGrijp/Full-Cycle-Desafio-Multithreading/internal/handler"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	slog.InfoContext(ctx, "Iniciando servidor na porta 8080")

	slog.InfoContext(ctx, "Rotas configuradas")
	r.Get("/cep", BuscarEnderecoConcorrente)

	slog.InfoContext(ctx, "Servidor iniciado na porta 12000")
	http.ListenAndServe(":12000", r)
}

func BuscarEnderecoConcorrente(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()

	slog.InfoContext(ctx, "Iniciando busca de endereço")

	slog.InfoContext(ctx, "Buscando CEP")
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		slog.ErrorContext(ctx, "CEP não informado")
		http.Error(w, "CEP não informado", http.StatusBadRequest)
		return
	}

	slog.InfoContext(ctx, "Buscando endereço", "cep", cep)

	resultado := make(chan *handler.Address)

	slog.InfoContext(ctx, "Buscando endereço via APIs")
	go buscaViaCep(cep, resultado)
	go buscaBrasilApi(cep, resultado)

	select {
	case address := <-resultado:
		slog.InfoContext(ctx, "Endereço encontrado", "cep", cep)
		slog.InfoContext(ctx, "Endereço encontrado", "endereco via_cep", address.ViaCep)
		slog.InfoContext(ctx, "Endereço encontrado", "endereco brasil_api", address.BrasilApi)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(address)
		return
	case <-ctx.Done():
		slog.ErrorContext(ctx, "Tempo limite excedido", "cep", cep)
		http.Error(w, "Tempo limite excedido", http.StatusGatewayTimeout)
		return
	}

}

func buscaViaCep(cep string, resultado chan<- *handler.Address) {
	address := &handler.Address{}
	resultado <- address.GetViaCep(cep)
}

func buscaBrasilApi(cep string, resultado chan<- *handler.Address) {
	address := &handler.Address{}
	resultado <- address.GetBrasilApi(cep)
}
