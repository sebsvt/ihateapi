package process_test

import (
	"os"
	"testing"

	"github.com/sebsvt/ihateapi/process"
)

func TestMergePDF(t *testing.T) {
	files := []string{"../assets/file1.pdf", "../assets/file2.pdf"}
	merged, err := process.MergePDF(files)
	if err != nil {
		t.Fatalf("Error merging PDF files: %v", err)
	}

	// save the merged PDF to a file
	err = os.WriteFile("merged.pdf", merged, 0644)
	if err != nil {
		t.Fatalf("Error saving merged PDF: %v", err)
	}
}
