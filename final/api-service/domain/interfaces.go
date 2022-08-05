package domain

//go:generate mockgen -source=interfaces.go -package=mocks -destination=../mocks/mock_interfaces.go

// Queue represents a queue for passing batches of data to the store service
type Queue interface {
	Add(batch []Product) error
}

// Storage is used for storing objects
type Storage interface {
	Save(product Product) (string, bool, error)
	Get(id string) (Product, bool, error)
	Update(id string, diff Product) (bool, error)
	Delete(id string) (bool, error)
}
