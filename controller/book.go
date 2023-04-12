package controller

import (
	"log"
	"net/http"
	"strconv"

	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"github.com/dimasyudhana/alterra-group-project-2/service/book"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type BookController struct {
	dig.In
	Service book.Service
	Dep     dependecy.Depend
}

type BookRequest struct {
	Title    string `json:"title"`
	Year     string `json:"year"`
	Author   string `json:"author"`
	Contents string `json:"contents"`
	Image    string `json:"image"`
}

type UpdateRequest struct {
	Title    string `json:"title"`
	Year     string `json:"year"`
	Author   string `json:"author"`
	Contents string `json:"contents"`
	Image    string `json:"image"`
}

func (bc *BookController) InsertBook(c echo.Context) error {

	// user_id, err := helper.DecodeJWT(c.Get("user").(*jwt.Token))
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
	// 		"code":    http.StatusUnauthorized,
	// 		"message": "Invalid or expired JWT",
	// 	})
	// }

	input := BookRequest{}
	if err := c.Bind(&input); err != nil {
		c.Logger().Error("terjadi kesalahan bind", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	if input.Title == "" || input.Year == "" || input.Author == "" || input.Contents == "" || input.Image == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request, Invalid Input",
		})
	}

	_, err := bc.Service.InsertBook(c.Request().Context(), entities.Core{
		Title:    input.Title,
		Year:     input.Year,
		Author:   input.Author,
		Contents: input.Contents,
		Image:    input.Image,
	})
	if err != nil {
		c.Logger().Error("terjadi kesalahan", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create a Book",
	})
}

func (bc *BookController) GetAllBooks(c echo.Context) error {

	books, err := bc.Service.GetAllBooks(c.Request().Context())
	if err != nil {
		c.Logger().Error("terjadi kesalahan", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"data": books,
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

func (bc *BookController) UpdateByBookID(c echo.Context) error {

	// create an instance of UpdateRequest to bind the request body
	update := new(UpdateRequest)
	if err := c.Bind(update); err != nil {
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

	if err := bc.Service.UpdateByBookID(c.Request().Context(), uint(bookID), update.ToEntity()); err != nil {
		c.Logger().Error("Failed to update book", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	book, err := bc.Service.GetBookByBookID(c.Request().Context(), uint(bookID))
	if err != nil {
		c.Logger().Error("Failed to get updated book", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	// create a response map with the updated book details
	response := map[string]interface{}{
		"title":    book.Title,
		"year":     book.Year,
		"author":   book.Author,
		"contents": book.Contents,
		"image":    book.Image,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success update a book",
		"data":    response,
	})
}

func (ur *UpdateRequest) ToEntity() entities.Book {
	return entities.Book{
		Title:    ur.Title,
		Year:     ur.Year,
		Author:   ur.Author,
		Contents: ur.Contents,
		Image:    []byte(ur.Image),
	}
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
