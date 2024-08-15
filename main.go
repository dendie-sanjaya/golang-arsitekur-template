package main

import (
	"den/config"
	"den/entity"
	"den/migrate"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	dsn := "host=" + config.PostgresHost + " port=" + config.PostgressPort + " user=" + config.PostgresUser + " password=" + config.PostgresPassword + " dbname= " + config.PostgresDB + " sslmode=" + config.PostgresSSLMode

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

func app() *fiber.App {
	app := fiber.New()
	db := initDatabase()

	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(entity.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		db.Create(&user)
		return c.Status(201).JSON(user)
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user entity.User
		if result := db.First(&user, id); result.Error != nil {
			return c.Status(404).JSON(result.Error)
		}
		return c.JSON(user)
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user entity.User
		if result := db.First(&user, id); result.Error != nil {
			return c.Status(404).JSON(result.Error)
		}
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		db.Save(&user)
		return c.JSON(user)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user entity.User
		if result := db.First(&user, id); result.Error != nil {
			return c.Status(404).JSON(result.Error)
		}
		db.Delete(&user)
		return c.SendStatus(204)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []entity.User
		limit := c.QueryInt("limit", 10) // Default limit to 10 if not provided
		order := c.Query("order", "asc") // Default order to ascending if not provided

		if result := db.Order("created_at " + order).Limit(limit).Find(&users); result.Error != nil {
			return c.Status(500).JSON(result.Error)
		}
		return c.JSON(users)
	})

	return app
}

func main() {
	migrate.Run_db()
	migrate.Run_seed()
	app := app()
	log.Fatal(app.Listen(":4000"))
}
