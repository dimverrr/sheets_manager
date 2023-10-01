package sheethandlers

import (
	"fmt"
	"log"
	"os"
	"sheets_manager/cmd/handlers/spreadsheets"
	"sheets_manager/setup/config"

	"github.com/spf13/cobra"
	"google.golang.org/api/sheets/v4"
)

var deleteSheetsCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete sheet from your spreadsheet.",
	Long: `Delete sheet from your spreadsheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		DeleteSheet()
	},
}

func init() {
	SheetsCmd.AddCommand(deleteSheetsCmd)
	deleteSheetsCmd.Flags().StringVarP(&sheetName, "name", "n", "", "name for sheet")
	deleteSheetsCmd.MarkFlagRequired("name")
}

func DeleteSheet() {
	srv := config.ApiConnect()
	id := spreadsheethandlers.CheckId()
	sheetsArr := GetSheets()

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
		
					_, err := srv.Spreadsheets.BatchUpdate(id, &batchUpdate).Do()
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