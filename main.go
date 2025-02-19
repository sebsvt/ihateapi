package main

import (
	"context"

	"github.com/sebsvt/ihateapi/infrastructure/adapter"
	"github.com/sebsvt/ihateapi/service"
)

func main() {
	pdfService := service.NewPDFService(adapter.NewPDFCpuAdapter())
	pdfService.Compress(context.Background(), "test.pdf")
}
