package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/snapiz/go-vue-portal-starter/services/auth/db/models"
	"github.com/snapiz/go-vue-portal-starter/services/auth/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type profile struct {
	ID          string
	DisplayName string
	Picture     string
	Email       string
}

type providerConfig struct {
	ClientID     string
	ClientSecret string
	Endpoint     oauth2.Endpoint
	Scopes       []string
	GetProfile   func(*oauth2.Token) (*profile, error)
}

var providers = map[string]*providerConfig{
	"google": &providerConfig{
		ClientID:     "GOOGLE_CLIENT_ID",
		ClientSecret: "GOOGLE_CLIENT_SECRET",
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		GetProfile: func(tok *oauth2.Token) (*profile, error) {
			response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken)
			if err != nil {
				return nil, fmt.Errorf("failed getting user info: %s", err.Error())
			}
			defer response.Body.Close()

			type googlePlusProfile struct {
				ID      string `json:"id"`
				Name    string `json:"name"`
				Picture string `json:"picture"`
				Email   string `json:"email"`
			}

			gPP := &googlePlusProfile{}
			json.NewDecoder(response.Body).Decode(&gPP)

			return &profile{
				ID:          gPP.ID,
				Email:       gPP.Email,
				DisplayName: gPP.Name,
				Picture:     gPP.Picture,
			}, nil
		},
	},
	"facebook": &providerConfig{
		ClientID:     "FACEBOOK_CLIENT_ID",
		ClientSecret: "FACEBOOK_CLIENT_SECRET",
		Endpoint:     facebook.Endpoint,
		Scopes:       []string{"public_profile", "email"},
		GetProfile: func(tok *oauth2.Token) (*profile, error) {
			response, err := http.Get("https://graph.facebook.com/me?fields=id,email,name,picture&access_token=" + tok.AccessToken)
			if err != nil {
				return nil, fmt.Errorf("failed getting user info: %s", err.Error())
			}
			defer response.Body.Close()

			type facebookPictureData struct {
				URL string `json:"url"`
			}

			type facebookPicture struct {
				Data facebookPictureData `json:"data"`
			}

			type facebookProfile struct {
				ID      string          `json:"id"`
				Name    string          `json:"name"`
				Picture facebookPicture `json:"picture"`
				Email   string          `json:"email"`
			}

			gPP := &facebookProfile{}
			json.NewDecoder(response.Body).Decode(&gPP)

			return &profile{
				ID:          gPP.ID,
				Email:       gPP.Email,
				DisplayName: gPP.Name,
				Picture:     gPP.Picture.Data.URL,
			}, nil
		},
	},
}

func createOauthConf(c utils.Context, provider *providerConfig) oauth2.Config {
	name := c.Param("provider")

	return oauth2.Config{
		ClientID:     os.Getenv(provider.ClientID),
		ClientSecret: os.Getenv(provider.ClientSecret),
		RedirectURL:  fmt.Sprintf("%s/auth/%s/callback", c.Host, name),
		Scopes:       provider.Scopes,
		Endpoint:     provider.Endpoint,
	}
}

func oauth2Handler(c utils.Context) (events.APIGatewayProxyResponse, error) {
	if ok := strings.HasSuffix(c.Request.URL.Path, "/callback"); ok {
		return oauth2CallbackHandler(c)
	}

	name := c.Param("provider")
	provider := providers[name]

	if provider == nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Unknown provider %s.", name))
	}

	conf := createOauthConf(c, provider)

	return c.Redirect(http.StatusTemporaryRedirect, conf.AuthCodeURL("state"))
}

func oauth2CallbackHandler(c utils.Context) (events.APIGatewayProxyResponse, error) {
	input := new(struct {
		Code        string `form:"code"`
		ClientID    string `form:"clientId"`
		RedirectURI string `form:"redirectUri"`
	})

	if proxyResponse, err := c.Bind(input, nil); proxyResponse.StatusCode != http.StatusOK {
		return proxyResponse, err
	}

	name := c.Param("provider")
	provider := providers[name]

	if provider == nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Unknown provider %s.", name))
	}

	conf := createOauthConf(c, provider)
	tok, err := conf.Exchange(oauth2.NoContext, input.Code)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to exchange")
	}

	p, err := provider.GetProfile(tok)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get profile")
	}

	users, err := models.Users(qm.Where("email = ?", p.Email)).AllG()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch user")
	}

	var user *models.User

	if users == nil {
		hash := md5.Sum([]byte(p.Email))
		user = &models.User{
			Email:        p.Email,
			EmailHash:    hex.EncodeToString(hash[:]),
			TokenVersion: null.Int64From(time.Now().Unix()),
			DisplayName:  null.StringFrom(p.DisplayName),
			Picture:      null.StringFrom(p.Picture),
		}

		err := user.InsertG(boil.Whitelist("email", "email_hash", "token_version", "display_name", "picture"))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to create user")
		}
	} else {
		user = users[0]
	}

	// Check account is disabled
	if user.State == models.UserStateDisable {
		return c.JSON(http.StatusUnauthorized, "errors.auth.accountIsDisabled")
	}

	userProviders, err := models.UserProviders(qm.Where("provider = ? AND provider_id = ? AND user_id = ?", name, p.ID, user.ID)).AllG()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch user provider")
	}

	if userProviders == nil {
		userProvider := &models.UserProvider{
			Provider:   name,
			ProviderID: p.ID,
			UserID:     user.ID,
		}

		err = userProvider.InsertG(boil.Whitelist("provider", "provider_id", "user_id"))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to create user provider")
		}
	}

	_, err = utils.SetToken(c, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to set token")
	}

	return c.Redirect(http.StatusTemporaryRedirect, c.Host)
}
