package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AlbenaveGitHub/twiittar/bd"
	"github.com/AlbenaveGitHub/twiittar/models"
)

/*Registro es la funcion para crear en la BD el registro de un usuario*/
func Registro(rw http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(rw, "El email del usuario es requerido ", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(rw, "El password debe ser de al menos 6 caracteres ", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(rw, "Ya existe un usuario con este email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar registrar el usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(rw, "No se logro registral el usuario"+err.Error(), 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
