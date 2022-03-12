package models

type Post struct {
	Base
	Title  string `json:"title"`
	Text   string `json:"text"`
	Points int    `json:"points"`
	UserId uint   `json:"user_id"`
	Vote   []Vote `json:"votes" gorm:"foreignKey:PostId"`
}
