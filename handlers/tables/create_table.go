package tables

import (
	"fmt"
	"log"
	"os"

	sheethandlers"sheets_manager/handlers/sheets"
	"sheets_manager/setup/config"

	"google.golang.org/api/sheets/v4"
)

func CreateTable() {
	srv := config.ApiConnect()
	var sheetName string
	var value string
	var column string
	var row *sheets.ValueRange

	row = &sheets.ValueRange{
		Values: [][]interface{}{},
	}

	sheetsArr := sheethandlers.GetSheets()
	fmt.Println("Enter sheet name where you want to create table")
	fmt.Scan(&sheetName)
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

	fmt.Println("Enter Column where you want to add data `Example : A1 `")
	fmt.Scan(&column)

	readRange := sheetName + `!` + column

	resp, err := srv.Spreadsheets.Values.Get("1anvMrlKNxPsu5QXwgd3BeVg4QNs8DZbwXWIlg18Qf7g", readRange).Do()
	if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("Insert header of column")
		fmt.Scan(&value)
		row.Values = append(row.Values, []interface{}{value})
	}


	for {
		fmt.Println("Insert new value or print `end` to write cells")
		fmt.Scan(&value)

		if value == "end" {
			break
		}else{
			row.Values = append(row.Values, []interface{}{value})
		}
	}
	
	_, err = srv.Spreadsheets.Values.Append("1anvMrlKNxPsu5QXwgd3BeVg4QNs8DZbwXWIlg18Qf7g", readRange, row).ValueInputOption("RAW").Do()
	if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

}