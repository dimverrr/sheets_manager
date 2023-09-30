package sheets

import (
	"fmt"
	"log"
	"sheets_manager/models"
	"sheets_manager/setup/config"
	"google.golang.org/api/sheets/v4"
)
func GetAllSheets() {
	srv := config.ApiConnect()

	res, err := srv.Spreadsheets.Get(models.SpreadsheetId).Do()
    if err != nil {
        log.Fatalf("Unable to get spreadsheet: %v", err)
    }

    // Get the list of sheets in the spreadsheet.
    sheets := res.Sheets

    // Print the names of all sheets in the spreadsheet.
    for _, sheet := range sheets {
        fmt.Println(sheet.Properties.Title)
    }

}

func GetSheets() []*sheets.Sheet{
	srv := config.ApiConnect()

	res, err := srv.Spreadsheets.Get(models.SpreadsheetId).Do()
    if err != nil {
        log.Fatalf("Unable to get spreadsheet: %v", err)
    }

    // Get the list of sheets in the spreadsheet.
    sheets := res.Sheets

    // Print the names of all sheets in the spreadsheet.
    return sheets
}