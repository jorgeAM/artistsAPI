package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jorgeAM/artistaAPI/models"
)

//DisplayMessage -> respuesta JSON
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al convertir el mensaje: %s", err)
	}

	w.WriteHeader(m.Code)
	w.Write(j)
}
