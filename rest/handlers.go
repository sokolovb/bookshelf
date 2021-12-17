package rest

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"github.com/sokolovb/bookshelf/db"
	"github.com/sokolovb/bookshelf/models"
)

func (s *Server) addNewBooks(c echo.Context) error {
	var books []models.Book
	if err := c.Bind(&books); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := models.Validate(books); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := s.storage.AddBooks(books); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (s *Server) listBooks(c echo.Context) error {
	var (
		name            = c.QueryParam("name")
		publishingHouse = c.QueryParam("publishing_house")
		isbn            = c.QueryParam("isbn")
		dateStr         = c.QueryParam("date")
	)

	if err := validateDate(dateStr); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	books, err := s.storage.ListBooks(
		db.FilterByName(name),
		db.FilterByPublishingHouse(publishingHouse),
		db.FilterByISBN(isbn),
		db.FilterByDate(dateStr),
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

func validateDate(dateStr string) error {
	if dateStr == "" {
		return nil
	}
	if _, err := time.Parse(time.RFC3339, dateStr); err != nil {
		return err
	}
	return nil
}
