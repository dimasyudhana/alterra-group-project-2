package controller

import (
	"fmt"
	"net/http"

	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/helper"
	"github.com/dimasyudhana/alterra-group-project-2/service/transaction"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Trx struct {
	dig.In
	Service transaction.TrxServiceInterface
	Dep     dependecy.Depend
}

func (u *Trx) Createtrx(c echo.Context) error {
	var req = struct {
		Data []int `json:"data"`
	}{}
	if err := c.Bind(&req); err != nil {
		u.Dep.Log.Errorf("error %v", err)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, err.Error(), nil))
	}
	fmt.Println(len(req.Data))
	if len(req.Data) == 0 {
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Data tidak boleh kosong", nil))
	}
	uid := helper.GetUid(c.Get("user").(*jwt.Token))
	err := u.Service.Create(c.Request().Context(), req.Data, uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.CreateWebResponse(http.StatusOK, "Successful Operation", nil))
}
