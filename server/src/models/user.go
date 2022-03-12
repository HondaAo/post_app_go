package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	IsAdmin  bool   `json:"-"`
	Posts    []Post `json:"posts" gorm:"foreignKey:UserId"`
	Votes    []Vote `json:"votes" gorm:"foreignKey:UserId"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
