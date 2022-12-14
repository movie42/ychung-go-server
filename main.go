package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/movie42/ychung-go-server/pkg/repository"
	"github.com/movie42/ychung-go-server/pkg/service"
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

	noticeService := databaseCollections("notices", db)
	worshipService := databaseCollections("worships", db)
	blogService := databaseCollections("blogs", db)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome first go application"))
	})

	// TODO: 마음에 안든다... 일단 구현하고 구조화를 해보자.
	api := app.Group("/api")

	router.NoticeRouter(api, noticeService)
	router.WorshipRouter(api, worshipService)
	router.BlogRouter(api, blogService)
	// router.EducationRouter(api, worshipService)

	defer cancel()
	log.Fatal(app.Listen(":4000"))
}

func databaseCollections(collectionType string, db *mongo.Database) service.Service {
	dbCollection := db.Collection(collectionType)
	dbReop := repository.NewRepo(dbCollection)
	dbService := service.NewService(dbReop)

	return dbService
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://localhost:27017").SetServerSelectionTimeout(
		5*time.Second))

	if err != nil {
		cancel()
		return nil, nil, err
	}

	db := client.Database("yangchung")
	return db, cancel, nil
}
