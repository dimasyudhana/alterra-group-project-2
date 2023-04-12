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
}
