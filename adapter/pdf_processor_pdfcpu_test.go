package adapter

import (
	"testing"

	"github.com/sebsvt/ihateapi/domain"
	"github.com/stretchr/testify/assert"
)

func TestMergeFunction(t *testing.T) {
	type testCase struct {
		name          string
		files         []string
		expectedError error
	}

	testcases := []testCase{
		{
			name:          "merge two files should not return an error",
			files:         []string{"../asset_temp/file1.pdf", "../asset_temp/file2.pdf"},
			expectedError: nil,
		},
		{
			name:          "can not merge when number of file is less than 2 and should return an error",
			files:         []string{"../asset_temp/file1.pdf"},
			expectedError: domain.ErrMergeNotEnoughFiles,
		},
	}

	pdfProcessor := NewPdfProcessorPdfcpu()

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := pdfProcessor.Merge(tc.files, "merged.pdf")
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			}
		})
	}
}
