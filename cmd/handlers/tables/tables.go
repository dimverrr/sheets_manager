package tables

import (
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/api/sheets/v4"
)
var (
	
	sheetName string
	value string
	column string
	decision string
	row *sheets.ValueRange
	column1 string
	column2 string
)
var TablesCmd = &cobra.Command{
	Use:   "tables",
	Short: "CRUD operations with your tables.",
	Long: `CRUD operations with your tables.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CRUD operations with your tables. Call `create`, `update`, `get` or `delete` command.")
	},
}

