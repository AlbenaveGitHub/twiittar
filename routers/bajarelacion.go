package routers

import (
	"net/http"

	"github.com/AlbenaveGitHub/twiittar/bd"
	"github.com/AlbenaveGitHub/twiittar/models"
)

/*BajaRelacion realiza el borrado de la relacion entre usuarios */
func BajaRelacion(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar insertar borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(rw, "No se ha logrado borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
