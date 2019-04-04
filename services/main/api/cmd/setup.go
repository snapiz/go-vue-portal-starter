package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Init db",
	Run: func(cmd *cobra.Command, args []string) {
		rgx := regexp.MustCompile(`dbname=([a-zA-Z_0-9]+)`)
		result := rgx.FindStringSubmatch(source)
		db, err := sql.Open("postgres", strings.Replace(source, " "+result[0], "", 1))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		db.Exec(fmt.Sprintf("CREATE DATABASE %s", result[1]))

		up(modelsCmd, []string{})
	},
}
