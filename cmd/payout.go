/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ccavenue/client"
	"ccavenue/config"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type PayoutSummaryDetails struct {
	client.PayoutSummary
}

type PayoutSummaryList struct {
	PayoutSummaryDetails `json:"payout_summary_details,omitempty"`
}

type PayoutSummaryResult struct {
	client.PayoutSummaryError
	PayoutSummaryList `json:"payout_summary_list,omitempty"`
}

type PayoutResponse struct {
	PayoutSummaryResult `json:"payout_summary_result,omitempty"`
}

// payoutCmd represents the payout command
var payoutCmd = &cobra.Command{
	Use:   "payout",
	Short: "Retrieve CCAvenue payouts by date",
	RunE: func(cmd *cobra.Command, args []string) error {
		dateStr, _ := cmd.Flags().GetString("date")
		if _, err := time.Parse("02-01-2006", dateStr); err != nil {
			return fmt.Errorf("date must be of the form dd-mm-yyyy: %w", err)
		}

		cfg, err := config.Configuration()
		if err != nil {
			return fmt.Errorf("error reading configuration: %w", err)
		}
		ccavClient, err := client.NewClient(cfg, "1.2")
		if err != nil {
			return fmt.Errorf("error creating client: %w", err)
		}

		filter := client.PayoutFilter{
			SettlementDate: dateStr,
		}

		jsonBytes, err := ccavClient.Post(filter)
		if err != nil {
			return fmt.Errorf("error from payout request: %w", err)
		}

		payout := PayoutResponse{}
		if err := json.Unmarshal(*jsonBytes, &payout); err != nil {
			return fmt.Errorf("error unmarshaling into payout: %w", err)
		}

		fmt.Printf("%+v", payout)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(payoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// payoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	payoutCmd.Flags().StringP("date", "d", "", "date to fetch payouts for")
	cobra.MarkFlagRequired(payoutCmd.Flags(), "date")

}
