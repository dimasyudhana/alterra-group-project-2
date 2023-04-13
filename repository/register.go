package repository

import (
	"github.com/dimasyudhana/alterra-group-project-2/repository/book"
	"github.com/dimasyudhana/alterra-group-project-2/repository/transaction"
	"github.com/dimasyudhana/alterra-group-project-2/repository/user"
	"go.uber.org/dig"
)

func Register(C *dig.Container) error {
	if err := C.Provide(user.NewUserRepo); err != nil {
		return err
	}
	if err := C.Provide(book.New); err != nil {
		return err
	}
	if err := C.Provide(transaction.Newtransaction); err != nil {
		return err
	}
	return nil
}
