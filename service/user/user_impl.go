package user

import (
	"context"
	"errors"

	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"github.com/dimasyudhana/alterra-group-project-2/err"
	"github.com/dimasyudhana/alterra-group-project-2/helper"
	usrrepo "github.com/dimasyudhana/alterra-group-project-2/repository/user"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

const (
	Url = "XXXXXXXXXXX"
)

type user struct {
	repo      usrrepo.UserRepoInterface
	dep       dependecy.Depend
	validator *validator.Validate
}

func NewUserService(repo usrrepo.UserRepoInterface, dep dependecy.Depend) UserServiceInterface {
	return &user{repo: repo, dep: dep, validator: validator.New()}
}
func (u *user) Login(ctx context.Context, req entities.UserReqLogin) (error, int) {
	if err1 := u.validator.Struct(req); err1 != nil {
		return err.NewErr(err1.Error()), 0
	}
	user, err1 := u.repo.FindByEmail(u.dep.Db.WithContext(ctx), req.Email)
	if errors.Is(err1, gorm.ErrRecordNotFound) {
		return err.NewErr(err1.Error()), 0
	}
	if !errors.Is(err1, gorm.ErrRecordNotFound) {
		return err1, 0
	}
	if err1 := helper.VerifyPassword(user.Password, req.Password); err1 != nil {
		return err.NewErr("Password Salah"), 0
	}
	return nil, user.Id
}
