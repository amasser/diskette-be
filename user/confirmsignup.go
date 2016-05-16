package user

import (
	"errors"
	"github.com/getdiskette/diskette/util"
	"net/http"

	"github.com/getdiskette/diskette/tokens"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// http POST localhost:5025/user/confirm token=<confirmation_token>
func (service *serviceImpl) ConfirmSignup(c echo.Context) error {
	var request struct {
		Token string `json:"token"`
	}
	c.Bind(&request)

	if request.Token == "" {
		return c.JSON(http.StatusBadRequest, util.CreateErrResponse(errors.New("Missing parameter 'token'")))
	}

	token, err := tokens.ParseConfirmationToken(service.jwtKey, request.Token)
	if err != nil || token.Key == "" {
		return c.JSON(http.StatusForbidden, util.CreateErrResponse(err))
	}

	err = service.userCollection.Update(
		bson.M{"confirmationKey": token.Key},
		bson.M{
			"$set": bson.M{
				"isConfirmed": true,
			},
		},
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.CreateErrResponse(err))
	}

	return c.JSON(http.StatusOK, util.CreateOkResponse(nil))
}
