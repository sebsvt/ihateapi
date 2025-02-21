package repository

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type fileStorageRepository struct {
	client *storage.Client
}

func NewFileStorageRepository(credentialsPath string) (FileStorageRepository, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create storage client: %v", err)
	}

	return &fileStorageRepository{
		client: client,
	}, nil
}

// Download implements FileStorageRepository.
func (repo *fileStorageRepository) Download(bucketName string, objectName string) ([]byte, error) {
	ctx := context.Background()
	bucket := repo.client.Bucket(bucketName)
	obj := bucket.Object(objectName)

	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader: %v", err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read object: %v", err)
	}

	return data, nil
}

// GetSignedURL implements FileStorageRepository.
func (repo *fileStorageRepository) GetSignedURL(bucketName string, objectName string, expirationMinutes int) (string, error) {
	opts := &storage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().Add(time.Duration(expirationMinutes) * time.Minute),
	}

	url, err := storage.SignedURL(bucketName, objectName, opts)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %v", err)
	}

	return url, nil
}

// ObjectExists implements FileStorageRepository.
func (repo *fileStorageRepository) ObjectExists(bucketName string, objectName string) (bool, error) {
	ctx := context.Background()
	bucket := repo.client.Bucket(bucketName)
	obj := bucket.Object(objectName)

	_, err := obj.Attrs(ctx)
	if err == storage.ErrObjectNotExist {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("failed to get object attributes: %v", err)
	}

	return true, nil
}

// Upload implements FileStorageRepository.
func (repo *fileStorageRepository) Upload(bucketName string, objectName string, data []byte) error {
	ctx := context.Background()
	bucket := repo.client.Bucket(bucketName)
	obj := bucket.Object(objectName)

	writer := obj.NewWriter(ctx)
	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("failed to write to object: %v", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %v", err)
	}

	return nil
}
