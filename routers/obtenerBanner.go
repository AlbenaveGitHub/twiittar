package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/AlbenaveGitHub/twiittar/bd"
)

/*ObtenerBanner envia el banner al HTTP */
func ObtenerBanner(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(rw, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banner" + perfil.Banner)

	if err != nil {
		http.Error(rw, "Imagen no encontrado", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(rw, openFile)
	if err != nil {
		http.Error(rw, "Error al copiar la imagen", http.StatusBadRequest)
	}

}
