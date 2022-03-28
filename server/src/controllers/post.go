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

	database.DB.Preload("Vote").Order("created_at").Find(&posts)

	return c.JSON(posts)
}

func PostIndex(c *fiber.Ctx) error {
	var posts []models.Post
	page, _ := strconv.Atoi(c.Params("page"))

	limit := page * 10
	realLimit := (page - 1) * 10

	database.DB.Preload("Vote").Order("created_at").Limit(limit).Find(&posts)

	posts = posts[realLimit:]

	return c.JSON(posts)
}

func SortPost(c *fiber.Ctx) {

}

func GetPost(c *fiber.Ctx) error {
	var post models.Post

	id, _ := strconv.Atoi(c.Params("id"))
	post.Id = uint(id)

	database.DB.Preload("Tags").Preload("Replies").Preload("Vote").Find(&post)

	return c.JSON(post)
}

func CreateTag(c *fiber.Ctx) error {
	var tag models.Tag

	if err := c.BodyParser(&tag); err != nil {
		return err
	}

	database.DB.Create(&tag)

	return c.JSON(tag)
}

func Tags(c *fiber.Ctx) error {
	var tags []models.Tag

	database.DB.Preload("Post").Find(&tags)

	return c.JSON(tags)
}

func GetTag(tag_id int) (models.Tag, error) {
	var tag models.Tag

	database.DB.Find(&tag, tag_id)

	return tag, nil
}

func AddTags(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var post models.Post

	id, _ := strconv.Atoi(c.Params("id"))
	tag_id, _ := strconv.Atoi(c.Params("tag_id"))
	post.Id = uint(id)

	tag, _ := GetTag(tag_id)

	database.DB.Find(&post).Association("Tags").Append(&tag)

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

func CreateComment(c *fiber.Ctx) error {
	var comment models.Reply

	if err := c.BodyParser(&comment); err != nil {
		return err
	}

	id, _ := middleware.GetUserId(c)
	comment.UserId = id

	post_id, _ := strconv.Atoi(c.Params("id"))
	comment.PostId = uint(post_id)

	database.DB.Create(&comment)

	return c.JSON(comment)

}
