package spreadsheethandlers

import (
	"fmt"
	"log"
	"os"
	"sheets_manager/setup/config"

	"github.com/spf13/cobra"
	"google.golang.org/api/sheets/v4"
)

var createSpreadsheetCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new spreadsheet.",
	Long: `Create new spreadsheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateSpreadsheet()
	},
}

func init() {
	SpreadsheetsCmd.AddCommand(createSpreadsheetCmd)
	createSpreadsheetCmd.Flags().StringVarP(&spreadsheetName, "name", "n", "", "name for spreadsheet")
	createSpreadsheetCmd.MarkFlagRequired("name")
}


func CreateSpreadsheet() {
	srv := config.ApiConnect()

	spreadsheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: spreadsheetName,
		},
	}

	res, err := srv.Spreadsheets.Create(spreadsheet).Do()
	if err != nil {
		log.Fatalf("Unable to create spreadsheet: %v", err)
	}

	err = os.WriteFile("id.txt", []byte(res.SpreadsheetId), 0644)
	if err != nil {
		log.Fatalf("Unable to write spreadsheet id to file: %v", err)
	}

	fmt.Printf("New sheet created with Title: %v\n", spreadsheetName)
}


func CheckId() string{
	file, err := os.ReadFile("id.txt")
	if err != nil {
		log.Fatalf("Run `spreadsheets create` command to create spreadsheet")
	}

	fileDecrypt := string(file)

	return fileDecrypt
}