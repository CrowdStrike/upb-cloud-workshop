package domain

//go:generate mockgen -source=storage.go -package=mocks -destination=mocks/storage.go

type Storage interface {
	GetContent(id string) (string, error)
	WriteContent(id string, content string) error
}
