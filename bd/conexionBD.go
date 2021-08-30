package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objecto de conexion a la bd
var MongoCN = ConectarBD()
var clienteOptions = options.Client().ApplyURI("mongodb+srv://root:roder@twittor-copy.4dzze.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//conectarBD() conexion a la bd
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la BD")
	return client
}

//ChequeoConexion() es el ping a la bd
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
