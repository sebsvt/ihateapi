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
func (p *pdfService) Compress(ctx context.Context, pdf string) error {
	return p.pdfPort.Compress(ctx, pdf, pdf)
}

// Merge implements PDFService.
func (p *pdfService) Merge(ctx context.Context, pdfs []string) error {
	return p.pdfPort.Merge(ctx, pdfs, "merged.pdf")
}

// Split implements PDFService.
func (p *pdfService) Split(ctx context.Context, pdf string) error {
	return p.pdfPort.Split(ctx, pdf, "split")
}
