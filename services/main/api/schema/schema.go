package schema

import "github.com/graphql-go/graphql"

// Schema defined in init()
var Schema graphql.Schema

func init() {
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"node": nodeDefinitions.NodeField,
				"me":   meQuery,
			},
		}),
	})
	if err != nil {
		// panic if there is an error in schema
		panic(err)
	}
}
