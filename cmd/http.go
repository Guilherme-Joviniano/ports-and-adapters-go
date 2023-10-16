/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Guilherme-Joviniano/go-hexagonal/adapters/web/fiber/server"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server := server.MakeNewWebServer()
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
