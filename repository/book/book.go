package book

import (
	"errors"
	"fmt"
	"log"
	"time"

	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"gorm.io/gorm"
)

type BookModel struct {
	dep dependecy.Depend
}

func New() Repository {
	return &BookModel{}
}

func (bm *BookModel) InsertBook(db *gorm.DB, book entities.Core) (entities.Core, error) {
	var insertBook entities.Book
	insertBook.Title = book.Title
	insertBook.Year = book.Year
	insertBook.Author = book.Author
	insertBook.Contents = book.Contents
	insertBook.Image = []byte(book.Image)

	err := db.Table("books").Create(&insertBook).Error
	if err != nil {
		log.Println("Terjadi error saat membuat daftar buku baru", err.Error())
		return entities.Core{}, err
	}
	return book, nil
}

func (bm *BookModel) GetAllBooks(db *gorm.DB) ([]entities.Core, error) {
	var books []entities.Book
	if err := db.Table("books").Where("deleted_at < ?", 0).Find(&books).Error; err != nil {
		log.Println("Terjadi error saat mengambil daftar buku", err.Error())
		return nil, err
	}

	var cores []entities.Core
	for _, book := range books {
		core := entities.Core{
			ID:       book.ID,
			Title:    book.Title,
			Year:     book.Year,
			Author:   book.Author,
			Contents: book.Contents,
			Image:    string(book.Image),
		}
		cores = append(cores, core)
	}
	return cores, nil
}

func (bm *BookModel) GetBookByBookID(db *gorm.DB, bookID uint) (entities.Core, error) {
	var book entities.Book
	if err := db.Table("books").Where("id = ? AND deleted_at < ?", bookID, 0).First(&book).Error; err != nil {
		log.Println("Terjadi error saat mengambil buku dengan ID", bookID, err.Error())
		return entities.Core{}, err
	}

	core := entities.Core{
		ID:       book.ID,
		Title:    book.Title,
		Year:     book.Year,
		Author:   book.Author,
		Contents: book.Contents,
		Image:    string(book.Image),
	}

	return core, nil
}

func (um *BookModel) UpdateByBookID(db *gorm.DB, bookID uint, updatedBook entities.Book) error {
	book := entities.Book{}
	if bookID == 0 {
		return fmt.Errorf("Terjadi kesalahan input ID")
	}
	if err := db.First(&book, bookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID buku %v tidak ditemukan", bookID)
		}
		log.Println("Terjadi error saat mengambil buku dengan ID", err)
		return err
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Year = updatedBook.Year
	book.Contents = updatedBook.Contents
	book.Image = updatedBook.Image
	book.UpdatedAt = time.Now()

	if err := db.Save(&book).Error; err != nil {
		log.Println("Terjadi error saat melakukan update daftar buku", err)
		return err
	}

	return nil
}

func (um *BookModel) DeleteByBookID(db *gorm.DB, bookID uint) error {
	book := entities.Book{}
	if bookID == 0 {
		return fmt.Errorf("Terjadi kesalahan input ID")
	}
	if err := db.First(&book, bookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID buku %v tidak ditemukan", bookID)
		}
		log.Println("Terjadi error saat mengambil buku dengan ID", err)
		return err
	}

	book.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := db.Save(&book).Error; err != nil {
		log.Println("Terjadi error saat melakukan delete buku", err)
		return err
	}

	return nil
}
