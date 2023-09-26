package handlersspreadsheets

import (
	"context"
	"log"
	"os"
	"sheets_manager/token"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func CreateSheet() string{

	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
	}
	
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := token.GetClient(config)
	
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
			log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	
	
	spreadsheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: "My New Spreadsheet1",
		},
	}
	
	
	resp , err := srv.Spreadsheets.Create(spreadsheet).Do()
	if err != nil {
		log.Fatalf("Unable to create spreadsheet: %v", err)
	}
	return resp.SpreadsheetId
	}

