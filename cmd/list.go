/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"traefik-acme/intetnal/resolver"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List domains and certificates",
	Run:   List,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func List(cmd *cobra.Command, args []string) {
	acmeFile, _ := rootCmd.PersistentFlags().GetString("input")
	// Load the ACME configuration from a JSON file
	certResolvers, err := resolver.Load(acmeFile)
	if err != nil {

		slog.Error("Error loading ACME configuration:", err)
		os.Exit(1)
	}

	for resolverName, r := range certResolvers {
		fmt.Println("Name:", resolverName)
		fmt.Println("email:", r.Account.Email)
		domainCrts := r.GetCerts()
		for _, crt := range domainCrts {
			fmt.Println("-> Domain:", crt.Domain)
		}
	}
}
