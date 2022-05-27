package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriptar el password*/
func EncriptarPassword(pass string) (string, error) {
	/*El consto es la cantidad de veces que se realiza el encriptado del password, el estandar es 8, se recomienda
	no bajar de 6 y para usuarios Admin se recomienda 10*/
	costo := 8 /*2 a la 8 es 256 veces que se encripta un password*/
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
