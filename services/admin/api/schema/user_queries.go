package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/snapiz/go-vue-portal-starter/services/admin/api/utils"
)

var meQuery = &graphql.Field{
	Type: userType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		c := utils.FromContext(p.Context)

		return c.User, nil
	},
}
