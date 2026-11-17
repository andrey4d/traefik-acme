/*
 *   Copyright (c) 2026 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package cmd

import (
	"fmt"
	"log/slog"

	"traefik-acme/internal/resolver"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get certificates from acme.json",
	RunE:  Get,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("dir", "d", "certs", "output certificates dir")
}

func Get(cmd *cobra.Command, args []string) error {
	acmeFile, _ := rootCmd.PersistentFlags().GetString("input")
	// Load the ACME configuration from a JSON file
	certResolvers, err := resolver.Load(acmeFile)
	if err != nil {
		return fmt.Errorf("load acme file %q: %w", acmeFile, err)
	}
	path, _ := cmd.Flags().GetString("dir")

	for resolverName, r := range certResolvers {
		slog.Info("Found resolver", "name", resolverName, "email", r.Account.Email)
		for _, crt := range r.GetCerts() {
			slog.Info("Certificate found", "domain", crt.Domain)

			if err := resolver.Save(crt, path); err != nil {
				return fmt.Errorf("save certificate for %q: %w", crt.Domain, err)
			}
		}
	}
	return nil
}
