package api

import (
	"encoding/json"
	"fmt"
	"github.com/feader02/online-shop/internal/entities"
	"log"
	"net/http"
	"strconv"
)

func (a *App) GetProductsList(w http.ResponseWriter, r *http.Request) {
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 7
	}

	pageNum, err := strconv.Atoi(r.URL.Query().Get("page_num"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	search := r.URL.Query().Get("search")
	prType := r.URL.Query().Get("type")
	priceRadius := r.URL.Query().Get("price_radius")

	products, err := a.GetProducts(pageNum, pageSize, search, prType, priceRadius)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("error: %v", err))
		return
	}

	sendResponse(w, products)
}

func (a *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Cannot parse id: %v", err))
		return
	}

	var products []entities.Product

	err = a.storage.Get(&products, "Product", "WHERE id = ?", id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	}

	for _, product := range products {
		sendResponse(w, product)
	}
}

func (a *App) Registration(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	User := map[string]interface{}{
		"login":    username,
		"password": password,
		"email":    email,
	}

	// Create the order in which data will be added
	order := []string{"login", "password", "email"}

	err := a.storage.Add(User, "Users", order)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendOk(w)
}

func (a *App) SignIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	ok, err := a.storage.ValidateLogin(username, password, "Users")
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !ok {
		sendError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	sendOk(w)
}

func sendError(w http.ResponseWriter, status int, text string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`{"status":"error","message":%s"}`, text)))
}

func sendOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func sendResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("writing response: %s", err)
	}
}
