package handlers

import (
	"log"

	"github.com/Samestora/WebOverflow/internal/models"
	"github.com/Samestora/WebOverflow/pkg/databases"
	"github.com/gofiber/fiber/v2"
)

// Get all user, for listing purposes only
// @Require nothing
func GetUsers(c *fiber.Ctx, db *databases.Database) error {
	type returner struct{
		Username	string
		Email		string
	}

	var users []returner;
	rows, err := db.Queryx("SELECT Username, Email FROM Users");
	if err != nil {
		log.Fatal(err);
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	defer rows.Close();

	for rows.Next() {
		var user returner;
		err := rows.StructScan(&user);

		if err != nil {
			log.Fatal(err);
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		users = append(users, user);
	}
	return c.Status(200).JSON(users);
}

// Get all user info via login route
// @Require Username, Password
func GetUserInfo(c *fiber.Ctx, db *databases.Database) error {
	m := c.Queries()
	if m["username"] == "" || m["password"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "invalid queries, must have username and password",
		});
	}
	userdata := new(models.User);

	err := db.Get(userdata, "SELECT * FROM Users WHERE Username = $1 AND Password = $2;", m["username"], m["password"]);
	if err != nil {
		return c.Status(403).JSON(fiber.Map{
			"message": "Invalid Password or Username",
		});
	}
	return c.Status(200).JSON(userdata);
}

// Make a user
// @Require Username, Password, Email
func CreateUser(c *fiber.Ctx, db *databases.Database) error {
	user := new(models.User);
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		});
	}

	if user != nil {
		_, err := db.NamedExec("INSERT INTO users (Username, Password, Email) VALUES (:username, :password, :email)", user);
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            	"message": "Failed to create user",
	        });
    	}
		return c.Status(200).JSON(fiber.Map{
			"message": "User successfully created!",
		});
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Invalid Parameter",
	})
}

// Delete a user
// @Require Username and Password
func DeleteUser (c *fiber.Ctx, db *databases.Database) error {
	user := new(models.User);

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		});
	}

	if user != nil {
		_, err := db.NamedExec("DELETE FROM users WHERE username = (:username) AND password = (:password)", user);
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            	"message": "Failed to delete user",
	        });
    	}
		return c.Status(200).JSON(fiber.Map{
			"message": "User is successfully deleted",
		});
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Invalid Parameter",
	})
}
