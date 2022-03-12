package main

import (
	"log"
	"new_go_app/src/database"
	"new_go_app/src/models"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()
	for i := 0; i < 30; i++ {
		post := models.Post{
			Title:  faker.Username(),
			Text:   faker.Word(),
			UserId: 1,
		}
		log.Print(post)

		database.DB.Create(&post)
	}
}
