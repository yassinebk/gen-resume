package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

// Custom template functions
var funcMap = template.FuncMap{
	"join": func(items []string, separator string) string {
		return strings.Join(items, separator)
	},
	"isLast": func(index, length int) bool {
		return index == length-1
	},
}

func main() {
	app := &cli.App{
		Name:  "resume-generator",
		Usage: "Generate a PDF resume from YAML data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Input YAML file",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output PDF file",
				Value:   "resume.pdf",
			},
			&cli.StringFlag{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "HTML template file",
				Value:   "templates/index.html",
			},
		},
		Action: generateResume,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func generateResume(c *cli.Context) error {
	// Read YAML file
	yamlData, err := os.ReadFile(c.String("input"))
	if err != nil {
		return fmt.Errorf("error reading YAML file: %v", err)
	}

	var resume Resume
	err = yaml.Unmarshal(yamlData, &resume)
	if err != nil {
		return fmt.Errorf("error parsing YAML: %v", err)
	}

	// Parse template

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles(filepath.Join(exPath, c.String("template")))
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	// Execute template
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, resume)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	// Initialize PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return fmt.Errorf("error creating PDF generator: %v", err)
	}

	// Set PDF options
	pdfg.Dpi.Set(300)
	pdfg.Cover.DisableSmartShrinking.Set(true)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Cover.Zoom.Set(1.0)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	// Add HTML content to PDF generator
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(buf.Bytes()))
	page.EnableLocalFileAccess.Set(true)
	pdfg.AddPage(page)

	// Generate PDF
	err = pdfg.Create()
	if err != nil {
		return fmt.Errorf("error generating PDF: %v", err)
	}

	// Write PDF to file
	err = pdfg.WriteFile(c.String("output"))
	if err != nil {
		return fmt.Errorf("error writing PDF file: %v", err)
	}

	fmt.Printf("Successfully generated resume at: %s\n", c.String("output"))
	return nil
}
