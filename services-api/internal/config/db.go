package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() *mongo.Database {
	// Asegurarse de que las variables de entorno estén cargadas
	LoadEnv()

	// Obtener la URI de MongoDB desde las variables de entorno
	MONGO_URI := GetEnv("MONGO_URI")
	if MONGO_URI == "" {
		log.Fatalf("MONGO_URI environment variable not set")
	}

	// Configurar las opciones de conexión
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGO_URI).SetServerAPIOptions(serverAPI)

	// Conectar al cliente de MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping a la base de datos para asegurarse de que la conexión está bien
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{
		{Key: "ping", Value: 1},
	}).Err(); err != nil {
		log.Fatalf("Failed pinging MongoDB: %v", err)
	}

	fmt.Println("Conexión exitosa a MongoDB")
	
	Client = client

	return client.Database("maintenance-service")
}
