package service

import (
	"curso-imersao-full-cycle/go-gateway-api/internal/domain"
	"curso-imersao-full-cycle/go-gateway-api/internal/dto"
)

/**
* Serviço de conta
* @param repository *domain.AccountRepository
* @return *AccountService
*/
type AccountService struct {
	repository domain.AccountRepository
}

/**
* Cria um novo serviço de conta
* @param repository *domain.AccountRepository
* @return *AccountService
*/
func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

/**
* Cria uma conta
* @param input dto.CreateAccountInput
* @return *dto.AccountOutput, error
*/
func (s *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {//se der erro e não for erro de conta não encontrada, retorna o erro
		return nil, err
	}

	if existingAccount != nil { // se a conta já existir, retorna o erro	
		return nil, domain.ErrAccountAlreadyExists
	}

	err = s.repository.Save(account) // salva a conta
	if err != nil { // se der erro, retorna o erro
		return nil, err
	}

	// retorna a conta criada
	return dto.FromAccount(account), nil
}

/**
* Atualiza o saldo de uma conta
* @param apiKey string
* @param amount float64
* @return *dto.AccountOutput, error
*/
func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil { // se der erro, retorna o erro
		return nil, err
	}

	account.AddBalance(amount)

	err = s.repository.UpdateBalance(account)
	if err != nil { // se der erro, retorna o erro
		return nil, err
	}

	// retorna a conta atualizada
	return dto.FromAccount(account), nil
}

/**
* Encontra uma conta pela API Key
* @param apiKey string
* @return *dto.AccountOutput, error
*/
func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil { // se der erro, retorna o erro
		return nil, err
	}

	return dto.FromAccount(account), nil
}

/**
* Encontra uma conta pelo ID
* @param id string
* @return *dto.AccountOutput, error
*/
func (s *AccountService) FindByID(id string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByID(id)
	if err != nil { // se der erro, retorna o erro
		return nil, err
	}

	return dto.FromAccount(account), nil
}	