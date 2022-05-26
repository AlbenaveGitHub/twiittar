package main

import (
	"log"

	"github.com/AlbenaveGitHub/twiittar/bd"
	"github.com/AlbenaveGitHub/twiittar/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la Base de Datos")
		return
	}
	handlers.Manejadores()

}
