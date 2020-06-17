package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN objeto de conexión a la base de datos
var MongoCN = ConectarBD()

//TODO: PONER PASSWORD
var clientOptions = options.Client().ApplyURI("mongodb+srv://Ironman:gAarfnVUpJB3ouMA@cluster0-6u08f.mongodb.net/test?retryWrites=true&w=majority")

//-----------------------//

//ConectarBD es la función que me permite conectar la base de datos
func ConectarBD() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión Exitosa con la BD")
	return client

}

//ChequeoConnection sirve para chequear la conexión a la base de datos
func ChequeoConecction() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
