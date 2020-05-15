package middlew

import (
	"net/http"

	"github.com/geronm12/twittor/bd"
)

//ChequeoBd middlewar que chequea la base de datos para verificar que esté funcional
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if bd.ChequeoConecction() == 0 {
			http.Error(w, "Conexión pérdida con la Base de Datos", 500)
			return
		}

		next.ServeHTTP(w, r)
	}

}
