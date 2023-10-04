package spreadsheethandlers

import (
	"fmt"
	"log"
	"os"
	"sheets_manager/setup/config"

	"github.com/spf13/cobra"
)



var setSpreadsheetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set spreadsheet id for create sheets and tables in it. Run this command again to choose another spreadsheet.",
	Long: `Set spreadsheet id for create sheets and tables in it. Run this command again to choose another spreadsheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		SetSpreadsheet()
	},
}

func init() {
	SpreadsheetsCmd.AddCommand(setSpreadsheetCmd)
	setSpreadsheetCmd.Flags().StringVarP(&spreadsheetId, "id", "i", "", "id of spreadsheet")
	setSpreadsheetCmd.MarkFlagRequired("i")
}

func SetSpreadsheet() {
	srv := config.ApiConnect()

	res, err := srv.Spreadsheets.Get(spreadsheetId).Do()
	if err != nil {
		log.Fatalf("Spreadsheet with this ID does not exist")
	}

	err = os.WriteFile("id.txt", []byte(res.SpreadsheetId), 0644)
	if err != nil {
		log.Fatalf("Unable to write spreadsheet id to file: %v", err)
	}

	fmt.Println("Spreadsheet was successfully set")
}

func CheckId() string{
	file, err := os.ReadFile("id.txt")
	if err != nil {
		log.Fatalf("Run `spreadsheets set` command to set spreadsheet ID")
	}

	fileDecrypt := string(file)

	return fileDecrypt
}