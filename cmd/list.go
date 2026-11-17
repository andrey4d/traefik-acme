/*
 *   Copyright (c) 2026 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package cmd

import (
	"fmt"

	"traefik-acme/internal/resolver"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List domains and certificates",
	RunE:  List,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func List(cmd *cobra.Command, args []string) error {
	acmeFile, _ := rootCmd.PersistentFlags().GetString("input")
	// Load the ACME configuration from a JSON file
	certResolvers, err := resolver.Load(acmeFile)
	if err != nil {
		return fmt.Errorf("load acme file %q: %w", acmeFile, err)
	}

	for resolverName, r := range certResolvers {
		fmt.Println("Name:", resolverName)
		fmt.Println("email:", r.Account.Email)
		domainCrts := r.GetCerts()
		for _, crt := range domainCrts {
			fmt.Println("-> Domain:", crt.Domain)
		}
	}
	return nil
}
