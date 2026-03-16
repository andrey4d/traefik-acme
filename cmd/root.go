/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "traefik-acme",
	Short: "Convert traefik acme.json to certificates",
	Long:  `Convert traefik acme.json to certificates.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.traefik-acme.yaml)")
	rootCmd.PersistentFlags().StringP("input", "i", "acme/acme.json", "Input file (acme/acme.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().StringP("input", "i", "acme/acme.json", "Input file (acme/acme.json)")

}
