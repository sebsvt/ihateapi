package service_test

import (
	"os"
	"testing"

	"github.com/sebsvt/ihateapi/service"
)

func TestPdfMerge(t *testing.T) {
	pdfService := service.NewPdfService()

	// read the files
	file1, err := os.ReadFile("../assets/file1.pdf")
	if err != nil {
		t.Fatal(err)
	}

	// read file 2
	file2, err := os.ReadFile("../assets/file2.pdf")
	if err != nil {
		t.Fatal(err)
	}

	// merge the files
	merged, err := pdfService.Merge([][]byte{file1, file2})
	if err != nil {
		t.Fatal(err)
	}

	// write the merged file
	if err := os.WriteFile("../assets/merged.pdf", merged, 0644); err != nil {
		t.Fatal(err)
	}
}

func TestPdfCompress(t *testing.T) {
	pdfService := service.NewPdfService()

	// read the file
	file, err := os.ReadFile("../assets/merged.pdf")
	if err != nil {
		t.Fatal(err)
	}

	// compress the file
	compressed, err := pdfService.Compress(file)
	if err != nil {
		t.Fatal(err)
	}

	// write the compressed file
	if err := os.WriteFile("../assets/compressed.pdf", compressed, 0644); err != nil {
		t.Fatal(err)
	}
}
