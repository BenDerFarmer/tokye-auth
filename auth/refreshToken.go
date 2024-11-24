package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/ChaotenHG/auth-server/db"
	"github.com/ChaotenHG/auth-server/model"
	"github.com/golang-jwt/jwt/v5"
)

func CreateRefreshToken(user model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"sub": user.ID,
		"exp": time.Now().Add(5 * 24 * time.Hour).Unix(),
	})

	var (
		err         error
		tokenString string
	)

	// Sign the token with the private key
	if tokenString, err = token.SignedString(refreshToken); err != nil {
		return tokenString, errors.Join(errors.New("Failed to sign rf-token:"), err)
	}

	if err = db.Rdb.SAdd(db.RedisContext, user.ID, tokenString).Err(); err != nil {
		return tokenString, errors.Join(errors.New("Failed to save rf-token:"), err)
	}

	return tokenString, nil

}

func VerifyRefreshToken(tokenString string) (jwt.Claims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return refreshToken, nil
	})

	var (
		claims jwt.MapClaims
		ok     bool
		sub    string
	)

	if err != nil {
		return claims, fmt.Errorf("Failed to parse token: %v", err)
	}

	// Check token validity and claims
	if claims, ok = parsedToken.Claims.(jwt.MapClaims); !ok && !parsedToken.Valid {
		return claims, fmt.Errorf("Token is invalid or claims are incorrect.")
	}

	if sub, err = claims.GetSubject(); err != nil {
		return claims, err
	}

	if ok, err = db.Rdb.SIsMember(db.RedisContext, sub, parsedToken.Raw).Result(); err != nil || !ok {
		return claims, fmt.Errorf("RefreshToken not found")
	}

	return claims, err
}

func RevokeRefreshToken(user model.User, tokenString string) error {
	return db.Rdb.SRem(db.RedisContext, user.ID, tokenString).Err()
}
