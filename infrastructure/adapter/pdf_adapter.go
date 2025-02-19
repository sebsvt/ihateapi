package adapter

import (
	"context"
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/sebsvt/ihateapi/infrastructure/port"
)

// PDFAdapter defines the methods that any PDF tool should implement
type pdfCpuAdapter struct {
	cfg *model.Configuration
}

func NewPDFCpuAdapter() port.PDFPort {
	cfg := model.NewDefaultConfiguration()
	return &pdfCpuAdapter{cfg: cfg}
}

// Compress implements port.PDFPort.
func (p *pdfCpuAdapter) Compress(ctx context.Context, inputFile string, outputFile string) error {
	fmt.Println("Compressing", inputFile, "to", outputFile)
	return nil
}

// Merge implements port.PDFPort.
func (p *pdfCpuAdapter) Merge(ctx context.Context, inputFiles []string, outputFile string) error {
	fmt.Println("Merging", inputFiles, "to", outputFile)
	return nil
}

// Split implements port.PDFPort.
func (p *pdfCpuAdapter) Split(ctx context.Context, inputFile string, outputDir string) error {
	fmt.Println("Splitting", inputFile, "to", outputDir)
	return nil
}
