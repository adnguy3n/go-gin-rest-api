package authenthication

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// Key used for generating JWT's.
// Normally should be stored outside of code.
var jwtKey = []byte("secretKey")

/*
 * Struct for JWT Claims
 */
type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

/*
 * Generate JWT. Takes in an e-mail and username as parameters.
 */
func GenerateJWT(email string, username string) (tokenString string, err error) {
	// Default Expiration Time is 1 hour.
	expirationTime := time.Now().Add(1 * time.Hour)

	// Generate a claim variable using the available data and expiration time.
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate the token using HS256 Signing Algorithms.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)

	return
}

/*
 * Validate Token.
 */
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("Could not parse claims.")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}
