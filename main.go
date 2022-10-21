package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/movie42/ychung-go-server/pkg/notice"
	"github.com/movie42/ychung-go-server/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database Connection Success!")
	noticeCollection := db.Collection("notices")
	noticeReop := notice.NewRepo(noticeCollection)
	noticeService := notice.NewService(noticeReop)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome first go application"))
	})

	api := app.Group("/api")
	router.ApiRouter(api, noticeService)
	defer cancel()
	log.Fatal(app.Listen(":3000"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetServerSelectionTimeout(5*time.Second))

	if err != nil {
		cancel()
		return nil, nil, err
	}

	db := client.Database("yanchung")
	return db, cancel, nil
}
