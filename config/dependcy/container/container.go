package container

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	"github.com/dimasyudhana/alterra-group-project-2/config"
	"github.com/dimasyudhana/alterra-group-project-2/pkg"
	"github.com/dimasyudhana/alterra-group-project-2/repository"
	"github.com/dimasyudhana/alterra-group-project-2/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

var (
	Container = dig.New()
)

func RunAll() {
	Container := Container
	if err := Container.Provide(config.InitConfiguration); err != nil {
		panic(err)
	}
	if err := Container.Provide(config.GetConnection); err != nil {
		panic(err)
	}
	if err := Container.Provide(echo.New); err != nil {
		panic(err)
	}
	if err := Container.Provide(NewStorage); err != nil {
		panic(err)
	}
	if err := repository.Register(Container); err != nil {
		panic(err)
	}
	if err := service.Register(Container); err != nil {
		panic(err)
	}
	if err := Container.Provide(NewLog); err != nil {
		panic(err)
	}

}

func NewStorage(cfg *config.Config) (*pkg.StorageGCP, error) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", cfg.GCP.Credential)
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}
	return &pkg.StorageGCP{
		ClG:        client,
		ProjectID:  cfg.GCP.PRJID,
		BucketName: cfg.GCP.BCKNM,
		Path:       cfg.GCP.Path,
	}, nil
}

func NewLog() (*log.Logger, error) {
	var logger = log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	return logger, nil
}
