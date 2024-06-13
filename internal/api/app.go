package api

import (
	"github.com/feader02/online-shop/internal/db"
	"net/http"
)

type App struct {
	storage *db.MySQLStorage
}

func NewApp(s *db.MySQLStorage) *App {
	app := App{
		storage: s,
	}

	return &app
}

func (a *App) GetHandle() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/products", a.GetProducts)

	return router
}
