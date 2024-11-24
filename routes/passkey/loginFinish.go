package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChaotenHG/auth-server/auth"
	"github.com/ChaotenHG/auth-server/db"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/labstack/echo/v4"
)

func Post_loginFinishPasskey(c echo.Context) error {

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

	credential, err := auth.WebAuthn.FinishLogin(&user, session, c.Request())
	if err != nil {
		return err
	}

	if credential.Authenticator.CloneWarning {
		return fmt.Errorf("[WARN] can't finish login: %s", "CloneWarning")
	}

	user.UpdateCredential(credential)

	if err = db.SaveUser(&user); err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "sid",
		Value:    "",
		Path:     "/passkey/loginFinish",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode, // TODO: SameSiteStrictMode maybe?
	})

	var tokenPair auth.TokenPair
	if tokenPair, err = auth.CreateTokenPair(user); err != nil {
		return err
	}

	return c.JSON(200, tokenPair)
}
