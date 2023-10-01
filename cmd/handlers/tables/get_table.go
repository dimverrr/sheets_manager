package tables

import (
	"fmt"
	"log"
	"os"
	sheethandlers "sheets_manager/cmd/handlers/sheets"
	spreadsheethandlers "sheets_manager/cmd/handlers/spreadsheets"
	"sheets_manager/setup/config"

	"github.com/spf13/cobra"
)

var tableGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get table from your sheet.",
	Long: `Get table from your sheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		GetTable()
	},
}

func init() {
	TablesCmd.AddCommand(tableGetCmd)

	tableGetCmd.Flags().StringVarP(&sheetName, "name", "n", "", "name for sheet")
	tableGetCmd.Flags().StringVarP(&column1, "column1", "s", "", "start column and cell for getting values `Example: A1 `")
	tableGetCmd.Flags().StringVarP(&column2, "column2", "e", "", "end column for getting values `Example: E `")

	tableGetCmd.MarkFlagRequired("name")
	tableGetCmd.MarkFlagRequired("column1")
	tableGetCmd.MarkFlagRequired("column2")
}

func GetTable() {
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

	resp, err := srv.Spreadsheets.Values.Get(id, readRange).Do()
	if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found")
		os.Exit(1)
	} else {
		for _, row := range resp.Values{
			fmt.Println(row...)
		}
	}

}
