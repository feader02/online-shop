package api

import (
	"encoding/json"
	"fmt"
	"github.com/feader02/online-shop/internal/entities"
	"log"
	"net/http"
)

func (a *App) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []entities.Product

	err := a.storage.GetProducts(&products, "Product", "")
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	}

	sendResponse(w, products)

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
