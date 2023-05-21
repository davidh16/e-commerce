package orm

type Methods[T any] interface {
	Create(T) (*T, error)
	Update(T) (*T, error)
	Delete(T) (*T, error)
	FindByUuid(uuid string) (*T, error)
}

func (p postgres[M]) Create(data M) (*M, error) {
	result := p.Db().Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil

}
func (p postgres[M]) Update(model M) (*M, error) {
	return nil, nil
}
func (p postgres[M]) Delete(model M) (*M, error) {
	return nil, nil
}
func (p postgres[M]) FindByUuid(uuid string) (*M, error) {
	return nil, nil
}
