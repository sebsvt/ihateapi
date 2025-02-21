package process_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/sebsvt/ihateapi/process"
)

func TestSplitPDF(t *testing.T) {
	pdfBytes, err := os.ReadFile("../assets/merged.pdf")
	if err != nil {
		t.Fatalf("Error reading PDF file: %v", err)
	}

	// Save in assets directory
	outputDir := "../assets/split_output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Fatalf("Error creating output directory: %v", err)
	}

	splits, err := process.SplitPDF(pdfBytes)
	if err != nil {
		t.Fatalf("Error splitting PDF file: %v", err)
	}

	// Check if we got multiple PDF files
	if len(splits) <= 1 {
		t.Fatalf("Expected multiple PDF files, but got %d files", len(splits))
	}

	// Save each split PDF to a separate file
	for i, split := range splits {
		filename := filepath.Join(outputDir, fmt.Sprintf("split_%d.pdf", i))
		err = os.WriteFile(filename, split, 0644)
		if err != nil {
			t.Fatalf("Error saving split PDF file %d: %v", i, err)
		}
		t.Logf("Saved split PDF to: %s", filename)
	}
}
