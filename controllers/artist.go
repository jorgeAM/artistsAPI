package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgeAM/artistaAPI/config"
	"github.com/jorgeAM/artistaAPI/models"
	"github.com/jorgeAM/artistaAPI/services"
)

//GetArtists -> conseguir todos los artistas
func GetArtists(w http.ResponseWriter, r *http.Request) {
	artists := []models.Artist{}
	m := models.Message{}
	db := config.GetConnection()
	defer db.Close()
	db.Find(&artists)
	j, err := json.Marshal(artists)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al convertir artistas a JSON"
		services.DisplayMessage(w, m)
		return
	}

	if len(artists) <= 0 {
		m.Code = http.StatusOK
		m.Message = "Aún no has registrado artistas"
		services.DisplayMessage(w, m)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetArtist -> conseguir todos los artista
func GetArtist(w http.ResponseWriter, r *http.Request) {
	artist := models.Artist{}
	m := models.Message{}
	params := mux.Vars(r)
	db := config.GetConnection()
	defer db.Close()
	db.Where("id = ?", params["id"]).Find(&artist)
	/*MOSTRAR LOS ALBUMS DE UN ARTISTA*/
	db.Model(&artist).Related(&artist.Albums)
	for i := range artist.Albums {
		db.Where("id = ?", params["id"]).Find(&artist.Albums[i].Artist)
	}
	j, err := json.Marshal(artist)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al convertir artista a JSON"
		services.DisplayMessage(w, m)
		return
	}

	if artist.ID == 0 {
		m.Code = http.StatusOK
		m.Message = "No se encontro al artista"
		services.DisplayMessage(w, m)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//SaveArtist -> registrar artista
func SaveArtist(w http.ResponseWriter, r *http.Request) {
	artist := models.Artist{}
	m := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&artist)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el artista a registrar: %s", err)
		m.Code = http.StatusBadRequest
		services.DisplayMessage(w, m)
		return
	}

	db := config.GetConnection()
	defer db.Close()
	err = db.Create(&artist).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al registrar artista: %s", err)
		services.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Se registro correctamente al artista"
	services.DisplayMessage(w, m)

}

//DeleteArtist -> borrar artista
func DeleteArtist(w http.ResponseWriter, r *http.Request) {
	artist := models.Artist{}
	m := models.Message{}
	params := mux.Vars(r)
	db := config.GetConnection()
	defer db.Close()
	err := db.Where("id = ?", params["id"]).Find(&artist).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "No se encontró al artista que se desea borrar"
		services.DisplayMessage(w, m)
		return
	}

	err = db.Delete(&artist).Error
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "No se pudó borrar al artista"
		services.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusOK
	m.Message = "Se eliminó correctamente al artista"
	services.DisplayMessage(w, m)
}
