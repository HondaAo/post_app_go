package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	Id        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	t := GenerateISOString()
	base.CreatedAt, base.UpdatedAt = t, t

	return nil
}

func (base *Base) AfterUpdate(tx *gorm.DB) error {
	base.UpdatedAt = GenerateISOString()
	return nil
}
