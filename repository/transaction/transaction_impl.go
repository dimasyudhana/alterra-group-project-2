package transaction

import (
	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"gorm.io/gorm"
)

type transaction struct {
	dep dependecy.Depend
}

func Newtransaction() TrasanctionRepoInterface {
	return &transaction{}
}

func (t *transaction) Create(db *gorm.DB, req entities.Transaction) (int, error) {

	row := db.Create(&req)
	if err := row.Error; err != nil {
		return 0, err
	}
	return req.Id, nil
}

func (t *transaction) InsertTrxBook(db *gorm.DB, req entities.TransactionBook) error {
	if err := db.Create(&req).Error; err != nil {
		t.dep.Log.Errorf("error DB %v : ", err)
		return err
	}
	return nil
}

func (t *transaction) FindMyTransaction(db *gorm.DB, uid int) ([]*entities.MyTransactionResponses, error) {
	var res []*entities.MyTransactionResponses
	rows := db.Model(&entities.Transaction{}).Select(`transactions.id
	,transactions.end_date, users.name, books.title, books.contents, books.image`).Where("transactions.borrower_id = ? && transactions.submited_date IS NULL", uid).Joins(`
	JOIN transaction_books on transaction_books.transaction_id = transactions.id 
	JOIN books on books.id=transaction_books.book_id JOIN users on users.id = books.user_id`).Scan(&res)
	if err := rows.Error; err != nil {
		t.dep.Log.Errorf("Error Database: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *transaction) GetAllAvailableBooks(db *gorm.DB) ([]*entities.AvailableBookResponses, error) {
	var res []*entities.AvailableBookResponses
	rows := db.Model(&entities.Book{}).Select(`books.id, users.name, books.name , books.contents,books.image`).Where("transactions.submited_date IS NOT NULL").Joins(`
	JOIN users on users.id = books.user_id 
	`).Scan(&res)
	if err := rows.Error; err != nil {
		t.dep.Log.Errorf("Error Database: %v", err)
		return nil, err
	}
	return res, nil
}

func (t *transaction) GetBorrowedBook(db *gorm.DB, uid int) ([]*entities.MyBookBorrowedResponses, error) {
	var res []*entities.MyBookBorrowedResponses
	rows := db.Model(&entities.Transaction{}).Where("transaction.submited_date IS NULL AND books.users_id = ?").Select(`books.id,transactions.end_date,users.name,books.image`).Joins(`
	JOIN transaction_books on transaction_books.transaction_id = transaction.id 
	JOIN books on books.id = transaction_books.book_id
	`).Scan(&res)
	if err := rows.Error; err != nil {
		t.dep.Log.Errorf("Error Database: %v", err)
		return nil, err
	}
	return res, nil
}
