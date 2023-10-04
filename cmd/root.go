package cmd

import (
	"os"
	sheethandlers "sheets_manager/cmd/handlers/sheets"
	spreadsheethandlers "sheets_manager/cmd/handlers/spreadsheets"
	"sheets_manager/cmd/handlers/tables"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sheets_manager",
	Short: "Sheets_manager will help you to manage spreadsheets, sheets and tables easily.",
	Long: `Sheets_manager allows you to create, update, delete and show your spreadsheets/sheets/tables. Call "sheets_manager help" to see available commands.`,

	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(spreadsheethandlers.SpreadsheetsCmd)
	rootCmd.AddCommand(sheethandlers.SheetsCmd)
	rootCmd.AddCommand(tables.TablesCmd)
}


