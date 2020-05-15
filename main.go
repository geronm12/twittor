package main

import (
	"log"

	"github.com/geronm12/twittor/bd"
	"github.com/geronm12/twittor/handlers"
)

func main() {

	if bd.ChequeoConecction() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}

	handlers.Manejadores()

}
