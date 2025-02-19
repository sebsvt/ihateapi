package adapter

import (
	"context"
	"fmt"

	"github.com/sebsvt/ihateapi/infrastructure/port"
)

type unidocAdapter struct {
}

func NewUnidocAdapter() port.PDFPort {
	return &unidocAdapter{}
}

// Compress implements port.PDFPort.
func (u *unidocAdapter) Compress(ctx context.Context, inputFile string, outputFile string, level string) error {
	fmt.Println("Compressing", inputFile, "to", outputFile, "with level", level)
	return nil
}

// Merge implements port.PDFPort.
func (u *unidocAdapter) Merge(ctx context.Context, inputFiles []string, outputFile string) error {
	fmt.Println("Merging", inputFiles, "to", outputFile)
	return nil
}

// Split implements port.PDFPort.
func (u *unidocAdapter) Split(ctx context.Context, inputFile string, outputDir string, pageNrs int) error {
	fmt.Println("Splitting", inputFile, "to", outputDir, "with page numbers", pageNrs)
	return nil
}
