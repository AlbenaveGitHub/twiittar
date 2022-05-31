package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AlbenaveGitHub/twiittar/bd"
)

/*VerPerfil permite extraer los valores del perfil */
func VerPerfil(wr http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(wr, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(wr, "Ocurrio un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	wr.Header().Set("context-type", "application/json")
	wr.WriteHeader(http.StatusCreated)
	json.NewEncoder(wr).Encode(perfil)
}
