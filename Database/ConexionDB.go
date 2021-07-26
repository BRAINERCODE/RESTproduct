package Database

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Se utiliza para crear un objeto singleton del cliente MongoDB.
Inicializado y expuesto a través de GetMongoClient (). */
var clientInstance *mongo.Client

// Se usa durante la creación del objeto de cliente singleton en GetMongoClient ().
var clientInstanceError error

// Se usa para ejecutar el procedimiento de creación de clientes solo una vez.
var mongoOnce sync.Once

// He usado las siguientes constantes solo para mantener las configuraciones requeridas de la base de datos.
const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "db_issue_manager"
	ISSUES           = "col_issues"
)

// GetMongoClient - Devuelve la conexión mongodb para trabajar con
func GetMongoClient() (*mongo.Client, error) {

	// Realice la operación de creación de conexión solo una vez.
	mongoOnce.Do(func() {
		// Establecer las opciones del cliente
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)

		// Conectarse a MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Verifica la conexión

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError

}
