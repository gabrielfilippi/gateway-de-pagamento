package domain

import "errors"

var (
	//Retorna erro se a conta não for encontrada
	ErrAccountNotFound = errors.New("account not found")

	//Retorna erro se a conta já existir
	ErrAccountAlreadyExists = errors.New("account already exists")
	
	//Retorna erro se a chave da API já existir
	ErrDuplicatedAPIKey = errors.New("duplicated api key")
	
	//Retorna erro se a fatura não for encontrada
	ErrInvoiceNotFound = errors.New("invoice not found")
	
	//Retorna erro se o acesso não for autorizado
	ErrUnauthorizedAccess = errors.New("unauthorized access")

	//Retorna erro se o saldo for insuficiente
	ErrInsufficientBalance = errors.New("insufficient balance")

	//Retorna erro se o cartão não for válido
	ErrInvalidCreditCard = errors.New("invalid credit card")

	//Retorna erro se o cartão não for válido
	ErrInvalidAmount = errors.New("invalid amount")

	//Retorna erro se o status não for válido
	ErrInvalidStatus = errors.New("invalid status")
)