package user

import (
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	Update(db *gorm.DB, req entities.User) error
	FindById(db *gorm.DB, id int) (*entities.User, error)
	FindByEmail(db *gorm.DB, email string) (*entities.User, error)
	Create(db *gorm.DB, req entities.User) error
}
