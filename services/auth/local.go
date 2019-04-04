package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/snapiz/go-vue-portal-starter/services/auth/db/models"
	"github.com/snapiz/go-vue-portal-starter/services/auth/db/services"
	"github.com/snapiz/go-vue-portal-starter/services/auth/utils"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/aws/aws-lambda-go/events"
)

type respJSON struct {
	ID          string      `json:"id"`
	Email       string      `json:"email"`
	Username    null.String `json:"username"`
	DisplayName null.String `json:"display_name"`
	Picture     null.String `json:"picture"`
	Role        string      `json:"role"`
	State       string      `json:"state"`
}

func localHandler(c utils.Context) (events.APIGatewayProxyResponse, error) {
	if ok := strings.HasSuffix(c.Request.URL.Path, "/new"); ok {
		return localNewHandler(c)
	}

	p := new(struct {
		Login    string `json:"login" validate:"required"`
		Password string `json:"password" validate:"required"`
	})

	proxyResponse, err := c.Bind(p, func(x validator.FieldError) string {
		return fmt.Sprintf("The %s field is required", strings.ToLower(x.Field()))
	})

	if proxyResponse.StatusCode != http.StatusOK {
		return proxyResponse, err
	}

	users, err := models.Users(qm.Where("email = ?", p.Login), qm.Or("username = ?", p.Login)).AllG()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Falied to request user")
	}

	if users == nil {
		return c.JSON(http.StatusUnauthorized, "errors.auth.invalidCredentials")
	}

	user := users[0]

	if user.Password.Ptr() == nil {
		return c.JSON(http.StatusUnauthorized, "errors.auth.invalidCredentials")
	}

	// Verify password
	if ok := services.VerifyUserPassword(user, p.Password); !ok {
		return c.JSON(http.StatusUnauthorized, "errors.auth.invalidCredentials")
	}

	// Check account is disabled
	if user.State == models.UserStateDisable {
		return c.JSON(http.StatusUnauthorized, "errors.auth.accountIsDisabled")
	}

	// write token into cookie
	_, err = utils.SetToken(c, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to write token")
	}

	return c.JSON(http.StatusOK, &respJSON{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Picture:     user.Picture,
		Role:        user.Role,
		State:       user.State,
	})
}

func localNewHandler(c utils.Context) (events.APIGatewayProxyResponse, error) {
	p := new(struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
		Password string `json:"password" validate:"required,min=8,max=20"`
	})

	proxyResponse, err := c.Bind(p, func(x validator.FieldError) string {
		switch x.Field() {
		case "Email":
			return "The email field must be an email"
		case "Username":
			return "The username field must be between 3 and 50 characters long"
		case "Password":
			return "The password field must be between 2 and 20 characters long"
		default:
			return x.Translate(nil)
		}
	})

	if proxyResponse.StatusCode != http.StatusOK {
		return proxyResponse, err
	}

	users, err := models.Users(qm.Where("email = ?", p.Email)).AllG()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch user by email")
	}

	if users != nil {
		return c.JSON(http.StatusBadRequest, "errors.auth.emailAlreadyExists")
	}

	users, err = models.Users(qm.Where("username = ?", p.Username)).AllG()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch user by username")
	}

	if users != nil {
		return c.JSON(http.StatusBadRequest, "errors.auth.usernameAlreadyExists")
	}

	hash := md5.Sum([]byte(p.Email))
	user := &models.User{
		Email:     p.Email,
		EmailHash: hex.EncodeToString(hash[:]),
		Username:  null.StringFrom(p.Username),
	}

	if err := services.SetUserPassword(user, p.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to set password")
	}

	err = user.InsertG(boil.Whitelist("email", "email_hash", "username", "password", "token_version"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	_, err = utils.SetToken(c, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to write token")
	}

	return c.JSON(http.StatusOK, &respJSON{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Picture:     user.Picture,
		Role:        user.Role,
		State:       user.State,
	})
}
