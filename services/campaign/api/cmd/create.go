package main

import (
	"log"

	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "goose create",
	Run: func(cmd *cobra.Command, args []string) {
		err := goose.Create(db, "api/db/migrations", args[0], "sql")
		if err != nil {
			log.Fatal(err)
		}
	},
}
