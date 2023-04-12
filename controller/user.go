package controller

import (
	"net/http"

	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"github.com/dimasyudhana/alterra-group-project-2/err"
	errr "github.com/dimasyudhana/alterra-group-project-2/err"
	"github.com/dimasyudhana/alterra-group-project-2/helper"
	"github.com/dimasyudhana/alterra-group-project-2/service/user"
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
		if err2, ok := err.(errr.BadRequest); ok {
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
		c.Logger().Errorf("Error: %v", err1)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Bad Request", nil))
	}
	file, err1 := c.FormFile("image")
	if err1 != nil {
		u.Dep.Log.Errorf("Controller : %v", err1)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Bad Request", nil))
	}
	if err1 := u.Service.Register(c.Request().Context(), req, file); err1 != nil {
		if err1, ok := err1.(err.BadRequest); ok {
			return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, err1.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.CreateWebResponse(http.StatusInternalServerError, "Internal Server Error", nil))
		}
	}
	return c.JSON(http.StatusCreated, helper.CreateWebResponse(http.StatusCreated, "Successful Operation", nil))
}
