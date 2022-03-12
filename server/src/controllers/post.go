package controllers

import (
	"new_go_app/src/database"
	"new_go_app/src/middleware"
	"new_go_app/src/models"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var post models.Post

	if err := c.BodyParser(&post); err != nil {
		return err
	}

	id, _ := middleware.GetUserId(c)

	post.UserId = id

	database.DB.Create(&post)

	return c.JSON(post)
}

func Posts(c *fiber.Ctx) error {
	var posts []models.Post

	database.DB.Preload("Vote").Find(&posts)

	return c.JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	var post models.Post

	id, _ := strconv.Atoi(c.Params("id"))
	post.Id = uint(id)

	database.DB.Find(&post)

	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	var post models.Post

	id, _ := strconv.Atoi(c.Params("id"))
	post.Id = uint(id)

	database.DB.Find(&post)

	userId, _ := middleware.GetUserId(c)

	if post.UserId != userId {
		return c.JSON(fiber.Map{
			"message": "You cannot delete this post.",
		})
	}

	database.DB.Delete(&post)

	return c.JSON(fiber.Map{
		"message": "Deleted.",
	})
}
