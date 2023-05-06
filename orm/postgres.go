package orm

import (
	"fmt"
	"gorm.io/gorm"
)

type Postgres[M any] interface {
	BaseInterface
	Methods[M]
}

type postgres[M any] struct {
	BaseInterface
}

func (p postgres[M]) Create(model M) error {
	fmt.Println("afkbakfnakfakfskfskfsfks")
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

func NewPostgres[M any](db gorm.DB) postgres[M] {

	b := NewBase(db)

	return postgres[M]{
		BaseInterface: b,
	}
}
