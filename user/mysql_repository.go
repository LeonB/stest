package user

type MysqlRepository struct {
}

func (r MysqlRepository) Save(user User) error {
	return nil
}

func (r MysqlRepository) FindByEmail(email string) (User, error) {
	return User{}, nil
}

func (r MysqlRepository) FindById(id int) (User, error) {
	return User{}, nil
}

func (r MysqlRepository) All() (CollectionRepository, error) {
	return &MysqlCollection{}, nil
}
