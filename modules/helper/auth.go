package helperfunctions

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TokenDetails - token generated for user
type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}

// GenerateToken - to generate token for user
func GenerateToken(userid int) (*TokenDetails, error) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	var AccessSecret = os.Getenv("ACCESS_SECRET")
	td := &TokenDetails{}
	td.AccessTokenExpires = time.Now().Add(time.Minute * 15).Unix()
	td.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AccessTokenExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(AccessSecret))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}

	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RefreshTokenExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(AccessSecret))
	if err != nil {
		return nil, err
	}
	return td, nil

}

// AuthMiddleware - Middleware for every call
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
		if strings.Contains(c.Request.URL.Path, "login") {
			c.Next()
			return
		}
		log.Println("AuthMiddleware=============================================================")
		token := c.Request.Header.Get("Authorization")
		sanitizedToken := strings.Replace(token, "Bearer ", "", 1)
		if tokenn, err := VerifyToken(sanitizedToken); err == nil {
			c.Set("user", tokenn)
		}

		c.Next()
	}
}

// VerifyToken - verify token for every call
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims["user_id"], claims["exp"])
		return claims, nil
	}
	return nil, nil
}
