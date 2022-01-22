package repository

type ProductRepository interface {
	FindById(id int) (Produc, error)
	FindAll() ([]Produc, error)
}
