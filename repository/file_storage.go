package repository

type FileStorageRepository interface {
	// Upload uploads a file to cloud storage
	Upload(bucketName, objectName string, data []byte) error

	// Download retrieves a file from cloud storage
	Download(bucketName, objectName string) ([]byte, error)

	// GetSignedURL generates a temporary signed URL for object access
	GetSignedURL(bucketName, objectName string, expirationMinutes int) (string, error)

	// ObjectExists checks if an object exists in a bucket
	ObjectExists(bucketName, objectName string) (bool, error)
}
