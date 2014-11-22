package user

import "time"

type MysqlCollection struct {
}

func (c MysqlCollection) AddActiveFilter() (CollectionRepository, error) {
	return c, nil
}

func (c MysqlCollection) AddCreatedAtFilter(from time.Time, to time.Time) (CollectionRepository, error) {
	return c, nil
}

func (c MysqlCollection) Get() ([]User, error) {
	return []User{}, nil
}

func (c MysqlCollection) Slice(offset int, length int) (CollectionRepository, error) {
	return c, nil
}
