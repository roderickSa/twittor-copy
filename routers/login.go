package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/roderickSa/twittor-copy/bd"
	"github.com/roderickSa/twittor-copy/jwt"
	"github.com/roderickSa/twittor-copy/models"
)

//Registro funcion para registrar usuario
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario o contraseña invalida "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El Email de usuario es requerido", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario o contraseña invalida", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Error al generar token "+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "appication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
