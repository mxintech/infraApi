package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TheGolurk/infraApi/models"
)

func DisplayMessage(w http.ResponseWriter, m models.Message) {
	JSON, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al convertir el mensaje: %s", err)
	}
	w.WriteHeader(m.Code)
	w.Write(JSON)
}
