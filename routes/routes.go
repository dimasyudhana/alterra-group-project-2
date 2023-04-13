package routes

import (
	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/controller"
	"github.com/dimasyudhana/alterra-group-project-2/helper"
	"go.uber.org/dig"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routes struct {
	dig.In
	Depend dependecy.Depend
	User   controller.User
	Book   controller.BookController
	Trx    controller.Trx
}

func (r *Routes) CSRFMiddlewareCustom(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Mode 1 : Dimana token digunakan selamat 24 jam tanpa generete ulang
		//Mode 2 : Dimana token akan digenerate ulang setiap Endpoint dengan method selain GET di Hit
		if r.Depend.Config.CSRFMode == "1" {
			return next(c)
		}
		helper.DeleteCookieCSRF(c)
		return next(c)
	}
}

func (r *Routes) RegisterRoutes() {
	ro := r.Depend.Echo
	// No Auth
	ro.Use(middleware.RemoveTrailingSlash())
	ro.Use(middleware.CORS())
	ro.Use(middleware.Logger())
	ro.Use(middleware.Recover())
	ro.POST("/auth/login", r.User.Login)
	ro.POST("/auth/register", r.User.Register)
	/// Auth
	rauth := ro.Group("", middleware.JWT([]byte(r.Depend.Config.JwtSecret)), middleware.CSRFWithConfig(middleware.CSRFConfig{TokenLength: uint8(r.Depend.Config.CSRFLength)}))
	rauth.POST("/books", r.Book.InsertBook, r.CSRFMiddlewareCustom)           //, middleware.JWT([]byte(config.JWTSecret)))
	rauth.GET("/books", r.Book.GetAllBooks)                                   //, middleware.JWT([]byte(config.JWTSecret)))
	rauth.GET("/books/:id", r.Book.GetBookByBookID)                           //, middleware.JWT([]byte(config.JWTSecret)))
	rauth.PUT("/books/:id", r.Book.UpdateBook, r.CSRFMiddlewareCustom)        //, middleware.JWT([]byte(config.JWTSecret)))
	rauth.DELETE("/books/:id", r.Book.DeleteByBookID, r.CSRFMiddlewareCustom) //, middleware.JWT([]byte(config.JWTSecret)))
	/// User Area
	rauth.GET("/users", r.User.GetById)
	rauth.PUT("/users", r.User.Update, r.CSRFMiddlewareCustom)
	/// Transactions
	rauth.POST("/transaction", r.Trx.Createtrx)
}
