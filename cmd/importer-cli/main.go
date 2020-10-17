package main

import (
	"log"

	"github.com/jeslopcru/email-customer-challenge/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "customer-cli"}
	rootCmd.AddCommand(cli.InitCustomerCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
