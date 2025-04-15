package handlers

import (
	"encoding/json"
	"net/http"

	"curso-imersao-full-cycle/go-gateway-api/internal/dto"
	"curso-imersao-full-cycle/go-gateway-api/internal/service"
)

/**
* Handler de conta
* @param accountService *service.AccountService
* @return *AccountHandler
*/
type AccountHandler struct {
	accountService *service.AccountService
}

/**
* Cria um novo handler de conta
* @param accountService *service.AccountService
* @return *AccountHandler
*/
func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

/**
* Endpoint HTTP para criar uma conta
* @param w http.ResponseWriter
* @param r *http.Request
*/	
func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	// cria uma variável para receber o corpo da requisição
	var input dto.CreateAccountInput

	// decodifica o corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil { // se der erro, retorna o erro
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// chama o serviço para criar a conta
	output, err := h.accountService.CreateAccount(input)
	if err != nil { // se der erro, retorna o erro
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// define o tipo de conteúdo da resposta
	w.Header().Set("Content-Type", "application/json")

	// define o status da resposta
	w.WriteHeader(http.StatusCreated)

	// retorna a resposta
	json.NewEncoder(w).Encode(output)
}

func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	// pega a API Key do header da requisição
	apiKey := r.Header.Get("X-API-Key")

	// verifica se a API Key está presente
	if apiKey == "" {
		http.Error(w, "API Key is required", http.StatusUnauthorized)
		return
	}

	// chama o serviço para encontrar a conta
	output, err := h.accountService.FindByAPIKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// define o tipo de conteúdo da resposta
	w.Header().Set("Content-Type", "application/json")

	// define o status da resposta
	w.WriteHeader(http.StatusOK)

	// retorna a resposta
	json.NewEncoder(w).Encode(output)
}