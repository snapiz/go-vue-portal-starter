package utils

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/snapiz/go-vue-portal-starter/services/auth/db/models"
	common "github.com/snapiz/go-vue-portal-starter/common/go"
	"github.com/volatiletech/sqlboiler/queries/qm"

	jwt "github.com/dgrijalva/jwt-go"
)

// SetToken set cookie token
func SetToken(c Context, u *models.User) (token string, err error) {
	expireDuration, err := time.ParseDuration(os.Getenv("SESSION_EXPIRE_DURATION"))

	if err != nil {
		return "", err
	}

	token, err = common.SignToken(jwt.StandardClaims{
		Id:      u.ID,
		Subject: strconv.FormatInt(*u.TokenVersion.Ptr(), 10),
		Issuer:  c.Host,
	})

	if err != nil {
		return "", err
	}

	cookie := &http.Cookie{
		Name:     os.Getenv("SESSION_KEY"),
		Value:    token,
		Expires:  time.Now().Add(expireDuration),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(c.Response, cookie)

	return token, nil
}

// RemoveToken unset token from cookie
func RemoveToken(c Context) {
	cookie := &http.Cookie{
		Name:     os.Getenv("SESSION_KEY"),
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Path:     "/",
	}

	c.Response.Header().Del("Set-Cookie")
	http.SetCookie(c.Response, cookie)
}

// GetUserFromCookie get user from cookie
func GetUserFromCookie(c Context) (*models.User, error) {
	authCookie, err := c.Request.Cookie(os.Getenv("SESSION_KEY"))

	if err != nil {
		return nil, err
	}

	claims, err := common.VerifyToken(authCookie.Value, c.Host)

	if err != nil {
		return nil, err
	}

	users, err := models.Users(qm.Where("id = ? AND token_version = ? AND state != 'disable'", claims.Id, claims.Subject)).AllG()

	if err != nil {
		return nil, err
	}

	if users == nil {
		RemoveToken(c)
		return nil, nil
	}

	u := users[0]
	SetToken(c, u)

	return u, nil
}
