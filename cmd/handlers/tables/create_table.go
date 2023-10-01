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

var tableCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create table in your sheet.",
	Long: `Create table in your sheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateTable()
	},
}

func init() {
	TablesCmd.AddCommand(tableCreateCmd)
	tableCreateCmd.Flags().StringVarP(&sheetName, "name", "n", "", "name for sheet")
	tableCreateCmd.Flags().StringVarP(&column, "column", "c", "", "column and cell for entering values `Example : A1 `")

	tableCreateCmd.MarkFlagRequired("name")
	tableCreateCmd.MarkFlagRequired("column")
}

func CreateTable() {
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

	readRange := sheetName + `!` + column

	resp, err := srv.Spreadsheets.Values.Get(id, readRange).Do()
	if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	row = &sheets.ValueRange{
		Values: [][]interface{}{},
	}
	if len(resp.Values) == 0 {
		fmt.Println("Enter header of column")
		fmt.Scan(&value)
		row.Values = append(row.Values, []interface{}{value})
	}


	for {
		fmt.Println("Enter new value or print `end` to write cells")
		fmt.Scan(&value)

		if value == "end" {
			break
		}else{
			row.Values = append(row.Values, []interface{}{value})
		}
	}
	
	_, err = srv.Spreadsheets.Values.Append(id, readRange, row).ValueInputOption("RAW").Do()
	if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

}