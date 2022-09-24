package middleware

import (
	"agmc/pkg/util/response"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// type ctxValue int

// const (
// 	ctxUserJWT ctxValue = iota + 1
// )

type UserClaims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

var mySigningKey = []byte(os.Getenv("APP_JWT_SECRET"))

func CreateToken(userID int64) (string, error) {
	claims := &UserClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func UserJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		keyFunc := func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
			}
			return mySigningKey, nil
		}

		_, err := jwt.Parse(auth, keyFunc)
		if err != nil {
			return response.ResponseWithJSON(c, err.Error(), nil, http.StatusUnauthorized)
		}
		// claims, ok := token.Claims.(jwt.MapClaims)
		// if !ok {
		// 	return util.ResponseWithJSON(c, "internal server error", nil, http.StatusInternalServerError)
		// }

		// ctx := context.WithValue(c.Request().Context(), ctxUserJWT, token)

		// return next(c.Echo().NewContext(c.Request().WithContext(ctx), c.Response()))
		return next(c)
	}
}
