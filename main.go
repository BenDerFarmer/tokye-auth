package main

import (
	"crypto/sha256"
	"log"
	"net/http"
	"time"

	"github.com/ChaotenHG/auth-server/auth"
	"github.com/ChaotenHG/auth-server/config"
	"github.com/ChaotenHG/auth-server/db"
	"github.com/ChaotenHG/auth-server/mail"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	cfg := config.New()
	e := echo.New()

	db.LoadSQLCredentials(cfg)
	db.InitialMigration()
	db.LoadRedisClient(cfg)
	auth.LoadKeys(cfg)
	auth.LoadPasskeyConfig(cfg)
	mail.LoadConfig(cfg)

	// Middleware
	if cfg.DebugMode {
		e.Use(middleware.Logger())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:                             []string{"*"},
			AllowCredentials:                         true,
			UnsafeWildcardOriginWithAllowCredentials: true,
		}))
	} else {
		e.Use(middleware.Recover())
		e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}))

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     cfg.CORSOrigns,
			AllowCredentials: true,
		}))

		e.Use(middleware.BodyLimit("1M"))

		config := middleware.RateLimiterConfig{
			Skipper: middleware.DefaultSkipper,
			Store: middleware.NewRateLimiterMemoryStoreWithConfig(
				middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(10), Burst: 30, ExpiresIn: 3 * time.Minute},
			),
			IdentifierExtractor: func(ctx echo.Context) (string, error) {

				ip := sha256.New()
				ip.Write([]byte(ctx.RealIP()))

				return string(ip.Sum(nil)), nil
			},
			ErrorHandler: func(context echo.Context, err error) error {
				return context.JSON(http.StatusForbidden, nil)
			},
			DenyHandler: func(context echo.Context, identifier string, err error) error {
				return context.JSON(http.StatusTooManyRequests, nil)
			},
		}

		e.Use(middleware.RateLimiterWithConfig(config))
	}

	registerRoutes(e)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
