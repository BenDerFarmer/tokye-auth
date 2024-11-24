package auth

import (
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/ChaotenHG/auth-server/db"
	"github.com/redis/go-redis/v9"
)

type MailAuth struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func GenerateOTP(maxDigits uint32) string {
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)
	if err != nil {
		log.Panic(err)
	}
	return fmt.Sprintf("%0*d", maxDigits, bi)
}

func SaveOTP(email string, code string) error {
	return db.Rdb.Set(db.RedisContext, email, code, 10*time.Minute).Err()
}

func SaveTimer(ip string, delay time.Duration) error {
	return db.Rdb.Set(db.RedisContext, ip, []byte{0}, delay).Err()
}

func VerifyTimer(ip string) error {
	err := db.Rdb.Get(db.RedisContext, ip).Err()

	if err != nil {

		if err == redis.Nil {
			return nil
		}

		return err
	}

	return fmt.Errorf("request limit")
}

func VerifyOTP(email string, code string) error {

	value, err := db.Rdb.Get(db.RedisContext, email).Result()
	if err != nil {
		return err
	}

	if code != value {
		return fmt.Errorf("OTP invalid")
	}

	return nil
}
