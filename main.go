package main

import (
	"Shiner-PST-2025/pstformat"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Shiner-Dist/shinerutils"
)

func main() {

	start := time.Now()

	fileName := shinerutils.GetCurrentFileName()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%s has stopped execution due to an error extracting the current working directory. See the following error message: %v", fileName, err)
	}

	parent := filepath.Dir(cwd)

	cfgPath := filepath.Join(parent, "config.json")

	cfg, err := shinerutils.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("%s has stopped execution due to an error loading the config file into memory. The file path for the config is %s. See the following error message: %v", fileName, cfgPath, err)
	}

	logDB, err := shinerutils.OpenLogDBConnection(cfg.LogConfig)
	if err != nil {
		log.Fatalf("The processPST function has stopped executing due to an error connecting to the logging database. See the following error: %v", err)
	}
	defer logDB.Close()

	logParams := shinerutils.CreateLogParams(fileName, "SQL-PST-0001", "Execute stored procedure to update [dbo].[fShipped PST OG]")

	execParams := shinerutils.CreateExecParams(cfg.ExecConfig, "WHServer", "Warehouse", 1433)

	queryParams := shinerutils.CreateQueryParamsFromExec(execParams, "EXEC [Warehouse].[dbo].[Update fShipped PST OG];")

	logger := shinerutils.NewLogger(start, logParams, logDB)

	rows, err := shinerutils.ExecSQLProcedure(queryParams, logger)

	if err != nil {
		log.Fatalf("An errror has occurred whilst executing the procedure to update [fShipped PST OG] table. The transaction has been rolled back. See the following error message: %v", err)
	}

	fmt.Printf("The [fShipped PST OG] table has been updated. %d rows were affected.\n", rows)

	// Update Merch PST

	queryParams = shinerutils.CreateQueryParamsFromExec(execParams, "EXEC [Warehouse].[dbo].[Update fShipped Qty];")

	logger.UpdateLogAction("Execute stored procedure to update [dbo].[fShipped Qty]")

	rows, err = shinerutils.ExecSQLProcedure(queryParams, logger)

	if err != nil {
		log.Fatalf("An errror has occurred whilst executing the procedure to update [fShipped Qty] table. The transaction has been rolled back. See the following error message: %v", err)
	}

	fmt.Printf("The [fShipped Qty] table has been updated. %d rows were affected.\n", rows)

	logger.UpdateLogAction("Read in Shiner Combined PST V2.sql query bytes.")

	queryBytes, err := os.ReadFile(filepath.Join("sql", "Shiner Combined PST V2.sql"))
	if err != nil {
		logger.Error(err, nil, "The processPST function has stopped becuase an error has occurred while reading in the sql file query text from Shiner Combined PST V2.sql. See the following error: %v", err)
	}

	queryText := string(queryBytes)

	queryParams = shinerutils.CreateQueryParamsFromExec(execParams, queryText)

	logger.UpdateLogAction("Execute sql select query statement to return the table of data for the Ltd & B.V combined pst.")

	combinedTable, err := shinerutils.FetchSQLData(queryParams, logger)

	if err != nil {
		log.Fatalf("An errror has occurred whilst fetching sql data for the combined PST. See the following error message: %v", err)
	}

	logger.UpdateLogAction("Read in Shiner LLC PST.sql query bytes.")

	queryBytes, err = os.ReadFile(filepath.Join("sql", "Shiner LLC PST.sql"))
	if err != nil {
		logger.Error(err, nil, "The processPST function has stopped becuase an error has occurred while reading in the sql file query text from Shiner LLC PST.sql. See the following error: %v", err)
	}

	queryText = string(queryBytes)

	queryParams = shinerutils.CreateQueryParamsFromExec(execParams, queryText)

	logger.UpdateLogAction("Execute sql select query statement to return the table of data for the LLC pst.")

	llcTable, err := shinerutils.FetchSQLData(queryParams, logger)

	if err != nil {
		log.Fatalf("An errror has occurred whilst fetching sql data for the Merch PST. See the following error message: %v", err)
	}

	logger.UpdateLogAction("Read in Shiner Merch PST.sql query bytes.")

	queryBytes, err = os.ReadFile(filepath.Join("sql", "Shiner Merch PST.sql"))
	if err != nil {
		logger.Error(err, nil, "The processPST function has stopped becuase an error has occurred while reading in the sql file query text from Shiner LLC PST.sql. See the following error: %v", err)
	}

	queryText = string(queryBytes)

	queryParams = shinerutils.CreateQueryParamsFromExec(execParams, queryText)

	logger.UpdateLogAction("Execute sql select query statement to return the table of data for the merchandising pst.")

	merchTable, err := shinerutils.FetchSQLData(queryParams, logger)

	if err != nil {
		log.Fatalf("An errror has occurred whilst fetching sql data for the Merch PST. See the following error message: %v", err)
	}

	sheet1 := shinerutils.NewSheetDefinition("Ltd & B.V", combinedTable, nil, pstformat.FormatCombinedPST)
	sheet2 := shinerutils.NewSheetDefinition("LLC", llcTable, nil, pstformat.FormatLLCPST)
	sheet3 := shinerutils.NewSheetDefinition("Merch", merchTable, nil, pstformat.FormatMerchPST)

	var worksheets []shinerutils.SheetDefinition

	worksheets = shinerutils.AddSheetDefinition(worksheets, sheet1)
	worksheets = shinerutils.AddSheetDefinition(worksheets, sheet2)
	worksheets = shinerutils.AddSheetDefinition(worksheets, sheet3)

	logger.UpdateLogAction("Write excel sheets data to a memory buffer.")

	excelBuffer, err := shinerutils.WriteExcelBufferWithSheets(worksheets, logger)

	if err != nil {
		log.Fatalf("The processPST function has stopped becuase an errror has occurred whilst writing data to an excel buffer. See the following error message: %v", err)
	}

	formattedTime := time.Now().Format("02/01/2006 15:04:05")
	fileTime := time.Now().Format("02.01.2006 15.04.05")

	logger.UpdateLogAction("Read in the email html template bytes.")

	tmplBytes, err := os.ReadFile("template.html")
	if err != nil {
		logger.Error(err, nil, "The processPST function has stopped becuase an error has occurred while reading in the html email template. See the following error: %v", err)
	}

	logger.UpdateLogAction("Create email template from bytes.")

	tmpl, err := template.New("email").Parse(string(tmplBytes))
	if err != nil {
		logger.Error(err, nil, "The processPST function has stopped becuase an error has occurred while parsing email template bytes and converting into string. See the following error: %v", err)
	}

	timeStamp := shinerutils.EmailTemplateData{
		DynamicText: formattedTime,
	}

	logger.UpdateLogAction("Create email html template variable for sending out.")

	var htmlBody bytes.Buffer
	err = tmpl.Execute(&htmlBody, timeStamp)
	if err != nil {
		logger.Error(err, nil, "The function has stopped becuase an error has occurred while creating html string variable. See the following error: %v", err)
	}

	emailBody := htmlBody.String()

	var attachments []shinerutils.BufferAttachment

	attachments = shinerutils.AddBufferAttachment(
		attachments,
		fmt.Sprintf("Shiner Potential Sell Through - %s.xlsx", fileTime),
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		excelBuffer.Bytes(),
	)

	logger.UpdateLogAction("Send out potential sell through report email with report attachment.")

	err = shinerutils.SendGraphEmail(cfg.GraphConfig, emailBody, []string{"potentialsellthrough@shinerltd.onmicrosoft.com"}, fmt.Sprintf("Shiner Potential Sell Through - %s", formattedTime), nil, attachments, nil, nil, logger)

	if err != nil {
		logger.Error(err, nil, "The processPST function has stopped becuase an error has occurred while parsing email template bytes and converting into string. See the following error: %v", err)
	}

	logger.Info(nil, "The function to create and send out the shiner pst report has successfully completed.")
}
