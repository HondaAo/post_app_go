package test

import (
	"new_go_app/src/models"
	"testing"
)

func CreatePostTest(t *testing.T) {
	post := models.Post{
		Title:  "test_title",
		UserId: 1,
	}
}
