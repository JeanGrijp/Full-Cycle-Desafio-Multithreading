package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/JeanGrijp/Full-Cycle-Desafio-Multithreading/internal/model"
)

type AddressSource string

type Address struct {
	Via       AddressSource            `json:"via"`
	BrasilApi *model.BrasilApiResponse `json:"brasil_api"`
	ViaCep    *model.ViaCepResponse    `json:"via_cep"`
}

func (a *Address) GetViaCep(cep string) *Address {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		slog.ErrorContext(ctx, "Erro ao criar requisição para ViaCep", "error", err)
		return nil
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.ErrorContext(ctx, "Erro ao fazer requisição para ViaCep", "error", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.ErrorContext(ctx, "Resposta inválida de ViaCep", "status_code", resp.StatusCode)
		return nil
	}

	var viaCep model.ViaCepResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCep); err != nil {
		slog.ErrorContext(ctx, "Erro ao decodificar resposta do ViaCep", "error", err)
		return nil
	}

	// Atualiza o struct atual
	a.Via = "viacep"
	a.ViaCep = &viaCep
	a.BrasilApi = nil

	return a

}

func (a *Address) GetBrasilApi(cep string) *Address {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		slog.ErrorContext(ctx, "Erro ao criar requisição para BrasilApi", "error", err)
		return nil
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.ErrorContext(ctx, "Erro ao fazer requisição para BrasilApi", "error", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.ErrorContext(ctx, "Resposta inválida de BrasilApi", "status_code", resp.StatusCode)
		return nil
	}

	var brasilApi model.BrasilApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&brasilApi); err != nil {
		slog.ErrorContext(ctx, "Erro ao decodificar resposta do BrasilApi", "error", err)
		return nil
	}

	a.Via = "brasil_api"
	a.BrasilApi = &brasilApi
	a.ViaCep = nil

	return a
}
