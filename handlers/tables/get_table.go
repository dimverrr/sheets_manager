package tables

import (
	"fmt"
	"log"
	"os"

	sheethandlers"sheets_manager/handlers/sheets"
	"sheets_manager/setup/config"

)

func GetTable() {
	srv := config.ApiConnect()
	var sheetName string
	var column1 string
	var column2 string

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

	fmt.Println("Enter start Column you want to get `Example : A1 `")
	fmt.Scan(&column1)

	fmt.Println("Enter end Column  you want to add get `Example : E `")
	fmt.Scan(&column2)

	readRange := sheetName + `!` + column1 + ":" + column2

	resp, err := srv.Spreadsheets.Values.Get("1anvMrlKNxPsu5QXwgd3BeVg4QNs8DZbwXWIlg18Qf7g", readRange).Do()
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
