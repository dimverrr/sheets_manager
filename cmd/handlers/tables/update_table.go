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

var tableUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update table in your sheet.",
	Long: `Update table in your sheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		UpdateTable()
	},
}

func init() {
	TablesCmd.AddCommand(tableUpdateCmd)

	tableUpdateCmd.Flags().StringVarP(&sheetName, "name", "n", "", "name of the sheet")
	tableUpdateCmd.Flags().StringVarP(&column, "column", "c", "", "column and cell for updating values `Example : A1 `")

	tableUpdateCmd.MarkFlagRequired("name")
	tableUpdateCmd.MarkFlagRequired("column")
}

func UpdateTable() {
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
		fmt.Println("There is no sheet with this name")
		os.Exit(1)
	}
	
	readRange := sheetName + `!` + column

	resp, err := srv.Spreadsheets.Values.Get(id, readRange).Do()
	if err != nil {
			log.Fatalf("Unable to retrieve data from the sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("This cell is empty. Do you want to write data to this column? Print `yes` or `no`.")
		fmt.Scan(&decision)
		if decision == "no"{
			os.Exit(1)
		}
	}

	row = &sheets.ValueRange{
		Values: [][]interface{}{},
	}

	for {
		fmt.Println("Enter new value or print `end` to write cells.")
		fmt.Scan(&value)


		if value == "end" {
			break
		}else{
			row.Values = append(row.Values, []interface{}{value})
		}
	}
	
	_, err = srv.Spreadsheets.Values.Update(id, readRange, row).ValueInputOption("RAW").Do()
	if err != nil {
			log.Fatalf("Unable to update data in the sheet: %v", err)
	}
	
	fmt.Println("Your table was updated successfully.")
	
}