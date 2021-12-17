package models

import (
	"errors"
	"time"
)

type (
	Book struct {
		Name            string `json:"name"`
		ISBN            string `json:"isbn"`
		Authors         string `json:"authors"`
		Genre           string `json:"genre"`
		PublishingHouse string `json:"publishing_house"`
		Year            string `json:"year"`
		PagesCount      int    `json:"pages_count"`
		Updated         string `json:"updated"`
	}

	BookInternal struct {
		Book
		UpdatedTime time.Time
	}
)

func Validate(books []Book) error {
	for _, book := range books {
		if book.ISBN == "" {
			return errors.New("ISBN cannot be empty")
		}
	}
	return nil
}
