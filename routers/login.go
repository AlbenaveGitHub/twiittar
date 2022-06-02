package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AlbenaveGitHub/twiittar/bd"
	"github.com/AlbenaveGitHub/twiittar/jwt"
	"github.com/AlbenaveGitHub/twiittar/models"
)

/*Login realiza el login*/
func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")

	var t models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(rw, "Usuario o contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(rw, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(rw, "Usuario o Contraseña invalidos", 400)
		return
	}

	jwtClave, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar generar el Token correspondiente"+err.Error(), 400)
		return
	}

	respuesta := models.RespuestaLogin{
		Token: jwtClave,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(respuesta)

	tiempoExpiracion := time.Now().Add(24 * time.Hour)
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   jwtClave,
		Expires: tiempoExpiracion,
	})

}
