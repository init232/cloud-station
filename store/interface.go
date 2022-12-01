package store

type Uploader interface {
	Upload(BucketName string, ObjectKey string, FileName string) error
}
