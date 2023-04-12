package dependecy

import (
	"github.com/dimasyudhana/alterra-group-project-2/config"
	"github.com/dimasyudhana/alterra-group-project-2/pkg"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type Depend struct {
	dig.In
	Db     *gorm.DB
	Config *config.Config
	Echo   *echo.Echo
	Log    *logrus.Logger
	Gcp    *pkg.StorageGCP
}
