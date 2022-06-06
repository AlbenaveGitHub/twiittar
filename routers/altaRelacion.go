package routers

import (
	"net/http"

	"github.com/AlbenaveGitHub/twiittar/bd"
	"github.com/AlbenaveGitHub/twiittar/models"
)

/*AltaRelacion realiza el registro de la relacion entre usuarios */
func AltaRelacion(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(rw, "No se ha logrado insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
