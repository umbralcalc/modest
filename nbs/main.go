package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/umbralcalc/modest/pkg/plotly"
	"github.com/umbralcalc/modest/pkg/renderer"
)

func plotSomething() {
	csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-02,32,321.31,54320
"United Kingdom",2012-02-03,17,18.2,12345
"United States",2012-02-04,32,321.31,54320
"United Kingdom",2012-02-05,NA,18.2,12345
"United States",2012-02-06,32,321.31,54320
"United States",2012-02-07,32,321.31,54320
"Spain",2012-02-08,66,555.42,00241
`

	// Load CSV data into dataframe
	df := dataframe.ReadCSV(strings.NewReader(csvStr))

	// Extract Date and Amount columns
	dateCol := df.Col("Date").Records()
	amountCol := df.Col("Amount").Float()

	// Convert dates and amounts to JavaScript arrays
	var dates []string
	var amounts []string
	for i, dateStr := range dateCol {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Fatalf("error parsing date: %v", err)
		}
		dates = append(dates, fmt.Sprintf(`"%s"`, date.Format("2006-01-02")))
		amounts = append(amounts, fmt.Sprintf(`%f`, amountCol[i]))
	}

	// Create the JavaScript code for Plotly
	plotCode := plotly.Plot{
		Dates:   dates,
		Amounts: amounts,
	}
	htmlContent := renderer.RenderDoc([]fmt.Stringer{plotCode})

	// Create a temporary HTML file to store the content
	tmpFile, err := os.CreateTemp("", "amount_over_time_*.html")
	if err != nil {
		log.Fatalf("error creating temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Ensure the temp file is deleted after the program ends

	// Write the HTML content to the temporary file
	_, err = tmpFile.Write([]byte(htmlContent))
	if err != nil {
		log.Fatalf("error writing HTML content: %v", err)
	}
	tmpFile.Close()

	// Open the HTML file in the default web browser
	err = exec.Command("open", tmpFile.Name()).Start() // For macOS & Linux; use "start" on Windows
	if err != nil {
		log.Fatalf("error opening HTML file in browser: %v", err)
	}

	// Keep the program running until the user presses Enter
	fmt.Println("Press Enter to exit...")
	fmt.Scanln() // Wait for Enter key press to exit the program
}

func main() {
	plotSomething()
}
