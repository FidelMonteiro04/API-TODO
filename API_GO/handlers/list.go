package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FidelMonteiro04/api_go/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetALL()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
