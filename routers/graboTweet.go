package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AlbenaveGitHub/twiittar/bd"
	"github.com/AlbenaveGitHub/twiittar/models"
)

/*GraboTweet permite grabar el tweet en la base de datos*/
func GraboTweet(rw http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(rw, "No se ha logrado insertar el Tweet", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)

}
