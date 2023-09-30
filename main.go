package main

import (
	"sheets_manager/handlers/sheets"
	"sheets_manager/handlers/spreadsheets"
)

func main() {
	spreadsheets.CreateSpreadsheet()
	sheets.CreateSheet()
	sheets.RenameSheet()
}
