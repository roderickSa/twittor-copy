package routers

import (
	"encoding/json"
	"net/http"

	"github.com/roderickSa/twittor-copy/bd"
	"github.com/roderickSa/twittor-copy/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al modificar el registro "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se logro modificar el registro "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
