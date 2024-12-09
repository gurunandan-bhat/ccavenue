/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ccavenue/client"
	"ccavenue/config"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// payoutCmd represents the payout command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt any encoded string received from CCAvenue",
	RunE: func(cmd *cobra.Command, args []string) error {

		fPath, _ := cmd.Flags().GetString("file")
		body, err := os.ReadFile(fPath)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", fPath, err)
		}
		encStr := strings.TrimSpace(string(body))

		cfg, err := config.Configuration()
		if err != nil {
			return fmt.Errorf("error reading configuration: %w", err)
		}
		ccavClient, err := client.NewClient(cfg, "1.2")
		if err != nil {
			return fmt.Errorf("error creating client: %w", err)
		}

		decStr, err := ccavClient.Decode(encStr)
		if err != nil {
			return fmt.Errorf("error decrypting string:\n%s\n%w", encStr, err)
		}

		values, err := url.ParseQuery(string(*decStr))
		if err != nil {
			return fmt.Errorf("error parsing query string %s: %w", string(*decStr), err)
		}
		jsonBytes, err := json.MarshalIndent(values, "", "\t")
		if err != nil {
			return fmt.Errorf("error marshaling query %#v: %w", values, err)
		}

		fmt.Println(string(jsonBytes))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// payoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	decryptCmd.Flags().StringP("file", "f", "", "full path of file that contains encrypted string")
	cobra.MarkFlagRequired(decryptCmd.Flags(), "file")

}
