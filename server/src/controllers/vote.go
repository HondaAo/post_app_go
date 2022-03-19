package controllers

import (
	"log"
	"new_go_app/src/database"
	"new_go_app/src/middleware"
	"new_go_app/src/models"

	fiber "github.com/gofiber/fiber/v2"
)

func Vote(c *fiber.Ctx) error {
	var data map[string]int

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	log.Print(data)
	userId, _ := middleware.GetUserId(c)
	postId := data["post_id"]
	value := data["value"]

	var vote models.Vote
	var post models.Post
	post.Id = uint(postId)
	database.DB.Find(&post)

	if count := database.DB.Where(&models.Vote{UserId: userId}, &models.Vote{PostId: uint(postId)}).First(&vote); count.RowsAffected > 0 {
		if vote.Value == value {
			return c.JSON(vote)
		}

		post.Points += value * 2

		vote.Value = value
		database.DB.Save(&vote)
		database.DB.Save(&post)
		return c.JSON(post)
	}

	vote.UserId = uint(userId)
	vote.PostId = uint(postId)
	vote.Value = value

	database.DB.Create(&vote)

	post.Points += value
	database.DB.Save(&post)

	return c.JSON(vote)
}
