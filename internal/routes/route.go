package routes

import (
	"github.com/Samestora/WebOverflow/internal/handlers"
	"github.com/Samestora/WebOverflow/pkg/databases"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *databases.Database) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json");
		return handlers.GetUsers(c, db);
	});

	app.Get("/user", func(c *fiber.Ctx) error {
		c.Accepts("application/json");
		return handlers.GetUserInfo(c, db);
	});

	app.Post("/createuser", func(c *fiber.Ctx) error {
		c.Accepts("application/json");
		return handlers.CreateUser(c, db);
	});

	app.Delete("/deleteuser", func(c *fiber.Ctx) error {
		c.Accepts("application/json");
		return handlers.DeleteUser(c, db);
	});
}
