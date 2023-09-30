package sheets

import (
	"fmt"
	"log"
	"os"
	"sheets_manager/models"
	"sheets_manager/setup/config"

	"google.golang.org/api/sheets/v4"
)

var newName string
func RenameSheet() {
	srv := config.ApiConnect()

	sheetsArr1 := GetSheets()
	fmt.Println("Enter sheet name you want to rename")
	fmt.Scan(&sheetName)
	found := false

	for _, sheet := range sheetsArr1{
		if sheet.Properties.Title == sheetName{
			found = true
			fmt.Println("Enter new name for sheet")
			fmt.Scan(&newName)
			batchUpdate := sheets.BatchUpdateSpreadsheetRequest{
				Requests: []*sheets.Request{{
					UpdateSheetProperties: &sheets.UpdateSheetPropertiesRequest{
						Properties: &sheets.SheetProperties{
							Title: newName,
							SheetId: sheet.Properties.SheetId,
						},
						Fields: "title",
					},
				}},
			}

			_, err := srv.Spreadsheets.BatchUpdate(models.SpreadsheetId, &batchUpdate).Do()
			if err != nil {
				log.Fatal("Impossible to rename sheet")
			}

			fmt.Println("Sheet was renamed successfully")
			break
		}
	}

	if !found {
		fmt.Println("There is no sheet with such name")
		os.Exit(1)
	}
	

}

