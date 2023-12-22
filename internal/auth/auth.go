package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GLOBAL, TREAT LIKE CONSTANT
var globalJWTSigningMethod = jwt.SigningMethodHS256

const (
	jwtTimeBeforeExpiration = time.Hour * 2
)

// Claims -
type Claims struct {
	UserID string
	jwt.StandardClaims
}

// MakeJWT -
func MakeJWT(userID string, tokenSecret string) (string, error) {
	signingKey := []byte(tokenSecret)

	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "learntocode.fyi",
			IssuedAt:  time.Now().UTC().Unix(),
			ExpiresAt: time.Now().UTC().Add(jwtTimeBeforeExpiration).Unix(),
		},
	}

	token := jwt.NewWithClaims(globalJWTSigningMethod, claims)
	return token.SignedString(signingKey)
}

// ValidateJWT -
func ValidateJWT(tokenString, tokenSecret string) (Claims, error) {
	claimsStruct := Claims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)
	if err != nil {
		return Claims{}, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return Claims{}, errors.New("invalid JWT")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return Claims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}
