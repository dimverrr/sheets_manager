package sheethandlers

import (
	"fmt"
	"log"
	"os"
	"sheets_manager/models"
	"sheets_manager/setup/config"

	"google.golang.org/api/sheets/v4"
)
var sheetName string
func DeleteSheet() {
	srv := config.ApiConnect()
	sheetsArr := GetSheets()

	fmt.Println("Enter sheet name you want to delete")
	fmt.Scan(&sheetName)
	found := false
	for _, sheet := range sheetsArr{
	
		if sheet.Properties.Title == sheetName{
					found = true
					batchUpdate := sheets.BatchUpdateSpreadsheetRequest{
						Requests: []*sheets.Request{{
							DeleteSheet: &sheets.DeleteSheetRequest{
								SheetId: sheet.Properties.SheetId,
							},
						}},
					}
		
					_, err := srv.Spreadsheets.BatchUpdate(models.SpreadsheetId, &batchUpdate).Do()
					if err != nil {
						log.Fatal("Impossible to delete sheet")
					}
		
					fmt.Println("Sheet was deleted successfully")
					break
		}
	}
	
	if !found {
		fmt.Println("There is no sheet with such name")
		os.Exit(1)
	}

}

	
	




