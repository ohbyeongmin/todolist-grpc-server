package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID       primitive.ObjectID `bson:"_id"`
	NickName string             `bson:"nickname"`
	Td       string             `bson:"todo"`
	Status   bool               `bson:"status"`
}

const MONGO_URI = "mongodb://127.0.0.1:27017"
const DATABSE_NAME = "todo_database_test"
const COLLECTION_NAME = "todo_list"

type DB interface {
	InsertTodo(string, string) string
	UpdateTodo(string, string) string
}

type MongoClient struct {
	client *mongo.Client
}

func connectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (c *MongoClient) getCollection() *mongo.Collection {
	return c.client.Database(DATABSE_NAME).Collection(COLLECTION_NAME)
}

func (c *MongoClient) InsertTodo(nick string, td string) string {
	todoCollection := c.getCollection()
	newTodo := Todo{
		ID:       primitive.NewObjectID(),
		NickName: nick,
		Td:       td,
		Status:   false,
	}
	result, err := todoCollection.InsertOne(context.Background(), newTodo)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", result.InsertedID)
}

func (c *MongoClient) UpdateTodo(targetId string, td string) string {
	todoCollection := c.getCollection()
	id, _ := primitive.ObjectIDFromHex(targetId)
	result, err := todoCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"todo": td}},
	)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", result.UpsertedID)
}

type Service struct {
	DBservice DB
}

var SVC Service

func init() {
	TodoConn := MongoClient{}
	TodoConn.client = connectDB()
	SVC = Service{
		DBservice: &TodoConn,
	}
}
