package controller

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

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

func (u *Trx) MyTransaction(c echo.Context) error {
	uid := helper.GetUid(c.Get("user").(*jwt.Token))
	res, err := u.Service.FindMyTransaction(c.Request().Context(), uid)
	if err != nil {
		u.Dep.Log.Errorf("controller : %v", err)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, err.Error(), nil))
	}
	page := c.QueryParam("page")
	perPage := c.QueryParam("per_page")

	total := len(res)
	pageInt := 1
	if page != "" || perPage == "" {
		pageInt, _ = strconv.Atoi(page)
	}
	perPageInt, _ := strconv.Atoi(perPage)

	totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

	startIndex := (pageInt - 1) * perPageInt
	endIndex := startIndex + perPageInt
	if endIndex > total {
		endIndex = total
	}

	data := res[startIndex:endIndex]

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        http.StatusOK,
		"page":        pageInt,
		"per_page":    perPageInt,
		"total_pages": totalPages,
		"data":        data,
	})
}

func (u *Trx) GetAllAvailableBooks(c echo.Context) error {
	res, err := u.Service.GetAllAvailableBooks(c.Request().Context())
	if err != nil {
		u.Dep.Log.Errorf("controller : %v", err)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, err.Error(), nil))
	}
	page := c.QueryParam("page")
	perPage := c.QueryParam("per_page")

	total := len(res)
	pageInt := 1
	if page != "" || perPage == "" {
		pageInt, _ = strconv.Atoi(page)
	}
	perPageInt, _ := strconv.Atoi(perPage)

	totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

	startIndex := (pageInt - 1) * perPageInt
	endIndex := startIndex + perPageInt
	if endIndex > total {
		endIndex = total
	}

	data := res[startIndex:endIndex]

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        http.StatusOK,
		"page":        pageInt,
		"per_page":    perPageInt,
		"total_pages": totalPages,
		"data":        data,
	})
}

func (u *Trx) GetAllBorrowedBook(c echo.Context) error {
	uid := helper.GetUid(c.Get("user").(*jwt.Token))
	res, err := u.Service.GetAllBorrowedBooks(c.Request().Context(), uid)
	if err != nil {
		u.Dep.Log.Errorf("controller : %v", err)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, err.Error(), nil))
	}
	page := c.QueryParam("page")
	perPage := c.QueryParam("per_page")

	total := len(res)
	pageInt := 1
	if page != "" || perPage == "" {
		pageInt, _ = strconv.Atoi(page)
	}
	perPageInt, _ := strconv.Atoi(perPage)

	totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

	startIndex := (pageInt - 1) * perPageInt
	endIndex := startIndex + perPageInt
	if endIndex > total {
		endIndex = total
	}

	data := res[startIndex:endIndex]

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        http.StatusOK,
		"page":        pageInt,
		"per_page":    perPageInt,
		"total_pages": totalPages,
		"data":        data,
	})
}
