package db

import (
	"database/sql"
	"go/build"
	"log"
	"os"

	common "github.com/snapiz/go-vue-portal-starter/common/go"
	"github.com/volatiletech/sqlboiler/boil"

	// import psql driver
	_ "github.com/lib/pq"
	testfixtures "gopkg.in/testfixtures.v2"
)

var (
	// DB from database/sql
	DB *sql.DB

	// Fixtures fake data
	Fixtures *testfixtures.Context

	// Source database source
	Source string
)

func init() {
	var err error

	Source = os.Getenv("DATABASE_SOURCE")
	if Source == "" {
		common.LoadEnv("services/main")
		Source = os.Getenv("DATABASE_SOURCE")
	}
	DB, err = sql.Open("postgres", Source)

	if err != nil {
		log.Fatal(err)
	}
	boil.SetDB(DB)

	if os.Getenv("GO_ENV") == "development" {
		boil.DebugMode = true
	}

	if os.Getenv("GO_ENV") == "test" {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}
		Fixtures, err = testfixtures.NewFolder(DB, &testfixtures.PostgreSQL{
			SkipResetSequences: true,
		}, gopath+"/src/github.com/snapiz/go-vue-portal-starter/services/main/api/db/fixtures")

		if err != nil {
			log.Fatal(err)
		}
	}
}
