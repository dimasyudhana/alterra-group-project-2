package user

import (
	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"gorm.io/gorm"
)

type user struct {
	dep dependecy.Depend
}

func NewUserRepo() UserRepoInterface {
	return &user{}
}

func (u *user) Create(db *gorm.DB, req entities.User) error {
	if err := db.Create(&req).Error; err != nil {
		u.dep.Log.Errorf("Error %v", err)
		return err
	}
	return nil
}
func (u *user) FindByEmail(db *gorm.DB, email string) (*entities.User, error) {
	var user entities.User
	if err := db.Where("email = ?", email).Find(&user).Error; err != nil {
		u.dep.Log.Errorf("Error %v", err)
		return nil, err
	}
	return &user, nil
}

func (u *user) FindById(db *gorm.DB, id int) (*entities.User, error) {
	var user entities.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		u.dep.Log.Errorf("Error %v", err)
		return nil, err
	}
	return &user, nil
}

func (u *user) Update(db *gorm.DB, req entities.User) error {
	if err := db.Save(&req).Error; err != nil {
		u.dep.Log.Errorf("Error %v", err)
		return err
	}
	return nil
}
