package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"

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
	if user == nil || user.Email == "" {
		return err.NewErr("Email belum Terdaftar"), 0
	}
	if err1 != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) || user.Id == 0 {
			return err.NewErr(err1.Error()), 0
		}
		if !errors.Is(err1, gorm.ErrRecordNotFound) {
			return err1, 0
		}
	}
	if err2 := helper.VerifyPassword(user.Password, req.Password); err2 != nil {
		return err.NewErr("Password Salah"), 0
	}
	return nil, int(user.Id)
}

func (u *user) Register(ctx context.Context, req entities.UserReqRegister, filehead *multipart.FileHeader) error {
	if err1 := u.validator.Struct(req); err1 != nil {
		u.dep.Log.Errorf("Error Service : %v", err1)
		return err.NewErr(err1.Error())
	}
	userd, err1 := u.repo.FindByEmail(u.dep.Db.WithContext(ctx), req.Email)
	if err1 != nil {
		if !errors.Is(err1, gorm.ErrRecordNotFound) {
			u.dep.Log.Errorf("Error Service : %v", err1)
			return err.NewErrInter("Gagal mencari data user")
		}
	}
	if userd.Id != 0 {
		return err.NewErr("Email sudah terdaftar!!!")
	}
	passhash, err1 := helper.HashPassword(req.Password)
	if err1 != nil {
		u.dep.Log.Errorf("Error Service : %v", err1)
		return err.NewErr("Gagal membuat akun")
	}
	file, err1 := filehead.Open()
	defer file.Close()
	if err1 != nil {
		u.dep.Log.Errorf("failed to open file", err1)
		return err1
	}
	filename := fmt.Sprintf("%s_%s", "User", filehead.Filename)

	if err1 := u.dep.Gcp.UploadFile(file, filename); err1 != nil {
		log.Print(err1)
		u.dep.Log.Errorf("Error Service : %v", err1)
		return err.NewErr("Gagal membuat pada saat mengupload gambar")
	}
	user := entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passhash,
		Image:    filename,
		Address:  req.Address,
	}
	if err1 := u.repo.Create(u.dep.Db.WithContext(ctx), user); err1 != nil {
		u.dep.Log.Errorf("Error Service : %v", err1)
		return err.NewErrInter("Terjadi kesalahan pada server")
	}
	return nil
}

func (u *user) GetById(ctx context.Context, id int) (*entities.User, error) {

	user, err1 := u.repo.FindById(u.dep.Db.WithContext(ctx), id)
	if err1 != nil {
		if !errors.Is(err1, gorm.ErrRecordNotFound) {
			u.dep.Log.Errorf("Error Service : %v", err1)
			return nil, err.NewErrInter("Gagal mencari data user")
		}
	}
	if user.Id == 0 {
		return nil, errors.New("Id tidak terdaftar")
	}

	return user, nil
}

func (u *user) Update(ctx context.Context, req entities.UserReqUpdate, filehead *multipart.FileHeader, id int) error {
	if err1 := u.validator.Struct(req); err1 != nil {
		u.dep.Log.Errorf("Error Service : %v", err1)
		return err.NewErr(err1.Error())
	}
	userd, err1 := u.repo.FindByEmail(u.dep.Db.WithContext(ctx), req.Email)
	if err1 != nil {
		if !errors.Is(err1, gorm.ErrRecordNotFound) {
			u.dep.Log.Errorf("Error Service : %v", err1)
			return err.NewErrInter("Gagal mencari data user")
		}
	}
	if userd.Id != 0 {
		return err.NewErr("Email sudah terdaftar!!!")
	}
	if userd.Password != req.Password {
		passhash, err1 := helper.HashPassword(req.Password)
		if err1 != nil {
			u.dep.Log.Errorf("Error Service : %v", err1)
			return err.NewErr("Gagal membuat akun")
		}
		req.Password = passhash
	}
	if filehead.Filename != req.Image {
		file, err1 := filehead.Open()
		defer file.Close()
		if err1 != nil {
			u.dep.Log.Errorf("failed to open file", err1)
			return err.NewErr("gagal memuat gambar")
		}
		filename := fmt.Sprintf("%s_%s", "User", filehead.Filename)

		if err1 := u.dep.Gcp.UploadFile(file, filename); err1 != nil {
			log.Print(err1)
			u.dep.Log.Errorf("Error Service : %v", err1)
			return err.NewErr("Gagal membuat pada saat mengupload gambar")
		}
		req.Image = filename
	}
	user := entities.User{
		Id:       uint(id),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Image:    req.Image,
		Address:  req.Address,
	}
	if err1 := u.repo.Update(u.dep.Db.WithContext(ctx), user); err1 != nil {
		u.dep.Log.Errorf("Error Service : %v", err1)
		return err.NewErrInter("Terjadi kesalahan pada server")
	}
	return nil
}
