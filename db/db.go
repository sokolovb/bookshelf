package db

import (
	"sync"
	"time"

	"github.com/sokolovb/bookshelf/models"
)

type (
	Storage interface {
		AddBooks(books []models.Book) error
		ListBooks(...Filter) ([]models.Book, error)
	}

	storage struct {
		*sync.Mutex
		books map[string]models.BookInternal
	}
)

func NewStorage() Storage {
	return &storage{
		&sync.Mutex{},
		map[string]models.BookInternal{},
	}
}

func (s *storage) AddBooks(books []models.Book) error {
	s.Lock()
	defer s.Unlock()

	for _, book := range books {
		updatedTime, err := time.Parse(time.RFC3339, book.Updated)
		if err != nil {
			return err
		}
		s.books[book.ISBN] = models.BookInternal{
			Book:        book,
			UpdatedTime: updatedTime,
		}
	}
	return nil
}

func (s *storage) ListBooks(filters ...Filter) ([]models.Book, error) {
	s.Lock()
	defer s.Unlock()

	var ret []models.Book
	for _, book := range s.books {
		suits := true
		for _, filter := range filters {
			if !filter(book) {
				suits = false
				break
			}
		}
		if suits {
			ret = append(ret, book.Book)
		}
	}
	return ret, nil
}
