package repositories

import "time"

type BaseModel struct {
	ID         uint `gorm:"primarykey"`
	CreateDate time.Time
	UpdateDate time.Time
}

type User struct {
	BaseModel
}
