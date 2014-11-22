package user

import "time"

type Repository interface {
	Save(user User) error
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	All() (CollectionRepository, error)
}

type CollectionRepository interface {
	AddActiveFilter() (CollectionRepository, error)
	AddCreatedAtFilter(time.Time, time.Time) (CollectionRepository, error)
	Get() ([]User, error)
	Slice(int, int) (CollectionRepository, error)
}
