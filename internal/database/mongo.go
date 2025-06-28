package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

// Connect establishes a connection to the MongoDB database
func Connect(uri string) {
    var err error
    
    // Set client options with longer timeout for Atlas
    clientOptions := options.Client().ApplyURI(uri).SetMaxPoolSize(10)
    
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    Client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("Failed to connect to MongoDB Atlas:", err)
    }

    // Ping with longer timeout for Atlas
    pingCtx, pingCancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer pingCancel()
    
    err = Client.Ping(pingCtx, nil)
    if err != nil {
        log.Fatal("Failed to ping MongoDB Atlas:", err)
    }

    Database = Client.Database("ctf_toolkit")
    log.Println("Successfully connected to MongoDB Atlas!")
}

// GetDatabase returns the MongoDB database instance
func GetDatabase() *mongo.Database {
    return Database
}

// Disconnect closes the connection to the MongoDB database
func Disconnect() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := Client.Disconnect(ctx); err != nil {
        log.Fatal("Failed to disconnect from MongoDB:", err)
    }

    log.Println("Disconnected from MongoDB!")
}