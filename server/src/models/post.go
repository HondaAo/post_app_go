package models

type Record struct {
	Base
	Text   string `json:"text"`
	UserId uint   `json:"user_id"`
}

type Post struct {
	Base
	Text    string  `json:"text"`
	UserId  uint    `json:"user_id"`
	Title   string  `json:"title"`
	Points  int     `json:"points"`
	Tags    []Tag   `json:"tags" gorm:"many2many:post_tags;"`
	Vote    []Vote  `json:"votes" gorm:"foreignKey:PostId"`
	Replies []Reply `json:"replies" gorm:"foreignKey:PostId"`
}

type Tag struct {
	Base
	Text string `json:"text"`
	Post []Post `json:"posts" gorm:"many2many:post_tags;"`
}

type Reply struct {
	Base
	Text   string `json:"text"`
	UserId uint   `json:"user_id"`
	PostId uint   `json:"post_id"`
}
