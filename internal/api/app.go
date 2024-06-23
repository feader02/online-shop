package api

import (
	"github.com/feader02/online-shop/internal/db"
	"github.com/feader02/online-shop/internal/middlewares"
	"github.com/rs/cors"
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
	router.HandleFunc("GET /api/products", a.GetProductsList)
	router.HandleFunc("GET /api/products/{id}", a.GetProduct)
	router.HandleFunc("POST /api/registration/user", a.Registration)
	router.HandleFunc("POST /api/sign-in/user", middlewares.CheckCookie(a.SignIn))
	router.HandleFunc("POST /api/sign-out", a.SignOut)

	c := cors.Default().Handler(router)

	return c
}
