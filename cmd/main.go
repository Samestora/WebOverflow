package main

import (
	"log"

	"github.com/Samestora/WebOverflow/internal/routes"
	"github.com/Samestora/WebOverflow/pkg/databases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := databases.Connect();
	if err != nil {
		log.Fatal("Failed to connect!");
	}
	app := fiber.New();
	
	routes.SetupRoutes(app,db);

	app.Listen(":3000");
}
