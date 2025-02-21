package process_test

import (
	"os"
	"testing"

	"github.com/sebsvt/ihateapi/process"
)

func TestCompressPDF(t *testing.T) {
	pdfBytes, err := os.ReadFile("../assets/file1.pdf")
	if err != nil {
		t.Fatalf("Error reading PDF file: %v", err)
	}
	compressed, err := process.CompressPDF(pdfBytes)
	if err != nil {
		t.Fatalf("Error compressing PDF file: %v", err)
	}

	// save the compressed PDF to a file
	err = os.WriteFile("compressed.pdf", compressed, 0644)
	if err != nil {
		t.Fatalf("Error saving compressed PDF file: %v", err)
	}
}
