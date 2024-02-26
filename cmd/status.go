/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ccavenue/client"
	"ccavenue/config"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

type StatusResponse struct {
	client.Order
	client.OrderValue
	client.OrderError
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Retrieve order status by OrderID",
	RunE: func(cmd *cobra.Command, args []string) error {
		orderID, _ := cmd.Flags().GetInt32("order-id")
		cfg, err := config.Configuration()
		if err != nil {
			return fmt.Errorf("error reading configuration: %w", err)
		}
		ccavClient, err := client.NewClient(cfg, "1.2")
		if err != nil {
			return fmt.Errorf("error creating client: %w", err)
		}

		filter := client.StatusFilter{
			OrderNo: strconv.Itoa(int(orderID)),
		}

		jsonBytes, err := ccavClient.Post(filter)
		if err != nil {
			return fmt.Errorf("error from orders request: %w", err)
		}

		status := StatusResponse{}
		if err := json.Unmarshal(*jsonBytes, &status); err != nil {
			return fmt.Errorf("error parsing status JSON: %w", err)
		}

		fmt.Printf("%+v\n", status)
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
