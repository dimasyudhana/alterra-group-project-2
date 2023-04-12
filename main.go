package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dimasyudhana/alterra-group-project-2/config"
	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/config/dependcy/container"
	"github.com/dimasyudhana/alterra-group-project-2/routes"
)

func main() {

	container.RunAll()
	err := container.Container.Invoke(func(depend dependecy.Depend, ro routes.Routes) {
		config.Migrate(depend.Config)
		var sig = make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		ro.RegisterRoutes()
		go func() {
			depend.Log.Infof("Starting server on port %s", depend.Config.Server.Port)
			if err := depend.Echo.Start(fmt.Sprintf(":%s", depend.Config.Server.Port)); err != nil {
				depend.Log.Errorf("Failed to start server: %v", err)
				sig <- syscall.SIGTERM
			}
		}()
		<-sig
		depend.Log.Info("Shutting down server")
	})
	if err != nil {
		log.Print(err)
	}

	"github.com/dimasyudhana/alterra-group-project-2/config"
	bookHandler "github.com/dimasyudhana/alterra-group-project-2/controller"
	bookRepo "github.com/dimasyudhana/alterra-group-project-2/repository/book"
	"github.com/dimasyudhana/alterra-group-project-2/routes"
	bookLogic "github.com/dimasyudhana/alterra-group-project-2/service/book"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.InitConfiguration()
	db, _ := config.GetConnection(cfg)
	config.Migrate(db)

	bookModel := bookRepo.New(db)
	bookService := bookLogic.New(bookModel)
	bookController := bookHandler.New(bookService)

	routes.Route(e, bookController)

	e.Start(":8080")
}
