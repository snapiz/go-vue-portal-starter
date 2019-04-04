package schema

import (
	"strings"

	common "github.com/snapiz/go-vue-portal-starter/common/go"
	"github.com/snapiz/go-vue-portal-starter/services/main/api/db/models"
	"github.com/snapiz/go-vue-portal-starter/services/main/api/utils"
	"github.com/graphql-go/graphql"
)

// UserRoleType role of user
var UserRoleType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserRole",
	Description: "The role of the user",
	Values: graphql.EnumValueConfigMap{
		strings.ToUpper(models.UserRoleAdmin): &graphql.EnumValueConfig{
			Value: models.UserRoleAdmin,
		},
		strings.ToUpper(models.UserRoleStaff): &graphql.EnumValueConfig{
			Value: models.UserRoleStaff,
		},
		strings.ToUpper(models.UserRoleUser): &graphql.EnumValueConfig{
			Value: models.UserRoleUser,
		},
	},
})

// UserStateType state of user
var UserStateType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserState",
	Description: "The state of the user",
	Values: graphql.EnumValueConfigMap{
		strings.ToUpper(models.UserStateEnable): &graphql.EnumValueConfig{
			Value: models.UserStateEnable,
		},
		strings.ToUpper(models.UserStateDisable): &graphql.EnumValueConfig{
			Value: models.UserStateDisable,
		},
		strings.ToUpper(models.UserStateMaintenance): &graphql.EnumValueConfig{
			Value: models.UserStateMaintenance,
		},
	},
})

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": utils.GlobalIDField("User"),
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "The email of the user.",
		},
		"displayName": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(*models.User)
				return u.DisplayName.Ptr(), nil
			},
		},
		"picture": &graphql.Field{
			Type:        graphql.String,
			Description: "The picture of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(*models.User)
				return u.Picture.Ptr(), nil
			},
		},
		"avatar": &graphql.Field{
			Type:        graphql.String,
			Description: "The avatar of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(*models.User)
				if u.Picture.Ptr() == nil {
					return common.GetAvatarURL(u.EmailHash), nil
				}

				return u.Picture.Ptr(), nil
			},
		},
		"role": &graphql.Field{
			Type: UserRoleType,
		},
		"state": &graphql.Field{
			Type: UserStateType,
		},
	},
	Interfaces: []*graphql.Interface{
		nodeDefinitions.NodeInterface,
	},
})
