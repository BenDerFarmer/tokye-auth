package routes

import (
	"crypto/sha256"
	"time"

	"github.com/ChaotenHG/auth-server/auth"
	"github.com/ChaotenHG/auth-server/mail"
	"github.com/ChaotenHG/auth-server/utils"
	"github.com/labstack/echo/v4"
)

func Post_otp(c echo.Context) error {

	ip := sha256.New()
	ip.Write([]byte(c.RealIP()))

	var err error

	if err = auth.VerifyTimer(string(ip.Sum(nil))); err != nil {
		return err
	}

	var credentials auth.MailAuth

	if err = utils.BodyToObject(c, &credentials); err != nil {
		return err
	}

	code := auth.GenerateOTP(6)

	if err = auth.SaveOTP(credentials.Email, code); err != nil {
		return err
	}

	if err = mail.SendOTPMail(credentials.Email, code); err != nil {
		return err
	}

	delay := 25 * time.Second

	if err = auth.SaveTimer(string(ip.Sum(nil)), delay); err != nil {
		return err
	}

	type Response struct {
		Message string        `json:"message"`
		Time    time.Duration `json:"time"`
	}

	return c.JSON(200, Response{Message: "Email to " + credentials.Email, Time: delay})

}
