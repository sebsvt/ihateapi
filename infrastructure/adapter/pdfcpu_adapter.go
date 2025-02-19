package adapter

import (
	"context"

	"github.com/pdfcpu/pdfcpu/pkg/api"
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

// Split implements port.PDFPort.
func (p *pdfCpuAdapter) Split(ctx context.Context, inputFile string, outputDir string, pageNrs int) error {
	return api.SplitFile(inputFile, outputDir, pageNrs, p.cfg)
}

// Merge implements port.PDFPort.
func (p *pdfCpuAdapter) Merge(ctx context.Context, inputFiles []string, outputFile string) error {
	return api.MergeCreateFile(inputFiles, outputFile, false, p.cfg)
}

// Compress implements port.PDFPort.
func (p *pdfCpuAdapter) Compress(ctx context.Context, inputFile string, outputFile string, level string) error {
	conf := model.NewDefaultConfiguration()
	conf.Optimize = true
	switch level {
	case "extreme":
		conf.OptimizeResourceDicts = true
		conf.OptimizeDuplicateContentStreams = true
	case "recommended":
		conf.OptimizeResourceDicts = true
		conf.OptimizeDuplicateContentStreams = false
	case "low":
		conf.OptimizeResourceDicts = false
		conf.OptimizeDuplicateContentStreams = false
	default:
		// if level is not recognized, use recommended
		conf.OptimizeResourceDicts = true
		conf.OptimizeDuplicateContentStreams = false

	}
	err := api.OptimizeFile(inputFile, outputFile, conf)
	return err
}
