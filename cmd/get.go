/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"
	"os"
	"traefik-acme/intetnal/resolver"

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
		slog.Error("Error loading ACME configuration:", "error", err)
		os.Exit(1)
	}
	path, _ := cmd.Flags().GetString("dir")

	for resolverName, r := range certResolvers {
		slog.Info("Found resolver", "name", resolverName, "email", r.Account.Email)
		for _, crt := range r.GetCerts() {
			slog.Info("Certificate found", "domain", crt.Domain)

			if err := resolver.Save(crt, path); err != nil {
				slog.Error("Write file", "path", path)
				return err
			}
		}
	}
	return nil
}
