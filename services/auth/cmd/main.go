package main

import (
	"database/sql"
	"log"
	"os"

	common "github.com/snapiz/go-vue-portal-starter/common/go"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gws-cli",
		Short: "Internal db migration tool",
	}
	db     *sql.DB
	source string
)

func init() {
	rootCmd.AddCommand(modelsCmd)
}

func main() {
	var err error
	common.LoadEnv("services/auth")
	source = os.Getenv("DATABASE_SOURCE")
	db, err = sql.Open("postgres", source)

	if err != nil {
		log.Fatal(err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
