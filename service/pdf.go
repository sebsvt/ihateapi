package service

type PdfService interface {
	Merge(files [][]byte) ([]byte, error)
	Split(file []byte) ([][]byte, error)
	Compress(file []byte) ([]byte, error)
}
