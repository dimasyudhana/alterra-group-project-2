package entities

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type (
	User struct {
		Id           uint          `json:"-" gorm:"type:int;primaryKey;autoIncrement"`
		Name         string        `json:"name" gorm:"type:varchar(50);not null"`
		Email        string        `json:"email" gorm:"type:varchar(50);not null"`
		Password     string        `json:"password" gorm:"type:varchar(80);not null"`
		Image        string        `json:"image" gorm:"type:varchar(30);not null"`
		Address      string        `json:"address" gorm:"type:varchar(255);not null"`
		Books        []Book        `json:"-"`
		Transactions []Transaction `json:"-" gorm:"foreignKey:BorrowerId;references:Id"`
	}
	UserReqLogin struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	UserReqRegister struct {
		Name     string `form:"name" validate:"required"`
		Email    string `form:"email" validate:"required"`
		Password string `form:"password" validate:"required"`
		Image    string
		Address  string `form:"address" validate:"required"`
	}
	UserReqUpdate struct {
		Name     string `form:"name" validate:"required"`
		Email    string `form:"email" validate:"required"`
		Password string `form:"password" validate:"required"`
		Image    string
		Address  string `form:"address" validate:"required"`
	}
	Transaction struct {
		Id               int            `gorm:"type:int;primaryKey;autoIncrement"`
		EndDate          string         `gorm:"type:timestamp;not null"`
		SubmitedDate     mysql.NullTime ` gorm:"type:timestamp;default:null"`
		BorrowerId       int
		TransactionBooks []TransactionBook
	}
	TransactionBook struct {
		TransactionId int `gorm:"not null"`
		BookId        int `gorm:"not null"`
	}
	MyTransactionResponses struct {
		Id              int       `json:"id"`
		EndDate         time.Time `json:"end_date"`
		BookOwner       string    `json:"book_owner"`
		BookName        string    `json:"book_name"`
		BookDescription string    `json:"book_description"`
		BookImage       string    `json:"book_image"`
	}
	MyBookBorrowedResponses struct {
		Id           int       `json:"id"`
		EndDate      time.Time `json:"end_date"`
		BorrowerName string    `json:"borrower_name"`
		BookName     string    `json:"book_name"`
		BookImage    string    `json:"book_image"`
	}
	AvailableBookResponses struct {
		Id           int    `json:"id"`
		BookOwner    string `json:"book_owner"`
		BookName     string `json:"book_name"`
		BookContents string `json:"contents"`
		BookImage    string `json:"book_image"`
	}
	WebResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}
)
