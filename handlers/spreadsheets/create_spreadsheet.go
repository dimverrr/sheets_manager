package spreadsheethandlers

import (
	"log"
	"sheets_manager/setup/config"
	"google.golang.org/api/sheets/v4"
)

func CreateSpreadsheet() {
	srv := config.ApiConnect()

	spreadsheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: "My New Spreadsheet1",
		},
	}

	_, err := srv.Spreadsheets.Create(spreadsheet).Do()
	if err != nil {
		log.Fatalf("Unable to create spreadsheet: %v", err)
	}


}
