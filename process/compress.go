package process

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func CompressPDF(input []byte) ([]byte, error) {
	// Configure PDF processor to handle version 1.4
	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationRelaxed
	conf.Reader15 = true // Enables compatibility with older PDF versions

	// Create a temp file for input
	inputFile, err := os.CreateTemp("", "input*.pdf")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	// Write input bytes to temp file
	if _, err := inputFile.Write(input); err != nil {
		return nil, fmt.Errorf("failed to write to temp input file: %v", err)
	}
	inputFile.Close()

	// Create a temp file for output
	outputFile, err := os.CreateTemp("", "output*.pdf")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp output file: %v", err)
	}
	defer os.Remove(outputFile.Name())
	outputFile.Close()
	// Optimize the PDF
	if err := api.OptimizeFile(inputFile.Name(), outputFile.Name(), conf); err != nil {
		return nil, fmt.Errorf("failed to optimize PDF: %v", err)
	}

	// Read the optimized file
	compressed, err := os.ReadFile(outputFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read compressed file: %v", err)
	}

	return compressed, nil
}
