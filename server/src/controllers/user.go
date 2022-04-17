package controllers

import (
	"new_go_app/server/src/database"
	"new_go_app/server/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	var user models.User

	user_id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Where("id = ?", user_id).Preload("Posts").Preload("Votes").First(&user)

	return c.JSON(user)
}

func ChangeBio(c *fiber.Ctx) error {
	var user models.User

	user_id, _ := strconv.Atoi(c.Params("id"))

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	database.DB.Where("id = ?", user_id).First(&user).Update("bio", data["bio"])

	return c.JSON(user)
}
