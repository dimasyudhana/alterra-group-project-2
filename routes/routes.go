package routes

import (
	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/controller"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
)

type Routes struct {
	dig.In
	Depend     dependecy.Depend
	Controller controller.User
}

func (r *Routes) RegisterRoutes() {
	ro := r.Depend.Echo
	ro.Use(middleware.RemoveTrailingSlash())
	ro.Use(middleware.CORS())
	ro.POST("/auth/login", r.Controller.Login)
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
}

func Route(e *echo.Echo, bc entities.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/books", bc.InsertBook())           //, middleware.JWT([]byte(config.JWTSecret)))
	e.GET("/books", bc.GetAllBooks())           //, middleware.JWT([]byte(config.JWTSecret)))
	e.GET("/books/:id", bc.GetBookByBookID())   //, middleware.JWT([]byte(config.JWTSecret)))
	e.PUT("/books/:id", bc.UpdateByBookID())    //, middleware.JWT([]byte(config.JWTSecret)))
	e.DELETE("/books/:id", bc.DeleteByBookID()) //, middleware.JWT([]byte(config.JWTSecret)))
}
