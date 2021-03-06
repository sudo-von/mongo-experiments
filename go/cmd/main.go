package main

import (
	"log"
	"os"

	"github.com/mongo-experiments/go/pkg/http"
	"github.com/mongo-experiments/go/pkg/mongo"
	"github.com/mongo-experiments/go/pkg/users"
)

func main() {

	// Gets environment variables.
	dbURL := os.Getenv("MONGOURL")
	dbName := os.Getenv("ME_CONFIG_MONGODB_AUTH_DATABASE")
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	// Handles MongoDB connection.
	log.Println("[main] error: Making a new connection to MongoDB...")
	db, err := mongo.NewStorage(dbURL, dbName, user, pass)
	if err != nil {
		log.Println("[main] error:", err)
		return
	}
	log.Println("[main]: MongoDB connection established")

	// Services.
	userService := users.NewService(db)
	// Controllers.
	controllers := http.Controller{UserService: *userService}
	// Http server.
	http.ListenAndServe(controllers)
}
