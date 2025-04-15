package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(apiKey string) (*Account, error)
	FindByID(accountID string) (*Account, error)
	UpdateBalance(account *Account) error
}
