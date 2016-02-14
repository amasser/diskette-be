package admin

import (
	"github.com/labstack/echo"
	"labix.org/v2/mgo"
)

type Service interface {
	GetUsers(c *echo.Context) error
	CreateUser(c *echo.Context) error
	ChangeUserPassword(c *echo.Context) error
	ChangeUserEmail(c *echo.Context) error
	ChangeUserRoles(c *echo.Context) error
}

type serviceImpl struct {
	userCollection *mgo.Collection
	jwtKey         []byte
}

func NewService(userCollection *mgo.Collection, jwtKey []byte) Service {
	return &serviceImpl{userCollection, jwtKey}
}
