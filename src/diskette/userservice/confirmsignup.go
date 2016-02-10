package userservice

import (
	"diskette/util"
	"errors"
	"net/http"
	"time"

	"diskette/tokens"

	"github.com/labstack/echo"
	"labix.org/v2/mgo/bson"
)

// http POST localhost:5025/public/confirm token=<confirmation_token>
func (self impl) ConfirmSignup(c *echo.Context) error {
	var request struct {
		Token string `json:"token"`
	}
	c.Bind(&request)

	if request.Token == "" {
		return c.JSON(http.StatusBadRequest, util.CreateErrResponse(errors.New("Missing parameter 'token'")))
	}

	token, err := tokens.ParseConfirmationToken(self.jwtKey, request.Token)
	if err != nil || token.Key == "" {
		return c.JSON(http.StatusForbidden, util.CreateErrResponse(err))
	}

	return self.userCollection.Update(
		bson.M{"confirmationKey": token.Key},
		bson.M{
			"$set": bson.M{
				"confirmedAt": time.Now(),
			},
		},
	)
}
