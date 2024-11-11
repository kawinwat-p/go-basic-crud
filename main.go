package main

import (
	"log"
	"os"
	"practice/configuration"
	"practice/domain/datasources"
	"practice/domain/repositories"
	"practice/src/gateways"
	"practice/src/middlewares"
	"practice/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	// // // remove this before deploy ###################
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// /// ############################################

	app := fiber.New(configuration.NewFiberConfiguration())
	middlewares.Logger(app)
	app.Use(recover.New())
	app.Use(cors.New())

	mongodb := datasources.NewMongoDB(10)

	artistMongo := repositories.NewArtistsRepository(mongodb)

	sv0 := services.NewArtistService(artistMongo)

	gateways.NewHTTPGateway(app,sv0)

	PORT := os.Getenv("DB_PORT_LOGIN")

	if PORT == ""{
		PORT = "5000"
	}

	app.Listen(":"+PORT)
}