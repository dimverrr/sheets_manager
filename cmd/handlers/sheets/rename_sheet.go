package sheethandlers

import (
	"fmt"
	"log"
	"os"
	spreadsheethandlers "sheets_manager/cmd/handlers/spreadsheets"
	"sheets_manager/setup/config"

	"github.com/spf13/cobra"
	"google.golang.org/api/sheets/v4"
)

var renameSheetCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename your sheet fromspreadsheet.",
	Long: `Rename your sheet fromspreadsheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		RenameSheet()
	},
}

func init() {
	SheetsCmd.AddCommand(renameSheetCmd)

	renameSheetCmd.Flags().StringVarP(&sheetName, "name", "n", "", "name of sheet")
	renameSheetCmd.Flags().StringVarP(&newName, "newName", "w", "", "new name for sheet")
	renameSheetCmd.MarkFlagRequired("name")
	renameSheetCmd.MarkFlagRequired("newName")
}

func RenameSheet() {

	srv := config.ApiConnect()
	id := spreadsheethandlers.CheckId()
	sheetsArr := GetSheets()

	found := false
	for _, sheet := range sheetsArr{
		if sheet.Properties.Title == sheetName{
			found = true
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

			_, err := srv.Spreadsheets.BatchUpdate(id, &batchUpdate).Do()
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