package utils

import (
	"encoding/base64"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/graphql-go/handler"
)

// NewHandler create graphql handler from APIGatewayProxyRequest
func NewHandler(config handler.Config, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	accessor := core.RequestAccessor{}
	r, err := accessor.ProxyEventToHTTPRequest(e)
	w := core.NewProxyResponseWriter()

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	GraphqlHandler(config, w, r)

	return w.GetProxyResponse()
}

// GraphqlHandler create graphql handler
func GraphqlHandler(config handler.Config, w http.ResponseWriter, r *http.Request) {
	h := handler.New(&config)

	ctx := NewContext(Context{
		Response: w,
		Request:  r,
	})

	h.ContextHandler(ctx, w, r)
}

// HandlerFunc http HandlerFun
func HandlerFunc(w http.ResponseWriter, r *http.Request, resp events.APIGatewayProxyResponse) {
	for key, value := range resp.Headers {
		w.Header().Set(key, value)
	}

	decodedBody := []byte(resp.Body)
	if resp.IsBase64Encoded {
		base64Body, err := base64.StdEncoding.DecodeString(resp.Body)
		if err != nil {
			panic(err)
		}
		decodedBody = base64Body
	}

	if resp.StatusCode < 300 || resp.StatusCode > 308 {
		w.WriteHeader(resp.StatusCode)
		w.Write(decodedBody)
	} else {
		http.Redirect(w, r, w.Header().Get("Location"), resp.StatusCode)
	}
}
