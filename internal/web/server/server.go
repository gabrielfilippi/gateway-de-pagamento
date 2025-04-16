package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"curso-imersao-full-cycle/go-gateway-api/internal/service"
	"curso-imersao-full-cycle/go-gateway-api/internal/web/handlers"
	"curso-imersao-full-cycle/go-gateway-api/internal/web/middleware"
)

/**
* Definição do servidor
* @param router *chi.Mux
* @param server *http.Server
* @param accountService *service.AccountService
* @param port string
*/
type Server struct {
	router *chi.Mux
	server *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port string
}

/**
* Cria um novo servidor
* @param accountService *service.AccountService
* @param port string
* @return *Server
*/
func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	return &Server{
		router: chi.NewRouter(), // cria um novo roteador
		accountService: accountService, // define o serviço de conta
		invoiceService: invoiceService, // define o serviço de fatura
		port: port, // define a porta
	}	
}

/**
* Configura as rotas do servidor
*/
func (s *Server) ConfigureRoutes() {
	// cria um novo handler de conta
	accountHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)

	// configura as rotas de conta
	s.router.Route("/accounts", func(r chi.Router) {
		r.Post("/", accountHandler.Create)
		r.Get("/", accountHandler.Get)
	})

	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)

		r.Route("/invoice", func(r chi.Router) {
			r.Post("/", invoiceHandler.Create)
			r.Get("/", invoiceHandler.ListByAccount)
			r.Get("/{id}", invoiceHandler.GetByID)
		})
	})
}

/**
* Inicia o servidor
* @return error
*/
func (s *Server) Start() error {
	// cria um novo servidor	
	s.server = &http.Server{
		Addr:    ":" + s.port, // define a porta
		Handler: s.router, // define o roteador
	}

	// inicia o servidor
	return s.server.ListenAndServe()
}