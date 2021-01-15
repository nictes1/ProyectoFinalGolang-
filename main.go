package main

import (
	"log"

	"github.com/nictes1/FinalProjectGO/bd"
	"github.com/nictes1/FinalProjectGO/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
