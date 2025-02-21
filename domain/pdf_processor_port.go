package domain

import "errors"

// Status represents the processing state of a PDF file
type Status string

const (
	StatusSuccess       Status = "file has been processed successfully"
	StatusWaiting       Status = "file is waiting to be processed"
	StatusWrongPassword Status = "file has not been processed because needed a password and was not provided or incorrect"
	StatusTimeout       Status = "file has not been processed correctly because it took more than your time limit to process it"
	StatusNotFound      Status = "file has not been processed because has not been found in the server"
	StatusDamaged       Status = "file has not been processed because it was damaged or we were unable to read it"
	StatusNoImages      Status = "file has not been processed because we couldn't find any images to extract. Maybe there are vectors?"
	StatusOutOfRange    Status = "file has not been processed because some of the ranges do not match the number of pages"
	StatusNonConformant Status = "PDF file validation has not passed against PDF/A conformance provided"
	StatusUnknown       Status = "unknown error"
)

var (
	ErrMergeNotEnoughFiles = errors.New("error merging files: the files provided are less than 2 files")
)

type PDFProcessorPort interface {
	Merge(files []string, outputFilename string) (string, error)
	Split(file string) (string, error)
	Compress(file string) (string, error)
	OCR(file string) (string, error)
	PDFtoJPG(file string) (string, error)
	ImageToPDF(file string) (string, error)
	Unlock(file string) (string, error)
	PageNumber(file string) (string, error)
	Watermark(file string, text string) (string, error)
	Repair(file string) (string, error)
	Rotate(file string, angle int) (string, error)
	Protect(file string) (string, error)
	PDFa(file string) (string, error)
	ValidatePDFa(file string) (string, error)
}
