package transaction

import (
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"gorm.io/gorm"
)

type TrasanctionRepoInterface interface {
	Create(db *gorm.DB, req entities.Transaction) (int, error)
	FindMyTransaction(db *gorm.DB, uid int) ([]*entities.MyTransactionResponses, error)
	GetAllAvailableBooks(db *gorm.DB) ([]*entities.AvailableBookResponses, error)
	GetBorrowedBook(db *gorm.DB, uid int) ([]*entities.MyBookBorrowedResponses, error)
	InsertTrxBook(db *gorm.DB, req entities.TransactionBook) error
}
