package main

import (
	"sheets_manager/handlers/spreadsheets"
	"sheets_manager/handlers/tables"
)

func main() {
	spreadsheets.CreateSpreadsheet()
	tables.DeleteTable()
}
