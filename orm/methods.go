package orm

type Methods[T any] interface {
	Create(T) error
	Update(T) error
	Delete(T) error
	FindByUuid(uuid string) (*T, error)
}
