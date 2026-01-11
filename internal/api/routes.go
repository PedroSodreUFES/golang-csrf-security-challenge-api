package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(
		middleware.RequestID,
		middleware.Recoverer,
		middleware.Logger,
		api.Sessions.LoadAndSave,
	)

	api.Router.Route("/conta", func(r chi.Router) {
		// rotas não autenticadas
		r.Post("/login", api.handleLogin)
		r.Post("/", api.handleCreateAccount)

		// Rotas autenticadas
		r.Group(func(r chi.Router) {
			r.Use(api.AuthMiddleware)
			r.Get("/{id}/saldo", api.handleGetBalance)
			r.Post("/{id}/deposito", api.handleDeposit)
			r.Post("/{id}/saque", api.handleWithdraw)
			r.Post("/transferencia", api.handleTransfer)
			r.Delete("/{id}", api.handleCloseAccount)
			r.Post("/logout", api.handleLogout)
		})
	})
}
