package process

import (
	"os"
	"testing"
)

func TestAddWatermark(t *testing.T) {
	// testing add text watermark and image watermark to file and save it to new file
	// read pdf file
	pdfBytes, err := os.ReadFile("../assets/merged.pdf")
	if err != nil {
		t.Fatalf("Failed to read pdf file: %v", err)
	}

	watermarkImage := []byte{}
	watermarkText := "Test Watermark"

	result, err := AddWatermark(pdfBytes, watermarkImage, watermarkText)
	if err != nil {
		t.Fatalf("Failed to add watermark: %v", err)
	}

	// save result to new file
	err = os.WriteFile("result.pdf", result, 0644)
	if err != nil {
		t.Fatalf("Failed to save result: %v", err)
	}

}
