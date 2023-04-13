package book

import (
	"context"

	"github.com/dimasyudhana/alterra-group-project-2/entities"
)

type Service interface {
	InsertBook(ctx context.Context, book entities.Core) error
	GetAllBooks(ctx context.Context) ([]entities.Core, error)
	GetBookByBookID(ctx context.Context, bookID uint) (entities.Core, error)
	UpdateByBookID(ctx context.Context, bookID uint, updatedBook entities.Book) error
	DeleteByBookID(ctx context.Context, bookID uint) error
}
