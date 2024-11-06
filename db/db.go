package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

// ConnectDB establece la conexión con MongoDB y devuelve el cliente
func ConnectDB() (*mongo.Client, error) {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://go-grpc-mongo-mongodb-1:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}
	collection = client.Database("argentina_office").Collection("personas")

	// Verifica la conexión
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, err
	}

	log.Println("Conexión a MongoDB establecida correctamente")
	return client, nil
}

// InsertDummyData inserta datos de ejemplo en la colección "personas"
func InsertDummyData() {
	personas := []interface{}{
		bson.M{"nombre": "Juan Pérez", "edad": 30, "antiguedad": 24},
		bson.M{"nombre": "María López", "edad": 28, "antiguedad": 18},
		bson.M{"nombre": "Pedro Gómez", "edad": 35, "antiguedad": 30},
		bson.M{"nombre": "Ana Torres", "edad": 22, "antiguedad": 6},
		bson.M{"nombre": "Luis Fernández", "edad": 40, "antiguedad": 60},
		bson.M{"nombre": "Sofía Ruiz", "edad": 25, "antiguedad": 12},
		bson.M{"nombre": "Diego Díaz", "edad": 33, "antiguedad": 36},
		bson.M{"nombre": "Laura Castro", "edad": 29, "antiguedad": 24},
		bson.M{"nombre": "Javier Morales", "edad": 27, "antiguedad": 15},
		bson.M{"nombre": "Paula Salazar", "edad": 31, "antiguedad": 48},
	}

	_, err := collection.InsertMany(context.TODO(), personas)
	if err != nil {
		log.Fatalf("Failed to insert dummy data: %v", err)
	}
}
