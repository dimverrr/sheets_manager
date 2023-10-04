package sheethandlers

import (
	"fmt"
	"log"
	spreadsheethandlers "sheets_manager/cmd/handlers/spreadsheets"
	"sheets_manager/setup/config"

	"github.com/spf13/cobra"
	"google.golang.org/api/sheets/v4"
)

var getSheetsCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all your sheets from spreadsheet.",
	Long: `Get all your sheets from spreadsheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		GetAllSheets()
	},
}

func init() {
	SheetsCmd.AddCommand(getSheetsCmd)
}

func GetAllSheets() {
	srv := config.ApiConnect()
	id := spreadsheethandlers.CheckId()

	res, err := srv.Spreadsheets.Get(id).Do()
    if err != nil {
        log.Fatalf("Unable to get sheets: %v", err)
    }

    sheets := res.Sheets

    for _, sheet := range sheets {
        fmt.Println(sheet.Properties.Title)
    }

}

func GetSheets() []*sheets.Sheet{
	srv := config.ApiConnect()
	id := spreadsheethandlers.CheckId()

	res, err := srv.Spreadsheets.Get(id).Do()
    if err != nil {
        log.Fatalf("Unable to get sheets: %v", err)
    }

    sheets := res.Sheets

    return sheets
}