package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"github.com/dimasyudhana/alterra-group-project-2/helper"
	"github.com/dimasyudhana/alterra-group-project-2/service/book"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type BookController struct {
	dig.In
	Service book.Service
	Dep     dependecy.Depend
}

type CreateRequest struct {
	Title    string `form:"title"`
	Year     string `form:"year"`
	Author   string `form:"author"`
	Contents string `form:"contents"`
	Image    string
}

type UpdateRequest struct {
	Title    string `form:"title"`
	Year     string `form:"year"`
	Author   string `form:"author"`
	Contents string `form:"contents"`
	Image    string
}

type BookFormatResponse struct {
	Title    string `json:"title"`
	Year     string `json:"year"`
	Author   string `json:"author"`
	Contents string `json:"contents"`
	Image    string
	Username string `json:"username"`
}

func (bc *BookController) InsertBook(c echo.Context) error {
	userID := getUserIdFromToken(c)
	input := CreateRequest{}
	if err := c.Bind(&input); err != nil {
		bc.Dep.Log.Errorf("Error Controller : %v", err)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Bad Request", nil))
	}

	FileHeader, err := c.FormFile("image")
	if err != nil {
		bc.Dep.Log.Errorf("Error Controller : %v", err)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Bad Request", nil))
	}

	file, err := FileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.CreateWebResponse(http.StatusInternalServerError, "Bad Request", nil))
	}

	err1 := bc.Dep.Gcp.UploadFile(file, FileHeader.Filename)
	if err1 != nil {
		bc.Dep.Log.Errorf("Error Controller : %v", err)
		return c.JSON(http.StatusInternalServerError, helper.CreateWebResponse(http.StatusInternalServerError, "Internal Server Error", nil))
	}

	book := entities.Core{
		Title:    input.Title,
		Year:     input.Year,
		Author:   input.Author,
		Contents: input.Contents,
		Image:    FileHeader.Filename,
		UserID:   userID,
	}

	if err := bc.Service.InsertBook(c.Request().Context(), book); err != nil {
		bc.Dep.Log.Errorf("Error Controller : %v", err)
		return c.JSON(http.StatusInternalServerError, helper.CreateWebResponse(http.StatusInternalServerError, "Internal Server Error", nil))
	}

	return c.JSON(http.StatusCreated, helper.CreateWebResponse(http.StatusCreated, "Success Create a Book", nil))
}

func getUserIdFromToken(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		log.Println("Invalid User ID in Token")
	}
	return uint(userID)
}

func (bc *BookController) UpdateBook(c echo.Context) error {
	input := new(UpdateRequest)
	if err := c.Bind(input); err != nil {
		log.Println("Failed to bind request body", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	bookID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Logger().Error("Failed to parse book ID", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	FileHeader, err := c.FormFile("image")
	if err != nil {
		bc.Dep.Log.Errorf("Error Controller : %v", err)
		return c.JSON(http.StatusBadRequest, helper.CreateWebResponse(http.StatusBadRequest, "Bad Request", nil))
	}

	var imageName string
	if FileHeader != nil {
		file, err := FileHeader.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.CreateWebResponse(http.StatusInternalServerError, "Bad Request", nil))
		}

		err1 := bc.Dep.Gcp.UploadFile(file, FileHeader.Filename)
		if err1 != nil {
			bc.Dep.Log.Errorf("Error Controller : %v", err1)
			return c.JSON(http.StatusInternalServerError, helper.CreateWebResponse(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		imageName = FileHeader.Filename
	}

	book := input.ToEntity()
	book.Image = imageName

	if err := bc.Service.UpdateByBookID(c.Request().Context(), uint(bookID), book); err != nil {
		c.Logger().Error("Failed to update book", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	book5, err := bc.Service.GetBookByBookID(c.Request().Context(), uint(bookID))
	if err != nil {
		c.Logger().Error("Failed to get updated book", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success update a book",
		"data":    book5,
	})
}

func (ur *UpdateRequest) ToEntity() entities.Book {
	return entities.Book{
		Title:    ur.Title,
		Year:     ur.Year,
		Author:   ur.Author,
		Contents: ur.Contents,
		Image:    ur.Image,
	}
}

func (bc *BookController) GetAllBooks(c echo.Context) error {
	books, err := bc.Service.GetAllBooks(c.Request().Context())
	if err != nil {
		c.Logger().Error("Failed to get all books", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	if len(books) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "No books found",
		})
	}

	formattedBooks := []BookFormatResponse{}
	for _, book := range books {
		formattedBook := BookFormatResponse{
			Title:    book.Title,
			Year:     book.Year,
			Author:   book.Author,
			Contents: book.Contents,
			Image:    book.Image,
		}
		formattedBooks = append(formattedBooks, formattedBook)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"data": formattedBooks,
	})
}

func (bc *BookController) GetBookByBookID(c echo.Context) error {
	inputID := c.Param("id")
	if inputID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	bookID, err := strconv.ParseUint(inputID, 10, 32)
	if err != nil {
		c.Logger().Error("terjadi kesalahan parse uint", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	book, err := bc.Service.GetBookByBookID(c.Request().Context(), uint(bookID))
	if err != nil {
		c.Logger().Error("terjadi kesalahan", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success get a book",
		"data":    book,
	})
}

func (bc *BookController) DeleteByBookID(c echo.Context) error {

	bookID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Logger().Error("Failed to parse book ID", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	if err := bc.Service.DeleteByBookID(c.Request().Context(), uint(bookID)); err != nil {
		c.Logger().Error("Failed to delete book", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success delete a book",
	})
}
