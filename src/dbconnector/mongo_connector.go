package dbconnector

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through getMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

const (
	// Collection names
	COLLECTION_CLOTHES_CONTAINER = "containers.clothes"
	COLLECTION_OILS_CONTAINER    = "containers.oils"

	// Error messages
	COLLECTION_NAME_EMPTY_ERROR = "collection name cannot be empty"
)

func InitMongoClient() {
	log.Println("Initializing MongoDB client")
	_, clientInstanceError = getMongoClient()
	if clientInstanceError != nil {
		panic(clientInstanceError)
	}
}

//GetMongoClient - Return mongodb connection to work with
func getMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(os.Getenv("DB_CONNECTION_URI"))
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
			clientInstanceError = err
		}
		clientInstance = client
		log.Println("Connected to MongoDB")
	})
	return clientInstance, clientInstanceError
}

func DbPing() bool {
	client, _ := getMongoClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Ping(ctx, nil)
	if err != nil {
		log.Println("Ping can't reach the database.", err)
		return false
	}
	return true
}

func getCollection(collectionName string) (*mongo.Collection, error) {
	client, err := getMongoClient()
	if err != nil {
		log.Println("Connection to database is lost.", err)
		return nil, err
	}
	return client.Database(os.Getenv("DB_NAME")).Collection(collectionName), nil
}

func ListDocuments[T any](collectionName string, filter *bson.D) (*[]*T, error) {
	if collectionName == "" {
		return nil, errors.New(COLLECTION_NAME_EMPTY_ERROR)
	}
	if filter == nil {
		filter = &bson.D{}
	}
	coll, err := getCollection(collectionName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	result := &[]*T{}
	if err = cursor.All(context.TODO(), result); err != nil {
		return nil, err
	}
	return result, nil
}

func InsertDocument[T any](collectionName string, document *T) (*mongo.InsertOneResult, error) {
	if collectionName == "" {
		return nil, errors.New(COLLECTION_NAME_EMPTY_ERROR)
	}
	coll, err := getCollection(collectionName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	inserted, err := coll.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func InsertDocuments[T any](collectionName string, documents *[]*T) (*mongo.InsertManyResult, error) {
	if collectionName == "" {
		return nil, errors.New(COLLECTION_NAME_EMPTY_ERROR)
	}
	var iDocuments []interface{}
	for _, t := range *documents {
		iDocuments = append(iDocuments, t)
	}
	coll, err := getCollection(collectionName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	inserted, err := coll.InsertMany(ctx, iDocuments)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func DeleteDocuments(collectionName string, filter *bson.D) (*mongo.DeleteResult, error) {
	if collectionName == "" {
		return nil, errors.New(COLLECTION_NAME_EMPTY_ERROR)
	}
	if filter == nil {
		filter = &bson.D{}
	}
	coll, err := getCollection(collectionName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return coll.DeleteMany(ctx, filter)
}
