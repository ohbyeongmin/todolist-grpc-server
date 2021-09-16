package db

import (
	"context"
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

type MongoClient struct {
	client *mongo.Client
}

var TodoConn *MongoClient

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

func (c *MongoClient) InsertTodo(nick string, td string) *mongo.InsertOneResult {
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
	return result
}

func (c *MongoClient) UpdateTodo(targetId string, td string) *mongo.UpdateResult {
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
	return result
}

func init() {
	TodoConn = &MongoClient{}
	TodoConn.client = connectDB()
}
