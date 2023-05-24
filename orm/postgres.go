package orm

import (
	"gorm.io/gorm"
)

type Postgres interface {
	BaseInterface
}

type postgres struct {
	BaseInterface
}

func NewPostgres(db *gorm.DB) postgres {

	b := NewBase(db)

	return postgres{
		BaseInterface: b,
	}
}
