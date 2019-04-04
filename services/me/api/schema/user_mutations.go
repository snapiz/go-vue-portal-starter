package schema

import (
	"context"
	"net/http"
	"time"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/snapiz/go-vue-portal-starter/services/me/api/db/models"
	"github.com/snapiz/go-vue-portal-starter/services/me/api/db/services"
	"github.com/snapiz/go-vue-portal-starter/services/me/api/utils"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	validator "gopkg.in/go-playground/validator.v9"
)

var updateUserMutation = relay.MutationWithClientMutationID(relay.MutationConfig{
	Name: "UpdateUser",
	InputFields: graphql.InputObjectConfigFieldMap{
		"displayName": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"picture": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"username": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
	OutputFields: graphql.Fields{
		"user": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if payload, ok := p.Source.(map[string]interface{}); ok {
					return payload["user"].(*models.User), nil
				}
				return nil, nil
			},
		},
	},
	MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
		c := utils.FromContext(ctx)

		c.EnsureIsAuthorized(nil)

		input := new(struct {
			Username    string `validate:"omitempty,alphanum,min=3,max=50"`
			DisplayName string `validate:"required,min=3,max=50"`
			Picture     string `validate:"url|len=0"`
		})

		c.Validate(inputMap, input, func(e validator.FieldError) string {
			switch e.Field() {
			case "DisplayName":
				return "The displayName field must be between 3 and 50 characters long"
			case "Picture":
				return "The picture field is invalid URL"
			case "Username":
				if e.Tag() == "alphanum" {
					return "The username field must be alphanum"
				}
				return "The username field must be between 3 and 50 characters long"
			}
			return e.Translate(nil)
		})

		if input.Username != "" {
			if c.User.Username.Ptr() != nil {
				c.Panic(http.StatusUnauthorized, "errors.user.usernameAlreadyDefined")
			}
			users, err := models.Users(qm.Where("username = ?", input.Username)).AllG()
			if err != nil {
				c.Panic(http.StatusInternalServerError, "errors.user.fetch")
			}
			if users != nil {
				c.Panic(http.StatusBadRequest, "errors.auth.usernameAlreadyExists")
			}
			c.User.Username = null.StringFrom(input.Username)
		}

		if input.Picture == "" {
			c.User.Picture = null.NewString("", false)
		} else {
			c.User.Picture = null.StringFrom(input.Picture)
		}

		c.User.DisplayName = null.StringFrom(input.DisplayName)
		c.User.UpdatedAt = null.TimeFrom(time.Now())
		c.User.UpdateG(boil.Whitelist("username", "display_name", "picture", "updated_at"))

		return map[string]interface{}{
			"user": c.User,
		}, nil
	},
})

var changePasswordMutation = relay.MutationWithClientMutationID(relay.MutationConfig{
	Name: "ChangePassword",
	InputFields: graphql.InputObjectConfigFieldMap{
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"currentPassword": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	OutputFields: graphql.Fields{
		"user": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if payload, ok := p.Source.(map[string]interface{}); ok {
					return payload["user"].(*models.User), nil
				}

				return nil, nil
			},
		},
	},
	MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
		c := utils.FromContext(ctx)

		c.EnsureIsAuthorized(nil)

		input := new(struct {
			Password        string `validate:"required,min=8,max=20"`
			CurrentPassword string `validate:"min=8|len=0,max=20"`
		})

		c.Validate(inputMap, input, func(e validator.FieldError) string {
			switch e.Field() {
			case "Password":
				return "The password field must be between 8 and 20 characters long"
			case "CurrentPassword":
				return "The currentPassword field must be between 8 and 20 characters long"
			}
			return e.Translate(nil)
		})

		if c.User.Password.Ptr() != nil && !services.VerifyUserPassword(c.User, input.CurrentPassword) {
			c.Panic(http.StatusBadRequest, "errors.user.badPassword")
		}

		if err := services.SetUserPassword(c.User, input.Password); err != nil {
			c.Panic(http.StatusInternalServerError, err.Error())
		}

		if _, err := c.User.UpdateG(boil.Whitelist("password", "token_version")); err != nil {
			c.Panic(http.StatusInternalServerError, "Failed to update user")
		}

		utils.SetToken(c, c.User)

		return map[string]interface{}{
			"user": c.User,
		}, nil
	},
})
