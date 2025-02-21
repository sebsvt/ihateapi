package process_test

import (
	"os"
	"testing"

	"github.com/sebsvt/ihateapi/process"
)

func TestImageToPDF(t *testing.T) {
	imageBytes, err := os.ReadFile("../assets/cat.jpg")
	if err != nil {
		t.Fatalf("Error reading image file: %v", err)
	}

	pdfBytes, err := process.ImageToPDF(imageBytes)
	if err != nil {
		t.Fatalf("Error converting image to PDF: %v", err)
	}

	os.WriteFile("../assets/test.pdf", pdfBytes, 0644)

	t.Logf("PDF file saved to: %s", "../assets/test.pdf")
}
