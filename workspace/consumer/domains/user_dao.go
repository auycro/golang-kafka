package domains

import "gorm.io/gorm"

type UserRepository interface {
	Save(user User) error
	FindAll() ([]User, error)
	//FindUser(int64) (*User, error)
	//UpdateUser(int64) error
	//DeleteUser(int64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	//db.Table("test_users").AutoMigrate(&User{})
	db.AutoMigrate(&User{})
	return userRepository{db}
}

func (obj userRepository) Save(user User) error {
	return obj.db.Save(user).Error
}

func (obj userRepository) FindAll() ([]User, error) {
	users := []User{}
	err := obj.db.Find(&users).Error
	return users, err
}
