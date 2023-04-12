package book

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"gorm.io/gorm"
)

type BookModel struct {
	db *gorm.DB
}

func New(d *gorm.DB) entities.Repository {
	return &BookModel{
		db: d,
	}
}

func (bm *BookModel) InsertBook(book entities.Core) (entities.Core, error) {
	var insertBook entities.Book
	insertBook.Title = book.Title
	insertBook.Year = book.Year
	insertBook.Author = book.Author
	insertBook.Contents = book.Contents
	insertBook.Image = []byte(book.Image)

	err := bm.db.Table("books").Create(&insertBook).Error
	if err != nil {
		log.Println("Terjadi error saat membuat daftar buku baru", err.Error())
		return entities.Core{}, err
	}
	return book, nil
}

func (bm *BookModel) GetAllBooks() ([]entities.Core, error) {
	var books []entities.Book
	if err := bm.db.Table("books").Where("deleted_at < ?", 0).Find(&books).Error; err != nil {
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

func (bm *BookModel) GetBookByBookID(bookID uint) (entities.Core, error) {
	var book entities.Book
	if err := bm.db.Table("books").Where("id = ? AND deleted_at < ?", bookID, 0).First(&book).Error; err != nil {
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

func (um *BookModel) UpdateByBookID(bookID uint, updatedBook entities.Book) error {
	book := entities.Book{}
	if bookID == 0 {
		return fmt.Errorf("Terjadi kesalahan input ID")
	}
	if err := um.db.First(&book, bookID).Error; err != nil {
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

	if err := um.db.Save(&book).Error; err != nil {
		log.Println("Terjadi error saat melakukan update daftar buku", err)
		return err
	}

	return nil
}

func (um *BookModel) DeleteByBookID(bookID uint) error {
	book := entities.Book{}
	if bookID == 0 {
		return fmt.Errorf("Terjadi kesalahan input ID")
	}
	if err := um.db.First(&book, bookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID buku %v tidak ditemukan", bookID)
		}
		log.Println("Terjadi error saat mengambil buku dengan ID", err)
		return err
	}

	book.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := um.db.Save(&book).Error; err != nil {
		log.Println("Terjadi error saat melakukan delete buku", err)
		return err
	}

	return nil
}
