package adapter

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/sebsvt/ihateapi/domain"
)

type pdfProcessorPdfcpu struct{}

func NewPdfProcessorPdfcpu() domain.PDFProcessorPort {
	return &pdfProcessorPdfcpu{}
}

// Merge implements domain.PDFProcessorPort.
// We will get the files from cloud and use this function to merge them and return the merged file back to the cloud.
func (p *pdfProcessorPdfcpu) Merge(files []string, outputFilename string) (string, error) {
	if len(files) < 2 {
		return "", domain.ErrMergeNotEnoughFiles
	}
	// merge the files
	if err := api.MergeCreateFile(files, outputFilename, false, nil); err != nil {
		return "", err
	}
	return outputFilename, nil
}

// Compress implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) Compress(file string) (string, error) {
	if err := api.OptimizeFile(file, "optimized.pdf", nil); err != nil {
		return "", err
	}
	return file, nil
}

// ImageToPDF implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) ImageToPDF(file string) (string, error) {
	// skip for now
	return "", nil
}

// OCR implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) OCR(file string) (string, error) {
	// skip for now
	return "", nil
}

// PDFa implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) PDFa(file string) (string, error) {
	// skip for now
	return "", nil
}

// PDFtoJPG implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) PDFtoJPG(file string) (string, error) {
	panic("unimplemented")
}

// PageNumber implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) PageNumber(file string) (string, error) {
	panic("unimplemented")
}

// Protect implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) Protect(file string) (string, error) {
	panic("unimplemented")
}

// Repair implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) Repair(file string) (string, error) {
	panic("unimplemented")
}

// Rotate implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) Rotate(file string, angle int) (string, error) {
	panic("unimplemented")
}

// Split implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) Split(file string) (string, error) {
	panic("unimplemented")
}

// Unlock implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) Unlock(file string) (string, error) {
	if err := api.DecryptFile(file, "unlocked.pdf", nil); err != nil {
		return "", err
	}
	return "unlocked.pdf", nil
}

// ValidatePDFa implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) ValidatePDFa(file string) (string, error) {
	panic("unimplemented")
}

// Watermark implements domain.PDFProcessorPort.
func (p *pdfProcessorPdfcpu) Watermark(file string, text string) (string, error) {
	panic("unimplemented")
}
