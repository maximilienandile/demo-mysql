package storage

import (
	"context"
	"github.com/maximilienandile/demo-mysql/internal/book"
)

type Storage interface {
	// TODO: why returning a book, we should return only an error
	Create(ctx context.Context, b book.Book) (book.Book, error)
}
