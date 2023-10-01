package sheethandlers

import (
	"fmt"

	"github.com/spf13/cobra"
)
var (
	sheetName string
	newName string

)
var SheetsCmd = &cobra.Command{
	Use:   "sheets",
	Short: "CRUD operations with your sheets.",
	Long: `CRUD operations with your sheets.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CRUD operations with your sheets. Call `create`, `rename`, `get` or `delete` command.")
	},
}
