package service_test

import (
	"context"
	"testing"

	"github.com/sebsvt/ihateapi/infrastructure/adapter"
	"github.com/sebsvt/ihateapi/service"
)

func TestSplitWithPDFCPUAdapterInPDFService(t *testing.T) {
	type testCase struct {
		name        string
		inputFile   string
		outputDir   string
		pageNrs     int
		expectedErr error
	}
	testcases := []testCase{
		{
			name:      "Splitting a PDF file with 1 page should not return an error",
			inputFile: "../assets/merged.pdf",
			outputDir: "../assets/out",
			pageNrs:   1,
		},
	}

	pdfCpuAdapter := adapter.NewPDFCpuAdapter()
	pdfSrv := service.NewPDFService(pdfCpuAdapter)
	for _, tc := range testcases {
		err := pdfSrv.Split(context.Background(), tc.inputFile, tc.outputDir, tc.pageNrs)
		if err != tc.expectedErr {
			t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
		}
	}
}
