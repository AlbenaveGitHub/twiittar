package bd

import (
	"github.com/AlbenaveGitHub/twiittar/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de acceso a la BD*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	if err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes); err != nil {
		return usu, false
	}
	return usu, true

}
