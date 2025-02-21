package service

import (
	"bytes"
	"errors"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type pdfService struct{}

func NewPdfService() PdfService {
	return &pdfService{}
}

// Compress implements PdfService.
func (srv *pdfService) Compress(file []byte) ([]byte, error) {
	// TODO: Implement compression
	var output bytes.Buffer
	reader := bytes.NewReader(file)

	config := model.NewDefaultConfiguration()
	config.Optimize = true
	config.OptimizeDuplicateContentStreams = true
	config.OptimizeResourceDicts = true

	if err := api.Optimize(reader, &output, config); err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

// Merge implements PdfService.
func (srv *pdfService) Merge(files [][]byte) ([]byte, error) {
	// TODO: Implement merge
	if len(files) < 2 {
		return nil, errors.New("at least 2 files are required to merge")
	}

	// convert all files to io.ReadSeeker
	var readers []io.ReadSeeker
	for _, file := range files {
		readers = append(readers, bytes.NewReader(file))
	}

	// create a output buffer
	var output bytes.Buffer

	// merge the files
	if err := api.MergeRaw(readers, &output, false, nil); err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

// Split implements PdfService.
func (srv *pdfService) Split(file []byte) ([][]byte, error) {
	return nil, nil
}
