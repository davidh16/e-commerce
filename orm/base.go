package orm

import "gorm.io/gorm"

type BaseInterface interface {
	Db() gorm.DB
}

type Base struct {
	db gorm.DB
}

func NewBase(db gorm.DB) Base {
	return Base{
		db: db,
	}
}

func (b Base) Db() gorm.DB {
	return b.db
}
