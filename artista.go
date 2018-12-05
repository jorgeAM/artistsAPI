package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jorgeAM/artistaAPI/migrations"
	"github.com/jorgeAM/artistaAPI/routes"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Ejecutar migraciones")
	flag.Parse()

	if migrate == "yes" {
		log.Println("Ejecutando migraciones...")
		migration.Migrate()
		log.Println("Migraciones ejecutadas")
	}

	s := &http.Server{
		Addr:    ":4000",
		Handler: routes.InitRoutes(),
	}

	log.Println("Servidor corriendo en http://localhost:4000 🦄")
	log.Println(s.ListenAndServe())
}
