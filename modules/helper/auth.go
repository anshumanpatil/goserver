package helperfunctions

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
