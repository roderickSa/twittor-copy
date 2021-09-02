package bd

import (
	"context"
	"time"

	"github.com/roderickSa/twittor-copy/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificoRegistro(t models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("twittor-copy")
	col := bd.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(t.Nombre) > 0 {
		registro["nombre"] = t.Nombre
	}
	if len(t.Apellidos) > 0 {
		registro["apellidos"] = t.Apellidos
	}
	registro["fechaNacimiento"] = t.FechaNacimiento
	if len(t.Avatar) > 0 {
		registro["avatar"] = t.Avatar
	}
	if len(t.Banner) > 0 {
		registro["banner"] = t.Banner
	}
	if len(t.Biografia) > 0 {
		registro["biografia"] = t.Biografia
	}
	if len(t.Ubicacion) > 0 {
		registro["ubicacion"] = t.Ubicacion
	}
	if len(t.SitioWeb) > 0 {
		registro["sitioWeb"] = t.SitioWeb
	}

	updateString := bson.M{
		"$set": registro,
	}
	objectID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objectID}}

	_, err := col.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
