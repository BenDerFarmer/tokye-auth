package routes

import (
	"errors"

	"github.com/ChaotenHG/auth-server/auth"
	"github.com/ChaotenHG/auth-server/db"
	"github.com/ChaotenHG/auth-server/model"
	"github.com/ChaotenHG/auth-server/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Post_Logout(c echo.Context) error {

	var (
		claims jwt.Claims
		err    error
		token  string = c.QueryParam("refresh_token")
	)

	if claims, err = auth.VerifyRefreshToken(token); err != nil {
		return errors.Join(err, c.JSON(401, utils.MSG{Message: "Not Valid"}))
	}

	var (
		sub  string
		user model.User
	)

	if sub, err = claims.GetSubject(); err != nil {
		return err
	}

	if user, err = db.FindUserByID(sub); err != nil {
		return err
	}

	if err = auth.RevokeRefreshToken(user, token); err != nil {
		return err
	}

	return c.JSON(200, utils.MSG{Message: "Successful"})
}
