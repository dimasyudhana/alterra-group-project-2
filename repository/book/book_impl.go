package book

import (
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"gorm.io/gorm"
)

type Repository interface {
	InsertBook(db *gorm.DB, book entities.Core) (entities.Core, error)
	GetAllBooks(db *gorm.DB) ([]entities.Core, error)
	GetBookByBookID(db *gorm.DB, bookID uint) (entities.Core, error)
	UpdateByBookID(db *gorm.DB, bookID uint, updatedBook entities.Book) error
	DeleteByBookID(db *gorm.DB, bookID uint) error
}
