package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/artistaAPI/controllers"
)

//InitRoutes -> iniciar rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/artistas", controllers.GetArtists).Methods("GET")
	router.HandleFunc("/artistas/{id}", controllers.GetArtist).Methods("GET")
	router.HandleFunc("/artistas", controllers.SaveArtist).Methods("POST")
	router.HandleFunc("/artistas/{id}", controllers.DeleteArtist).Methods("DELETE")
	router.HandleFunc("/albums", controllers.GetAlbums).Methods("GET")
	router.HandleFunc("/albums/{id}", controllers.GetAlbum).Methods("GET")
	router.HandleFunc("/albums", controllers.SaveAlbum).Methods("POST")
	router.HandleFunc("/albums/{id}", controllers.DeleteAlbum).Methods("DELETE")
	/*RUTA PARA SUBIR IMAGEN*/
	router.HandleFunc("/{tipo}/{id}/upload", controllers.UploadFile).Methods("POST")
	return router
}
