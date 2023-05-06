package orm

import (
	"gorm.io/gorm"
)

type Postgres[M any] interface {
	BaseInterface
	Methods[M]
}

type postgres[M any] struct {
	BaseInterface
}

func NewPostgres[M any](db gorm.DB) postgres[M] {

	b := NewBase(db)

	return postgres[M]{
		BaseInterface: b,
	}
}
