package repository

import (
	"database/sql"
	"time"

	"curso-imersao-full-cycle/go-gateway-api/internal/domain"
)

/*
* Repositório de contas
* @param db *sql.DB
* @return *AccountRepository
*/
type AccountRepository struct {
	db *sql.DB
}

/**
* Cria um novo repositório de contas
* @param db *sql.DB
* @return *AccountRepository
*/
func NewAccountRepository(db *sql.DB) *AccountRepository{
	return &AccountRepository{
		db: db,
	}
}

/**
* Salva uma conta
* @param account *domain.Account
* @return error
*/
func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare("INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	_, err = stmt.Exec(account.ID, account.Name, account.Email, account.APIKey, account.Balance, account.CreatedAt, account.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

/**
* Encontra uma conta pela chave da API
* @param apiKey string
* @return *domain.Account, error
*/
func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	row := r.db.QueryRow("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = $1", apiKey)
	
	err := row.Scan(&account.ID, &account.Name, &account.Email, &account.APIKey, &account.Balance, &createdAt, &updatedAt)
	
	if err == sql.ErrNoRows { // se não houver resultado, retorna um erro
		return nil, domain.ErrAccountNotFound
	}
	
	if err != nil { // se houver erro, retorna o erro
		return nil, err
	}
	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}

/**
* Encontra uma conta pelo ID
* @param accountID string
* @return *domain.Account, error
*/
func (r *AccountRepository) FindByID(accountID string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	row := r.db.QueryRow("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE id = $1", accountID)
	
	err := row.Scan(&account.ID, &account.Name, &account.Email, &account.APIKey, &account.Balance, &account.CreatedAt, &account.UpdatedAt)
	
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	
	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}

/**
* Atualiza o saldo da conta
* @param account *domain.Account
* @return error
*/
func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()

	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1 FOR UPDATE", account.ID).Scan(&currentBalance)
	
	if err == sql.ErrNoRows {//se não houver resultado, retorna um erro
		return domain.ErrAccountNotFound
	}

	if err != nil {//se houver erro, retorna o erro
		return err
	}

	if currentBalance < account.Balance {//se o saldo atual for menor que o saldo a ser atualizado, retorna um erro
		return domain.ErrInsufficientBalance
	}

	_, err = tx.Exec("UPDATE accounts SET balance = $1, updated_at = $2 WHERE id = $3", account.Balance, time.Now(), account.ID)
	if err != nil {//se houver erro, retorna o erro
		return err
	}

	//se não houver erro, commita (salva) a transação
	return tx.Commit()
}