package utils

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
)

func BodyToObject(c echo.Context, obj any) error {

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request().Body)
	}
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return json.Unmarshal(bodyBytes, obj)

}
