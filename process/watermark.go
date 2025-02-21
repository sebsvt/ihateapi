package process

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func AddWatermark(pdfBytes []byte, watermarkImage []byte, watermarkText string) ([]byte, error) {
	conf := model.NewDefaultConfiguration()
	tempDir, err := os.MkdirTemp("", "watermark_*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create input PDF file
	inputFile, err := os.CreateTemp(tempDir, "input*.pdf")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp input file: %v", err)
	}
	inputPath := inputFile.Name()
	if _, err := inputFile.Write(pdfBytes); err != nil {
		return nil, fmt.Errorf("failed to write input PDF: %v", err)
	}
	inputFile.Close()

	outputFile := inputPath + "_out.pdf"

	// If watermark image is provided
	if len(watermarkImage) > 0 {
		// Create temp watermark image file
		imgFile, err := os.CreateTemp(tempDir, "watermark*.png")
		if err != nil {
			return nil, fmt.Errorf("failed to create temp watermark file: %v", err)
		}
		if _, err := imgFile.Write(watermarkImage); err != nil {
			return nil, fmt.Errorf("failed to write watermark image: %v", err)
		}
		imgFile.Close()

		if err := api.AddImageWatermarksFile(inputPath, outputFile, []string{imgFile.Name()}, false, "", "", conf); err != nil {
			return nil, fmt.Errorf("failed to add image watermark: %v", err)
		}
		// Update input path for next operation
		inputPath = outputFile
	}

	// If watermark text is provided
	if watermarkText != "" {
		if err := api.AddTextWatermarksFile(inputPath, outputFile, []string{watermarkText}, true, "points:72", "center:t", conf); err != nil {
			return nil, fmt.Errorf("failed to add text watermark: %v", err)
		}
	}

	// Read final output
	result, err := os.ReadFile(outputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read output file: %v", err)
	}

	return result, nil
}
