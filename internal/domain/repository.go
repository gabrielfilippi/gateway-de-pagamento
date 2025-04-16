package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(apiKey string) (*Account, error)
	FindByID(accountID string) (*Account, error)
	UpdateBalance(account *Account) error
}

type InvoiceRepository interface {
	Save(invoice *Invoice) error
	FindByID(invoiceID string) (*Invoice, error)
	FindByAccountID(accountID string) ([]*Invoice, error)
	UpdateStatus(invoice *Invoice) error
}