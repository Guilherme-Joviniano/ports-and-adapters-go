/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Guilherme-Joviniano/go-hexagonal/adapters/cli"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/main/factories"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productStatus string
var productPrice float32

// cliProductCmd represents the cliProduct command
var cliProductCmd = &cobra.Command{
	Use:   "cli-product",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(factories.MakeProductService(), action, productName, productId, productPrice)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliProductCmd)
	cliProductCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable / Disable a Product")
	cliProductCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	cliProductCmd.Flags().StringVarP(&productName, "product", "n", "", "Product ID")
	cliProductCmd.Flags().Float32VarP(&productPrice, "price", "p", float32(0), "Product ID")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliProductCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliProductCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
