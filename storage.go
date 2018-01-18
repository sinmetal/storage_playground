package main

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
)

// StorageService is Cloud Storage Serviceを提供するstruct
type StorageService struct {
	C *storage.Client
}

// NewStorageService is Cloud Storage Serviceを作成
func NewStorageService(ctx context.Context) (*StorageService, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client")
	}

	return &StorageService{
		client,
	}, nil
}

// Upload is ファイルをCloud StorageにUploadする
func (s *StorageService) Upload(ctx context.Context, encryptKey []byte, bucketName string, objectName string, file []byte) (int, error) {
	bucket := s.C.Bucket(bucketName)
	obj := bucket.Object(objectName).Key(encryptKey)
	w := obj.NewWriter(ctx)

	size, err := w.Write(file)
	if err != nil {
		return 0, errors.Wrapf(err, "file write error")
	}

	if err := w.Close(); err != nil {
		return 0, errors.Wrapf(err, "file writer close error")
	}

	return size, nil
}

// Download is ファイルをCloud StorageからDownloadする
func (s *StorageService) Download(ctx context.Context, encryptKey []byte, bucketName string, objectName string) ([]byte, error) {
	obj := s.C.Bucket(bucketName).Object(objectName)
	rc, err := obj.Key(encryptKey).NewReader(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "file new reader error")
	}

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, errors.Wrap(err, "file read error")
	}

	if err := rc.Close(); err != nil {
		return nil, errors.Wrap(err, "file reader close error")
	}
	return data, nil
}
