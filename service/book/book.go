package book

import (
	"errors"
	"fmt"

	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type BookModel struct {
	repository entities.Repository
}

func New(br entities.Repository) entities.Service {
	return &BookModel{
		repository: br,
	}
}

func (bm *BookModel) InsertBook(book entities.Core) (entities.Core, error) {
	result, err := bm.repository.InsertBook(book)
	if err != nil {
		log.Errorf("terjadi kesalahan input buku: %v", err)
		return entities.Core{}, errors.New("terdapat masalah pada server")
	}
	return result, nil
}

func (bm *BookModel) GetAllBooks() ([]entities.Core, error) {
	books, err := bm.repository.GetAllBooks()
	if err != nil {
		log.Errorf("terjadi kesalahan saat mengambil data buku: %v", err)
		return []entities.Core{}, errors.New("terdapat masalah pada server")
	}
	return books, nil
}

func (bm *BookModel) GetBookByBookID(bookID uint) (entities.Core, error) {
	book, err := bm.repository.GetBookByBookID(bookID)
	if err != nil {
		log.Errorf("terjadi kesalahan saat mengambil data buku dengan ID %d: %v", bookID, err)
		return entities.Core{}, errors.New("terdapat masalah pada server")
	}
	return book, nil
}

func (bm *BookModel) UpdateByBookID(bookID uint, updatedBook entities.Book) error {
	book, err := bm.repository.GetBookByBookID(bookID)
	if err != nil {
		log.Errorf("terjadi kesalahan saat mengambil data buku dengan ID %d: %v", bookID, err)
		return errors.New("terdapat masalah pada server")
	}

	book.Title = updatedBook.Title
	book.Year = updatedBook.Year
	book.Author = updatedBook.Author
	book.Contents = updatedBook.Contents
	book.Image = string(updatedBook.Image)

	if err := bm.repository.UpdateByBookID(bookID, updatedBook); err != nil {
		log.Errorf("terjadi kesalahan saat update data buku dengan ID %d: %v", bookID, err)
		return errors.New("terdapat masalah pada server")
	}

	return nil
}

func (bm *BookModel) DeleteByBookID(bookID uint) error {
	if bookID == 0 {
		return fmt.Errorf("ID buku tidak valid")
	}
	err := bm.repository.DeleteByBookID(bookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("buku dengan ID %v tidak ditemukan", bookID)
		}
		log.Errorf("terjadi kesalahan saat menghapus data buku dengan ID %d: %v", bookID, err)
		return errors.New("terdapat masalah pada server")
	}

	return nil
}
