package main

import (
	"net/http"
	"os"

	"github.com/snapiz/go-vue-portal-starter/services/main/api/utils"
	"github.com/gorilla/mux"

	common "github.com/snapiz/go-vue-portal-starter/common/go"
	_ "github.com/snapiz/go-vue-portal-starter/services/main/api/db"
	"github.com/snapiz/go-vue-portal-starter/services/main/api/schema"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/handler"
)

func apiHandler(e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return utils.NewHandler(handler.Config{
		Schema:   &schema.Schema,
		GraphiQL: true,
	}, e)
}

func init() {
	common.LoadEnv("services/main")
}

func main() {
	if os.Getenv("GO_ENV") == "dev" {
		rtr := mux.NewRouter()
		rtr.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			utils.GraphqlHandler(handler.Config{
				Schema:   &schema.Schema,
				GraphiQL: true,
			}, w, r)
		}).Methods("POST")
		rtr.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			utils.GraphqlHandler(handler.Config{
				Schema:   &schema.Schema,
				GraphiQL: true,
			}, w, r)
		}).Methods("GET")

		http.Handle("/", rtr)

		if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
			panic(err)
		}
	} else {
		lambda.Start(apiHandler)
	}
}
