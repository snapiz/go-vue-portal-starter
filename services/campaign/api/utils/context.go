package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin/binding"

	"github.com/aws/aws-lambda-go/events"
	"github.com/snapiz/go-vue-portal-starter/services/campaign/api/db/models"
	validator "gopkg.in/go-playground/validator.v9"
)

type (
	key int

	// Context graphql context
	Context struct {
		Request  *http.Request
		Response http.ResponseWriter
		User     *models.User
		Params   map[string]string
		Host     string
		Binder   binding.Binding
	}
)

// Validator instance
var Validator = validator.New()

// Panic error
func (c *Context) Panic(code int, message string) {
	c.Response.WriteHeader(code)
	panic(message)
}

// Param get route param by key
func (c *Context) Param(key string) string {
	return c.Params[key]
}

// EnsureIsAuthorized user is authorized
func (c *Context) EnsureIsAuthorized(cb func(*models.User) bool) {
	if c.User == nil {
		c.Panic(http.StatusUnauthorized, "Anonymous access is denied")
	}

	if c.User.State == models.UserStateMaintenance || (cb != nil && !cb(c.User)) {
		c.Panic(http.StatusForbidden, "Access is denied")
	}
}

// Validate struct
func (c *Context) Validate(inputMap map[string]interface{}, s interface{}, cb func(err validator.FieldError) string) {
	jsonString, _ := json.Marshal(inputMap)
	json.Unmarshal(jsonString, &s)

	if err := Validator.Struct(s); err != nil {
		e := err.(validator.ValidationErrors)[0].(validator.FieldError)

		if cb == nil {
			c.Panic(http.StatusBadRequest, e.Translate(nil))
		} else {
			c.Panic(http.StatusBadRequest, cb(e))
		}
	}
}

// Bind struct
func (c *Context) Bind(s interface{}, cb func(err validator.FieldError) string) (events.APIGatewayProxyResponse, error) {
	if err := c.Binder.Bind(c.Request, s); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to bind")
	}

	if err := Validator.Struct(s); err != nil {
		e := err.(validator.ValidationErrors)[0].(validator.FieldError)
		message := ""
		if cb == nil {
			message = e.Translate(nil)
		} else {
			message = cb(e)
		}

		return c.JSON(http.StatusBadRequest, message)
	}

	return c.JSON(http.StatusOK, "")
}

// JSON APIGatewayProxyResponse
func (c *Context) JSON(code int, v interface{}) (events.APIGatewayProxyResponse, error) {
	jsonString, err := json.Marshal(v)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	headers := map[string]string{
		"Content-type": "application/json",
	}
	for name, value := range c.Response.Header() {
		headers[name] = value[0]
	}

	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       string(jsonString[:]),
		Headers:    headers,
	}, nil
}

// Redirect redirect to url
func (c *Context) Redirect(code int, url string) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Location": url,
	}

	for name, value := range c.Response.Header() {
		headers[name] = value[0]
	}

	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers:    headers,
	}, nil
}

const responseKeyID key = 0
const requestKeyID key = 1
const userKeyID key = 2
const hostKeyID key = 3
const paramsKeyID key = 4
const binderKeyID key = 5

// NewContext create new context
func NewContext(c Context) context.Context {
	if c.Request == nil {
		c.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
		c.Request.Header.Set("Content-Type", "application/json")
	}
	if c.Response == nil {
		c.Response = httptest.NewRecorder()
	}
	if c.Params == nil {
		c.Params = map[string]string{}
	}

	r := c.Request
	scheme := "https"
	referer := r.Referer()
	u, err := url.Parse(referer)

	if r.TLS == nil {
		scheme = "http"
	}

	if referer == "" || err != nil {
		c.Host = fmt.Sprintf("%s://%s", scheme, r.Host)
	} else {
		c.Host = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	}

	c.Binder = binding.Default(r.Method, r.Header.Get("Content-Type"))

	if c.User == nil {
		c.User, _ = GetUserFromCookie(c)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, responseKeyID, c.Response)
	ctx = context.WithValue(ctx, requestKeyID, c.Request)
	ctx = context.WithValue(ctx, userKeyID, c.User)
	ctx = context.WithValue(ctx, hostKeyID, c.Host)
	ctx = context.WithValue(ctx, paramsKeyID, c.Params)
	ctx = context.WithValue(ctx, binderKeyID, c.Binder)

	return ctx
}

// FromContext convert context.Context to Context
func FromContext(c context.Context) Context {
	return Context{
		Response: c.Value(responseKeyID).(http.ResponseWriter),
		Request:  c.Value(requestKeyID).(*http.Request),
		User:     c.Value(userKeyID).(*models.User),
		Host:     c.Value(hostKeyID).(string),
		Params:   c.Value(paramsKeyID).(map[string]string),
		Binder:   c.Value(binderKeyID).(binding.Binding),
	}
}
