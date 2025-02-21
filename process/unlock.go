package process

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func UnlockPdf(pdfBytes []byte, password string) ([]byte, error) {
	// Create temp files
	tempDir, err := os.MkdirTemp("", "unlock_*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Write input PDF
	inputFile, err := os.CreateTemp(tempDir, "input*.pdf")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %v", err)
	}
	if _, err := inputFile.Write(pdfBytes); err != nil {
		return nil, fmt.Errorf("failed to write PDF: %v", err)
	}
	inputFile.Close()

	outputFile := inputFile.Name() + "_unlocked.pdf"

	// Configure with password
	conf := model.NewDefaultConfiguration()
	conf.UserPW = password
	conf.OwnerPW = password

	// Decrypt the PDF
	if err := api.DecryptFile(inputFile.Name(), outputFile, conf); err != nil {
		return nil, fmt.Errorf("failed to decrypt PDF: %v", err)
	}

	// Read the decrypted file
	result, err := os.ReadFile(outputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read decrypted file: %v", err)
	}

	return result, nil
}
