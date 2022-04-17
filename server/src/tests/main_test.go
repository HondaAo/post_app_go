package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http/httptest"
	"new_go_app/server/src/database"
	"new_go_app/server/src/models"
	"new_go_app/server/src/routes"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	utils "github.com/gofiber/utils"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/assert.v1"
)

func TestMain(m *testing.M) {
	database.Conn()

	os.Setenv("ENV", "Test")

	os.Exit(m.Run())
}

func TestFiber(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, World!")
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}

func migrateUserTable() error {
	database.Conn()
	err := database.DB_TEST.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Print("DR ERROR")
	}
	err = database.DB_TEST.AutoMigrate(&models.User{})
	if err != nil {
		log.Print("MG Error")
	}
	log.Printf("Successfully migrate table")
	return nil
}

func migratePostTable() error {
	database.Conn()
	err := database.DB_TEST.Migrator().DropTable(&models.Post{})
	if err != nil {
		log.Print("DR ERROR")
	}
	err = database.DB_TEST.AutoMigrate(&models.Post{})
	if err != nil {
		log.Print("MG Error")
	}
	return nil
}

func SeedOneUser() (models.User, error) {
	err := migrateUserTable()

	if err != nil {
		log.Fatal(err)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("test123"), 12)

	user := models.User{
		Username: "test",
		Email:    "test@test.com",
		Password: password,
	}

	err = database.DB_TEST.Model(&models.User{}).Create(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func SeedOnePost() (models.Post, error) {
	err := migratePostTable()

	if err != nil {
		log.Fatal()
	}

	post := models.Post{
		Title: "title",
		Text:  "text",
	}

	err = database.DB_TEST.Model(&models.Post{}).Create(&post).Error

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func TestLogin(t *testing.T) {
	_, err := SeedOneUser()

	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}

	app := fiber.New()

	routes.Routes(app)

	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"email": "test@test.com", "password": "test123"}`,
			statusCode: 200,
		},
	}

	for _, v := range samples {
		req := httptest.NewRequest("POST", "/api/user/login", bytes.NewBufferString(v.inputJSON))
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)

		utils.AssertEqual(t, nil, err)
		utils.AssertEqual(t, 200, res.StatusCode)
	}

}

func TestGetPost(t *testing.T) {
	_, err := SeedOnePost()

	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	routes.Routes(app)

	req := httptest.NewRequest("GET", "/api/post/1", nil)
	res, err := app.Test(req)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	assert.Equal(t, 200, res.StatusCode)
}
