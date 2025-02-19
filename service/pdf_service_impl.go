package service

import (
	"context"

	"github.com/sebsvt/ihateapi/infrastructure/port"
)

type pdfService struct {
	pdfPort port.PDFPort
}

func NewPDFService(pdfPort port.PDFPort) PDFService {
	return &pdfService{pdfPort: pdfPort}
}

// Compress implements PDFService.
func (p *pdfService) Compress(ctx context.Context, pdf string, outputFile string, level string) error {
	return p.pdfPort.Compress(ctx, pdf, outputFile, level)
}

// Merge implements PDFService.
func (p *pdfService) Merge(ctx context.Context, pdfs []string, outputFile string) error {
	return p.pdfPort.Merge(ctx, pdfs, outputFile)
}

// Split implements PDFService.
func (p *pdfService) Split(ctx context.Context, pdf string, outputDir string, pageNrs int) error {
	return p.pdfPort.Split(ctx, pdf, outputDir, pageNrs)
}
