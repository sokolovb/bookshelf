package rest

import (
	"github.com/labstack/echo"

	"github.com/sokolovb/bookshelf/db"
)

const (
	prefixV1 = "/bookshelf/v1"
)

type (
	Server struct {
		port    string
		storage db.Storage
	}
)

func NewServer(port string) Server {
	return Server{
		port:    port,
		storage: db.NewStorage(),
	}
}

func (s Server) Start() {
	e := echo.New()

	v1 := e.Group(prefixV1)
	v1.GET("/books", s.listBooks)
	v1.PUT("/books", s.addNewBooks)

	e.Logger.Fatal(e.Start(":" + s.port))
}
