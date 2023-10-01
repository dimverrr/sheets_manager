package sheethandlers

import (
	"fmt"
	"log"
	"sheets_manager/setup/config"
	"google.golang.org/api/sheets/v4"
)

func CreateSheet() {
	srv := config.ApiConnect()
	sheet := sheets.SheetProperties{
		Title: "Sheet2",
	}

	batchUpdate := sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{{
			AddSheet: &sheets.AddSheetRequest{
				Properties: &sheet,
			},
		}},
	}

	resp, err := srv.Spreadsheets.BatchUpdate("1anvMrlKNxPsu5QXwgd3BeVg4QNs8DZbwXWIlg18Qf7g", &batchUpdate).Do()
	if err != nil {
		log.Fatalf("Unable to create new sheet: %v", err)
	}

	fmt.Printf("New sheet created with ID: %v\n", resp.Replies[0].AddSheet.Properties.SheetId)

}
