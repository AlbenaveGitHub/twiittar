package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AlbenaveGitHub/twiittar/bd"
)

func LeoTweetsSeguidores(rw http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {

		http.Error(rw, "Deve enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(rw, "Debe enviar el parametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respusta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(rw, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(respusta)
}
