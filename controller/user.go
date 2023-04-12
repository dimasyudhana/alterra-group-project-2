package controller

import (
	"net/http"

	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"github.com/dimasyudhana/alterra-group-project-2/err"
	errr "github.com/dimasyudhana/alterra-group-project-2/err"
	"github.com/dimasyudhana/alterra-group-project-2/helper"
	"github.com/dimasyudhana/alterra-group-project-2/service/user"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type User struct {
	dig.In
	Service user.UserServiceInterface
	Dep     dependecy.Depend
}

func (u *User) Login(c echo.Context) error {
	var req entities.UserReqLogin
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, err.Error(), nil))
	}
	err, id := u.Service.Login(c.Request().Context(), req)
	if err != nil {
		u.Dep.Log.Errorf("Controller : %v", err)
		if err2, ok := err.(errr.BadRequest); ok {
			u.Dep.Log.Errorf("Controller : %v", err2)
			return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, err2.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.CreateWebResponse(http.StatusInternalServerError, err.Error(), nil))
	}
	token := helper.GenerateJWT(id, u.Dep)
	return c.JSON(http.StatusOK, helper.CreateWebResponse(http.StatusOK, "Successful Operation", map[string]interface{}{"Token": token}))
}

func (u *User) Register(c echo.Context) error {
	var req entities.UserReqRegister
	if err1 := c.Bind(&req); err1 != nil {
		u.Dep.Log.Errorf("Controller : %v", err1)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Bad Request", nil))
	}
	file, err1 := c.FormFile("image")
	if err1 != nil {
		u.Dep.Log.Errorf("Controller : %v", err1)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Bad Request", nil))
	}
	if err1 := u.Service.Register(c.Request().Context(), req, file); err1 != nil {
		if err1, ok := err1.(err.BadRequest); ok {
			u.Dep.Log.Errorf("Controller : %v", err1)
			return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, err1.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.CreateWebResponse(http.StatusInternalServerError, "Internal Server Error", nil))
		}
	}
	return c.JSON(http.StatusCreated, helper.CreateWebResponse(http.StatusCreated, "Successful Operation", nil))
}

func (u *User) GetById(c echo.Context) error {
	uid := helper.GetUid(c.Get("user").(*jwt.Token))
	user, err1 := u.Service.GetById(c.Request().Context(), uid)
	var res = struct {
		Csrf string
		Data any
	}{Csrf: c.Get("csrf").(string), Data: user}
	if err1 != nil {
		u.Dep.Log.Errorf("Controller : %v", err1)
		if _, ok := err1.(err.BadRequest); ok {
			return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Bad Request", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.CreateWebResponse(http.StatusInternalServerError, "Status Internal Error", nil))
		}
	}
	return c.JSON(http.StatusOK, helper.CreateWebResponse(http.StatusOK, "Successful Operation", res))
}
func (u *User) Update(c echo.Context) error {
	var req entities.UserReqUpdate
	uid := helper.GetUid(c.Get("user").(*jwt.Token))
	if err1 := c.Bind(&req); err1 != nil {
		c.Logger().Errorf("Error: %v", err1)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Bad Request", nil))
	}
	file, err1 := c.FormFile("image")
	if err1 != nil {
		u.Dep.Log.Errorf("Controller : %v", err1)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Bad Request", nil))
	}
	if err1 := u.Service.Update(c.Request().Context(), req, file, uid); err1 != nil {
		u.Dep.Log.Errorf("Controller : %v", err1)
		if err1, ok := err1.(err.BadRequest); ok {
			u.Dep.Log.Errorf("Controller : %v", err1)
			return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, err1.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.CreateWebResponse(http.StatusInternalServerError, "Internal Server Error", nil))
		}
	}

	return c.JSON(http.StatusCreated, helper.CreateWebResponse(http.StatusCreated, "Successful Operation", nil))
}
