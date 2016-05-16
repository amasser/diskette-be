package user

import (
	"errors"
	"github.com/getdiskette/diskette/tokens"
	"github.com/getdiskette/diskette/util"
	"net/http"
	"time"

	"github.com/satori/go.uuid"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// http POST localhost:5025/user/forgot-password email=joe.doe@gmail.com
func (service *serviceImpl) ForgotPassword(c echo.Context) error {
	var request struct {
		Email string `json:"email"`
	}
	c.Bind(&request)

	if request.Email == "" {
		return c.JSON(http.StatusBadRequest, util.CreateErrResponse(errors.New("Missing parameter 'email'")))
	}

	resetKey := uuid.NewV4().String()

	err := service.userCollection.Update(
		bson.M{"email": request.Email},
		bson.M{
			"$set": bson.M{
				"resetKey":         resetKey,
				"requestedResetAt": time.Now(),
			},
		},
	)
	if err != nil {
		return c.JSON(http.StatusNotFound, util.CreateErrResponse(errors.New("The user doesn't exist.")))
	}

	token := tokens.ResetToken{Key: resetKey}
	tokenStr, err := token.ToString(service.jwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.CreateErrResponse(err))
	}

	return c.JSON(http.StatusOK, util.CreateOkResponse(bson.M{"ResetToken": tokenStr}))
}
