package middlew

import (
	"net/http"

	"github.com/AlbenaveGitHub/twiittar/bd"
)

/*ChequeoBD es el middlew que permite conocer el estado de la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(wr, "Conexion perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(wr, r)
	}
}
