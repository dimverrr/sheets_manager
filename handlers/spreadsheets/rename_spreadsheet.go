package spreadsheethandlers

import (
	"fmt"
	"log"
	"sheets_manager/setup/config"
	"google.golang.org/api/sheets/v4"
)

func UpdateSpreadsheet() {
	
	srv := config.ApiConnect()
	
    // Create a BatchUpdateSpreadsheetRequest object.
    batchUpdateSpreadsheetRequest := &sheets.BatchUpdateSpreadsheetRequest{
        Requests: []*sheets.Request{{
            UpdateSpreadsheetProperties: &sheets.UpdateSpreadsheetPropertiesRequest{
                Properties: &sheets.SpreadsheetProperties{
                    Title: "New spreadsheet title",
                },
				Fields: "title",
            },
        }},
    }

    // Call the BatchUpdate() method on the SheetsService object, passing in the BatchUpdateSpreadsheetRequest object.
    _, err := srv.Spreadsheets.BatchUpdate("1anvMrlKNxPsu5QXwgd3BeVg4QNs8DZbwXWIlg18Qf7g", batchUpdateSpreadsheetRequest).Do()
    if err != nil {
        log.Fatalf("Unable to update spreadsheet title: %v", err)
    }

    fmt.Println("Spreadsheet title updated successfully!")
}