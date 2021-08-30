package middlew

import (
	"net/http"

	"github.com/roderickSa/twittor-copy/bd"
)

//ChequeoBD middleware verifica si existe la conexion con la bd antes de ejecutar el endpoint
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
