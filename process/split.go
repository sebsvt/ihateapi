package process

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func SplitPDF(input []byte) ([][]byte, error) {
	// Configure PDF processor
	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationRelaxed
	conf.Reader15 = true

	// Create temp directory for split files
	tempDir, err := os.MkdirTemp("", "split_*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create input file
	inputFile, err := os.CreateTemp(tempDir, "input*.pdf")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp input file: %v", err)
	}

	// Write input bytes
	if _, err := inputFile.Write(input); err != nil {
		return nil, fmt.Errorf("failed to write to temp input file: %v", err)
	}
	inputFile.Close()

	// Split the PDF - this will create files in the temp directory
	if err := api.SplitFile(inputFile.Name(), tempDir, 1, conf); err != nil {
		return nil, fmt.Errorf("failed to split PDF: %v", err)
	}

	// Read all split files
	var splitFiles [][]byte
	files, err := filepath.Glob(filepath.Join(tempDir, "*.pdf"))
	if err != nil {
		return nil, fmt.Errorf("failed to find split PDF files: %v", err)
	}

	for _, file := range files {
		splitBytes, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read split file: %v", err)
		}
		splitFiles = append(splitFiles, splitBytes)
	}

	return splitFiles, nil
}
