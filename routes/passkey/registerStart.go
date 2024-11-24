package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ChaotenHG/auth-server/auth"
	"github.com/ChaotenHG/auth-server/db"
	"github.com/ChaotenHG/auth-server/model"
	"github.com/labstack/echo/v4"
	"github.com/nrednav/cuid2"
)

func Post_registerStartPasskey(c echo.Context) error {

	jwt := c.Request().Header.Get("Authorization")
	claims, err := auth.VerifyToken(jwt)
	if err != nil {
		return err
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return err
	}

	var (
		user model.User
	)

	if user, err = db.FindUser(sub); err != nil {
		return err
	}

	options, session, err := auth.WebAuthn.BeginRegistration(&user)
	if err != nil {
		return err
	}

	sid := cuid2.Generate()

	obj, err := json.Marshal(session)
	if err != nil {
		return err
	}

	if err = db.Rdb.Set(db.RedisContext, sid, string(obj), 10*time.Minute).Err(); err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "sid",
		Value:    sid,
		Path:     "/passkey/registerFinish",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode, // TODO: SameSiteStrictMode maybe?
	})

	return c.JSON(200, options)
}
