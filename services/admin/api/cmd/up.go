package main

import (
	"log"

	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

func up(cmd *cobra.Command, args []string) {
	err := goose.Up(db, "api/db/migrations")
	if err != nil {
		log.Fatal(err)
	}
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "goose up",
	Run:   up,
}
