package tables

import (
	"fmt"
	"log"
	"os"
	sheethandlers "sheets_manager/cmd/handlers/sheets"
	spreadsheethandlers "sheets_manager/cmd/handlers/spreadsheets"
	"sheets_manager/setup/config"

	"github.com/spf13/cobra"
	"google.golang.org/api/sheets/v4"
)

var tableDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete table from your sheet.",
	Long: `Delete table from your sheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		DeleteTable()
	},
}

func init() {
	TablesCmd.AddCommand(tableDeleteCmd)

	tableDeleteCmd.Flags().StringVarP(&sheetName, "name", "n", "", "name of sheet")
	tableDeleteCmd.Flags().StringVarP(&column1, "column1", "s", "", "start column and cell for deleting values`Example: A1 `")
	tableDeleteCmd.Flags().StringVarP(&column2, "column2", "e", "", "end column for deleting values`Example: E `")

	tableDeleteCmd.MarkFlagRequired("name")
	tableCreateCmd.MarkFlagRequired("column1")
	tableCreateCmd.MarkFlagRequired("column2")
}

func DeleteTable() {
	srv := config.ApiConnect()
	id := spreadsheethandlers.CheckId()
	sheetsArr := sheethandlers.GetSheets()

	found := false
	for _, sheet := range sheetsArr {
		if sheet.Properties.Title == sheetName {
			found = true
			break
		}
	}
	if !found {
		fmt.Println("There is no sheet with such name")
		os.Exit(1)
	}

	readRange := sheetName + `!` + column1 + ":" + column2
	batchUpdate := sheets.ClearValuesRequest{}
	
	_, err := srv.Spreadsheets.Values.Clear(id, readRange, &batchUpdate).Do()
	if err != nil {
			log.Fatalf("Unable to delete data from sheet: %v", err)
	}

	fmt.Println("Your table was deleted successfully")

}