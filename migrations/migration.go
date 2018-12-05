package migration

import (
	"github.com/jorgeAM/artistaAPI/config"
	"github.com/jorgeAM/artistaAPI/models"
)

//Migrate -> hacer migraciones de las tablas por consola
func Migrate() {
	db := config.GetConnection()
	defer db.Close()

	/*TABLAS*/
	db.CreateTable(&models.Artist{})
	db.CreateTable(&models.Album{})

	/*RELACIONES*/
	db.Model(&models.Album{}).AddForeignKey("artist_id", "artists(id)", "RESTRICT", "RESTRICT")
}
