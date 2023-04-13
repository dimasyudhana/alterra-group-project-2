package book

import (
	"context"
	"errors"
	"fmt"

	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"github.com/dimasyudhana/alterra-group-project-2/repository/book"
	"github.com/go-playground/validator"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type BookModel struct {
	repo      book.Repository
	dep       dependecy.Depend
	validator *validator.Validate
}

func New(repo book.Repository, dep dependecy.Depend) Service {
	return &BookModel{
		repo:      repo,
		dep:       dep,
		validator: validator.New(),
	}
}

func (bm *BookModel) InsertBook(ctx context.Context, book entities.Core) error {
	_, err := bm.repo.InsertBook(bm.dep.Db.WithContext(ctx), book)
	if err != nil {
		log.Errorf("terjadi kesalahan input buku: %v", err)
		return errors.New("terdapat masalah pada server")
	}
	return nil
}

func (bm *BookModel) GetAllBooks(ctx context.Context) ([]entities.Core, error) {
	books, err := bm.repo.GetAllBooks(bm.dep.Db.WithContext(ctx))
	if err != nil {
		log.Errorf("terjadi kesalahan saat mengambil data buku: %v", err)
		return []entities.Core{}, errors.New("terdapat masalah pada server")
	}
	return books, nil
}

func (bm *BookModel) GetBookByBookID(ctx context.Context, bookID uint) (entities.Core, error) {
	book, err := bm.repo.GetBookByBookID(bm.dep.Db.WithContext(ctx), bookID)
	if err != nil {
		log.Errorf("terjadi kesalahan saat mengambil data buku dengan ID %d: %v", bookID, err)
		return book, errors.New("terdapat masalah pada server")
	}
	return book, nil
}

func (bm *BookModel) UpdateByBookID(ctx context.Context, bookID uint, updatedBook entities.Book) error {
	book, err := bm.repo.GetBookByBookID(bm.dep.Db.WithContext(ctx), bookID)
	if err != nil {
		log.Errorf("terjadi kesalahan saat mengambil data buku dengan ID %d: %v", bookID, err)
		return errors.New("terdapat masalah pada server")
	}

	book.Title = updatedBook.Title
	book.Year = updatedBook.Year
	book.Author = updatedBook.Author
	book.Contents = updatedBook.Contents
	book.Image = string(updatedBook.Image)

	if err := bm.repo.UpdateByBookID(bm.dep.Db.WithContext(ctx), bookID, updatedBook); err != nil {
		log.Errorf("terjadi kesalahan saat update data buku dengan ID %d: %v", bookID, err)
		return errors.New("terdapat masalah pada server")
	}

	return nil
}

func (bm *BookModel) DeleteByBookID(ctx context.Context, bookID uint) error {
	if bookID == 0 {
		return fmt.Errorf("ID buku tidak valid")
	}
	err := bm.repo.DeleteByBookID(bm.dep.Db.WithContext(ctx), bookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("buku dengan ID %v tidak ditemukan", bookID)
		}
		log.Errorf("terjadi kesalahan saat menghapus data buku dengan ID %d: %v", bookID, err)
		return errors.New("terdapat masalah pada server")
	}

	return nil
}
