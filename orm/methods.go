package orm

import "fmt"

type Methods[T any] interface {
	Create(T) error
	Update(T) error
	Delete(T) error
	FindByUuid(uuid string) (*T, error)
}

func (p postgres[M]) Create(model M) error {
	fmt.Println(model)
	return nil
}
func (p postgres[M]) Update(model M) error {
	return nil
}
func (p postgres[M]) Delete(model M) error {
	return nil
}
func (p postgres[M]) FindByUuid(uuid string) (*M, error) {
	return nil, nil
}
