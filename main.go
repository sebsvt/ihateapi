package main

import (
	"context"

	"github.com/sebsvt/ihateapi/infrastructure/adapter"
	"github.com/sebsvt/ihateapi/service"
)

func main() {
	// Adapters
	pdfCpuAdapter := adapter.NewPDFCpuAdapter()
	pdfUnidocAdapter := adapter.NewPDFUnidocAdapter()
	_ = pdfCpuAdapter

	// Services

	pdfSrv := service.NewPDFService(pdfUnidocAdapter)
	// pdfSrv := service.NewPDFService(pdfCpuAdapter)

	// please do not forget to change the path to the input file
	pdfSrv.Split(context.Background(), "../assets/merged.pdf", "../assets/out", 1)
}
