package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/roderickSa/twittor-copy/bd"
	"github.com/roderickSa/twittor-copy/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var t models.Tweet
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: t.Mensaje,
		Fecha:   time.Now(),
	}

	var status bool
	_, status, err = bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Error al insertar tweet "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Tweet no puso ser registrado ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
