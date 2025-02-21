package repository

import "fmt"

type fileStorageRepositoryMock struct {
	storage map[string]map[string][]byte
}

func NewFileStorageRepositoryMock() FileStorageRepository {
	return &fileStorageRepositoryMock{
		storage: make(map[string]map[string][]byte),
	}
}

// Download implements FileStorageRepository.
func (repo *fileStorageRepositoryMock) Download(bucketName string, objectName string) ([]byte, error) {
	if bucket, has := repo.storage[bucketName]; has {
		if data, has := bucket[objectName]; has {
			return data, nil
		}
	}
	return nil, fmt.Errorf("object not found")
}

// GetSignedURL implements FileStorageRepository.
func (repo *fileStorageRepositoryMock) GetSignedURL(bucketName string, objectName string, expirationMinutes int) (string, error) {
	return fmt.Sprintf("https://console.aiselena.com/%s/%s?expiry=%d", bucketName, objectName, expirationMinutes), nil
}

// ObjectExists implements FileStorageRepository.
func (repo *fileStorageRepositoryMock) ObjectExists(bucketName string, objectName string) (bool, error) {
	if bucket, has := repo.storage[bucketName]; has {
		_, exists := bucket[objectName]
		return exists, nil
	}
	return false, fmt.Errorf("bucket not found")
}

// Upload implements FileStorageRepository.
func (repo *fileStorageRepositoryMock) Upload(bucketName string, objectName string, data []byte) error {
	if _, has := repo.storage[bucketName]; !has {
		repo.storage[bucketName] = make(map[string][]byte)
	}
	repo.storage[bucketName][objectName] = data
	return nil
}
