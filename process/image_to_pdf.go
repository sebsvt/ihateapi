package process

import (
	"bytes"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func ImageToPDF(imageBytes []byte) ([]byte, error) {
	// convert image to pdf by using pdfcpu
	pdfBytes := &bytes.Buffer{}
	// Create a single reader for the image
	imageReader := bytes.NewReader(imageBytes)
	err := api.ImportImages(nil, pdfBytes, []io.Reader{imageReader}, nil, nil)
	if err != nil {
		return nil, err
	}

	return pdfBytes.Bytes(), nil
}
