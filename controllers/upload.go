package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/jorgeAM/artistaAPI/models"
	"github.com/jorgeAM/artistaAPI/services"
)

//UploadFile -> subir imagen
func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("image")
	if err != nil {
		log.Fatal("Hubo un error al leer el archivo: ", err)
		return
	}

	defer file.Close()
	mineType := handle.Header.Get("Content-type")

	switch mineType {
	case "image/jpeg":
		saveFile(w, file, handle)
	case "image/png":
		saveFile(w, file, handle)
	default:
		m := models.Message{}
		m.Code = http.StatusBadRequest
		m.Message = "Sólo puedes subir imagenes"
		services.DisplayMessage(w, m)
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Hubo un error al conseguir el archivo: ", err)
		return
	}

	err = ioutil.WriteFile("./pictures/"+handle.Filename, data, 0666)
	if err != nil {
		log.Fatal("Hubo un error al guardar el archivo: ", err)
		return
	}

	m := Message{}
	m.Code = http.StatusOK
	m.Message = "Everything goes well"

	j, err := json.Marshal(m)
	if err != nil {
		log.Fatal("Hubo un error al convertir a JSON: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
