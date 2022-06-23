package middlew

import (
	"net/http"
	"github.com/luiscarlosjo157/redSocial/bd"
)

/*ChequeoDB es el middlew que me permite conocer el estado de la DB */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if  bd.ChequeoConnection() == 0 {
			http.Error(w, "conexion perdida con la base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}