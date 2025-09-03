package pstformat

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	_ "image/png"
	"os"
	"strconv"
	"time"
	"path/filepath"
)

func FormatLLCPST(f *excelize.File, sheetName string) (err error) {
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
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#2DAB66"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 5},
			{Type: "bottom", Color: "#087329", Style: 5},
			{Type: "left", Color: "#087329", Style: 5},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		fmt.Println(err)
	}

	midHeaderStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Color: "#FFFFFF", Size: 12},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#2DAB66"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 5},
			{Type: "bottom", Color: "#087329", Style: 5},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		fmt.Println(err)
	}

	rightHeaderStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Color: "#FFFFFF", Size: 12},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#2DAB66"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 5},
			{Type: "bottom", Color: "#087329", Style: 5},
			{Type: "right", Color: "#087329", Style: 5},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		fmt.Println(err)
	}

	rightMostBorder, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#087329", Style: 5},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	bottomMostBorder, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 5},
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
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "left", Color: "#087329", Style: 3},
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
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "right", Color: "#087329", Style: 3},
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
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "left", Color: "#087329", Style: 3},
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
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "right", Color: "#087329", Style: 3},
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
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	rowStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	itemCodeStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#E2EFDA"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Color: "#087329",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "right", Color: "#087329", Style: 3},
			{Type: "left", Color: "#087329", Style: 5},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	dividerStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "right", Color: "#087329", Style: 3},
			{Type: "left", Color: "#087329", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	dividerLeftStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "left", Color: "#087329", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	dividerRightStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "right", Color: "#087329", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	usdNumFormat := "_-[$$-en-US]* #,##0.00_ ;_-[$$-en-US]* -#,##0.00 ;_-[$$-en-US]* \"-\"??_ ;_-@_ "

	usdLeftStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &usdNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "left", Color: "#087329", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	usdRightStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &usdNumFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "right", Color: "#087329", Style: 3},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	decimalFormat := "#,##0.00"

	decimalRightStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &decimalFormat,
		Border: []excelize.Border{
			{Type: "top", Color: "#087329", Style: 1},
			{Type: "bottom", Color: "#087329", Style: 1},
			{Type: "right", Color: "#087329", Style: 3},
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

	err = f.SetCellStyle(sheetName, "B1", "AS1", midHeaderStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AT1", "AT1", rightHeaderStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below section formats the item code column on the far left.
	err = f.SetCellStyle(sheetName, "A2", "A"+lastRow, itemCodeStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below section of code sets the far right and bottom bottom border.
	err = f.SetCellStyle(sheetName, "AU1", "AU"+lastRow, rightMostBorder)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "A"+lastRowPlus1, "AS"+lastRowPlus1, bottomMostBorder)
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets the main row standard format.
	err = f.SetCellStyle(sheetName, "B2", "AS"+lastRow, rowStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets the dashed divider column styles grouping columns together.
	err = f.SetCellStyle(sheetName, "U2", "W"+lastRow, dividerStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "P2", "P"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "Q2", "Q"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "X2", "AI"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AJ2", "AJ"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AK2", "AK"+lastRow, dividerRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AL2", "AL"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AM2", "AM"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AO2", "AO"+lastRow, dividerStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AP2", "AP"+lastRow, dividerStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AQ2", "AQ"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AS2", "AS"+lastRow, dividerLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below code sets the currency format for the item prices columns.
	err = f.SetCellStyle(sheetName, "S2", "S"+lastRow, usdLeftStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "T2", "T"+lastRow, usdRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below section of code sets the decimal format column styles.
	err = f.SetCellStyle(sheetName, "AK2", "AK"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AM2", "AM"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AR2", "AR"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle(sheetName, "AT2", "AT"+lastRow, decimalRightStyle)
	if err != nil {
		fmt.Println(err)
	}

	// The below code starts a loop through all the rows on the worksheet for condition based operations.
	for rowIndex := range rows {

		spLinkCell, err := excelize.CoordinatesToCellName(18, rowIndex+2)
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
		for i := 1; i <= 2; i++ {
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
			if val == "TRUE" && i == 2 {
				err = f.SetCellStyle(sheetName, cell, cell, trueRightStyle)
				if err != nil {
					fmt.Println(err)
				}
			} else if val == "FALSE" && i == 2 {
				err = f.SetCellStyle(sheetName, cell, cell, falseRightStyle)
				if err != nil {
					fmt.Println(err)
				}
			}
		}

		for i := 1; i <= 3; i += 2 {
			cell, err := excelize.CoordinatesToCellName(42+i, rowIndex+2)
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

	err = f.SetColWidth(sheetName, "C", "C", 18.29+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "D", "D", 14.57+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "E", "E", 14.57+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "F", "H", 32+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "I", "I", 7.86+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "J", "J", 7.71+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "K", "K", 8+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "L", "M", 9+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "N", "N", 7.14+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "O", "O", 10+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "P", "P", 14.57+colWidthAdjust)
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

	err = f.SetColWidth(sheetName, "AA", "AM", 4.29+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "AK", "AK", 13.29+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "AM", "AM", 12+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "AR", "AR", 12+colWidthAdjust)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColWidth(sheetName, "AT", "AT", 12+colWidthAdjust)
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

	logoFile, err := os.ReadFile(filepath.Join(parent, "Shiner-PST-2025", "logos", "SHINER_LOGO_BLK_LLC.png"))
	if err != nil {
		fmt.Println(err)
	}

	err = f.AddPictureFromBytes(sheetName, "B1", &excelize.Picture{
		Extension: ".png",
		File:      logoFile,
		Format: &excelize.GraphicOptions{
			AltText: "Shiner LLC Logo",
			ScaleX:  0.05,
			ScaleY:  0.05,
			OffsetX: 11,
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
					Text: "Shiner LLC PST " + formattedTime,
					Font: &excelize.Font{
						Bold:   true,
						Italic: true,
						Size:   14,
						Color:  "#087329",
					},
				},
			},
			Width:  415,
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

	// The below section of code creates outlining column groups on the worksheet for pricing columns.
	err = f.SetColOutlineLevel(sheetName, "T", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "U", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "V", 2)
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetColOutlineLevel(sheetName, "W", 2)
	if err != nil {
		fmt.Println(err)
	}

	// The below section of code creates outlining column groups on the worksheet for 30d shipped qtys.
	err = f.SetColOutlineLevel(sheetName, "S", 2)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetColOutlineLevel(sheetName, "T", 2)
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

	err = f.SetColOutlineLevel(sheetName, "AE", 2)
	if err != nil {
		fmt.Println(err)
	}

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
	index := 11
	err = f.SetSheetProps(sheetName, &excelize.SheetPropsOptions{
		TabColorIndexed: &index,
	})
	if err != nil {
		fmt.Println(err)
	}

	return nil
}