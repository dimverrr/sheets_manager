package spreadsheethandlers

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	spreadsheetName string
	spreadsheetId string
)

var SpreadsheetsCmd = &cobra.Command{
	Use:   "spreadsheets",
	Short: "Set operation for spreadsheets.",
	Long: `Set operation for spreadsheets.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Set operation for spreadsheets. Call `set` command to set new spreadsheet")
	},
}

