package models

import "time"

type BaseModel struct {
	ID         uint `gorm:"primarykey"`
	CreateDate *time.Time
	UpdateDate *time.Time
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

type Produc struct {
	BaseModel
	Name            string
	Alias           string
	SortDescription string
	FullDescription string
	CategoryID      uint
	BrandID         uint
}

type ProducDetail struct {
	BaseModel
	Name  string
	Value string
}

type ProductImage struct {
	BaseModel
	Name     string
	Url      string
	Path     string
	ProducID uint
}

type Category struct {
	BaseModel
	Name  string
	Alias string
	Image string
}

type Brand struct {
	BaseModel
	Name string
	Logo string
}
