package routes

import (
	"github.com/ChaotenHG/auth-server/auth"
	"github.com/labstack/echo/v4"
)

func Get_publicKey(c echo.Context) error {
	return c.String(200, auth.PublicKeyString)
}
