package repositories

import "time"

type BaseModel struct {
	ID         uint `gorm:"primarykey"`
	CreateDate time.Time
	UpdateDate time.Time
}

type User struct {
	BaseModel
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
}

type Role struct {
	BaseModel
	Code string
	Name string
	Desc string
}
