package controllers

import (
	"new_go_app/server/src/database"
	"new_go_app/server/src/models"

	fiber "github.com/gofiber/fiber/v2"
)

func GetReply(c *fiber.Ctx) error {
	var replys []models.Reply

	database.DB.Where("reply_id = ?", c.Params("id")).Find(&replys)

	return c.JSON(replys)
}
