package transaction

import (
	"context"

	"github.com/dimasyudhana/alterra-group-project-2/entities"
)

type TrxServiceInterface interface {
	Create(ctx context.Context, reqs []int, uid int) error
	FindMyTransaction(ctx context.Context, uid int) ([]*entities.MyTransactionResponses, error)
	GetAllAvailableBooks(ctx context.Context) ([]*entities.AvailableBookResponses, error)
	GetAllBorrowedBooks(ctx context.Context, uid int) ([]*entities.MyBookBorrowedResponses, error)
}
