package main

import (
	"log"

	"github.com/roderickSa/twittor-copy/bd"
	"github.com/roderickSa/twittor-copy/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexion a la bd")
		return
	}
	handlers.Manejadores()
}
