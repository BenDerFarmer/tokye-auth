package routes

import (
	"github.com/ChaotenHG/auth-server/auth"
	"github.com/ChaotenHG/auth-server/db"
	"github.com/ChaotenHG/auth-server/model"
	"github.com/ChaotenHG/auth-server/utils"
	"github.com/labstack/echo/v4"
)

func Post_registerMail(c echo.Context) error {

	var credentials auth.MailAuth

	if err := utils.BodyToObject(c, &credentials); err != nil {
		return err
	}

	if auth.VerifyOTP(credentials.Email, credentials.Code) != nil {
		return c.JSON(401, utils.MSG{Message: "invalid code or email is already in use"})
	}

	if err := db.CreateUser(credentials.Email); err != nil {

		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
			c.JSON(401, utils.MSG{Message: "invalid code or email is already in use"})
		}

		return err
	}

	var (
		TokenPair auth.TokenPair
		user      model.User
		err       error
	)

	if user, err = db.FindUser(credentials.Email); err != nil {
		return err
	}

	if TokenPair, err = auth.CreateTokenPair(user); err != nil {
		return err
	}

	return c.JSON(200, TokenPair)

}
