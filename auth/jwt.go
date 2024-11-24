package auth

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ChaotenHG/auth-server/config"
	"github.com/ChaotenHG/auth-server/model"
	"github.com/golang-jwt/jwt/v5"

	"os"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func CreateTokenPair(user model.User) (TokenPair, error) {

	var (
		token   string
		rfToken string
		err     error
	)

	if token, err = CreateUserToken(user); err != nil {
		return TokenPair{}, err
	}

	if rfToken, err = CreateRefreshToken(user); err != nil {
		return TokenPair{}, err
	}

	return TokenPair{AccessToken: token, RefreshToken: rfToken}, err
}

func CreateUserToken(user model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
		"userid": user.ID,
		"sub":    user.Email,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(30 * time.Minute).Unix(),
	})

	// Sign the token with the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		errors.Join(errors.New("Failed to sign token:"), err)
	}

	return tokenString, nil

}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	var (
		claims jwt.MapClaims
		ok     bool
	)

	if err != nil {
		return claims, fmt.Errorf("Failed to parse token: %v", err)
	}

	// Check token validity and claims
	if claims, ok = parsedToken.Claims.(jwt.MapClaims); !ok && !parsedToken.Valid {
		return claims, fmt.Errorf("Token is invalid or claims are incorrect.")
	}

	return claims, err
}

// Function to read a PEM-encoded file (private key or public key)
func readKeyFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

var privateKey *ecdsa.PrivateKey
var publicKey *ecdsa.PublicKey
var PublicKeyString string
var refreshToken []byte

func LoadKeys(cfg *config.Config) {
	privateKeyData, err := readKeyFromFile("private.pem")
	if err != nil {
		log.Fatalf("Failed to read private key: %v", err)
	}

	privateKey, err = jwt.ParseECPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	publicKeyData, err := readKeyFromFile("public.pem")
	if err != nil {
		log.Fatalf("Failed to read public key: %v", err)
	}

	PublicKeyString = string(publicKeyData)

	publicKey, err = jwt.ParseECPublicKeyFromPEM(publicKeyData)
	if err != nil {
		log.Fatalf("Failed to parse public key: %v", err)
	}

	refreshToken = []byte(cfg.Secret.RefreshToken)
}
