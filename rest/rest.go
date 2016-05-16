package rest

import (
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
)

type Service interface {
	Get(c echo.Context) error
	Post(c echo.Context) error
	Put(c echo.Context) error
	Delete(c echo.Context) error
}

type serviceImpl struct {
	db *mgo.Database
}

func NewService(db *mgo.Database) Service {
	return &serviceImpl{db}
}
