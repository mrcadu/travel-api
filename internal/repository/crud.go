package repository

type CRUD[model any] interface {
	Get(id string) (model, error)
	Create(model) (model, error)
	Update(model) (model, error)
	Delete(id string) (string, error)
}
