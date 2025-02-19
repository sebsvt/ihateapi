package port

import "context"

// PDFPort defines the methods that any PDF tool should implement
type PDFPort interface {
	Merge(ctx context.Context, inputFiles []string, outputFile string) error
	Split(ctx context.Context, inputFile, outputDir string, pageNrs int) error
	Compress(ctx context.Context, inputFile, outputFile, level string) error
}
