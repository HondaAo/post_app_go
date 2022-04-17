package main

import (
	"new_go_app/server/src/database"
	"new_go_app/server/src/models"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()
	for i := 0; i < 30; i++ {
		post := models.Post{
			Title:  faker.Sentence(),
			Text:   faker.Paragraph(),
			UserId: 1,
		}
		database.DB.Create(&post)
	}
}
