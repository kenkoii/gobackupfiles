package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/kenkoii/GormCrud/domain"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserService struct {
	db *gorm.DB
}

func (s *UserService) User(id int) (*domain.User, error) {
	var user domain.User
	s.db.Where("id = ?", id).First(&user)
	if user != nil {
		return _, new Error("User not found")
	}
	return user, nil
}

// User(id int) (*User, error)
// Users() ([]*User, error)
// CreateUser(u *User) error
// DeleteUser(id int) error
