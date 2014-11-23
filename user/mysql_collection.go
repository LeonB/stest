package user

import "time"

type MysqlCollection struct {
}

func (c *MysqlCollection) AddActiveFilter() error {
	return nil
}

func (c *MysqlCollection) AddCreatedAtFilter(from time.Time, to time.Time) error {
	return nil
}

func (c *MysqlCollection) Get() ([]User, error) {
	return []User{
		{id: 1, email: "test@gmail.com", password: "test", createdAt: time.Now()},
		{id: 2, email: "test2@gmail.com", password: "test2", createdAt: time.Now()},
	}, nil
}

func (c *MysqlCollection) Slice(offset int, length int) error {
	return nil
}
