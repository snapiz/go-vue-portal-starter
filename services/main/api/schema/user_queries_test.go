package schema

import (
	"log"
	"os"
	"testing"

	"github.com/snapiz/go-vue-portal-starter/services/main/api/db"
	"github.com/snapiz/go-vue-portal-starter/services/main/api/db/models"
	"github.com/snapiz/go-vue-portal-starter/services/main/api/utils"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var john *models.User
var albert *models.User

func TestMain(m *testing.M) {
	if err := db.Fixtures.Load(); err != nil {
		log.Fatal(err)
	}

	john, _ = models.Users(qm.Where("id = ?", "4c9f32df-c112-4d02-a459-3493fac49ea9")).OneG()
	albert, _ = models.Users(qm.Where("id = ?", "04231367-deef-4444-bc41-529281445b5f")).OneG()

	os.Exit(m.Run())
}

func TestMeQuery_AnonymousShouldNotQueryMe(t *testing.T) {
	query := `
        query MeQuery {
          me {
            id
			email
			role
          }
        }
	  `

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"me": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: query,
		Context:       utils.NewContext(utils.Context{}),
	})

	assert.Equal(t, expected, result, "Me should be nil")
}

func TestMeQuery_ShouldQueryMe(t *testing.T) {
	query := `
        query MeQuery {
          me {
            id
			email
			role
          }
        }
	  `

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"me": map[string]interface{}{
				"id":    "VXNlcjphNXRScjQ1SnU4RFYydWdZZ2JVeGo=",
				"email": "albert@dupont.com",
				"role":  "STAFF",
			},
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: query,
		Context:       utils.NewContext(utils.Context{User: albert}),
	})

	assert.Equal(t, expected, result, "Me should be albert")
}
