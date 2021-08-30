package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/roderickSa/twittor-copy/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("twittor-copy")
	col := bd.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = "" //para no devolver la contraseña
	if err != nil {
		fmt.Print("Registro no encontrado " + err.Error())
		return perfil, err
	}
	return perfil, nil
}
