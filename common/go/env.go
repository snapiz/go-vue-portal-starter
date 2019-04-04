package common

import (
	"go/build"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv load environement variables
func LoadEnv(dir string) {
	goEnv := os.Getenv("GO_ENV")

	if goEnv == "" {
		return
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	if goEnv == "test" {
		godotenv.Load(gopath + "/src/github.com/snapiz/go-vue-portal-starter/" + dir + "/.env.test")
	} else if goEnv == "dev" {
		godotenv.Load(gopath + "/src/github.com/snapiz/go-vue-portal-starter/" + dir + "/.env.local")
	} else if goEnv == "prod" {
		godotenv.Load(gopath + "/src/github.com/snapiz/go-vue-portal-starter/" + dir + "/.env.prod.local")
	}

	godotenv.Load(gopath + "/src/github.com/snapiz/go-vue-portal-starter/" + dir + "/.env")
}
