package routes

import (
	"errors"
	"log"

	"github.com/ChaotenHG/auth-server/auth"
	"github.com/ChaotenHG/auth-server/db"
	"github.com/ChaotenHG/auth-server/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Post_RefreshToken(c echo.Context) error {

	var (
		claims jwt.Claims
		err    error
		token  string = c.QueryParam("refresh_token")
	)

	if claims, err = auth.VerifyRefreshToken(token); err != nil {
		return errors.Join(err, c.String(401, "Not Valid"))
	}

	var (
		sub       string
		user      model.User
		tokenPair auth.TokenPair
	)

	if sub, err = claims.GetSubject(); err != nil {
		return err
	}

	if user, err = db.FindUserByID(sub); err != nil {
		return err
	}

	if tokenPair, err = auth.CreateTokenPair(user); err != nil {
		return err
	}

	if err = auth.RevokeRefreshToken(user, token); err != nil {
		log.Printf("Could not revoke refresh token: %v\n", err)
	}

	return c.JSON(200, tokenPair)
}
