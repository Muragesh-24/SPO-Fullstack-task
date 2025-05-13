package config

import (
    "context"
    "log"
    "time"
  "os"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
    uri := os.Getenv("MongoString")
client, err := mongo.NewClient(options.Client().ApplyURI(uri))

    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    DB = client
}

func GetCollection(collectionName string) *mongo.Collection {
    return DB.Database("mydb").Collection(collectionName)
}
