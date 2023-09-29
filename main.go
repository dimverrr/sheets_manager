package main

import (
	"fmt"
	"sheets_manager/handlers/sheets"
	"sheets_manager/handlers/spreadsheets"
	"sheets_manager/models"
)

func main() {
	spreadsheets.CreateSpreadsheet()
	fmt.Println(models.SpreadsheetId)
	spreadsheets.UpdateSpreadsheet()
	sheets.CreateSheet()
}
