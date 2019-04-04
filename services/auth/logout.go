package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/snapiz/go-vue-portal-starter/services/auth/utils"
)

func logoutHandler(c utils.Context) (events.APIGatewayProxyResponse, error) {

	utils.RemoveToken(c)

	return c.JSON(http.StatusOK, "ok")
}
