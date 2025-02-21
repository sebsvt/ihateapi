package process

import (
	"bytes"
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergePDF(files []string) ([]byte, error) {
	// create a buffer to store the merged PDF
	var buf bytes.Buffer
	err := api.Merge(files[0], files[1:], &buf, nil, false)
	if err != nil {
		log.Println("Error merging PDF files:", err)
		return nil, err
	}

	// save the merged PDF to a file
	// err = os.WriteFile("merged.pdf", buf.Bytes(), 0644)
	return buf.Bytes(), nil
}
