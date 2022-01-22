package repository

import "time"

type BaseModel struct {
	ID         uint       `gorm:"primarykey"`
	CreateDate *time.Time `gorm:"column:create_date"`
	UpdateDate *time.Time `gorm:"column:update_date"`
}

type User struct {
	BaseModel
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	Email     string `gorm:"column:email"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
}

type Role struct {
	BaseModel
	Code string `gorm:"column:code"`
	Name string `gorm:"column:name"`
	Desc string `gorm:"column:desc"`
}

type Produc struct {
	BaseModel
	Name       string `gorm:"column:name"`
	Alias      string `gorm:"column:alias"`
	ShortDesc  string `gorm:"column:short_descr"`
	FullDesc   string `gorm:"column:full_desc"`
	CategoryID uint   `gorm:"column:category_id"`
	BrandID    uint   `gorm:"column:brand_id"`
}

type ProducDetail struct {
	BaseModel
	Name  string `gorm:"column:name"`
	Value string `gorm:"column:value"`
}

type ProductImage struct {
	BaseModel
	Name     string `gorm:"column:name"`
	Url      string `gorm:"column:url"`
	Path     string `gorm:"column:path"`
	ProducID uint   `gorm:"column:product_id"`
}

type Category struct {
	BaseModel
	Name  string `gorm:"column:name"`
	Alias string `gorm:"column:alias"`
	Image string `gorm:"column:image"`
}

type Brand struct {
	BaseModel
	Name string `gorm:"column:name"`
	Logo string `gorm:"column:logo"`
}
