package controller

import (
	"den/entity"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserControler(app *fiber.App, db *gorm.DB) *fiber.App {
	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(entity.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		db.Create(&user)

		return c.JSON(fiber.Map{
			"status": "success",
			"code":   200,
			"data":   user,
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []entity.User
		limit := c.QueryInt("limit", 10) // Default limit to 10 if not provided
		order := c.Query("order", "asc") // Default order to ascending if not provided

		if result := db.Order("created_at " + order).Limit(limit).Find(&users); result.Error != nil {
			return c.Status(500).JSON(result.Error)
		}

		return c.JSON(fiber.Map{
			"status": "success",
			"code":   200,
			"data":   users,
		})
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user entity.User
		if result := db.First(&user, id); result.Error != nil {
			return c.Status(404).JSON(result.Error)
		}

		return c.JSON(fiber.Map{
			"status": "success",
			"code":   200,
			"data":   user,
		})
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

		return c.JSON(fiber.Map{
			"status": "success",
			"code":   200,
			"data":   user,
		})
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user entity.User
		if result := db.First(&user, id); result.Error != nil {
			return c.Status(404).JSON(result.Error)
		}
		db.Delete(&user)

		return c.JSON(fiber.Map{
			"status": "success",
			"code":   200,
			"data":   user,
		})
	})

	app.Get("/users2", func(c *fiber.Ctx) error {
		var users []entity.User
		limit := c.QueryInt("limit", 10) // Default limit to 10 if not provided
		order := c.Query("order", "asc") // Default order to ascending if not provided

		if result := db.Order("created_at " + order).Limit(limit).Find(&users); result.Error != nil {
			return c.Status(500).JSON(result.Error)
		}

		return c.JSON(fiber.Map{
			"status":   "success",
			"code":     200,
			"message:": "Hello World",
			"data":     users,
		})
	})
	return app
}
