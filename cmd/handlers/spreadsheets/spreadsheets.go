package spreadsheethandlers

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	spreadsheetName string
	SpreadsheetId string
)

var SpreadsheetsCmd = &cobra.Command{
	Use:   "spreadsheets",
	Short: "Crate operation for spreadsheets.",
	Long: `Crate operation for spreadsheets.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Crate operation for spreadsheets. Call `create` command to create new spreadsheet")
	},
}

