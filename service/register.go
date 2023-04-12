package service

import (
	"github.com/dimasyudhana/alterra-group-project-2/service/user"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	if err := c.Provide(user.NewUserService); err != nil {
		return err
	}
	return nil
}
