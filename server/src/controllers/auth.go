package controllers

import (
	"context"
	"log"
	"new_go_app/src/database"
	"new_go_app/src/middleware"
	"new_go_app/src/models"
	"strconv"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match.",
		})
	}

	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		IsAdmin:  false,
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	// utils.AuthMail(utils.MailBody{})

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User not found.",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password is wrong.",
		})
	}

	token, err := middleware.GenerateJWT(user.Id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User not found.",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Me(c *fiber.Ctx) error {

	id, _ := middleware.GetUserId(c)

	var user models.User

	database.DB.Where("id = ?", id).Preload("Posts").Preload("Votes").First(&user)

	if user.Id == 0 {
		return c.JSON(fiber.Map{
			"message": "not authed",
		})
	}

	return c.JSON(user)
}

func ForgetPassword(c *fiber.Ctx) string {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return "err"
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return "user not found"
	}

	token, _ := uuid.NewRandom()

	redis := database.SetupRedis()

	context := context.Background()

	err := redis.Set(context, token.String(), user.Id, 1000*60*60*24)

	if err != nil {
		panic(err)
	}

	// utils.ChangePasswordMail(data["email"], `<a href="http://localhost:3000/change-password/{token}">reset password</a>`)

	return token.String()

}
func ChangePassowrd(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	redis := database.SetupRedis()

	user_id, err := redis.Get(context.Background(), data["taken"]).Result()
	if err != nil {
		log.Fatal(err)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	var user models.User

	userId, _ := strconv.ParseInt(user_id, 10, 64)

	database.DB.Where("Id = ?", uint(userId)).First(&user)

	user.Password = []byte(data["newPassword"])

	database.DB.Save(&user)

	return c.JSON(fiber.Map{
		"message": "password was changed.",
	})
}
