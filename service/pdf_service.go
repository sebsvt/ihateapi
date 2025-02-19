package service

import "context"

type PDFService interface {
	Merge(ctx context.Context, pdfs []string, outputFile string) error
	Split(ctx context.Context, pdf string, outputDir string, pageNrs int) error
	Compress(ctx context.Context, pdf string, outputFile string, level string) error
}
