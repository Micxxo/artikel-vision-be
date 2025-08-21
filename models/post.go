package models

import "time"

type Post struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(200);not null" json:"title" validate:"required,min=20"`
	Content   string    `gorm:"type:text;not null" json:"content" validate:"required,min=200"`
	Category  string    `gorm:"type:varchar(100);not null" json:"category" validate:"required,min=3"`
	Status    string    `gorm:"type:enum('publish','draft','thrash');default:'draft'" json:"status" validate:"required,oneof=publish draft thrash"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
