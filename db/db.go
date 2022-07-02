package db

import (
	"context"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB_URI = os.Getenv("DB_URI")

type MongoDB struct {
	Ctx    context.Context
	Client *mongo.Client
	Err    error
	Cancel context.CancelFunc
}

func (db *MongoDB) Init() (*mongo.Client, error) {
	if DB_URI == "" {
		log.Println("No DB_URI provided, using localhost instead")
		DB_URI = "mongodb://localhost:27017"
	}
	log.Println("⏳ Connecting to MongoDB 🥭 ...")
	db.Ctx, db.Cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer db.Cancel()
	db.Client, db.Err = mongo.Connect(db.Ctx, options.Client().ApplyURI(DB_URI))
	if err := db.Client.Ping(db.Ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("✅ Connected to MongoDB 🥭")
	return db.Client, db.Err
}

func (db *MongoDB) Disconnect() {
	log.Println("🔌 Disconnecting from MongoDB ...")
	if db.Err = db.Client.Disconnect(db.Ctx); db.Err != nil {
		panic(db.Err)
	}
	log.Println("🔌 Disconnected from MongoDB")
}

func InitMGM() error {
	if DB_URI == "" {
		log.Println("No DB_URI provided, using localhost instead")
		DB_URI = "mongodb://localhost:27017"
	}
	log.Println("⏳ Initializing MGM ... 🗺")
	err := mgm.SetDefaultConfig(nil, "prosperity-game", options.Client().ApplyURI(DB_URI))
	if err != nil {
		log.Println("❌ Failed to connect to DB")
		panic(err)
	}
	log.Println("🔰 Looks like MGM is initialized")
	return err
}
