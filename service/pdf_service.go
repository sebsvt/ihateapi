package service

import "context"

type PDFService interface {
	Merge(ctx context.Context, pdfs []string) error
	Split(ctx context.Context, pdf string) error
	Compress(ctx context.Context, pdf string) error
}
