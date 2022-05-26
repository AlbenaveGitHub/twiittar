package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de la conexion a la Base de Datos*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://mongoAlbenave:Alb3rto$45@twiittar.ozitdwf.mongodb.net/?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conectar la Base de Datos*/
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
	log.Println("Conexion Exitosa con la Base de Datos")

	return client
}

/*ChequeoConnection es el Ping a la Base de Datos*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
