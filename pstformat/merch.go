package pstformat

import (
	"fmt"
	_ "image/png"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

func FormatMerchPST(f *excelize.File, sheetName string) (err error) {
	/*
		The below section of code defines the styles to be used in the worksheet formatting.

			The below list shows the index values of different types of excel border styles.

			   Style Code	Description	Visual Effect
			   0	No border	No visible border
			   1	Thin	A thin line
			   2	Medium	A thicker line
			   3	Dashed	Small dashes
			   4	Dotted	Dotted line
			   5	Thick	Very thick line
			   6	Double	Two thin lines with a small gap (classic Excel double)
			   7	Hair	Extremely thin line
			   8	Medium dashed	Medium weight dashed line
			   9	Dash dot	Dash-dot pattern
			   10	Medium dash dot	Thicker dash-dot
			   11	Dash dot dot	Dash-dot-dot pattern
			   12	Medium dash dot dot	Medium weight dash-dot-dot
			   13	Slanted dash dot	Diagonal dash-dot line
	*/
	leftHeaderStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Color: "#FFFFFF", Size: 12},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#757171"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 5},
			{Type: "bottom", Color: "#000000", Style: 5},
			{Type: "left", Color: "#000000", Style: 5},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		fmt.Println(err)
	}

	midHeaderStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Color: "#FFFFFF", Size: 12},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#757171"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 5},
			{Type: "bottom", Color: "#000000", Style: 5},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		fmt.Println(err)
	}

	rightHeaderStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Color: "#FFFFFF", Size: 12},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#757171"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 5},
			{Type: "bottom", Color: "#000000", Style: 5},
			{Type: "right", Color: "#000000", Style: 5},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		fmt.Println(err)
	}

	rightMostBorder, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 5},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	bottomMostBorder, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 5},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	trueStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#C6EFCE"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#006100",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	trueLeftStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#C6EFCE"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#006100",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	trueRightStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#C6EFCE"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#006100",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	falseStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFC7CE"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#9C0006",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	falseLeftStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFC7CE"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#9C0006",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	falseRightStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFC7CE"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#9C0006",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	linkStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:      false,
			Underline: "single",
			Color:     "#305496",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	rowStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	itemCodeStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#AEAAAA"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#FFFFFF",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
			{Type: "left", Color: "#000000", Style: 5},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	dividerStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	dividerLeftStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	dividerRightStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	gbpNumFormat := "_-£* #,##0.00_-;-£* #,##0.00_-;_-£* \"-\"??_-;_-@_-"

	gbpLeftStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &gbpNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	gbpRightStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &gbpNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	gbpMidStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &gbpNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	eurNumFormat := "_-[$€-x-euro2] * #,##0.00_-;-[$€-x-euro2] * #,##0.00_-;_-[$€-x-euro2] * \"-\"??_-;_-@_-"

	eurLeftStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &eurNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	eurRightStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &eurNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	eurMidStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &eurNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	usdNumFormat := "_-[$$-en-US]* #,##0.00_ ;_-[$$-en-US]* -#,##0.00 ;_-[$$-en-US]* \"-\"??_ ;_-@_ "

	usdLeftStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &usdNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	usdRightStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &usdNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	usdMidStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &usdNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	decimalFormat := "#,##0.00"

	decimalLeftStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#EAEAEA"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#000000",
		},
		CustomNumFmt: &decimalFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	decimalRightStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#EAEAEA"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#000000",
		},
		CustomNumFmt: &decimalFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	decimalLtdLeftStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#D1EFFA"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#003DAB",
		},
		CustomNumFmt: &decimalFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	decimalLtdRightStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#D1EFFA"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#003DAB",
		},
		CustomNumFmt: &decimalFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	decimalBVLeftStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FCEDD6"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#F96F00",
		},
		CustomNumFmt: &decimalFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	decimalBVRightStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FCEDD6"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#F96F00",
		},
		CustomNumFmt: &decimalFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	decimalLLCLeftStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#E2EFDA"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#087329",
		},
		CustomNumFmt: &decimalFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "left", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	decimalLLCRightStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#E2EFDA"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#087329",
		},
		CustomNumFmt: &decimalFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	// The below code starts to actually set the worksheet formatting.
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
	}

	lastRow := strconv.Itoa(len(rows))
	lastRowPlus1 := strconv.Itoa(len(rows) + 1)

	// The below section of code formats the header for the report.
	err = f.SetCellStyle(sheetName, "A1", "A1", leftHeaderStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "B1", "EI1", midHeaderStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "EJ1", "EJ1", rightHeaderStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below section formats the item code column on the far left.
	err = f.SetCellStyle(sheetName, "A2", "A"+lastRow, itemCodeStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below section of code sets the far right and bottom bottom border.
	err = f.SetCellStyle(sheetName, "EK1", "EK"+lastRow, rightMostBorder)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "A"+lastRowPlus1, "EJ"+lastRowPlus1, bottomMostBorder)
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets the main row standard format.
	err = f.SetCellStyle(sheetName, "B2", "EI"+lastRow, rowStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets the dashed divider column styles grouping columns together.
	err = f.SetCellStyle(sheetName, "R2", "R"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "T2", "T"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "U2", "U"+lastRow, dividerStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AD2", "AD"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AE2", "AE"+lastRow, gbpLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AF2", "AF"+lastRow, gbpMidStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AG2", "AG"+lastRow, gbpRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AH2", "AH"+lastRow, eurLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AI2", "AI"+lastRow, eurMidStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AJ2", "AJ"+lastRow, eurRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AK2", "AK"+lastRow, usdLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AL2", "AL"+lastRow, usdMidStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AM2", "AM"+lastRow, usdRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BA2", "BA"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BB2", "BB"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BC2", "BC"+lastRow, decimalLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BD2", "BD"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BE2", "BE"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BF2", "BF"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BG2", "BI"+lastRow, dividerStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BK2", "BK"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BM2", "BM"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AN2", "AN"+lastRow, dividerStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AO2", "AO"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AZ2", "AZ"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BA2", "BA"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BB2", "BB"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BC2", "BC"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BD2", "BD"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BE2", "BE"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BF2", "BF"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BG2", "BI"+lastRow, dividerStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BK2", "BK"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BM2", "BM"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below section starts to layout the Ltd specific columns and formats
	err = f.SetCellStyle(sheetName, "BN2", "BN"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BY2", "BY"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "BZ2", "BZ"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CA2", "CA"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CB2", "CB"+lastRow, decimalLtdLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CC2", "CC"+lastRow, decimalLtdRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CD2", "CD"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CE2", "CE"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CF2", "CF"+lastRow, dividerStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CJ2", "CJ"+lastRow, decimalLtdRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CL2", "CL"+lastRow, decimalLtdRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below section starts to layout the B.V specific columns and formats
	err = f.SetCellStyle(sheetName, "CM2", "CM"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CX2", "CX"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CY2", "CY"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "CZ2", "CZ"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DA2", "DA"+lastRow, decimalBVLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DB2", "DB"+lastRow, decimalBVRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DC2", "DC"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DD2", "DD"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DE2", "DE"+lastRow, dividerStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DI2", "DI"+lastRow, decimalBVRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DK2", "DK"+lastRow, decimalBVRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below section starts to layout the LLC specific columns and formats
	err = f.SetCellStyle(sheetName, "DL2", "DL"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DW2", "DW"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DX2", "DX"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DY2", "DY"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "DZ2", "DZ"+lastRow, decimalLLCLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "EA2", "EA"+lastRow, decimalLLCRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "EB2", "EB"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "EC2", "EC"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "ED2", "ED"+lastRow, dividerStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "EH2", "EH"+lastRow, decimalLLCRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "EJ2", "EJ"+lastRow, decimalLLCRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below code starts a loop through all the rows on the worksheet for condition based operations.
	for rowIndex := range rows {

		spLinkCell, err := excelize.CoordinatesToCellName(21, rowIndex+2)
		if err != nil {
			fmt.Println(err)
		}

		linkVal, err := f.GetCellValue(sheetName, spLinkCell)
		if err != nil {
			fmt.Println(err)
		}

		itemCodeCell, err := excelize.CoordinatesToCellName(1, rowIndex+2)
		if err != nil {
			fmt.Println(err)
		}

		itemCodeVal, err := f.GetCellValue(sheetName, itemCodeCell)
		if err != nil {
			fmt.Println(err)
		}

		// The below section of code converts the sharepoint image url into a clickable hyperlink.
		if linkVal != "" {
			err = f.SetCellFormula(sheetName, spLinkCell, "=HYPERLINK(\""+linkVal+"\", \""+itemCodeVal+"\")")
			if err != nil {
				fmt.Println(err)
			}

			err = f.SetCellStyle(sheetName, spLinkCell, spLinkCell, linkStyle)
			if err != nil {
				fmt.Println(err)
			}
		}

		// The below section of code applys green/red colour formats for boolean report columns.
		for i := 1; i <= 9; i++ {
			cell, err := excelize.CoordinatesToCellName(21+i, rowIndex+2)
			if err != nil {
				fmt.Println(err)
			}
			val, err := f.GetCellValue(sheetName, cell)
			if err != nil {
				fmt.Println(err)
			}
			if val == "TRUE" && i == 1 {
				err = f.SetCellStyle(sheetName, cell, cell, trueLeftStyle)
				if err != nil {
					fmt.Println(err)
				}
			} else if val == "FALSE" && i == 1 {
				err = f.SetCellStyle(sheetName, cell, cell, falseLeftStyle)
				if err != nil {
					fmt.Println(err)
				}
			}
			if val == "TRUE" && i == 9 {
				err = f.SetCellStyle(sheetName, cell, cell, trueRightStyle)
				if err != nil {
					fmt.Println(err)
				}
			} else if val == "FALSE" && i == 9 {
				err = f.SetCellStyle(sheetName, cell, cell, falseRightStyle)
				if err != nil {
					fmt.Println(err)
				}
			} else if val == "TRUE" && i > 1 && i < 9 {
				err = f.SetCellStyle(sheetName, cell, cell, trueStyle)
				if err != nil {
					fmt.Println(err)
				}
			} else if val == "FALSE" && i > 1 && i < 9 {
				err = f.SetCellStyle(sheetName, cell, cell, falseStyle)
				if err != nil {
					fmt.Println(err)
				}
			}

			for i := 1; i <= 3; i += 2 {
				cell, err := excelize.CoordinatesToCellName(61+i, rowIndex+2)
				if err != nil {
					fmt.Println(err)
				}
				val, err := f.GetCellValue(sheetName, cell)
				if err != nil {
					fmt.Println(err)
				}
				if val == "TRUE" && i == 1 {
					err = f.SetCellStyle(sheetName, cell, cell, trueLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				} else if val == "FALSE" && i == 1 {
					err = f.SetCellStyle(sheetName, cell, cell, falseLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				}
				if val == "TRUE" && i == 3 {
					err = f.SetCellStyle(sheetName, cell, cell, trueLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				} else if val == "FALSE" && i == 3 {
					err = f.SetCellStyle(sheetName, cell, cell, falseLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
			for i := 1; i <= 3; i += 2 {
				cell, err := excelize.CoordinatesToCellName(86+i, rowIndex+2)
				if err != nil {
					fmt.Println(err)
				}
				val, err := f.GetCellValue(sheetName, cell)
				if err != nil {
					fmt.Println(err)
				}
				if val == "TRUE" && i == 1 {
					err = f.SetCellStyle(sheetName, cell, cell, trueLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				} else if val == "FALSE" && i == 1 {
					err = f.SetCellStyle(sheetName, cell, cell, falseLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				}
				if val == "TRUE" && i == 3 {
					err = f.SetCellStyle(sheetName, cell, cell, trueLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				} else if val == "FALSE" && i == 3 {
					err = f.SetCellStyle(sheetName, cell, cell, falseLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
			for i := 1; i <= 3; i += 2 {
				cell, err := excelize.CoordinatesToCellName(111+i, rowIndex+2)
				if err != nil {
					fmt.Println(err)
				}
				val, err := f.GetCellValue(sheetName, cell)
				if err != nil {
					fmt.Println(err)
				}
				if val == "TRUE" && i == 1 {
					err = f.SetCellStyle(sheetName, cell, cell, trueLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				} else if val == "FALSE" && i == 1 {
					err = f.SetCellStyle(sheetName, cell, cell, falseLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				}
				if val == "TRUE" && i == 3 {
					err = f.SetCellStyle(sheetName, cell, cell, trueLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				} else if val == "FALSE" && i == 3 {
					err = f.SetCellStyle(sheetName, cell, cell, falseLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
			for i := 1; i <= 3; i += 2 {
				cell, err := excelize.CoordinatesToCellName(136+i, rowIndex+2)
				if err != nil {
					fmt.Println(err)
				}
				val, err := f.GetCellValue(sheetName, cell)
				if err != nil {
					fmt.Println(err)
				}
				if val == "TRUE" && i == 1 {
					err = f.SetCellStyle(sheetName, cell, cell, trueLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				} else if val == "FALSE" && i == 1 {
					err = f.SetCellStyle(sheetName, cell, cell, falseLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				}
				if val == "TRUE" && i == 3 {
					err = f.SetCellStyle(sheetName, cell, cell, trueLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				} else if val == "FALSE" && i == 3 {
					err = f.SetCellStyle(sheetName, cell, cell, falseLeftStyle)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}

	// The below section of code adjusts column widths where required.
	const colWidthAdjust = 0.78

	err = f.SetColWidth(sheetName, "A", "A", 14.57+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "B", "B", 26.86+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "C", "C", 13.33+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "D", "D", 18.29+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "E", "E", 16.86+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "F", "F", 13.14+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "G", "G", 23.14+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "H", "I", 32+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "J", "J", 7.86+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "K", "K", 7.71+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "L", "L", 8+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "M", "N", 9+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "O", "O", 7.14+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "P", "P", 10+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "Q", "Q", 14.57+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "Q", "Q", 30+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "R", "R", 14.57+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "S", "S", 14.57+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "T", "T", 30+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "U", "U", 13.67+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "AE", "AM", 8.78+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "AO", "AZ", 6.44+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "BN", "BY", 5.89+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "CM", "CX", 6.11+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "DL", "DW", 6.11+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}
	// 	The below section of code inserts an empty column in the first position.
	err = f.InsertCols(sheetName, "A", 1)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "A", "A", 1+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	// The below section of code inserts 2 empty rows above the header row.
	err = f.InsertRows(sheetName, 1, 2)
	if err != nil {
		fmt.Println(err)
	}

	// The below section of code inserts the shiner group logo at the top left of the worksheet.
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	parent := filepath.Dir(cwd)

	logoFile, err := os.ReadFile(filepath.Join(parent, "Shiner-PST-2025", "logos", "SHINER_LOGO_BLK_GEN.png"))
	if err != nil {
		fmt.Println(err)
	}

	err = f.AddPictureFromBytes(sheetName, "B1", &excelize.Picture{
		Extension: ".png",
		File:      logoFile,
		Format: &excelize.GraphicOptions{
			AltText: "Shiner Group Logo",
			ScaleX:  0.05,
			ScaleY:  0.05,
			OffsetX: 20,
			OffsetY: 7,
		}})

	if err != nil {
		fmt.Println(err)
	}

	// The below section of code creates the textbox report title.
	formattedTime := time.Now().Format("02/01/2006 15:04:05")

	err = f.AddShape(sheetName,
		&excelize.Shape{
			Cell: "C1",
			Type: "rect",
			Paragraph: []excelize.RichTextRun{
				{
					Text: "Shiner Ltd, B.V & LLC PST " + formattedTime,
					Font: &excelize.Font{
						Bold:   true,
						Italic: true,
						Size:   14,
						Color:  "#000000",
					},
				},
			},
			Width:  450,
			Height: 23,
			Format: excelize.GraphicOptions{
				OffsetX: 3,
				OffsetY: 3,
			},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	// Set the outlining for true/false columns.
	err = f.SetColOutlineLevel(sheetName, "W", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "X", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "Y", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "Z", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AA", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AB", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AC", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AD", 2)
	if err != nil {
		fmt.Println(err)
	}

	// Set the outlining for pricing columns.
	err = f.SetColOutlineLevel(sheetName, "AF", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AG", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AH", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AI", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AJ", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AK", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AL", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AM", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AN", 2)
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets the outlining the group shipped columns.
	err = f.SetColOutlineLevel(sheetName, "AP", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AQ", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AR", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AS", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AT", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AU", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AV", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AW", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AX", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AY", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "AZ", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BA", 2)
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets the outlining the Ltd shipped columns.
	err = f.SetColOutlineLevel(sheetName, "BO", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BP", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BQ", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BR", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BS", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BS", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BT", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BU", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BV", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BW", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BX", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BY", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "BZ", 2)
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets the outlining the B.V shipped columns.
	err = f.SetColOutlineLevel(sheetName, "CN", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CO", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CP", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CQ", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CR", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CS", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CT", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CU", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CV", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CW", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CX", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "CY", 2)
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets the outlining the B.V shipped columns.
	err = f.SetColOutlineLevel(sheetName, "DM", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DN", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DO", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DP", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DQ", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DR", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DS", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DT", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DU", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DV", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DW", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "DX", 2)
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets frozen panes for the worksheet.
	err = f.SetPanes(sheetName, &excelize.Panes{
		Freeze:      true,
		Split:       false,
		XSplit:      2,
		YSplit:      3,
		TopLeftCell: "C4",
		ActivePane:  "bottomRight",
	})
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets the worksheet tab colour.
	index := 5
	err = f.SetSheetProps(sheetName, &excelize.SheetPropsOptions{
		TabColorIndexed: &index,
	})
	if err != nil {
		fmt.Println(err)
	}

	return nil

}
