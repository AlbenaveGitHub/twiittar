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
func Login(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Add("content-type", "application/json")

	var t models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(wr, "Usuario o contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(wr, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(wr, "Usuario o Contraseña invalidos", 400)
		return
	}

	jwtClave, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(wr, "Ocurrio un error al intentar generar el Token correspondiente"+err.Error(), 400)
		return
	}

	respuesta := models.RespuestaLogin{
		Token: jwtClave,
	}

	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusCreated)
	json.NewEncoder(wr).Encode(respuesta)

	tiempoExpiracion := time.Now().Add(24 * time.Hour)
	http.SetCookie(wr, &http.Cookie{
		Name:    "token",
		Value:   jwtClave,
		Expires: tiempoExpiracion,
	})

}
