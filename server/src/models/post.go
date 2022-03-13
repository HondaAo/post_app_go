package models

type Record struct {
	Base
	Text   string `json:"text"`
	UserId uint   `json:"user_id"`
}

type Post struct {
	Record
	Title   string  `json:"title"`
	Points  int     `json:"points"`
	Vote    []Vote  `json:"votes" gorm:"foreignKey:PostId"`
	Replies []Reply `json:"replies" gorm:"foreignKey:PostId"`
}

type Reply struct {
	Record
	PostId uint `json:"post_id"`
}
