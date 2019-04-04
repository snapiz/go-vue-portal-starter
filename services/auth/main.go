package main

import (
	"net/http"
	"os"

	common "github.com/snapiz/go-vue-portal-starter/common/go"
	_ "github.com/snapiz/go-vue-portal-starter/services/auth/db"
	"github.com/snapiz/go-vue-portal-starter/services/auth/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/gorilla/mux"
)

func init() {
	common.LoadEnv("services/auth")
}

func handler(e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/* provider := e.PathParameters["provider"] */
	accessor := core.RequestAccessor{}
	r, err := accessor.ProxyEventToHTTPRequest(e)
	w := core.NewProxyResponseWriter()

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	c := utils.NewContext(utils.Context{
		Response: w,
		Request:  r,
		Params:   e.PathParameters,
	})
	ctx := utils.FromContext(c)

	switch r.URL.Path {
	case "/auth/logout":
		return logoutHandler(ctx)
	case "/auth/local":
		return localHandler(ctx)
	case "/auth/local/new":
		return localHandler(ctx)
	}

	return oauth2Handler(ctx)
}

func handleFunc(h func(c utils.Context) (events.APIGatewayProxyResponse, error), w http.ResponseWriter, r *http.Request) {
	c := utils.NewContext(utils.Context{
		Response: w,
		Request:  r,
		Params:   mux.Vars(r),
	})
	ctx := utils.FromContext(c)

	resp, err := h(ctx)
	if err != nil {
		panic(err)
	}

	utils.HandlerFunc(w, r, resp)
}

func main() {
	if os.Getenv("GO_ENV") == "dev" {
		rtr := mux.NewRouter()
		rtr.HandleFunc("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
			handleFunc(logoutHandler, w, r)
		}).Methods("GET", "POST")
		rtr.HandleFunc("/auth/local", func(w http.ResponseWriter, r *http.Request) {
			handleFunc(localHandler, w, r)
		}).Methods("POST")
		rtr.HandleFunc("/auth/local/new", func(w http.ResponseWriter, r *http.Request) {
			handleFunc(localHandler, w, r)
		}).Methods("POST")
		rtr.HandleFunc("/auth/{provider:google|facebook}", func(w http.ResponseWriter, r *http.Request) {
			handleFunc(oauth2Handler, w, r)
		}).Methods("GET")
		rtr.HandleFunc("/auth/{provider:google|facebook}/callback", func(w http.ResponseWriter, r *http.Request) {
			handleFunc(oauth2Handler, w, r)
		}).Methods("GET")

		http.Handle("/", rtr)

		if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
			panic(err)
		}
	} else {
		lambda.Start(handler)
	}
}
