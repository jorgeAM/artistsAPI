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

//GetAlbums -> conseguir todos los albums
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums := []models.Album{}
	//artist := models.Artist{}
	m := models.Message{}
	db := config.GetConnection()
	defer db.Close()
	db.Find(&albums)
	for i := range albums {
		db.Model(&albums[i]).Related(&albums[i].Artist)
	}
	j, err := json.Marshal(albums)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al convertir albums a JSON"
		services.DisplayMessage(w, m)
		return
	}

	if len(albums) <= 0 {
		m.Code = http.StatusOK
		m.Message = "Aún no has registrado albums"
		services.DisplayMessage(w, m)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetAlbum -> conseguir un album especifico
func GetAlbum(w http.ResponseWriter, r *http.Request) {
	album := models.Album{}
	m := models.Message{}
	params := mux.Vars(r)
	db := config.GetConnection()
	defer db.Close()
	db.Where("id = ?", params["id"]).Find(&album)
	db.Model(&album).Related(&album.Artist)
	j, err := json.Marshal(album)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al convertir album a JSON"
		services.DisplayMessage(w, m)
		return
	}

	if album.ID == 0 {
		m.Code = http.StatusOK
		m.Message = "No se encontró el album"
		services.DisplayMessage(w, m)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//SaveAlbum -> registrar artista
func SaveAlbum(w http.ResponseWriter, r *http.Request) {
	album := models.Album{}
	m := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el album a registrar: %s", err)
		m.Code = http.StatusBadRequest
		services.DisplayMessage(w, m)
		return
	}

	db := config.GetConnection()
	defer db.Close()
	err = db.Create(&album).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al registrar album: %s", err)
		services.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Se registro correctamente el album"
	services.DisplayMessage(w, m)

}

//DeleteAlbum -> borrar album
func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	album := models.Album{}
	m := models.Message{}
	params := mux.Vars(r)
	db := config.GetConnection()
	defer db.Close()
	err := db.Where("id = ?", params["id"]).Find(&album).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "No se encontró el album que se desea borrar"
		services.DisplayMessage(w, m)
		return
	}

	err = db.Delete(&album).Error
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "No se pudó borrar el album"
		services.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusOK
	m.Message = "Se eliminó correctamente el album"
	services.DisplayMessage(w, m)
}
