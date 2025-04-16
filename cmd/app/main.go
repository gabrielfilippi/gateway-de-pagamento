package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"curso-imersao-full-cycle/go-gateway-api/internal/repository"
	"curso-imersao-full-cycle/go-gateway-api/internal/service"
	"curso-imersao-full-cycle/go-gateway-api/internal/web/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database")
	}
	defer db.Close()

	// Inicializa o repositório e o serviço de conta
	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	// Inicializa o repositório e o serviço de fatura
	invoiceRepository := repository.NewInvoiceRepository(db)
	invoiceService := service.NewInvoiceService(invoiceRepository, *accountService)

	// Inicializa o servidor
	server := server.NewServer(accountService, invoiceService, os.Getenv("HTTP_PORT"))
	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server")
	}
}