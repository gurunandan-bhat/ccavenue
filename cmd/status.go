/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ccavenue/client"
	"ccavenue/config"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Retrieve order status by OrderID",
	RunE: func(cmd *cobra.Command, args []string) error {
		orderID, _ := cmd.Flags().GetInt32("order-id")
		cfg, err := config.Configuration()
		if err != nil {
			log.Fatalf("Error reading configuration: %s", err)
		}
		ccavClient, err := client.NewClient(cfg, "1.2")
		if err != nil {
			log.Fatal("Error creating client: ", err)
		}

		filter := client.StatusFilter{
			OrderNo: strconv.Itoa(int(orderID)),
		}

		jsonStr, err := ccavClient.Post(filter)
		if err != nil {
			log.Fatal("Error from orders request: ", err)
		}

		fmt.Println("JSON: ", string(*jsonStr))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	statusCmd.Flags().Int32P("order-id", "o", 0, "order-id to fetch status for")
	cobra.MarkFlagRequired(statusCmd.Flags(), "order-id")
}
