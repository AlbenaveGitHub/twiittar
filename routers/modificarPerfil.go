package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AlbenaveGitHub/twiittar/bd"
	"github.com/AlbenaveGitHub/twiittar/models"
)

func ModificarPerfil(rw http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(rw, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	if status, err := bd.ModificoRegistro(t, IDUsuario); err != nil {
		http.Error(rw, "Ocurrio un error al intentar modificar el registro.  Reintente nuevamente "+err.Error(), 400)
		return
	} else if status == false {
		http.Error(rw, "no se ha logrado modificar el registro del usuario ", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)

}
