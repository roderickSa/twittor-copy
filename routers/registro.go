package routers

import (
	"encoding/json"
	"net/http"

	"github.com/roderickSa/twittor-copy/bd"
	"github.com/roderickSa/twittor-copy/models"
)

//Registro funcion para registrar usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El Email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El Password debe tener como minimo 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error durante el registro de usuario "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se pudo registrar el usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
