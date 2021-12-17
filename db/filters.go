package db

import (
	"strings"
	"time"

	"github.com/sokolovb/bookshelf/models"
)

type Filter func(book models.BookInternal) bool

func FilterByName(name string) Filter {
	return func(book models.BookInternal) bool {
		if name == "" {
			return true
		}
		return strings.EqualFold(book.Name, name)
	}
}

func FilterByPublishingHouse(publishingHouse string) Filter {
	return func(book models.BookInternal) bool {
		if publishingHouse == "" {
			return true
		}
		return strings.EqualFold(book.PublishingHouse, publishingHouse)
	}
}

func FilterByISBN(isbn string) Filter {
	return func(book models.BookInternal) bool {
		if isbn == "" {
			return true
		}
		return strings.EqualFold(book.ISBN, isbn)
	}
}

func FilterByDate(dateStr string) Filter {
	return func(book models.BookInternal) bool {
		if dateStr == "" {
			return true
		}

		date, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			return false
		}

		return book.UpdatedTime.After(date)
	}
}
