package main

import (
	"log"

	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "goose down",
	Run: func(cmd *cobra.Command, args []string) {
		err := goose.Down(db, "api/db/migrations")
		if err != nil {
			log.Fatal(err)
		}
	},
}
