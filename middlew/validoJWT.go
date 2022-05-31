package middlew

import (
	"net/http"

	"github.com/AlbenaveGitHub/twiittar/routers"
)

/*ValidoJWT permite validar el JWT que nos viene en la peticion*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(wr, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(wr, r)
	}
}
