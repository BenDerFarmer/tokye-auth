package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ChaotenHG/auth-server/auth"
	"github.com/ChaotenHG/auth-server/db"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/labstack/echo/v4"
)

func Post_registerFinishPasskey(c echo.Context) error {

	sid, err := c.Cookie("sid")
	if err != nil {
		return err
	}

	sessionData, err := db.Rdb.Get(db.RedisContext, sid.Value).Result()
	if err != nil {
		return err
	}

	var session webauthn.SessionData

	json.Unmarshal([]byte(sessionData), &session)

	user, err := db.FindUserByID(string(session.UserID))
	if err != nil {
		return err
	}

	credential, err := auth.WebAuthn.FinishRegistration(&user, session, c.Request())
	if err != nil {

		c.SetCookie(&http.Cookie{
			Name:     "sid",
			Value:    "",
			Path:     "/passkey/registerFinish",
			MaxAge:   3600,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode, // TODO: SameSiteStrictMode maybe?
		})

		return err
	}
	user.AddCredential(credential)

	if err = db.SaveUser(&user); err != nil {
		return err
	}

	var tokenPair auth.TokenPair
	if tokenPair, err = auth.CreateTokenPair(user); err != nil {
		return err
	}

	return c.JSON(200, tokenPair)
}
