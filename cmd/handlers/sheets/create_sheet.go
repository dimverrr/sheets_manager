package sheethandlers

import (
	"fmt"
	"log"
	spreadsheethandlers "sheets_manager/cmd/handlers/spreadsheets"
	"sheets_manager/setup/config"

	"github.com/spf13/cobra"
	"google.golang.org/api/sheets/v4"
)

var createSheetCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new sheet in your spreadsheet.",
	Long: `Create new sheet in your spreadsheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateSheet()
	},
}

func init() {
	SheetsCmd.AddCommand(createSheetCmd)
	createSheetCmd.Flags().StringVarP(&sheetName, "name", "n", "", "name for sheet")
	createSheetCmd.MarkFlagRequired("name")
}

func CreateSheet() {
	srv := config.ApiConnect()
	id := spreadsheethandlers.CheckId()
	
	sheet := sheets.SheetProperties{
		Title: sheetName,
	}

	batchUpdate := sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{{
			AddSheet: &sheets.AddSheetRequest{
				Properties: &sheet,
			},
		}},
	}
	resp, err := srv.Spreadsheets.BatchUpdate(id, &batchUpdate).Do()
	if err != nil {
		log.Fatal("Sheet with this name already exists")
	}

	fmt.Printf("New sheet was created with Title: %v\n", resp.Replies[0].AddSheet.Properties.Title)

}
