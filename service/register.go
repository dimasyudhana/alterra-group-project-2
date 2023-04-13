package service

import (
	"github.com/dimasyudhana/alterra-group-project-2/service/book"
	"github.com/dimasyudhana/alterra-group-project-2/service/transaction"
	"github.com/dimasyudhana/alterra-group-project-2/service/user"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	if err := c.Provide(user.NewUserService); err != nil {
		return err
	}
	if err := c.Provide(book.New); err != nil {
		return err
	}
	if err := c.Provide(transaction.NewTrxService); err != nil {
		return err
	}
	return nil
}
