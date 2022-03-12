package models

type Vote struct {
	Base
	UserId uint `json:"user_id"`
	PostId uint `json:"post_id"`
	Value  int  `json:"value"`
}
