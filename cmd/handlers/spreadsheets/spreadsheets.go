/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package spreadsheethandlers

import (
	"fmt"

	"github.com/spf13/cobra"
)

// spreadsheetsCmd represents the spreadsheets command
var SpreadsheetsCmd = &cobra.Command{
	Use:   "spreadsheets",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("spreadsheets called")
	},
}

